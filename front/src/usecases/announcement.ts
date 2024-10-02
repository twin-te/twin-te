import {
	Code,
	type ConnectError,
	type PromiseClient,
	type Transport,
	createPromiseClient,
} from "@connectrpc/connect";
import type { Announcement } from "~/domain/announcement";
import {
	type InternalServerError,
	type NetworkError,
	UnauthenticatedError,
} from "~/domain/error";
import { fromPBAnnouncement } from "~/infrastructure/api/converters/announcementv1";
import { toPBUUID } from "~/infrastructure/api/converters/shared";
import { AnnouncementService } from "~/infrastructure/api/gen/announcement/v1/service_connect";
import { handleError } from "~/infrastructure/api/utils";

export interface IAnnouncementUseCase {
	getAnnouncements(): Promise<
		Announcement[] | NetworkError | InternalServerError
	>;

	readAnnouncements(
		ids: string[],
	): Promise<null | UnauthenticatedError | NetworkError | InternalServerError>;
}

export class AnnouncementUseCase implements IAnnouncementUseCase {
	#client: PromiseClient<typeof AnnouncementService>;

	constructor(transport: Transport) {
		this.#client = createPromiseClient(AnnouncementService, transport);
	}

	async getAnnouncements(): Promise<
		Announcement[] | NetworkError | InternalServerError
	> {
		return this.#client
			.getAnnouncements({})
			.then((res) => res.announcements.map(fromPBAnnouncement))
			.catch((error) => handleError(error));
	}

	async readAnnouncements(
		ids: string[],
	): Promise<null | UnauthenticatedError | NetworkError | InternalServerError> {
		return this.#client
			.readAnnouncements({ ids: ids.map(toPBUUID) })
			.then(() => null)
			.catch((error) => {
				return handleError(error, (connectError: ConnectError) => {
					if (connectError.code === Code.Unauthenticated) {
						return new UnauthenticatedError();
					}

					throw error;
				});
			});
	}
}
