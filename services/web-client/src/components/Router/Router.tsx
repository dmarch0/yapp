import React from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

const HomePage = React.lazy(() => import('pages/HomePage'));

const router = createBrowserRouter([
    {
        path: "/",
        element: <HomePage />
    }
])

const Router: React.FunctionComponent = () => <RouterProvider router={router}/>

export default Router;