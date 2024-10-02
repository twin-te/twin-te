import {
	Code,
	type ConnectError,
	type PromiseClient,
	type Transport,
	createPromiseClient,
} from "@connectrpc/connect";
import {
	type InternalServerError,
	type NetworkError,
	UnauthenticatedError,
} from "~/domain/error";
import type { User } from "~/domain/user";
import { fromPBUser } from "~/infrastructure/api/converters/authv1";
import { assurePresence } from "~/infrastructure/api/converters/utils";
import { AuthService } from "~/infrastructure/api/gen/auth/v1/service_connect";
import { handleError } from "~/infrastructure/api/utils";

export interface IAuthUseCase {
	getMe(): Promise<
		User | UnauthenticatedError | NetworkError | InternalServerError
	>;

	deleteUser(): Promise<
		null | UnauthenticatedError | NetworkError | InternalServerError
	>;
}

export class AuthUseCase implements IAuthUseCase {
	#client: PromiseClient<typeof AuthService>;

	constructor(transport: Transport) {
		this.#client = createPromiseClient(AuthService, transport);
	}

	async getMe(): Promise<
		User | UnauthenticatedError | NetworkError | InternalServerError
	> {
		return this.#client
			.getMe({})
			.then((res) => fromPBUser(assurePresence(res.user)))
			.catch((error) => {
				return handleError(error, (connectError: ConnectError) => {
					if (connectError.code === Code.Unauthenticated) {
						return new UnauthenticatedError();
					}

					throw error;
				});
			});
	}

	async deleteUser(): Promise<
		null | UnauthenticatedError | NetworkError | InternalServerError
	> {
		return this.#client
			.deleteAccount({})
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
