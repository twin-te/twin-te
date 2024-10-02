import type {
	InternalServerError,
	NetworkError,
	UnauthenticatedError,
} from "~/domain/error";
import {
	type Setting,
	getInitialSetting,
	settingProps,
} from "~/domain/setting";
import { LocalStorage } from "~/infrastructure/localstorage";
import { deepCopy, getKeysFromObj } from "~/utils";

export interface ISettingUseCase {
	getSetting(): Promise<Setting | NetworkError | InternalServerError>;

	updateSetting(
		data: Partial<Setting>,
	): Promise<
		Setting | UnauthenticatedError | NetworkError | InternalServerError
	>;
}

export class SettingUseCase implements ISettingUseCase {
	#localStorage: LocalStorage;
	#setting: Setting;

	constructor() {
		this.#localStorage = LocalStorage.getInstance();
		this.#setting = settingProps.reduce<Setting>((setting, prop) => {
			const value = this.#localStorage.get(prop);
			return value !== undefined ? { ...setting, [prop]: value } : setting;
		}, getInitialSetting());
	}

	async getSetting(): Promise<Setting | NetworkError | InternalServerError> {
		return deepCopy(this.#setting);
	}

	async updateSetting(
		data: Partial<Setting>,
	): Promise<
		Setting | UnauthenticatedError | NetworkError | InternalServerError
	> {
		getKeysFromObj(data).forEach((prop) => {
			this.#localStorage.set(prop, data[prop]);
		});

		this.#setting = { ...this.#setting, ...data };

		return deepCopy(this.#setting);
	}
}
