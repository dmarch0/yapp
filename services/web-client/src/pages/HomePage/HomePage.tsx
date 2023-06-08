import React from 'react';
import { Page } from 'components/Page';
import withAccessControl from 'helpers/components/withAccessControl';
import { isNonAuthUser, toProfile } from 'const/AccessControl';
import { LoginForm } from 'components/LoginForm';

const HomePage: React.FunctionComponent = () => {
    return (
        <Page>
            <LoginForm />
        </Page>
    )
}

export default withAccessControl(HomePage, isNonAuthUser, toProfile);