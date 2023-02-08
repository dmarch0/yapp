import React, { useMemo } from 'react';
import type { InputProps } from './Input.types';

function Input<T>(props: InputProps<T>, ref: React.ForwardedRef<HTMLInputElement>) {
    const { label } = props;

    const inputProps = useMemo(() => {
        if ("register" in props) {
            const formProps = props.register(props.name);
            if (ref) {
                formProps.ref = (el) => {
                    ref = el;
                    formProps.ref(el);
                }
            }
            return formProps;
        }

        if ("onChange" in props) {
            const { 
                onChange, 
                value,
             } = props;

            return {
                onChange,
                value,
                ref,
            }
        }

        return {};
    }, [props]);

    if (label) {
        return (
            <label>
                {label}
                <input {...inputProps}/>
            </label>
        )
    }

    return (
        <input {...inputProps} />
    )
};

type ExportAssertionType = <T>(
    props: InputProps<T> & { ref?: React.ForwardedRef<HTMLInputElement>  }
) => ReturnType<typeof Input>;

export default React.forwardRef(Input) as ExportAssertionType;