import dayjs from "dayjs";
import { computed, ref } from "vue";
import { isResultError } from "~/domain/error";
import {
  BaseModule,
  SchoolCalendarModule,
  isBaseModule,
} from "~/domain/module";
import { schoolCalendarUseCase } from "~/usecases";

const module = ref<BaseModule>("SpringA");
const currentModule = ref<BaseModule>("SpringA");
const moduleInitialized = ref<boolean>(false);

const setModule = (newModule: BaseModule) => {
  module.value = newModule;
};

const setToCurrentModule = () => {
  module.value = currentModule.value;
};

const schoolCalendarModuleToBaseModule = (
  schoolCalendarModule: SchoolCalendarModule
): BaseModule => {
  if (isBaseModule(schoolCalendarModule)) return schoolCalendarModule;

  const now = dayjs();

  switch (schoolCalendarModule) {
    case "SummerVacation":
      return "SpringC";
    case "WinterVacation":
      if (now.month() === 11) return "FallB";
      return "FallC";
    case "SpringVacation":
      if (now.month() === 2) return "FallC";
      return "SpringA";
  }
};

const initializeModule = async () => {
  if (moduleInitialized.value) return;
  return schoolCalendarUseCase.getCurrentModule().then((result) => {
    if (isResultError(result)) throw result;

    const baseModule: BaseModule = schoolCalendarModuleToBaseModule(result);

    module.value = currentModule.value = baseModule;
    moduleInitialized.value = true;
  });
};

const useModule = () => {
  return {
    module: computed(() => module.value),
    currentModule: computed(() => currentModule.value),
    isCurrentModule: computed(() => module.value == currentModule.value),
    setModule,
    setToCurrentModule,
    initializeModule,
  };
};

export default useModule;
