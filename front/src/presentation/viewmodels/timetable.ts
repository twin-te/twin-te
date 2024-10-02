import type { NormalDay, SpecialDay } from "~/domain/day";
import type { Module } from "~/domain/module";
import type { Period } from "~/domain/period";
import type { DisplayRegisteredCourse } from "./course";

export type DiaplayNormalTimetable<
	M extends Module,
	D extends NormalDay,
	P extends Period,
> = Record<M, Record<D, Record<P, DisplayRegisteredCourse[]>>>;

export type DisplaySpecialTimetable<D extends SpecialDay> = Record<
	D | "Others",
	DisplayRegisteredCourse[]
>;

export type DisplayTimetable<
	M extends Module,
	ND extends NormalDay,
	SD extends SpecialDay,
	P extends Period,
> = {
	normal: DiaplayNormalTimetable<M, ND, P>;
	special: DisplaySpecialTimetable<SD>;
};
