import React, { ChangeEventHandler } from 'react';
import { FieldValues, UseFormRegister } from 'react-hook-form';

type UncontrolledInputProps<T extends FieldValues> = {
    register: UseFormRegister<T>,
}

type ControlledInputProps = {
    onChange: ChangeEventHandler<HTMLInputElement>,
    value: string,
}

type InputProps<T = void> = T extends FieldValues ? UncontrolledInputProps<T> : ControlledInputProps

const Input = React.forwardRef<HTMLInputElement, InputProps>((props, ref) => (
    <input ref={ref} {...props} />
  ));

export default Input;