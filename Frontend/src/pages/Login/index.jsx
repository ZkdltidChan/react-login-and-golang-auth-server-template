import React, {
    // 
} from "react";
import {
    useColorModeValue,
    Container,
    Flex,
    Heading,
    Input,
    FormControl,
    FormLabel,
    VStack,
    Button,
    Divider,
} from "@chakra-ui/react";
import { FaTwitter } from 'react-icons/fa'
import { FcGoogle } from 'react-icons/fc'
import { useState } from "react";
import { loginUser, useAuthDispatch, useAuthState } from "../../context";
import { useNavigate } from "react-router-dom";

const LoginFrom = ({ onApply }) => {
    const dispatch = useAuthDispatch()
    const { loading } = useAuthState()
    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const navigate = useNavigate()

    const handleLogin = async (e) => {
        e.preventDefault()
        let payload = { email, password }
        console.log(payload)
        try {
            const response = await loginUser(dispatch, payload)
            if(response){
                navigate("/")
            }
            // if login success, router page
        } catch (e) {
            console.error(e)
        }
        onApply()
    }

    return (
        < VStack w="100%" spacing={4} >
            <FormControl id="email">
                <FormLabel >Email address</FormLabel>
                <Input type="email" onChange={(e) => setEmail(e.target.value)} />
            </FormControl>
            <FormControl id="password">
                <FormLabel>Password</FormLabel>
                <Input type="password" onChange={(e) => setPassword(e.target.value)} />
            </FormControl>
            <Button colorScheme="linkedin" onClick={handleLogin} isLoading={loading}>Login</Button>
        </VStack >
    )
}


// TODO: third party login api
const ThirdPartyLogin = () => {
    const LoginWithGoogle = () => (
        <Button bg="white" variant="outline" leftIcon={<FcGoogle />} onClick={() => console.log("login with google")}>
            Login With Google
        </Button>
    )
    const LoginWithTwitter = () => (
        <Button colorScheme='twitter' leftIcon={<FaTwitter />} onClick={() => console.log("login with twitter")}>
            Login With Twitter
        </Button>
    )
    return (
        <VStack w="100%">
            <LoginWithGoogle />
            <LoginWithTwitter />
        </VStack>
    )
}

const Login = () => {
    return (
        <>
            <Container p={3} py={5}>
                <VStack>
                    <Flex
                        shadow="xl"
                        rounded="lg"
                        p={8}
                        maxW="lg"
                        minW={{ base: "100%", md: "md" }}
                        justify="center"
                        bg={useColorModeValue('white', 'gray.700')}
                    >
                        <VStack w="100%" spacing={5}>
                            <Heading textAlign="center">
                                Login
                            </Heading>
                            <LoginFrom onApply={() => console.log("yo")} />
                            <Divider />
                            <ThirdPartyLogin />
                        </VStack>
                    </Flex>
                </VStack>
            </Container>
        </>
    )
}
export default Login;