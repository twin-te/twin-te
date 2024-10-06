import pandas as pd


def read_csv(path: str) -> pd.DataFrame:
    return pd.read_csv(path, keep_default_na=False)


def to_csv(df: pd.DataFrame, path: str) -> None:
    df.to_csv(
        path,
        encoding="utf-8",
        index=False,
    )


def migrate_users():
    df = read_csv("data/raw/user_users.csv")
    df.rename(
        columns={"createdAt": "created_at", "deletedAt": "deleted_at"}, inplace=True
    )
    to_csv(df, "data/processed/users.csv")


def migrate_user_authentications():
    df = read_csv("data/raw/user_user_authentications.csv")
    df = df[df["user_id"] != "null"]
    df.drop(columns=["id"], inplace=True)
    to_csv(df, "data/processed/user_authentications.csv")


def migrate_courses():
    df = read_csv("data/raw/course_courses.csv")
    df.rename(
        columns={"instructor": "instructors", "last_update": "last_updated_at"},
        inplace=True,
    )
    to_csv(df, "data/processed/courses.csv")


def migrate_course_methods():
    df = read_csv("data/raw/course_course_methods.csv")
    df = df[df["course_id"] != "null"]
    df.drop(columns=["id"], inplace=True)
    to_csv(df, "data/processed/course_methods.csv")


def migrate_course_recommended_grades():
    df = read_csv("data/raw/course_course_recommended_grades.csv")
    df = df[df["course_id"] != "null"]
    df.drop(columns=["id"], inplace=True)
    to_csv(df, "data/processed/course_recommended_grades.csv")


def migrate_course_schedules():
    df = read_csv("data/raw/course_course_schedules.csv")
    df = df[df["course_id"] != "null"]
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

    to_csv(df, "data/processed/course_schedules.csv")


def migrate_registered_courses():
    df = read_csv("data/raw/timetables_registered_courses.csv")
    df.rename(columns={"instractor": "instructors"}, inplace=True)
    to_csv(df, "data/processed/registered_courses.csv")


def migrate_registered_course_tags():
    df = read_csv("data/raw/timetables_registered_course_tags.csv")
    df.rename(
        columns={"registered_course": "registered_course_id", "tag": "tag_id"},
        inplace=True,
    )
    to_csv(df, "data/processed/registered_course_tags.csv")


def migrate_tags():
    df = read_csv("data/raw/timetables_tags.csv")
    to_csv(df, "data/processed/tags.csv")


def migrate_sessions():
    df = read_csv("data/raw/session_session.csv")
    to_csv(df, "data/processed/sessions.csv")


def migrate_payment_users():
    df = read_csv("data/raw/donation_payment_users.csv")
    df.rename(columns={"twinte_user_id": "user_id"}, inplace=True)
    to_csv(df, "data/processed/payment_users.csv")


def main():
    migrate_users()
    migrate_user_authentications()

    migrate_sessions()

    migrate_courses()
    migrate_course_methods()
    migrate_course_recommended_grades()
    migrate_course_schedules()

    migrate_registered_courses()
    migrate_registered_course_tags()

    migrate_tags()

    migrate_payment_users()


if __name__ == "__main__":
    main()
