import { RegisteredCourse } from "~/domain/course";
import { Tag } from "~/domain/tag";
import { getDisplayIcalTags } from "../presenters/tag";

const createRegisteredCourse = (
  id: string,
  tagIds: string[]
): RegisteredCourse => ({
  id,
  year: 2026,
  name: `授業${id}`,
  instructors: [],
  credit: 1,
  methods: [],
  schedules: [],
  rooms: [],
  memo: "",
  attendance: 0,
  absence: 0,
  late: 0,
  tagIds,
});

const tags: Tag[] = [
  { id: "tag-a", name: "必修", order: 0 },
  { id: "tag-b", name: "専門", order: 1 },
  { id: "tag-c", name: "教職", order: 2 },
];

describe(getDisplayIcalTags.name, () => {
  it("counts the courses of each tag.", () => {
    const courses = [
      createRegisteredCourse("1", ["tag-a"]),
      createRegisteredCourse("2", ["tag-a", "tag-b"]),
      createRegisteredCourse("3", []),
    ];
    expect(getDisplayIcalTags(courses, tags)).toEqual([
      { id: "tag-a", name: "必修", courseCount: 2 },
      { id: "tag-b", name: "専門", courseCount: 1 },
      { id: "tag-c", name: "教職", courseCount: 0 },
    ]);
  });

  it("keeps the order of the given tags.", () => {
    expect(
      getDisplayIcalTags([], [tags[2], tags[0]]).map(({ id }) => id)
    ).toEqual(["tag-c", "tag-a"]);
  });
});
