import { CourseLocationInfo } from "~/domain/courseLocation";
import { isContained } from "~/utils";

export type LocalStoragePropType = {
  darkMode: boolean;
  saturdayCourseMode: boolean;
  nightPeriodMode: boolean;
  timeLabelMode: boolean;
  displayYear: number;
  showRooms: boolean;
  courseLocationInfo: CourseLocationInfo;
};

export type LocalStorageHandler<T> = {
  getter: () => T | undefined;
  setter: (value: T | undefined) => void;
};

export type LocalStorageConfig = {
  [K in keyof LocalStoragePropType]: LocalStorageHandler<
    LocalStoragePropType[K]
  >;
};

export const createLocalStorageBooleanHandler = (
  prop: string
): LocalStorageHandler<boolean> => ({
  getter: () => {
    const value = localStorage.getItem(prop);
    if (isContained<"true" | "false">(value, ["true", "false"]))
      return value === "true";
    else return undefined;
  },
  setter: (value) => {
    if (value == undefined) localStorage.removeItem(prop);
    else localStorage.setItem(prop, String(value));
  },
});

export const createLocalStorageNumberHandler = (
  prop: string
): LocalStorageHandler<number> => ({
  getter: () => {
    const value = localStorage.getItem(prop);
    if (isNaN(Number(value))) return undefined;
    else return Number(value);
  },
  setter: (value) => {
    if (value == undefined) localStorage.removeItem(prop);
    else localStorage.setItem(prop, String(value));
  },
});

export const createLocalStorageObjectHandler = <T>(
  prop: string
): LocalStorageHandler<T> => ({
  getter: () => {
    const value = localStorage.getItem(prop);
    if (!value) return undefined;
    try {
      return JSON.parse(value) as T;
    } catch {
      return undefined;
    }
  },
  setter: (value) => {
    if (value == undefined) localStorage.removeItem(prop);
    else localStorage.setItem(prop, JSON.stringify(value));
  },
});

export const localStorageConfig: LocalStorageConfig = {
  darkMode: createLocalStorageBooleanHandler("darkMode"),
  saturdayCourseMode: createLocalStorageBooleanHandler("saturdayCourseMode"),
  nightPeriodMode: createLocalStorageBooleanHandler("nightPeriodMode"),
  timeLabelMode: createLocalStorageBooleanHandler("timeLabelMode"),
  displayYear: createLocalStorageNumberHandler("displayYear"),
  showRooms: createLocalStorageBooleanHandler("showRooms"),
  courseLocationInfo: createLocalStorageObjectHandler("courseLocationInfo"),
};
