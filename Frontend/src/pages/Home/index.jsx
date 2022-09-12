import React, {
    // 
} from "react";
import {
    Heading,
    Box,
    HStack,
    VStack,
} from "@chakra-ui/react";

import { Drawer } from "../../components/Buttons/Drawer";
import { Popover } from "../../components/Buttons/Popover";
import { Modal } from "../../components/Buttons/Modal";
import { FaHands } from "react-icons/fa";
import { HamburgerIcon } from "@chakra-ui/icons";
const Home = () => {
    return (
        <>
            <Heading>
                Home
            </Heading>
            <HStack>
                <Drawer
                    triggerButtonText="Drawer"
                >
                    <Heading>test</Heading>
                </Drawer>
                <Popover 
                    triggerButtonText="Popover">
                    <Heading>Body</Heading>
                </Popover>
                <Modal
                    triggerButtonText="Modal"
                >
                    <Heading>test</Heading>
                </Modal>
            </HStack>
        </>
    )
}
export default Home;