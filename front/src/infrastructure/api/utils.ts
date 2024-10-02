import { Code, ConnectError } from "@connectrpc/connect";

import { InternalServerError, NetworkError } from "~/domain/error";

export const handleError = <T>(
	error: unknown,
	callback: (connectError: ConnectError) => T = () => {
		throw error;
	},
) => {
	const connectError = ConnectError.from(error);

	if (
		connectError.cause instanceof TypeError &&
		connectError.cause.message === "Failed to fetch"
	) {
		return new NetworkError();
	}

	if (connectError.code === Code.Internal) {
		return new InternalServerError();
	}

	return callback(connectError);
};
