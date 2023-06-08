import React from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import ProfilePage from "../../pages/ProfilePage";

const HomePage = React.lazy(() => import('pages/HomePage'));

const router = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />
    },
    {
        path: "/profile",
        element: <ProfilePage />
    }
])

const Router: React.FunctionComponent = () => <RouterProvider router={router}/>

export default Router;