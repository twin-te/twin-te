import json
import sys
import datetime
from pathlib import Path
from typing import TypedDict, Literal, Final


class ModuleDetailJson(TypedDict):
    year: int
    module: str
    start: str
    end: str


class ModuleDetail:
    def __init__(self, json_: ModuleDetailJson):
        self.year = json_['year']
        self.module = json_['module']
        self.start = datetime.date.fromisoformat(json_['start'])
        self.end = datetime.date.fromisoformat(json_['end'])

    year: int
    module: str
    start: datetime.date
    end: datetime.date


class EventJson(TypedDict):
    date: str
    type: str
    description: str
    changeTo: str


class Event:
    def __init__(self, json_: EventJson):
        self.date = datetime.date.fromisoformat(json_['date'])
        self.type = json_['type']
        self.description = json_['description']
        self.changeTo = json_.get('changeTo')

    date: datetime.date
    type: str
    description: str
    changeTo: str | None


def error_exit():
    sys.stderr.write(f'Usage: {sys.argv[0]} {{year}}\n')
    exit(1)


def check_year(date: datetime.date, year: int):
    start = datetime.date(year, 4, 1)
    end = datetime.date(year + 1, 3, 31)
    return start <= date <= end


def verify_module(year: int, path: Path):
    if not path.exists():
        sys.stderr.write('module_detail: json ファイルが見つかりません。\n')
        exit(1)

    data = [ModuleDetail(m) for m in json.load(path.open())]
    last_date = datetime.date(year, 4, 1)
    for i, part in enumerate(data):
        if part.year != year:
            print(f'module_detail: [{i}].year が指定された年度と一致しません。')

        if not check_year(part.start, year):
            print(f'module_detail: [{i}].start が年度の期間外です。')
        elif last_date != part.start:
            print(f'module_detail: [{i}].start の指定が誤っています。')

        if not check_year(part.end, year):
            print(f'module_detail: [{i}].end が年度の期間外です。')

        last_date = part.end + datetime.timedelta(days=1)

    if last_date != datetime.date(year + 1, 4, 1):
        print(f'module_detail: 期間が3月31日で終了していません。')


types: Final = ['Holiday', 'PublicHoliday', 'Exam', 'SubstituteDay', 'Other']
weekdays: Final = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday']


def verify_event(year: int, path: Path):
    if not path.exists():
        sys.stderr.write('event: json ファイルが見つかりません。\n')
        exit(1)

    data = [Event(m) for m in json.load(path.open())]
    for i, part in enumerate(data):
        if not check_year(part.date, year):
            print(f'event: [{i}].date が年度の期間外です。')

        if part.type not in types:
            print(f'event: [{i}].type が正しく指定されていません。')

        if part.type == 'SubstituteDay' and part.changeTo not in weekdays:
            print(f'event: [{i}].changeTo の指定が誤っています。')


def main():
    if len(sys.argv) < 2:
        error_exit()
        return

    try:
        year = int(sys.argv[1])
    except ValueError:
        error_exit()
        return

    module_detail_path = Path('module_detail', f'{year}.json')
    event_path = Path('event', f'{year}.json')

    verify_module(year, module_detail_path)
    verify_event(year, event_path)
    print('検証完了')


if __name__ == '__main__':
    main()
