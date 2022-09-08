import React from "react";
import { Link as RouterLink, Outlet } from "react-router-dom";
import {
    Box,
    Flex,
    HStack,
    IconButton,
    Stack,
    Link as ChakraLink,
    useColorModeValue,
    useDisclosure,
    Drawer,
    DrawerOverlay,
    DrawerContent,
    DrawerCloseButton,
    DrawerHeader,
    DrawerBody,
    Avatar,
    Menu as ChakraMenu,
    MenuItem as ChakraMenuItem,
    MenuButton as ChakraMenuButton,
    MenuList as ChakraMenuList,
} from '@chakra-ui/react';
import {ColorModeSwitcher} form "../../../ColorModeSwitcher.js"
import {
    HamburgerIcon,
    CloseIcon,
} from '@chakra-ui/icons';
import Logo from "../Logo";

const Links = [
    { to: "embed_page_example", text: "Embed Page Example" },
    { to:"contact_us", text: "Contact Us"},
]


const NavLink = ({ to, children, ...rest }) => (
    <ChakraLink
        as={RouterLink}
        px={2}
        py={1}
        rounded={'md'}
        _hover={{
            textDecoration: 'none',
            bg: useColorModeValue('gray.200', 'gray.700'),
        }}
        {...rest}
        to={to}>
        {children}
    </ChakraLink>
);

const NavLogo = () => (
    <Flex as={NavLink} justify="center" to="/">
        <Logo />
    </Flex>
)

function MenuDrawer({ Header, Links }) {
    const { isOpen, onOpen, onClose } = useDisclosure()
    const btnRef = React.useRef()
    return (
        <>
            <IconButton
                bg={useColorModeValue('white', 'gray.800')}
                color={useColorModeValue('gray.600', 'white')}
                size={'md'}
                onClick={isOpen ? onClose : onOpen}
                display={{ base: 'flex', md: 'none' }}
                icon={isOpen ? <CloseIcon /> : <HamburgerIcon />}
            />
            <Drawer
                isOpen={isOpen}
                placement='top'
                onClose={onClose}
                finalFocusRef={btnRef}
            >
                <DrawerOverlay />
                <DrawerContent>
                    <DrawerCloseButton />
                    <DrawerHeader>{Header}</DrawerHeader>
                    <DrawerBody>
                        <Box pb={4} display={{ md: 'none' }}>
                            <Stack as={'nav'} onClick={onClose} spacing={4}>
                                {Links.map((link, index) => (
                                    <NavLink key={index} to={link.to} >{link.text}</NavLink>
                                ))}
                            </Stack>
                        </Box>
                    </DrawerBody>
                </DrawerContent>
            </Drawer>
        </>
    )
}


const MenuList = ({ Links }) => (
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
    </HStack>
)

const AuthMenu = ({ auth }) => (
    <HStack
        as={'nav'}
        spacing={4}
    >
        {auth ?
            <ChakraMenu>
                <ChakraMenuButton as={Avatar} size={"sm"}>
                    <ChakraMenuList>
                        <ChakraMenuItem>Setting</ChakraMenuItem>
                        <ChakraMenuItem>Logout</ChakraMenuItem>
                    </ChakraMenuList>
                </ChakraMenuButton>
            </ChakraMenu>
            :
            <HStack>
                <NavLink to="/login" >Login</NavLink>
                <NavLink
                    display={{
                        base: 'none',
                        md: 'flex'
                    }} to="/signup">Sign up</NavLink>
            </HStack>
        }
    </HStack>
)

export default function NavBar() {
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
                <MenuDrawer Header={<NavLogo />} Links={Links} />
                <Box>
                    <HStack>
                        <Box>
                            {<NavLogo />}
                        </Box>
                        <ColorModeSwitcher/>
                        <MenuList Links={Links} />
                    </HStack>
                </Box>
                {/* <AuthMenu auth={true} /> */}
                <AuthMenu auth={false} />
            </Flex>
            <Outlet />
        </>
    );
}