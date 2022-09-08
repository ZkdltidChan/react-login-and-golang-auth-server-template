import React, { useState } from 'react';
import {
    useColorModeValue,
    Button,
} from "@chakra-ui/react";
import { FaTwitter } from 'react-icons/fa'


function TwitterLoginButton({ }) {
    function onClick() {
        // TODO
        console.log("TODO: twitter login")
    }
    return (
        <Button colorScheme='twitter' leftIcon={<FaTwitter />} onClick={onClick}>
            Login With Twitter
        </Button>
    )
}

export default TwitterLoginButton;