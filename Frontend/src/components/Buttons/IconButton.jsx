import React, { } from "react";
import {
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
                    variant="ghost"
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
                                variant="ghost"
                                className="ChakraButton"
                                leftIcon={icon}
                                {...rest}
                            >
                                {children}
                            </ChakraButton>
                            :
                            <ChakraIconButton
                                variant="ghost"
                                className="ChakraIconButton"
                                borderRadius="full"
                                icon={icon}
                                {...rest}
                            />
                    }
                </>
            }
        </>
    )
}

