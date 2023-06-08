import React from "react";
import { Navigate } from "react-router-dom";
import { useRecoilValue } from "recoil";
import { userState } from "state/user/user.state";
import type { ControlCallback, RedirectCallback } from "./withAccessControl.types";

export default function withAccessControl(Component: React.FunctionComponent, controlCallback: ControlCallback, redirectCallback: RedirectCallback): React.FunctionComponent {
    return (props) => {
        const user = useRecoilValue(userState);
        console.log('USER', user);
        if (!controlCallback(user)) {
            return <Navigate to={redirectCallback(user)} />
        }

        return <Component {...props}/>
    }
}