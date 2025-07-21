import "core-js/features/array/at";
import { createGtm } from "@gtm-support/vue-gtm";
import * as Sentry from "@sentry/vue";
import { createHead } from "@vueuse/head";
import { createApp } from "vue";
import VueClickAway from "vue3-click-away";
import App from "./app/App.vue";
import { router } from "./route";
import "./styles/_index.scss";

const app = createApp(App);

Sentry.init({
  app,
  dsn: String(import.meta.env.VITE_APP_SENTRY_URL ?? ""),
  environment: import.meta.env.DEV ? "development" : undefined,
  integrations: [
    Sentry.browserTracingIntegration({
      router,
    }),
    Sentry.replayIntegration({
      maskAllText: false,
    }),
  ],
  tracePropagationTargets: ["localhost", /^https:\/\/app\.twinte\.net/],
  tracesSampleRate: 1.0,
  replaysSessionSampleRate: 0.01,
  replaysOnErrorSampleRate: 1.0,
});

const head = createHead();

app
  .use(router)
  .use(VueClickAway)
  .use(head)
  .use(
    createGtm({
      id: "GTM-PHSLD8B",
      vueRouter: router,
      enabled: import.meta.env.PROD,
      debug: import.meta.env.DEV,
    })
  )
  .mount("#app");
