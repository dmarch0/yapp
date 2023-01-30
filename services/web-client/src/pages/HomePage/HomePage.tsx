import React from 'react';
import { Page } from 'components/Page';
import { useUserActions } from 'state/user/user.state';
import withAccessControl from 'helpers/components/withAccessControl';
import { isNonAuthUser, toProfile } from 'const/AccessControl';
import { LoginForm } from 'components/LoginForm';

const HomePage: React.FunctionComponent = () => {
    const userActions = useUserActions();

    return (
        <Page>
            <LoginForm />
        </Page>
    )
}

export default withAccessControl(HomePage, isNonAuthUser, toProfile);