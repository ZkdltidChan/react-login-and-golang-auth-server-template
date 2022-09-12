import React from "react";
import {
    Box,
    HStack,
    Avatar,
    VStack,
    Button,
} from '@chakra-ui/react';
// import NavLink from "../NavLink";
import { logout, useAuthDispatch } from "../../../hook/auth";
import { Popover } from "../../Buttons/Popover";
import { IconButton } from "../../Buttons/IconButton";
import { BellIcon } from "@chakra-ui/icons";
import NavLink from "../NavLink";

export default function AuthMenu({ user }) {
    const dispatch = useAuthDispatch()
    const handleLogout = async () => {
        // await logout(dispatch)
        console.log("LLL")
        // navigator('/')
    }

    return (
        <HStack
            as={'nav'}
            spacing={4}
        >
            <>
                <Box>
                    <IconButton icon={<BellIcon />} />
                </Box>
                <Popover
                    triggerButtonIcon={<Avatar src={user.img}></Avatar>}
                    header="User Setting"
                >
                    <VStack align="start">
                        <NavLink to="/user_menu1" >User Menu1</NavLink>
                        <NavLink to="/user_menu2" >User Menu2</NavLink>
                        <NavLink to="/user_menu3" >User Menu3</NavLink>
                    </VStack>
                    <VStack align="end" p={3}>
                        <HStack>
                            <Button onClick={handleLogout} >
                                <NavLink to="/">
                                    Logout
                                </NavLink>
                            </Button>
                        </HStack>
                    </VStack>
                </Popover>
            </>
        </HStack>)
}

