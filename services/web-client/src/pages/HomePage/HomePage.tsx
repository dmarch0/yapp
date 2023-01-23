import React from 'react';
import { Page } from 'components/Page';
import { useUserActions } from 'state/user/user.state';
import withAccessControl from 'helpers/components/withAccessControl';
import { isNonAuthUser, toProfile } from 'const/AccessControl';

const HomePage: React.FunctionComponent = () => {
    const userActions = useUserActions();

    return (
        <Page>
            <button onClick={() => userActions.login({ email: "test@test.test", password: "123" })}>
                login
            </button>
        </Page>
    )
}

export default withAccessControl(HomePage, isNonAuthUser, toProfile);