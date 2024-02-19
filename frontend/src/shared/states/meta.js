import { atom } from "recoil";

export const MetaState = atom({
  key: "appMetaData",
  default: {
    name: null,
    versionText: null,
  },
});
