import React from 'react';
import type { ButtonTypes } from './Button.types';
interface Props {
    children: React.ReactNode,
    onClick: React.MouseEventHandler<HTMLButtonElement>,
    type: ButtonTypes
}

const Button: React.FunctionComponent<Props> = ({ children, onClick }) => {
    return (
        <button onClick={onClick}>
            {children}
        </button>
    )
}

export default Button;