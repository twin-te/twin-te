import json

import pandas as pd


def read_csv(path: str) -> pd.DataFrame:
    return pd.read_csv(path, keep_default_na=False)


def to_csv(df: pd.DataFrame, path: str) -> None:
    df.to_csv(
        path,
        encoding="utf-8",
        index=False,
    )


def list_active_user_ids() -> list[str]:
    df = read_csv("data/raw/user_users.csv")
    active_user_ids: list[str] = df[df["deletedAt"] == "null"]["id"].tolist()
    return active_user_ids


def migrate_users(active_user_ids: list[str]):
    df = read_csv("data/raw/user_users.csv")
    df.rename(
        columns={"createdAt": "created_at", "deletedAt": "deleted_at"}, inplace=True
    )
    df = df[df["id"].isin(active_user_ids)]
    to_csv(df, "data/processed/users.csv")


def migrate_user_authentications(active_user_ids: list[str]):
    df = read_csv("data/raw/user_user_authentications.csv")
    df = df[df["user_id"] != "null"]
    df.drop(columns=["id"], inplace=True)
    df = df[df["user_id"].isin(active_user_ids)]
    to_csv(df, "data/processed/user_authentications.csv")


def migrate_sessions(active_user_ids: list[str]):
    df = read_csv("data/raw/session_session.csv")
    df = df[df["user_id"].isin(active_user_ids)]
    to_csv(df, "data/processed/sessions.csv")


def migrate_payment_users():
    df = read_csv("data/raw/donation_payment_users.csv")
    df.rename(columns={"twinte_user_id": "user_id"}, inplace=True)
    to_csv(df, "data/processed/payment_users.csv")


def migrate_registered_courses(active_user_ids: list[str]):
    df = read_csv("data/raw/timetables_registered_courses.csv")
    df.rename(columns={"instractor": "instructors"}, inplace=True)

    df_tar = df[df["schedules"] != "null"]
    for index, row in df_tar.iterrows():
        schedules = json.loads(row["schedules"])
        for schedule in schedules:
            schedule["locations"] = schedule.pop("room")
        df.at[index, "schedules"] = json.dumps(schedules, ensure_ascii=False)

    df = df[df["user_id"].isin(active_user_ids)]

    to_csv(df, "data/processed/registered_courses.csv")


def migrate_registered_course_tags():
    df = read_csv("data/raw/timetables_registered_course_tags.csv")
    df.rename(
        columns={"registered_course": "registered_course_id", "tag": "tag_id"},
        inplace=True,
    )
    to_csv(df, "data/processed/registered_course_tags.csv")


def migrate_tags(active_user_ids: list[str]):
    df = read_csv("data/raw/timetables_tags.csv")
    df = df[df["user_id"].isin(active_user_ids)]
    to_csv(df, "data/processed/tags.csv")


def prepare_existing_courses():
    df = read_csv("data/raw/course_courses.csv")
    existing_courses = [
        {
            "id": row["id"],
            "year": row["year"],
            "code": row["code"],
        }
        for _, row in df.iterrows()
    ]
    with open("data/processed/existing_courses.json", "w") as f:
        json.dump(existing_courses, f)


def main():
    active_user_ids = list_active_user_ids()

    migrate_users(active_user_ids)
    migrate_user_authentications(active_user_ids)

    migrate_sessions(active_user_ids)

    migrate_registered_courses(active_user_ids)
    migrate_registered_course_tags()

    migrate_tags(active_user_ids)

    migrate_payment_users()

    prepare_existing_courses()


if __name__ == "__main__":
    main()
