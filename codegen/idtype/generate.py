import pathlib
import re
import shutil
import string


def convert_pascal_case_to_snake_case(s: str) -> str:
    """Convert PascalCase to snake_case.

    cf. https://stackoverflow.com/questions/1175208/elegant-python-function-to-convert-camelcase-to-snake-case
    """
    s = re.sub("(.)([A-Z][a-z]+)", r"\1_\2", s)
    s = re.sub("([a-z0-9])([A-Z])", r"\1_\2", s).lower()
    return s


def validate_id_names(id_names: list[str]) -> None:
    p = re.compile(r"[A-Z][A-Za-z]+ID")
    for id_name in id_names:
        if p.fullmatch(id_name) is None:
            raise ValueError(f"invalid id name : '{id_name}'")


def get_output_file_path(output_dir_path: pathlib.Path, id_name: str) -> pathlib.Path:
    file_name = f"{convert_pascal_case_to_snake_case(id_name)}_gen.go"
    output_file_path = output_dir_path.joinpath(file_name)
    return output_file_path


def generate(
    template_path: pathlib.Path,
    definition_path: pathlib.Path,
    output_dir_path: pathlib.Path,
) -> None:
    with definition_path.open(mode="r") as f:
        id_names = f.read().splitlines()

    validate_id_names(id_names)

    with open(template_path, "r") as f:
        template = string.Template(f.read())

    for id_name in id_names:
        output_file_path = get_output_file_path(output_dir_path, id_name)
        output_str = template.substitute({"id_name": id_name})
        with output_file_path.open(mode="w") as f:
            f.write(output_str)
        print(f"output to {output_file_path}")


def main():
    dir_path = pathlib.Path(__file__).parent

    int_template_path = dir_path.joinpath("int_template.txt")
    int_definition_path = dir_path.joinpath("int_definition.txt")

    string_template_path = dir_path.joinpath("string_template.txt")
    string_definition_path = dir_path.joinpath("string_definition.txt")

    uuid_template_path = dir_path.joinpath("uuid_template.txt")
    uuid_definition_path = dir_path.joinpath("uuid_definition.txt")

    output_dir_path = (
        pathlib.Path(__file__).parents[2].joinpath("back/module/shared/domain/idtype")
    )
    if output_dir_path.exists():
        shutil.rmtree(output_dir_path)
    output_dir_path.mkdir(parents=True)

    generate(int_template_path, int_definition_path, output_dir_path)
    generate(string_template_path, string_definition_path, output_dir_path)
    generate(uuid_template_path, uuid_definition_path, output_dir_path)


if __name__ == "__main__":
    main()
