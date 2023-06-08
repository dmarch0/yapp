import React from 'react';
import { Page } from 'components/Page';
import withAccessControl from 'helpers/components/withAccessControl';
import {isAuthUser, toHome} from 'const/AccessControl';

const ProfilePage: React.FunctionComponent = () => {
  return (
    <Page>
      kek
    </Page>
  )
}

export default withAccessControl(ProfilePage, isAuthUser, toHome);