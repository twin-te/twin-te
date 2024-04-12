class BaseResultError extends Error {
  readonly type = "err";
}

export class NotFoundError extends BaseResultError {
  readonly name = "NotFoundError";
}

export class UnauthenticatedError extends BaseResultError {
  readonly name = "UnauthenticatedError";
}

export class NetworkError extends BaseResultError {
  readonly name = "NetworkError";
}

export class InternalServerError extends BaseResultError {
  readonly name = "InternalServerError";
}

export type ResultError =
  | NotFoundError
  | UnauthenticatedError
  | NetworkError
  | InternalServerError;

export const isResultError = <T>(
  result: T | ResultError
): result is ResultError => {
  return result instanceof BaseResultError;
};

export const throwResultError = <T>(result: T | ResultError): T => {
  if (isResultError(result)) throw result;
  return result;
};
