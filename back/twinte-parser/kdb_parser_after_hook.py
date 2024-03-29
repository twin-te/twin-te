import dataclasses
import json

import kdb_parser


class Module(str):
    SpringA = "SpringA"
    SpringB = "SpringB"
    SpringC = "SpringC"
    FallA = "FallA"
    FallB = "FallB"
    FallC = "FallC"
    SummerVacation = "SummerVacation"
    SpringVacation = "SpringVacation"


class Day(str):
    Sun = "Sun"
    Mon = "Mon"
    Tue = "Tue"
    Wed = "Wed"
    Thu = "Thu"
    Fri = "Fri"
    Sat = "Sat"
    Intensive = "Intensive"
    Appointment = "Appointment"
    AnyTime = "AnyTime"


@dataclasses.dataclass
class Schedule:
    module: Module
    day: Day
    period: int
    rooms: str


class CourseMethod(str):
    OnlineAsynchronous = "OnlineAsynchronous"
    OnlineSynchronous = "OnlineSynchronous"
    FaceToFace = "FaceToFace"
    Others = "Others"


@dataclasses.dataclass
class Course:
    code: str
    name: str
    instructors: str
    credit: str
    overview: str
    remarks: str
    last_updated_at: str
    has_parse_error: bool
    is_annual: bool
    recommended_grades: list[int]
    methods: list[CourseMethod]
    schedules: list[Schedule]


def convert_module(m: kdb_parser.Module) -> Module:
    if m == kdb_parser.Module.SpringA:
        return Module.SpringA

    if m == kdb_parser.Module.SpringB:
        return Module.SpringB

    if m == kdb_parser.Module.SpringC:
        return Module.SpringC

    if m == kdb_parser.Module.FallA:
        return Module.FallA

    if m == kdb_parser.Module.FallB:
        return Module.FallB

    if m == kdb_parser.Module.FallC:
        return Module.FallC

    if m == kdb_parser.Module.SummerVacation:
        return Module.SummerVacation

    if m == kdb_parser.Module.SpringVacation:
        return Module.SpringVacation

    raise ValueError("invalid module", m)


def convert_day(d: kdb_parser.Day) -> Day:
    if d == kdb_parser.Day.Sun:
        return Day.Sun

    if d == kdb_parser.Day.Mon:
        return Day.Mon

    if d == kdb_parser.Day.Tue:
        return Day.Tue

    if d == kdb_parser.Day.Wed:
        return Day.Wed

    if d == kdb_parser.Day.Thu:
        return Day.Thu

    if d == kdb_parser.Day.Fri:
        return Day.Fri

    if d == kdb_parser.Day.Sat:
        return Day.Sat

    if d == kdb_parser.Day.Intensive:
        return Day.Intensive

    if d == kdb_parser.Day.AnyTime:
        return Day.AnyTime

    if d == kdb_parser.Day.Appointment:
        return Day.Appointment

    raise ValueError("invalid day", d)


def is_special_day(d: Day) -> bool:
    return d in [Day.Intensive, Day.AnyTime, Day.Appointment]


def convert_schedule(s: dict) -> Schedule:
    module = convert_module(s["module"])
    day = convert_day(s["day"])

    if is_special_day(day):
        return Schedule(module=module, day=day, period=0, rooms=s["room"])

    period = s["period"]

    if 1 <= period and period <= 8:
        return Schedule(module=module, day=day, period=period, rooms=s["room"])

    raise ValueError("invalid schedule", s)


def convert_recommended_grade(rg: int) -> int:
    if 1 <= rg and rg <= 6:
        return rg
    raise ValueError("invalid recommended grade", rg)


def convert_credit(c: int | float) -> str:
    c = float(c)
    return f"{c:.1f}"


def extract_course_methods(remarks: str) -> list[CourseMethod]:
    ret = []

    if "対面" in remarks:
        ret.append(CourseMethod.FaceToFace)
    if "オンデマンド" in remarks:
        ret.append(CourseMethod.OnlineAsynchronous)
    if "双方向" in remarks:
        ret.append(CourseMethod.OnlineSynchronous)
    if "その他" in remarks:
        ret.append(CourseMethod.Others)

    return ret


def convert(row: dict) -> Course:
    has_parse_error = row["error"]

    credit = convert_credit(row["credits"])
    methods = extract_course_methods(row["remarks"])

    recommended_grades: list[int] = []
    for rg in row["recommendedGrade"]:
        try:
            recommended_grades.append(convert_recommended_grade(rg))
        except ValueError:
            has_parse_error = True

    is_annual = False
    schedules: list[Schedule] = []
    for s in row["schedules"]:
        if s["module"] == kdb_parser.Module.Annual:
            is_annual = True
            for m in [
                kdb_parser.Module.SpringA,
                kdb_parser.Module.SpringB,
                kdb_parser.Module.SpringC,
                kdb_parser.Module.FallA,
                kdb_parser.Module.FallB,
                kdb_parser.Module.FallC,
            ]:
                try:
                    schedules.append(convert_schedule({**s, "module": m}))
                except ValueError:
                    has_parse_error = True
        else:
            try:
                schedules.append(convert_schedule(s))
            except ValueError:
                has_parse_error = True

    return Course(
        code=row["code"],
        name=row["name"],
        instructors=row["instructor"],
        credit=credit,
        overview=row["overview"],
        remarks=row["remarks"],
        last_updated_at=row["lastUpdate"],
        has_parse_error=has_parse_error,
        is_annual=is_annual,
        recommended_grades=recommended_grades,
        methods=methods,
        schedules=schedules,
    )


def parse_after_hook(json_data: str) -> str:
    courses: list[Course] = []
    for row in json.loads(json_data):
        courses.append(convert(row))

    for_json = []
    for course in courses:
        for_json.append(
            {
                "code": course.code,
                "name": course.name,
                "instructors": course.instructors,
                "credit": course.credit,
                "overview": course.overview,
                "remarks": course.remarks,
                "lastUpdatedAt": course.last_updated_at,
                "hasParseError": course.has_parse_error,
                "isAnnual": course.is_annual,
                "recommendedGrades": course.recommended_grades,
                "methods": course.methods,
                "schedules": [
                    {
                        "module": s.module,
                        "day": s.day,
                        "period": s.period,
                        "rooms": s.rooms,
                    }
                    for s in course.schedules
                ],
            }
        )

    ret = json.dumps(for_json, ensure_ascii=False, separators=(",", ":"))
    return ret
