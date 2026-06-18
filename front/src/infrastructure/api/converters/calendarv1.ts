import { IcalSubscriptionMode } from "~/domain/calendar";
import * as CalendarV1PB from "~/infrastructure/api/gen/calendar/v1/service_pb";

export const fromPBIcalSubscriptionMode = (
  pbMode: CalendarV1PB.IcalSubscriptionMode
): IcalSubscriptionMode => {
  switch (pbMode) {
    // url が未設定のとき mode は UNSPECIFIED になりうるので sync 扱いにする。
    case CalendarV1PB.IcalSubscriptionMode.UNSPECIFIED:
    case CalendarV1PB.IcalSubscriptionMode.SYNC:
      return "sync";
    case CalendarV1PB.IcalSubscriptionMode.EXCLUDE:
      return "exclude";
    case CalendarV1PB.IcalSubscriptionMode.TRANSPARENT:
      return "transparent";
  }
  throw Error(`invalid enum ${pbMode}`);
};

export const toPBIcalSubscriptionMode = (
  mode: IcalSubscriptionMode
): CalendarV1PB.IcalSubscriptionMode => {
  switch (mode) {
    case "sync":
      return CalendarV1PB.IcalSubscriptionMode.SYNC;
    case "exclude":
      return CalendarV1PB.IcalSubscriptionMode.EXCLUDE;
    case "transparent":
      return CalendarV1PB.IcalSubscriptionMode.TRANSPARENT;
  }
};
