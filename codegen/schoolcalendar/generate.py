import json
import os
from pathlib import Path

dir_path = Path(__file__).parent
root_dir_path = Path(__file__).parents[2]


def generate(resource_name: str):
    resource_dir_path = dir_path.joinpath(resource_name)
    data: list[dict] = []

    # Get json file names
    file_names = os.listdir(resource_dir_path)
    file_names.sort()

    # Load and concat jsons
    for file_name in file_names:
        file_path = resource_dir_path.joinpath(file_name)
        with file_path.open(mode="r", encoding="utf-8") as f:
            data += json.load(f)

    # Assign ids
    for i in range(len(data)):
        data[i] = {
            "id": i + 1,
            **data[i],
        }

    # Output
    output_file_path = root_dir_path.joinpath(
        f"back/module/schoolcalendar/data/{resource_name}_gen.json"
    )
    output_file_path.parent.mkdir(parents=True, exist_ok=True)
    with output_file_path.open(mode="w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)
    print(f"output to {output_file_path}")


def main():
    generate("event")
    generate("module_detail")


if __name__ == "__main__":
    main()
