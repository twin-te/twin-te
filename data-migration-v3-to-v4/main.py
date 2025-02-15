import csv
import json
from datetime import datetime, timezone

import pandas as pd

now = datetime.now(timezone.utc).strftime("%Y-%m-%dT%H:%M:%S.%f")


def read_csv(path: str) -> pd.DataFrame:
    return pd.read_csv(path, keep_default_na=False, na_values="null")


def to_csv(df: pd.DataFrame, path: str) -> None:
    df.to_csv(
        path,
        encoding="utf-8",
        index=False,
        quoting=csv.QUOTE_NONNUMERIC,
        escapechar='"',
    )


def migrate_users() -> list[str]:
    df = read_csv("data/raw/user_users.csv")
    df = df[df["deletedAt"].isna()]
    df.rename(columns={"createdAt": "created_at"}, inplace=True)
    df["updated_at"] = now
    df = df[["id", "created_at", "updated_at"]]
    to_csv(df, "data/processed/users.csv")
    active_user_ids = df["id"].tolist()
    return active_user_ids


def migrate_user_authentications(active_user_ids: list[str]):
    df = read_csv("data/raw/user_user_authentications.csv")
    df = df[df["user_id"].isna()]
    df = df[df["user_id"].isin(active_user_ids)]
    df = df[["user_id", "provider", "social_id"]]
    to_csv(df, "data/processed/user_authentications.csv")


def migrate_sessions(active_user_ids: list[str]):
    df = read_csv("data/raw/session_session.csv")
    df = df[df["user_id"].isin(active_user_ids)]
    df["created_at"] = now
    df["updated_at"] = now
    df = df[["id", "user_id", "expired_at", "created_at", "updated_at"]]
    to_csv(df, "data/processed/sessions.csv")


def migrate_payment_users():
    df = read_csv("data/raw/donation_payment_users.csv")
    df.rename(columns={"twinte_user_id": "user_id"}, inplace=True)
    df["created_at"] = now
    df["updated_at"] = now
    df = df[["id", "user_id", "display_name", "link", "created_at", "updated_at"]]
    to_csv(df, "data/processed/payment_users.csv")


def migrate_tags(active_user_ids: list[str]) -> list[str]:
    df = read_csv("data/raw/timetables_tags.csv")
    df = df[df["user_id"].isin(active_user_ids)]
    df.rename(columns={"position": "order"}, inplace=True)
    df["created_at"] = now
    df["updated_at"] = now
    df = df[["id", "user_id", "name", "order", "created_at", "updated_at"]]
    to_csv(df, "data/processed/tags.csv")
    active_tag_ids = df["id"].tolist()
    return active_tag_ids


def migrate_registered_courses(active_user_ids: list[str]) -> list[str]:
    df = read_csv("data/raw/timetables_registered_courses.csv")
    df.rename(columns={"instractor": "instructors"}, inplace=True)

    target_indices = df[~df["schedules"].isna()].index.values
    for index in target_indices:
        schedules = json.loads(df.at[index, "schedules"])
        for schedule in schedules:
            schedule["locations"] = schedule.pop("room")
        df.at[index, "schedules"] = json.dumps(schedules, ensure_ascii=False)

    df = df[df["user_id"].isin(active_user_ids)]

    df["created_at"] = now
    df["updated_at"] = now
    df = df[
        [
            "id",
            "user_id",
            "year",
            "course_id",
            "name",
            "instructors",
            "credit",
            "methods",
            "schedules",
            "memo",
            "attendance",
            "absence",
            "late",
            "created_at",
            "updated_at",
        ]
    ]

    to_csv(df, "data/processed/registered_courses.csv")

    active_registered_course_ids = df["id"].tolist()
    return active_registered_course_ids


def migrate_registered_course_tag_ids(
    active_registered_course_ids: list[str], active_tag_ids: list[str]
):
    df = read_csv("data/raw/timetables_registered_course_tags.csv")
    df.rename(
        columns={"registered_course": "registered_course_id", "tag": "tag_id"},
        inplace=True,
    )
    df = df[["registered_course_id", "tag_id"]]
    df = df[df["registered_course_id"].isin(active_registered_course_ids)]
    df = df[df["tag_id"].isin(active_tag_ids)]
    to_csv(df, "data/processed/registered_course_tag_ids.csv")


def load_all_kdb_courses():
    years = [2019, 2020, 2021, 2022, 2023, 2024]
    all_kdb_courses = []
    for year in years:
        with open(f"data/parsed/{year}.json") as f:
            kdb_courses = json.load(f)
            for kdb_course in kdb_courses:
                kdb_course["year"] = year
            all_kdb_courses += kdb_courses
    return all_kdb_courses


def migrate_course_aggregate_found() -> set[str]:
    # 既存のDBに保存してある講義のうち現在のKdBに存在する講義はKdBのデータを使用する
    # 存在しない講義は既存のDBに保存してある講義データを修正して使用する

    all_kdb_courses = load_all_kdb_courses()
    df_existing_courses = read_csv("data/raw/course_courses.csv")

    year_and_code_to_course_id: dict[tuple[int, str], str] = {}
    for _, row in df_existing_courses.iterrows():
        year_and_code_to_course_id[(row["year"], row["code"])] = row["id"]

    courses_data = {
        "id": [],
        "year": [],
        "code": [],
        "name": [],
        "instructors": [],
        "credit": [],
        "overview": [],
        "remarks": [],
        "last_updated_at": [],
        "has_parse_error": [],
        "is_annual": [],
    }

    course_methods_data = {
        "course_id": [],
        "method": [],
    }

    course_recommended_grades_data = {
        "course_id": [],
        "recommended_grade": [],
    }

    course_schedules_data = {
        "course_id": [],
        "module": [],
        "day": [],
        "period": [],
        "locations": [],
    }

    course_ids_found: list[str] = []

    for kdb_course in all_kdb_courses:
        year_and_code = (kdb_course["year"], kdb_course["code"])
        if year_and_code not in year_and_code_to_course_id:
            continue

        course_id = year_and_code_to_course_id[year_and_code]
        course_ids_found.append(course_id)

        courses_data["id"].append(course_id)
        courses_data["year"].append(kdb_course["year"])
        courses_data["code"].append(kdb_course["code"])
        courses_data["name"].append(kdb_course["name"])
        courses_data["instructors"].append(kdb_course["instructors"])
        courses_data["credit"].append(kdb_course["credit"])
        courses_data["overview"].append(kdb_course["overview"])
        courses_data["remarks"].append(kdb_course["remarks"])
        courses_data["last_updated_at"].append(kdb_course["lastUpdatedAt"])
        courses_data["has_parse_error"].append(kdb_course["hasParseError"])
        courses_data["is_annual"].append(kdb_course["isAnnual"])

        for method in kdb_course["methods"]:
            course_methods_data["course_id"].append(course_id)
            course_methods_data["method"].append(method)

        for recommended_grade in kdb_course["recommendedGrades"]:
            course_recommended_grades_data["course_id"].append(course_id)
            course_recommended_grades_data["recommended_grade"].append(
                recommended_grade
            )

        for schedule in kdb_course["schedules"]:
            course_schedules_data["course_id"].append(course_id)
            course_schedules_data["module"].append(schedule["module"])
            course_schedules_data["day"].append(schedule["day"])
            course_schedules_data["period"].append(schedule["period"])
            course_schedules_data["locations"].append(schedule["locations"])

    df_courses_found = pd.DataFrame(data=courses_data)
    df_course_schedules_found = pd.DataFrame(data=course_schedules_data)
    df_course_methods_found = pd.DataFrame(data=course_methods_data)
    df_course_recommended_grades_found = pd.DataFrame(
        data=course_recommended_grades_data
    )

    df_courses_found["created_at"] = now
    df_courses_found["updated_at"] = now

    df_courses_found = df_courses_found[
        [
            "id",
            "year",
            "code",
            "name",
            "instructors",
            "credit",
            "overview",
            "remarks",
            "last_updated_at",
            "has_parse_error",
            "is_annual",
            "created_at",
            "updated_at",
        ]
    ]
    df_course_schedules_found = df_course_schedules_found[
        [
            "course_id",
            "module",
            "day",
            "period",
            "locations",
        ]
    ]
    df_course_methods_found = df_course_methods_found[
        [
            "course_id",
            "method",
        ]
    ]
    df_course_recommended_grades_found = df_course_recommended_grades_found[
        [
            "course_id",
            "recommended_grade",
        ]
    ]

    to_csv(df_courses_found, "data/processed/courses_found.csv")
    to_csv(df_course_schedules_found, "data/processed/course_schedules_found.csv")
    to_csv(df_course_methods_found, "data/processed/course_methods_found.csv")
    to_csv(
        df_course_recommended_grades_found,
        "data/processed/course_recommended_grades_found.csv",
    )

    course_id_set_not_found = set(df_existing_courses["id"].tolist()).difference(
        course_ids_found
    )
    return course_id_set_not_found


def migrate_course_not_found(course_id_set_not_found: set[str]) -> set[str]:
    df = read_csv("data/raw/course_courses.csv")
    df.rename(
        columns={"instructor": "instructors", "last_update": "last_updated_at"},
        inplace=True,
    )
    df = df[df["id"].isin(course_id_set_not_found)]
    df["last_updated_at"] = df["last_updated_at"].str[:-3]
    df["created_at"] = now
    df["updated_at"] = now
    df = df[
        [
            "id",
            "year",
            "code",
            "name",
            "instructors",
            "credit",
            "overview",
            "remarks",
            "last_updated_at",
            "has_parse_error",
            "is_annual",
            "created_at",
            "updated_at",
        ]
    ]
    to_csv(df, "data/processed/courses_not_found.csv")


def migrate_course_schedules_not_found(course_id_set_not_found: set[str]):
    df = read_csv("data/raw/course_course_schedules.csv")
    df = df[df["course_id"].isna()]
    df = df[df["course_id"].isin(course_id_set_not_found)]
    df.drop(columns=["id"], inplace=True)

    df_annual = df[df["module"] == "Annual"]

    modules: list[str] = []
    days: list[str] = []
    periods: list[str] = []
    rooms: list[str] = []
    course_ids: list[str] = []

    for _, row in df_annual.iterrows():
        for module in [
            "SpringA",
            "SpringB",
            "SpringC",
            "FallA",
            "FallB",
            "FallC",
        ]:
            modules.append(module)
            days.append(row["day"])
            periods.append(row["period"])
            rooms.append(row["room"])
            course_ids.append(row["course_id"])

    df_to_add = pd.DataFrame(
        data={
            "module": modules,
            "day": days,
            "period": periods,
            "room": rooms,
            "course_id": course_ids,
        }
    )

    df = df[df["module"] != "Annual"]
    df = pd.concat([df, df_to_add], axis=0)

    df.rename(columns={"room": "locations"}, inplace=True)

    df = df[
        [
            "course_id",
            "module",
            "day",
            "period",
            "locations",
        ]
    ]

    to_csv(df, "data/processed/course_schedules_not_found.csv")


def migrate_course_recommended_grades_not_found(course_id_set_not_found: set[str]):
    df = read_csv("data/raw/course_course_recommended_grades.csv")
    df = df[df["course_id"].isna()]
    df = df[df["course_id"].isin(course_id_set_not_found)]
    df.drop(columns=["id"], inplace=True)
    df.rename(columns={"grade": "recommended_grade"}, inplace=True)
    df = df[["course_id", "recommended_grade"]]
    to_csv(df, "data/processed/course_recommended_grades_not_found.csv")


def migrate_course_methods_not_found(course_id_set_not_found: set[str]):
    df = read_csv("data/raw/course_course_methods.csv")
    df = df[df["course_id"].isna()]
    df = df[df["course_id"].isin(course_id_set_not_found)]
    df.drop(columns=["id"], inplace=True)
    df = df[["course_id", "method"]]
    to_csv(df, "data/processed/course_methods_not_found.csv")


def main():
    active_user_ids = migrate_users()
    migrate_user_authentications(active_user_ids)

    migrate_sessions(active_user_ids)

    migrate_payment_users()

    active_tag_ids = migrate_tags(active_user_ids)

    active_registered_course_ids = migrate_registered_courses(active_user_ids)
    migrate_registered_course_tag_ids(active_registered_course_ids, active_tag_ids)

    course_id_set_not_found = migrate_course_aggregate_found()

    migrate_course_not_found(course_id_set_not_found)
    migrate_course_schedules_not_found(course_id_set_not_found)
    migrate_course_methods_not_found(course_id_set_not_found)
    migrate_course_recommended_grades_not_found(course_id_set_not_found)


if __name__ == "__main__":
    main()
