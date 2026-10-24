export type User = {
  id: string;
  providers: Provider[];
};

export type Provider = "google" | "apple" | "twitter";
