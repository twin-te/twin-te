import { computed, ref } from "vue";

const creditYear = ref<number>(0); // 0 means all year.

const setCreditYear = (year: number) => {
	creditYear.value = year;
};

const setCreditYearToAll = () => {
	creditYear.value = 0;
};

const useCreditYear = () => {
	return {
		creditYear: computed(() => creditYear.value),
		setCreditYear,
		setCreditYearToAll,
	};
};

export default useCreditYear;
