import React from 'react';
import { useForm } from 'react-hook-form';
import { joiResolver } from '@hookform/resolvers/joi'
import styles from './LoginForm.module.scss';
import Joi from 'joi';
import { Input } from 'components/Input';
import type { FormValues } from './LoginForm.types';
import { Button } from 'components/Button';
import {useUserActions} from "state/user/user.state";
import {redirect} from "react-router-dom";

const LoginSchema = Joi.object<FormValues>({
    email: Joi.string().email({ tlds: { allow: false } }).required(),
    password: Joi.string().required(),
})

const LoginForm: React.FunctionComponent = () => {
    const { login } = useUserActions();
    const { register, handleSubmit } = useForm<FormValues>({
        resolver: joiResolver(LoginSchema)
    });


    const onSubmit = handleSubmit(async (data) => {
        const result = await login(data);
        if (result.status === 200) {
            await redirect("/profile");
        }
    });

    return (
        <div className={styles.wrapper}>
            <form onSubmit={onSubmit}>
                <Input<FormValues> register={register} name="email"/>
                <Input<FormValues> register={register} name="password"/>
                <Button type="submit">
                    less go
                </Button>
            </form>
        </div>
    )
}

export default LoginForm;