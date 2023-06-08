import {atom, useSetRecoilState} from "recoil"
import client from "helpers/http/axios"
import type { LoginProps, UserActions, UserState, PostLoginResponse } from "./user.types";
import {AxiosResponse} from "axios";

export const userState = atom<UserState>({
    key: "user",
    default: null,
});

export function useUserActions(): UserActions {
    const setUser = useSetRecoilState(userState);

    async function login(props: LoginProps): Promise<AxiosResponse<PostLoginResponse>> {
        const result = await client({
            method: "POST",
            url: "user/login",
            data: {
                email: props.email,
                password: props.password,
            }
        });
        setUser(result.data.data);
        return result;
    }

    return {
        login
    }
}