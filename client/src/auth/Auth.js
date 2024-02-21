import {Button, Center, Input, InputGroup, InputRightElement} from '@chakra-ui/react'
import React from "react";
import axios, {AxiosHeaders} from "axios";

export function AuthScreen() {
    const [show, setShow] = React.useState(false)
    const [token, setToken] = React.useState("")
    const handleClick = () => setShow(!show)

    const handleClickEnter = () => {
        axios.post('http://localhost:8080/auth',  {
            token: token
        }, {
            headers: new AxiosHeaders("Access-Control-Allow-Origin: *"),
        }).then((response) => {
            localStorage.setItem("Session", response.data)
        }, (error) => {
            console.log(error);
        });
    }

    return (
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
    )
}
