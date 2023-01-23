import { atom } from "recoil"
import client from "helpers/http/axios"
import type { LoginProps, UserActions, UserState } from "./user.types";

export const userState = atom<UserState>({
    key: "user",
    default: null,
});


export function useUserActions(): UserActions {

    async function login(props: LoginProps) {
        const result = await client({
            method: "POST",
            url: "user/login",
            data: {
                email: props.email,
                password: props.password,
            }
        });
        console.log(result)
    }

    return {
        login
    }
}