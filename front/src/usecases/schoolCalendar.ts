import {
	type PromiseClient,
	type Transport,
	createPromiseClient,
} from "@connectrpc/connect";
import dayjs, { type Dayjs } from "dayjs";
import type { InternalServerError, NetworkError } from "~/domain/error";
import type { Event } from "~/domain/event";
import type { SchoolCalendarModule } from "~/domain/module";
import {
	fromPBEvent,
	fromPBModule,
} from "~/infrastructure/api/converters/schoolcalendarv1";
import { toPBRFC3339FullDate } from "~/infrastructure/api/converters/shared";
import { SchoolCalendarService } from "~/infrastructure/api/gen/schoolcalendar/v1/service_connect";
import { handleError } from "~/infrastructure/api/utils";

export interface ISchoolCalendarUseCase {
	getEventByDate(
		date: Dayjs,
	): Promise<Event | null | NetworkError | InternalServerError>;

	getCurrentModule(): Promise<
		SchoolCalendarModule | NetworkError | InternalServerError
	>;
}

export class SchoolCalendarUseCase implements ISchoolCalendarUseCase {
	#client: PromiseClient<typeof SchoolCalendarService>;

	constructor(transport: Transport) {
		this.#client = createPromiseClient(SchoolCalendarService, transport);
	}

	async getEventByDate(
		date: Dayjs,
	): Promise<Event | null | NetworkError | InternalServerError> {
		return this.#client
			.getEventsByDate({ date: toPBRFC3339FullDate(date) })
			.then((res) => {
				const events = res.events.map((pbEvent) => fromPBEvent(pbEvent));
				if (events.length === 0) return null;
				return events[0];
			})
			.catch((error) => handleError(error));
	}

	async getCurrentModule(): Promise<
		SchoolCalendarModule | NetworkError | InternalServerError
	> {
		return this.#client
			.getModuleByDate({ date: toPBRFC3339FullDate(dayjs()) })
			.then((res) => fromPBModule(res.module))
			.catch((error) => handleError(error));
	}
}
