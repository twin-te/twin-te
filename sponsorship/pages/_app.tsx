import '../styles/globals.scss';
import type { AppProps } from 'next/app';
import '@stripe/stripe-js';
import * as dayjs from 'dayjs';
import 'dayjs/locale/ja';
import { DefaultSeo } from 'next-seo';
import SEO from '../next-seo.config';
import * as bulmaToast from 'bulma-toast';
import GoogleTagManager, { GoogleTagManagerId } from '@/components/GoogleTagManager';
import { googleTagManagerId } from '../utils/gtm';
import Layout from '@/components/Layout';

dayjs.locale('ja');

bulmaToast.setDefaults({
	message: "",
	position: 'top-center'
});

function MyApp({ Component, pageProps }: AppProps) {
	return (
		<Layout>
			<GoogleTagManager googleTagManagerId={googleTagManagerId as GoogleTagManagerId} />
			<DefaultSeo {...SEO} />
			<Component {...pageProps} />
		</Layout>
	);
}

export default MyApp;
