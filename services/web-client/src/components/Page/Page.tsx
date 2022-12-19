import React from 'react';

import styles from './Page.module.scss'

interface Props {
    children: React.ReactNode
}

const Page: React.FunctionComponent<Props> = ({ children }) => {
    return (
        <div className={styles.wrapper}>
            {children}
        </div>
    )
}

export default Page;