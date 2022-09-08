import React, {  } from 'react';
import {
    useColorModeValue,
    Button,
} from "@chakra-ui/react";
import { FaDiscord } from 'react-icons/fa'


function DiscordLoginButton() {
    const scope = "identify"
    const clientId=process.env.REACT_APP_DISCORD_CLIENT_ID
    const redirectUri=process.env.REACT_APP_AUTH_SERVER
    const handleMessage = (e) => {
        console.log(e)
    }

    function onClick() {
        const dPopup = window.open(`https://discord.com/api/oauth2/authorize?response_type=token&client_id=${clientId}&scope=${scope}&redirect_uri=${redirectUri}`, 'discordLogin', 'height=700,width=500')
        window.addEventListener('message', handleMessage, false)
    }
    return (
        <Button colorScheme="telegram" leftIcon={<FaDiscord />} onClick={onClick}>
            Login with Discord
        </Button>
        )
}

export default DiscordLoginButton;