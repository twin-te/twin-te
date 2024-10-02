import "../styles/globals.scss";
import type { AppProps } from "next/app";
import "@stripe/stripe-js";
import * as dayjs from "dayjs";
import "dayjs/locale/ja";
import GoogleTagManager, {
	type GoogleTagManagerId,
} from "@/components/GoogleTagManager";
import Layout from "@/components/Layout";
import * as bulmaToast from "bulma-toast";
import { DefaultSeo } from "next-seo";
import SEO from "../next-seo.config";
import { googleTagManagerId } from "../utils/gtm";

dayjs.locale("ja");

bulmaToast.setDefaults({
	message: "",
	position: "top-center",
});

function MyApp({ Component, pageProps }: AppProps) {
	return (
		<Layout>
			<GoogleTagManager
				googleTagManagerId={googleTagManagerId as GoogleTagManagerId}
			/>
			<DefaultSeo {...SEO} />
			<Component {...pageProps} />
		</Layout>
	);
}

export default MyApp;
