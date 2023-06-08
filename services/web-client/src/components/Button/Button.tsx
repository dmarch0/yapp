import React from 'react';
import type { ButtonTypes } from './Button.types';

interface Props {
    children: React.ReactNode,
    onClick?: React.MouseEventHandler<HTMLButtonElement>,
    type: ButtonTypes
}

const Button: React.FunctionComponent<Props> = ({ children, onClick, type }) => {
    return (
        <button onClick={onClick} type={type}>
            {children}
        </button>
    )
}

export default Button;