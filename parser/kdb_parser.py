import datetime
import io
import re

import pandas as pd


class Module:
    SpringA = "春A"
    SpringB = "春B"
    SpringC = "春C"
    FallA = "秋A"
    FallB = "秋B"
    FallC = "秋C"
    SummerVacation = "夏季休業中"
    SpringVacation = "春季休業中"
    Annual = "通年"
    Unknown = "不明"


class Day:
    Sun = "日"
    Mon = "月"
    Tue = "火"
    Wed = "水"
    Thu = "木"
    Fri = "金"
    Sat = "土"
    Intensive = "集中"
    Appointment = "応談"
    AnyTime = "随時"
    NT = "NT"  # NT
    Unknown = "不明"


def parse_recommended_grade(s: str) -> list[int]:
    if re.fullmatch(r"\d", s) is not None:
        return [int(s)]
    elif re.search(r"\d・\d", s) is not None:
        return list(map(int, s.split("・")))
    elif re.fullmatch(r"\d - \d", s) is not None:
        start_grade, end_grade = int(s[0]), int(s[-1])
        return [grade for grade in range(start_grade, end_grade + 1)]
    else:
        return []


def parse_credits(s: str) -> int | float:
    if re.fullmatch(r"\d+", s) is not None:
        return int(s)

    if re.fullmatch(r"\d+\.0", s) is not None:
        return int(s[:-2])

    if re.fullmatch(r"\d+\.5", s) is not None:
        return float(s)

    return 0


def analyze_module(s: str) -> list[Module]:
    res: list[Module] = []

    if re.findall(r"春[ABC]*A", s, re.MULTILINE):
        res.append(Module.SpringA)
    if re.findall(r"春[ABC]*B", s, re.MULTILINE):
        res.append(Module.SpringB)
    if re.findall(r"春[ABC]*C", s, re.MULTILINE):
        res.append(Module.SpringC)

    if s.find("春学期") != -1:
        res += [
            Module.SpringA,
            Module.SpringB,
            Module.SpringC,
            Module.SummerVacation,
        ]

    if s.find(Module.SummerVacation) != -1:
        res.append(Module.SummerVacation)

    if re.findall(r"秋[ABC]*A", s, re.MULTILINE):
        res.append(Module.FallA)
    if re.findall(r"秋[ABC]*B", s, re.MULTILINE):
        res.append(Module.FallB)
    if re.findall(r"秋[ABC]*C", s, re.MULTILINE):
        res.append(Module.FallC)

    if s.find("秋学期") != -1:
        res += [
            Module.FallA,
            Module.FallB,
            Module.FallC,
            Module.SpringVacation,
        ]

    if s.find(Module.SpringVacation) != -1:
        res.append(Module.SpringVacation)

    if s.find(Module.Annual) != -1:
        res.append(Module.Annual)

    if s != "" and len(res) == 0:
        res.append(Module.Unknown)

    return res


def analyze_day_and_period(s: str) -> list[dict]:
    result: list[dict] = []

    for string in re.split(r",(?=[日月火水木金土])", s):
        for day in Day.__dict__.values():
            for i in range(1, 9):
                if re.search(f"{day}.*{i}", string) is not None:
                    result.append(
                        {
                            "day": day,
                            "period": i,
                        }
                    )

            _match = re.search(f"([{day}]).*(\d)-(\d)", string)
            if _match is not None:
                for i in range(int(_match.groups()[1]), int(_match.groups()[2]) + 1):
                    flag = True
                    for r in result:
                        if r["day"] == day and r["period"] == i:
                            flag = False
                            break
                    if flag:
                        result.append({"day": day, "period": i})

    # 集中、応談、随時の検出
    if Day.Intensive in s:
        result.append({"day": Day.Intensive, "period": 0})

    if Day.Appointment in s:
        result.append({"day": Day.Appointment, "period": 0})

    if Day.AnyTime in s:
        result.append({"day": Day.AnyTime, "period": 0})

    # NTの検出
    if Day.NT in s:
        result.append({"day": Day.NT, "period": 0})

    # どのテストにも合格しなかったが空文字でなければ仮にUnknownとする
    if s != "" and not result:
        result.append({"day": Day.Unknown, "period": 0})

    return result


def analyze_row(row: list[str]) -> dict:
    course_data = {
        "code": row[0],
        "name": row[1],
        "credits": parse_credits(row[3]),
        "type": int(row[2]),
        "overview": row[9],
        "remarks": row[10],
        "recommendedGrade": parse_recommended_grade(row[4]),
        "schedules": [],
        "instructor": row[8],
        "error": False,
        "lastUpdate": datetime.datetime.fromisoformat(row[18] + "+09:00")
        .astimezone(datetime.timezone.utc)
        .isoformat(timespec="milliseconds")
        .replace("+00:00", "Z"),  # JST保証
    }

    module_str = row[5]
    period_str = row[6]
    location_str = row[7]

    module_array = re.split("\r\n", module_str)
    period_array = re.split("\r\n", period_str)
    location_array = re.split("\r\n", location_str)

    count = max([len(module_array), len(period_array), len(location_array)])

    if not (
        (len(module_array) == count or len(module_array) == 1)
        and (len(period_array) == count or len(period_array) == 1)
        and (len(location_array) == count or len(location_array) == 1)
    ):
        course_data["error"] = True

    for i in range(count):
        modules = analyze_module(
            module_array[0]
            if len(module_array) == 1
            else (
                module_array[i]
                if len(module_array) > i and module_array[i]
                else "unknown"
            )
        )
        when = analyze_day_and_period(
            period_array[0]
            if len(period_array) == 1
            else (
                period_array[i]
                if len(period_array) > i and period_array[i]
                else "unknown"
            )
        )

        for mod in modules:
            for w in when:
                course_data["schedules"].append(
                    {
                        "module": mod,
                        "period": w["period"],
                        "day": w["day"],
                        "location": (
                            location_array[0]
                            if len(location_array) == 1
                            else (
                                location_array[i]
                                if len(location_array) > i and location_array[i]
                                else ""
                            )
                        ),
                    }
                )

    return course_data


def parse(path_or_bytes: str | io.BytesIO) -> list:
    df = pd.read_excel(path_or_bytes, skiprows=[0, 1, 2, 3])
    courses = [
        analyze_row(row.tolist())
        for _, row in df.fillna("").iterrows()
        if row.iloc[0] != ""
    ]
    return courses
