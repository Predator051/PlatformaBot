import './App.css';
import { AuthScreen } from "./auth/Auth";
import { ChakraProvider } from '@chakra-ui/react'
import {MainScreen} from "./main/MainScreen";
import {createBrowserRouter, redirect, RouterProvider, useNavigate} from "react-router-dom";

function AuthCheck({component: Component}) {
    const navigate = useNavigate()

    if (!localStorage.getItem('Session')) {
        return navigate('/auth')
    }

    return (<Component/>)
}

function App() {
    const router = createBrowserRouter([
        {
            path: '/',
            element: <AuthCheck component={MainScreen}></AuthCheck>
        },
        {
            path: '/auth',
            element: <AuthScreen></AuthScreen>
        }]);
  return (
      <ChakraProvider>
          <div className="App">
                <RouterProvider router={router}></RouterProvider>
          </div>
      </ChakraProvider>
  );
}

export default App;
