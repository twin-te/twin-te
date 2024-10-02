import { useSwitch } from "../hooks/useSwitch";

const [isVisibleSidebar, openSidebar, closeSidebar, toggleSidebar, setSidebar] =
	useSwitch(false);

const useSidebar = () => {
	return {
		isVisibleSidebar,
		openSidebar,
		closeSidebar,
		toggleSidebar,
		setSidebar,
	};
};

export default useSidebar;
