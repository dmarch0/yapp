import React from 'react';
import { BrowserRouter, createBrowserRouter, Route, RouterProvider } from 'react-router-dom';
import { Home } from '../../pages/Home';

const router = createBrowserRouter([
    {
        path: "/",
        element: <Home />
    }
])

const Router: React.FunctionComponent = () => <RouterProvider router={router}/>

export default Router;