export interface UserActions {
    login(props: LoginProps): Promise<void>
}

export interface User {
    id: string
}

export type UserState = User | null

export interface LoginProps {
    email: string;
    password: string;
}