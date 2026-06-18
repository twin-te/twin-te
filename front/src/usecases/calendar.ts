import {
  createPromiseClient,
  ConnectError,
  Code,
  PromiseClient,
  Transport,
} from "@connectrpc/connect";
import { IcalSubscriptionMode } from "~/domain/calendar";
import {
  InternalServerError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import {
  fromPBIcalSubscriptionMode,
  toPBIcalSubscriptionMode,
} from "~/infrastructure/api/converters/calendarv1";
import { fromPBUUID, toPBUUID } from "~/infrastructure/api/converters/shared";
import { CalendarService } from "~/infrastructure/api/gen/calendar/v1/service_connect";
import { handleError } from "~/infrastructure/api/utils";

export interface ICalendarUseCase {
  getIcalSubscriptionUrl(): Promise<
    | { url: string; mode: IcalSubscriptionMode; targetTagIds: string[] }
    | { url: null }
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  >;

  enableIcalSubscription(): Promise<
    { url: string } | UnauthenticatedError | NetworkError | InternalServerError
  >;

  disableIcalSubscription(): Promise<
    null | UnauthenticatedError | NetworkError | InternalServerError
  >;

  updateIcalSubscription(
    mode: IcalSubscriptionMode,
    targetTagIds: string[]
  ): Promise<null | UnauthenticatedError | NetworkError | InternalServerError>;
}

export class CalendarUseCase implements ICalendarUseCase {
  #client: PromiseClient<typeof CalendarService>;

  constructor(transport: Transport) {
    this.#client = createPromiseClient(CalendarService, transport);
  }

  async getIcalSubscriptionUrl(): Promise<
    | { url: string; mode: IcalSubscriptionMode; targetTagIds: string[] }
    | { url: null }
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#client
      .getIcalSubscriptionUrl({})
      .then((res) => {
        const url = res.url;
        return url
          ? {
              url,
              mode: fromPBIcalSubscriptionMode(res.mode),
              targetTagIds: res.targetTagIds.map(fromPBUUID),
            }
          : { url: null };
      })
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }
          throw error;
        });
      });
  }

  async enableIcalSubscription(): Promise<
    { url: string } | UnauthenticatedError | NetworkError | InternalServerError
  > {
    return this.#client
      .enableIcalSubscription({})
      .then((res) => ({ url: res.url }))
      .catch((error) => {
        return handleError(error, (connectError: ConnectError) => {
          if (connectError.code === Code.Unauthenticated) {
            return new UnauthenticatedError();
          }
          throw error;
        });
      });
  }

  async disableIcalSubscription(): Promise<
    null | UnauthenticatedError | NetworkError | InternalServerError
  > {
    return this.#client
      .disableIcalSubscription({})
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

  async updateIcalSubscription(
    mode: IcalSubscriptionMode,
    targetTagIds: string[]
  ): Promise<null | UnauthenticatedError | NetworkError | InternalServerError> {
    return this.#client
      .updateIcalSubscription({
        mode: toPBIcalSubscriptionMode(mode),
        targetTagIds: targetTagIds.map(toPBUUID),
      })
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
