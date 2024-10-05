import { Dayjs } from "dayjs";

export type Announcement = {
  id: string;
  title: string;
  content: string;
  publishedAt: Dayjs;
  isRead?: boolean;
};
