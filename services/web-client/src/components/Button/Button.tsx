import React from 'react';

interface Props {
    children: React.ReactNode,
    onClick: React.MouseEventHandler<HTMLButtonElement>
}

const Button: React.FunctionComponent<Props> = ({ children, onClick }) => {
    return (
        <button onClick={onClick}>
            {children}
        </button>
    )
}

export default Button;