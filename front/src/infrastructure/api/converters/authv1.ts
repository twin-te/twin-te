import { Provider, User } from "~/domain/user";
import * as AuthV1PB from "~/infrastructure/api/gen/auth/v1/type_pb";
import { fromPBUUID } from "./shared";
import { assurePresence } from "./utils";

export const fromPBProvider = (pbProvider: AuthV1PB.Provider): Provider => {
  switch (pbProvider) {
    case AuthV1PB.Provider.GOOGLE:
      return "google";
    case AuthV1PB.Provider.APPLE:
      return "apple";
    case AuthV1PB.Provider.TWITTER:
      return "twitter";
  }
  throw Error(`invalid enum ${pbProvider}`);
};

export const toPBProvider = (provider: Provider): AuthV1PB.Provider => {
  switch (provider) {
    case "google":
      return AuthV1PB.Provider.GOOGLE;
    case "apple":
      return AuthV1PB.Provider.APPLE;
    case "twitter":
      return AuthV1PB.Provider.TWITTER;
  }
};

export const fromPBUser = (pbUser: AuthV1PB.User): User => {
  return {
    id: fromPBUUID(assurePresence(pbUser.id)),
    providers: pbUser.authentications.map((pbUserAuthentication) =>
      fromPBProvider(pbUserAuthentication.provider)
    ),
  };
};
