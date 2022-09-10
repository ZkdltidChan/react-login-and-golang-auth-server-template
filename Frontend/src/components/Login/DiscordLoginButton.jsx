import React, { useState } from 'react';
import {
    Button,
} from "@chakra-ui/react";
import { FaDiscord } from 'react-icons/fa'
import { useNavigate } from "react-router-dom";
import OAuth2Login from "react-simple-oauth2-login";
import { loginUserWithDiscord } from '../../hook/auth/actions';
import { useAuthDispatch }  from "../../hook/auth";

function DiscordLoginButton() {
    const dispatch = useAuthDispatch()
    const navigate = useNavigate()
    const [isLoading, setIsLoadign] = useState();
    const onSuccess = async(code) => {
        console.log(code);
        try {
            const response = await loginUserWithDiscord(dispatch, code)
            if (response) {
                navigate("/")
            }
        } catch (e) {
            console.error(e)
        }
        setIsLoadign(false)
    }
    const onFailure = response => {
        console.error(response);
        setIsLoadign(false)
    }

    return (
        <Button
            as={OAuth2Login}
            colorScheme="telegram"
            leftIcon={<FaDiscord />
            }
            authorizationUrl="https://discord.com/api/oauth2/authorize"
            responseType="code"
            clientId={process.env.REACT_APP_DISCORD_CLIENT_ID}
            redirectUri={process.env.REACT_APP_AUTH_SERVER}
            onSuccess={onSuccess}
            onFailure={onFailure}
            scope={"identify"}
            isLoading={isLoading}
            onRequest={()=>setIsLoadign(true)}
        >
            Login with Discord
        </Button>
    )
}

export default DiscordLoginButton;