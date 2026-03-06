import dayjs, { Dayjs } from "dayjs";
import _ from "lodash";

export const getAcademicYear = (date: Dayjs) => {
  return date.month() < 3 ? date.year() - 1 : date.year();
};

export const initialAcademicYear = 2019;

export const currentAcademicYear = getAcademicYear(dayjs());

const MAX_FUTURE_YEARS = 1;

export const academicYears: number[] = _.range(
  initialAcademicYear,
  currentAcademicYear + MAX_FUTURE_YEARS + 1
);

export const validateAcademicYear = (year: number): boolean => {
  return (
    initialAcademicYear <= year &&
    year <= currentAcademicYear + MAX_FUTURE_YEARS
  );
};

export const isFutureYear = (year: number): boolean =>
  year > currentAcademicYear;
