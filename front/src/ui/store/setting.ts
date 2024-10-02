import { useDark } from "@vueuse/core";
import { computed, ref } from "vue";
import { isResultError } from "~/domain/error";
import { type Setting, getInitialSetting } from "~/domain/setting";
import { currentAcademicYear, validateAcademicYear } from "~/domain/year";
import { settingUseCase } from "~/usecases";
import { deepCopy } from "~/utils";

const appliedYear = ref<number>(currentAcademicYear);

const isDark = useDark({
	selector: "body",
});

const setting = ref<Setting>(getInitialSetting());

const setAppliedYear = (displayYear: number) => {
	appliedYear.value = validateAcademicYear(displayYear)
		? displayYear
		: currentAcademicYear;
};

const updateSetting = (data: Partial<Setting>) => {
	return settingUseCase.updateSetting(data).then((result) => {
		if (isResultError(result)) throw result;
		setting.value = result;
		if ("darkMode" in data) {
			isDark.value = result.darkMode;
		}
		if ("displayYear" in data) {
			setAppliedYear(result.displayYear);
		}
	});
};

const initializeSetting = () => {
	return settingUseCase.getSetting().then((result) => {
		if (isResultError(result)) throw result;
		setting.value = result;
		isDark.value = result.darkMode;
		setAppliedYear(result.displayYear);
	});
};

const useSetting = () => {
	return {
		appliedYear: computed(() => appliedYear.value),
		isDark: computed(() => isDark.value),
		setting: computed(() => deepCopy(setting.value)),
		updateSetting,
		initializeSetting,
	};
};

export default useSetting;
