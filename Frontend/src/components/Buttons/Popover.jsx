import React, { } from "react";
import {
    Button,
    Box,
    Popover as ChakraPopover,
    PopoverTrigger as ChakraPopoverTrigger,
    PopoverContent as ChakraPopoverContent,
    PopoverHeader as ChakraPopoverHeader,
    PopoverArrow as ChakraPopoverArrow,
    PopoverCloseButton as ChakraPopoverCloseButton,
} from '@chakra-ui/react';

import { IconButton } from "./IconButton";
export function Popover({
    triggerButtonIcon,
    triggerButtonText,
    header = "Plz Set you Header",
    children
}) {
    return (
        <>
            <ChakraPopover>
                <ChakraPopoverTrigger>
                    <Box>
                        <IconButton
                            icon={triggerButtonIcon}
                        >
                            {triggerButtonText}
                        </IconButton>
                    </Box>
                    {/* <Button>?</Button> */}
                </ChakraPopoverTrigger>
                <ChakraPopoverContent>
                    <ChakraPopoverArrow />
                    <ChakraPopoverCloseButton />
                    <ChakraPopoverHeader>
                        {header}
                    </ChakraPopoverHeader>
                    {children}
                </ChakraPopoverContent>
            </ChakraPopover>
        </>
    )
}
