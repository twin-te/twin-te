import {
  createPromiseClient,
  ConnectError,
  PromiseClient,
  Transport,
  Code,
} from "@connectrpc/connect";
import { NetworkError, UnauthenticatedError } from "~/domain/error";
import { News } from "~/domain/news";
import { fromPBAnnouncement } from "~/infrastructure/api/converters/announcementv1";
import { toPBUUID } from "~/infrastructure/api/converters/shared";
import { AnnouncementService } from "~/infrastructure/api/gen/announcement/v1/service_connect";
import { handleError } from "~/infrastructure/api/utils";

export interface INewsUseCase {
  getNews(): Promise<News[] | UnauthenticatedError | NetworkError>;

  readNews(ids: string[]): Promise<null | UnauthenticatedError | NetworkError>;
}

export class NewsUseCase implements INewsUseCase {
  #client: PromiseClient<typeof AnnouncementService>;

  constructor(transport: Transport) {
    this.#client = createPromiseClient(AnnouncementService, transport);
  }

  async getNews(): Promise<News[] | UnauthenticatedError | NetworkError> {
    return this.#client
      .getAnnouncements({})
      .then((res) => {
        if (res.announcements.some(({ isRead }) => isRead == undefined)) {
          return new UnauthenticatedError();
        }
        return res.announcements.map(fromPBAnnouncement);
      })
      .catch((error) => handleError(error));
  }

  readNews(ids: string[]): Promise<null | UnauthenticatedError | NetworkError> {
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
