import {
  createPromiseClient,
  ConnectError,
  Code,
  PromiseClient,
  Transport,
} from "@connectrpc/connect";
import {
  InternalServerError,
  NetworkError,
  UnauthenticatedError,
} from "~/domain/error";
import { CalendarService } from "~/infrastructure/api/gen/calendar/v1/service_connect";
import { handleError } from "~/infrastructure/api/utils";

export interface ICalendarUseCase {
  getIcalSubscriptionUrl(): Promise<
    | { url: string }
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
}

export class CalendarUseCase implements ICalendarUseCase {
  #client: PromiseClient<typeof CalendarService>;

  constructor(transport: Transport) {
    this.#client = createPromiseClient(CalendarService, transport);
  }

  async getIcalSubscriptionUrl(): Promise<
    | { url: string }
    | { url: null }
    | UnauthenticatedError
    | NetworkError
    | InternalServerError
  > {
    return this.#client
      .getIcalSubscriptionUrl({})
      .then((res) => {
        const url = res.url;
        return url ? { url } : { url: null };
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
}
