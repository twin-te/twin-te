import { RegisteredCourse } from "~/domain/course";
import { Tag, TagColor } from "~/domain/tag";
import { initializeObject } from "~/utils";
import { DisplayCourseTag, DisplayCreditTag } from "../viewmodels/tag";
import { creditToDisplay } from "./credit";

export const getDisplayCourseTags = (
  course: RegisteredCourse,
  tags: Tag[]
): DisplayCourseTag[] => {
  return tags.map(({ id, name, color }) => ({
    id,
    name,
    assign: course.tagIds.findIndex((tagId) => tagId === id) !== -1,
    color,
  }));
};

export const getDisplayCreditTag = (
  courses: RegisteredCourse[],
  tag: Tag,
  years: number[]
): DisplayCreditTag => {
  const credit = courses
    .filter(({ year }) => years.includes(year))
    .reduce<number>((credit, course) => {
      if (course.tagIds.includes(tag.id)) credit += course.credit;
      return credit;
    }, 0);

  return {
    id: tag.id,
    name: tag.name,
    credit: creditToDisplay(credit),
    color: tag.color,
  };
};

export const getDisplayCreditTags = (
  courses: RegisteredCourse[],
  tags: Tag[],
  years: number[]
): DisplayCreditTag[] => {
  const tagIdToCredit: Record<string, number> = initializeObject(
    tags.map(({ id }) => id),
    0
  );

  courses
    .filter(({ year }) => years.includes(year))
    .forEach(({ credit, tagIds }) => {
      tagIds.forEach((tagId) => {
        tagIdToCredit[tagId] += credit;
      });
    });

  return tags.map(({ id, name, color }) => ({
    id,
    name,
    credit: creditToDisplay(tagIdToCredit[id]),
    color,
  }));
};

type PresetColor = Exclude<TagColor, `#${string}`>;

function isPresetColor(color: TagColor): color is PresetColor {
  return !color.startsWith("#");
}

type TagColorVariable = `--tag-${PresetColor | "default"})`;

export function tagColorToCss(
  color: TagColor | null
): `#${string}` | `var(${TagColorVariable}` {
  if (color && !isPresetColor(color)) return color;
  return `var(--tag-${color ?? "default"})`;
}
