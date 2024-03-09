import {
    Alert,
    AlertIcon,
    Box,
    Button,
    Center,
    Input,
    InputGroup,
    InputRightElement,
    Spacer,
    Stack
} from '@chakra-ui/react'
import React from "react";
import axios, {AxiosHeaders} from "axios";
import {redirect, redirectDocument, useNavigate} from "react-router-dom";

export function AuthScreen() {
    const [show, setShow] = React.useState(false)
    const [token, setToken] = React.useState("")
    const [showError, setShowError] = React.useState(false)
    const handleClick = () => setShow(!show)

    const navigate = useNavigate()

    console.log("dfd", process.env.REACT_APP_PB_SERVER_IP);
    const handleClickEnter = () => {

        axios.post(`http://${process.env.REACT_APP_PB_SERVER_IP}:8080/auth`,  {
            token: token
        }, {
            headers: new AxiosHeaders("Access-Control-Allow-Origin: *"),
        }).then((response) => {
            localStorage.setItem("Session", response.data);
            setShowError(false);
            navigate('/');
        }, (error) => {
            setShowError(true);
            console.log(error);
        });
    }

    return (
        <Center>
            <Box>
                <Stack spacing={3}>
                <Center>
                    <InputGroup size='md'>
                        <Input
                            pr='4.5rem'
                            type={show ? 'text' : 'password'}
                            placeholder='Enter token'
                            value={token}
                            onChange={(event) => setToken(event.target.value)}
                        />
                        <InputRightElement width='4.5rem'>
                            <Button h='1.75rem' size='sm' onClick={handleClick}>
                                {show ? 'Hide' : 'Show'}
                            </Button>
                        </InputRightElement>
                    </InputGroup>
                    <Button colorScheme='blue' onClick={() => handleClickEnter()}>Enter</Button>
                </Center>
                <Alert status='error' color={"Black"} borderRadius={5} visibility={showError ? "visible" : "hidden"}>
                    <AlertIcon />
                    There was an error processing your request
                </Alert>
                </Stack>
            </Box>
        </Center>
    )
}
