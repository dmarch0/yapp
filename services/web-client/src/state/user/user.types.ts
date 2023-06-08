import {AxiosResponse} from "axios";

export interface PostLoginResponse {
    data: UserState,
}

export interface UserActions {
    login(props: LoginProps): Promise<AxiosResponse<PostLoginResponse>>
}

export interface User {
    id: string
}

export type UserState = User | null

export interface LoginProps {
    email: string;
    password: string;
}