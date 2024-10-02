import { createConnectTransport } from "@connectrpc/connect-web";
import { AnnouncementUseCase, type IAnnouncementUseCase } from "./announcement";
import { AuthUseCase, type IAuthUseCase } from "./auth";
import { FeedbackUseCase, type IFeedbackUseCase } from "./feedback";
import {
	type ISchoolCalendarUseCase,
	SchoolCalendarUseCase,
} from "./schoolCalendar";
import { type ISettingUseCase, SettingUseCase } from "./setting";
import { type ITimetableUseCase, TimetableUseCase } from "./timetable";

const transport = createConnectTransport({
	baseUrl: import.meta.env.VITE_API_URL,
	useBinaryFormat: true,
	credentials: "include",
	useHttpGet: true,
});

export const announcementUseCase: IAnnouncementUseCase =
	new AnnouncementUseCase(transport);
export const authUseCase: IAuthUseCase = new AuthUseCase(transport);
export const feedbackUseCase: IFeedbackUseCase = new FeedbackUseCase();
export const schoolCalendarUseCase: ISchoolCalendarUseCase =
	new SchoolCalendarUseCase(transport);
export const settingUseCase: ISettingUseCase = new SettingUseCase();
export const timetableUseCase: ITimetableUseCase = new TimetableUseCase(
	transport,
);
