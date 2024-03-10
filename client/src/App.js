import './App.css';
import { AuthScreen } from "./auth/Auth";
import {Box, ChakraProvider} from '@chakra-ui/react'
import {MainScreen} from "./main/MainScreen";
import {createBrowserRouter, Navigate, redirect, RouterProvider, useNavigate} from "react-router-dom";

function AuthCheck({component: Component}) {
    const navigate = useNavigate()

    if (!localStorage.getItem('Session')) {
        console.log("No session");
        return <Navigate to={'/auth'}></Navigate>
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
          <Box className="App" h={'100vh'}>
                <RouterProvider router={router}></RouterProvider>
          </Box>
      </ChakraProvider>
  );
}

export default App;
