import dayjs from "dayjs";
import { computed, ref } from "vue";
import { isResultError } from "~/domain/error";
import { eventToDisplay } from "~/presentation/presenters/event";
import { displayNormalEvent } from "~/presentation/viewmodels/event";
import { schoolCalendarUseCase } from "~/usecases";

const event = ref<string>(displayNormalEvent);

const initializeEvent = async () => {
	const result = await schoolCalendarUseCase.getEventByDate(dayjs());
	if (isResultError(result)) throw result;
	if (result !== null) event.value = eventToDisplay(result);
};

const useEvent = () => {
	return {
		event: computed(() => event.value),
		initializeEvent,
	};
};

export default useEvent;
