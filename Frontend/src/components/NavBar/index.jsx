import React from "react";
import { Outlet } from "react-router-dom";
import {
    Box,
    Flex,
    HStack,
    Stack,
    useColorModeValue,
} from '@chakra-ui/react';
import { ColorModeSwitcher } from './../../ColorModeSwitcher';
import {
    HamburgerIcon,
} from '@chakra-ui/icons';
import Logo from "../Logo";
import AuthMenu from "./AuthMenu";
import NavLink from "./NavLink";
import { Drawer } from "../Buttons/Drawer";
const Links = [
    { to: "embed_page_example", text: "Embed Page Example" },
    { to: "contact_us", text: "Contact Us" },
    { to:"mock_example", text: "Mock Example"}
]

const NavLogo = () => (
    <Flex as={NavLink} linkHover={false} justify="center" to="/">
        <Logo />
    </Flex>
)

const MenuList = ({ Links, logo, user }) => {
    const MobileMenuList = () => {
        return (
            <Drawer
                display={{
                    base: 'flex',
                    md: 'none'
                    // mobile 不顯示, desktop 顯示
                }}
                header={<>
                    <HStack justify="space-between">
                        {logo}
                    </HStack>
                </>
                }
                triggerButtonOpenIcon={<HamburgerIcon />}
            >
                <Stack as={'nav'} spacing={3}>
                    {Links.map((link, index) => (
                        <NavLink key={index} to={link.to} >{link.text}</NavLink>
                    ))}
                    <Box alignSelf="end" p={1}>
                        <ColorModeSwitcher borderRadius="full" />
                    </Box>
                </Stack>

            </Drawer>
        )
    }
    return (
        <HStack>
            <HStack
                as={'nav'}
                spacing={4}
                display={{
                    base: 'none',
                    md: 'flex'
                    // mobile 不顯示, desktop 顯示
                }}
            >
                {Links.map((link, index) => (
                    <NavLink key={index} to={link.to} >{link.text}</NavLink>
                ))}
                {!user ?
                    <>
                        <NavLink to="/login" >Login</NavLink>
                        <NavLink
                            display={{
                                base: 'none',
                                md: 'flex'
                            }} to="/signup">Sign up</NavLink>
                    </>
                    : <></>}
                <ColorModeSwitcher borderRadius="full" />
            </HStack>
            {user ? <AuthMenu user={user} /> : <></>}
            <MobileMenuList />
        </HStack>
    )
}

const FAKE_USER = {
    username: "a",
    img: 'https://bit.ly/dan-abramov',
}

export default function NavBar({ user }) {
    // const isDesktop = useBreakpointValue({
    //     base: false,
    //     lg: true,
    //   })
    return (
        <>
            <Flex
                shadow="md"
                bg={useColorModeValue('white', 'gray.800')}
                color={useColorModeValue('gray.600', 'white')}
                borderStyle={'solid'}
                borderColor={useColorModeValue('gray.200', 'gray.900')}
                align={'center'}
                w="100%"
                h={16}
                p={3}
                direction="row"
                justify="space-between"
            >
                <Box>
                    {<NavLogo />}
                </Box>
                <MenuList Links={Links} logo={<NavLogo />} />
                {/* <MenuList Links={Links} logo={<NavLogo />} user={FAKE_USER} /> */}
            </Flex>
            <Outlet />
        </>
    );
}