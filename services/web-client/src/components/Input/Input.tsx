import React, { ChangeEventHandler } from 'react';

interface Props {
    onChange: ChangeEventHandler<HTMLInputElement>
}

const Input: React.FunctionComponent<Props> = (props) => {
    const { onChange } = props;
    return (
        <input onChange={onChange}/>
    )
}

export default Input;