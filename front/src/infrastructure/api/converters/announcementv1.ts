import { Announcement } from "~/domain/announcement";
import * as AnnouncementV1PB from "~/infrastructure/api/gen/announcement/v1/type_pb";
import { fromPBRFC3339DateTime, fromPBUUID } from "./shared";
import { assurePresence } from "./utils";

export const fromPBAnnouncement = (
  pbAnnouncement: AnnouncementV1PB.Announcement
): Announcement => {
  return {
    id: fromPBUUID(assurePresence(pbAnnouncement.id)),
    title: pbAnnouncement.title,
    content: pbAnnouncement.content,
    publishedAt: fromPBRFC3339DateTime(
      assurePresence(pbAnnouncement.publishedAt)
    ),
    isRead: pbAnnouncement.isRead,
  };
};
