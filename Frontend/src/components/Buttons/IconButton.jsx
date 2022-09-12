import React, { } from "react";
import {
    Box,
    IconButton as ChakraIconButton,
    Button as ChakraButton,
} from '@chakra-ui/react';


export function IconButton({
    icon,
    children,
    ...rest
}) {
    return (
        <>
            {!icon && !children ?
                <ChakraButton
                    className="ChakraButton"
                    leftIcon={icon}
                    {...rest}
                >
                    Button
                </ChakraButton>
                :
                <>
                    {
                        children ?
                            <ChakraButton
                                className="ChakraButton"
                                leftIcon={icon}
                                {...rest}
                            >
                                {children}
                            </ChakraButton>
                            :
                            <ChakraIconButton
                                className="ChakraIconButton"
                                icon={icon}
                                {...rest}
                            />
                    }
                </>
            }
        </>
    )
}

