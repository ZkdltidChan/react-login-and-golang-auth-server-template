import React, { } from "react";
import {
    useDisclosure,
    Drawer as ChakraDrawer,
    DrawerOverlay as ChakraDrawerOverlay,
    DrawerContent as ChakraDrawerContent,
    DrawerCloseButton as ChakraDrawerCloseButton,
} from '@chakra-ui/react';
import {
    HamburgerIcon,
    CloseIcon,
} from '@chakra-ui/icons';
import { IconButton } from "./IconButton";

export function Drawer({
    triggerButtonOpenIcon,
    triggerButtonCloseIcon,
    triggerButtonText,
    placement = "top",
    children, ...rest
}) {
    const { isOpen, onOpen, onClose } = useDisclosure()
    const btnRef = React.useRef()
    return (
        <>
            <IconButton
                size={'md'}
                onClick={isOpen ? onClose : onOpen}
                icon={isOpen ? triggerButtonCloseIcon : triggerButtonOpenIcon}
            >
                {triggerButtonText}
            </IconButton>

            <ChakraDrawer
                isOpen={isOpen}
                placement={placement}
                onClose={onClose}
                finalFocusRef={btnRef}
            >
                <ChakraDrawerOverlay />
                <ChakraDrawerContent>
                    <ChakraDrawerCloseButton />
                    {children}
                </ChakraDrawerContent>
            </ChakraDrawer>
        </>
    )
}
