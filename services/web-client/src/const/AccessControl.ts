import { UserState } from "state/user/user.types";
import { Paths } from "./Paths";

export const isAuthUser = (user: UserState) => !!user;
export const isNonAuthUser = (user: UserState) => !user;

export const toProfile = () => Paths.Profile
export const toHome = () => Paths.Home