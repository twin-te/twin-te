import { Method } from "./method";
import { Room } from "./room";
import { Schedule } from "./schedule";

export type Course = {
  id: string;
  year: number;
  code: string;
  name: string;
  instructors: string[];
  credit: number;
  overview: string;
  remarks: string;
  recommendedGrades: number[];
  methods: Method[];
  schedules: Schedule[];
  rooms: Room[];
};

export type RegisteredCourse = {
  id: string;
  year: number;
  code?: string; // If the course is added by manual, code is null.
  name: string;
  instructors: string[];
  credit: number;
  methods: Method[];
  schedules: Schedule[];
  rooms: Room[];
  memo: string;
  attendance: number;
  absence: number;
  late: number;
  tagIds: string[];
};
