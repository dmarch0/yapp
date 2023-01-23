import { Paths } from "const/Paths";
import { UserState } from "state/user/user.types";

export type ControlCallback = (user: UserState) => boolean
export type RedirectCallback = (user: UserState) => Paths