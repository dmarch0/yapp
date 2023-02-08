import { ChangeEventHandler } from 'react';
import { FieldValues, Path, UseFormRegister } from 'react-hook-form';

type CommonInputProps = {
    label?: React.ReactNode
}

type UncontrolledInputProps<T extends FieldValues> = {
    register: UseFormRegister<T>,
    name: Path<T>
}

type ControlledInputProps = {
    onChange: ChangeEventHandler<HTMLInputElement>,
    value: string,
}

export type InputProps<T = void> = CommonInputProps & (T extends FieldValues ? UncontrolledInputProps<T> : ControlledInputProps)