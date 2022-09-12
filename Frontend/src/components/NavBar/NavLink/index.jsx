import React from "react";
import { Link as RouterLink } from "react-router-dom";
import {
    Link as ChakraLink,
} from '@chakra-ui/react';

export default function NavLink({ to, linkHover = true, children, ...rest }) {
    return (
        <ChakraLink
            as={RouterLink}
            px={2}
            py={1}
            rounded={'md'}
            _hover={!linkHover ? {
                textDecoration: 'none',
            } : null}
            {...rest}
            to={to}>
            {children}
        </ChakraLink>
    )
}