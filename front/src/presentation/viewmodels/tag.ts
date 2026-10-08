import { TagColor } from "~/domain/tag";

export type DisplayCourseTag = {
  id: string;
  name: string;
  assign: boolean;
  color: TagColor | null;
};

export type DisplayCreditTag = {
  id: string;
  name: string;
  color: TagColor | null;
  credit: string;
};
