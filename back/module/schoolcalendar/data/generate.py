import json
import os
import pathlib

data_dir_path = pathlib.Path(__file__).parent


def generate(resource_name: str):
    raw_path = data_dir_path.joinpath(f"{resource_name}/raw")
    data: list[dict] = []

    # Get json file names
    file_names = os.listdir(raw_path)
    file_names.sort()

    # Load and concat jsons
    for file_name in file_names:
        file_path = raw_path.joinpath(file_name)
        with file_path.open(mode="r", encoding="utf-8") as f:
            data += json.load(f)

    # Assign ids
    for i in range(len(data)):
        data[i] = {
            "id": i + 1,
            **data[i],
        }

    # Output
    output_file_path = data_dir_path.joinpath(f"{resource_name}/prod_gen.json")
    output_file_path.parent.mkdir(parents=True, exist_ok=True)
    with output_file_path.open(mode="w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)


def main():
    generate("event")
    generate("module_detail")


if __name__ == "__main__":
    main()
