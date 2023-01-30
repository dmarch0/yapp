import React from 'react';
import styles from './LoginForm.module.scss';

interface Props {
    children: React.ReactNode,
    onClick: React.MouseEventHandler<HTMLButtonElement>
}


const LoginForm: React.FunctionComponent = () => {
    return (
        <div className={styles.wrapper}>
            
        </div>
    )
}

export default LoginForm;