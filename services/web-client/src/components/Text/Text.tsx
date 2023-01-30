import { TextComponents } from "const/TextProps";
import React from "react";

interface Props {
    children: React.ReactNode,
    component: TextComponents
}

const Text: React.FunctionComponent<Props> = (props) => {
    const { 
        component,
        children,
     } = props;

    return React.createElement(component, {}, children)
}

Text.defaultProps = {
    component: TextComponents.span,
}

export default Text;