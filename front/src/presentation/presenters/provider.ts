import { Provider } from "~/domain/user";
import { DisplayProvider } from "../viewmodels/provider";

export const providerMap: Record<Provider, DisplayProvider> = {
  google: "Google",
  apple: "Apple",
  twitter: "X",
};
