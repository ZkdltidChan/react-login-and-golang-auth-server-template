import React, {  } from 'react';
import {
    useColorModeValue,
    Button,
} from "@chakra-ui/react";
import { FcGoogle } from 'react-icons/fc'


function GoogleLoginButton({ clientId, scope = 'identify', redirectUri }) {
    const handleMessage = (e) => {
        console.log(e)
    }

    function onClick() {
        // TODO
        console.log("TODO: google login")
    }
    return (
        <Button bg="white" variant="outline" leftIcon={<FcGoogle />} onClick={onClick}>
            Login With Google
        </Button>
        )
}

export default GoogleLoginButton;