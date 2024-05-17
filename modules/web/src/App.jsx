import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import {
  HomeLayout,
  Landing,
  Error,
  Login,
  Register,
} from './pages'

const router = createBrowserRouter([
  {
    path: '/',
    element: <HomeLayout/>,
    errorElement: <Error/>,
    children: [
      {
      index: true,
      element: <Landing/>
      },
    ],
  }, 
  {
    path: '/login',
    element: <Login/>,
    errorElement: <Error/>,
  },
  {
    path: '/register',
    element: <Register/>,
    errorElement: <Error/>,
  }

])

const App = () => {
  return <RouterProvider router={router}/>
};
export default App;