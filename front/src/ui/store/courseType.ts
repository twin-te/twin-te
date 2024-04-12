import { computed, ref } from "vue";
import { deepCopy } from "~/utils";

type CourseType = "Normal" | "Special";

const courseType = ref<CourseType>("Normal");

const toggleCourseType = () => {
  courseType.value = courseType.value === "Normal" ? "Special" : "Normal";
};

const useCourseType = () => {
  return {
    courseType: computed(() => deepCopy(courseType.value)),
    toggleCourseType,
  };
};

export default useCourseType;
