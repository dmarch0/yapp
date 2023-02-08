import React from 'react';
import { useForm } from 'react-hook-form';
import { joiResolver } from '@hookform/resolvers/joi'
import styles from './LoginForm.module.scss';
import Joi from 'joi';
import { Input } from 'components/Input';
import type { FormValues } from './LoginForm.types';

const LoginSchema = Joi.object<FormValues>({
    email: Joi.string().email({ tlds: { allow: false } }).required(),
    password: Joi.string().required(),
})

const LoginForm: React.FunctionComponent = () => {
    const { register, formState, watch } = useForm<FormValues>({
        resolver: joiResolver(LoginSchema)
    });

    return (
        <div className={styles.wrapper}>
            <Input<FormValues> register={register} name="email"/>
            <Input<FormValues> register={register} name="password"/>
        </div>
    )
}

export default LoginForm;