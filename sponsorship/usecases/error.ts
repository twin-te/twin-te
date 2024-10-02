import { Code, ConnectError } from "@connectrpc/connect";

export class NetworkError extends Error {
	readonly name = "NetworkError";
}

export class NotFoundError extends Error {
	readonly name = "NotFoundError";
}

export class UnauthenticatedError extends Error {
	readonly name = "UnauthenticatedError";
}

export const isNetworkError = (v: unknown): v is NetworkError => {
	return v instanceof NetworkError;
};

export const isNotFoundError = (v: unknown): v is NotFoundError => {
	return v instanceof NotFoundError;
};

export const isUnauthenticatedError = (
	v: unknown,
): v is UnauthenticatedError => {
	return v instanceof UnauthenticatedError;
};

export const ConvertAPIError = (error: unknown) => {
	const connectError = ConnectError.from(error);

	if (
		connectError.cause instanceof TypeError &&
		connectError.cause.message === "Failed to fetch"
	) {
		throw new NetworkError();
	}

	if (connectError.code === Code.Unauthenticated) {
		throw new UnauthenticatedError();
	}

	if (connectError.code === Code.NotFound) {
		throw new NotFoundError();
	}

	throw error;
};
