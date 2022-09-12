import React, { } from "react";
import {
    Box,
    Popover as ChakraPopover,
    PopoverTrigger as ChakraPopoverTrigger,
    PopoverContent as ChakraPopoverContent,
    PopoverBody as ChakraPopoverBody,
    PopoverHeader as ChakraPopoverHeader,
    PopoverArrow as ChakraPopoverArrow,
    PopoverCloseButton as ChakraPopoverCloseButton,
} from '@chakra-ui/react';

import { IconButton } from "./IconButton";
export function Popover({
    triggerButtonIcon,
    triggerButtonText,
    header = "Plz Set you Header",
    children,
    ...rest
}) {
    return (
        <>
            <ChakraPopover>
                <ChakraPopoverTrigger>
                    <Box>
                        <IconButton
                            icon={triggerButtonIcon}
                            {...rest}
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
                    <ChakraPopoverBody>
                        {children}
                    </ChakraPopoverBody>
                </ChakraPopoverContent>
            </ChakraPopover>
        </>
    )
}
