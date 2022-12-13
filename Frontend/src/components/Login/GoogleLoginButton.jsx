import React, {  } from 'react';
import {
    Button,
} from "@chakra-ui/react";
import { FcGoogle } from 'react-icons/fc'


function GoogleLoginButton() {
    function onClick() {
        // TODO
        console.log("TODO: google login")
    }
    return (
        <Button disabled={true} bg="gray" variant="outline" leftIcon={<FcGoogle />} onClick={onClick}>
            Login With Google
        </Button>
        )
}

export default GoogleLoginButton;