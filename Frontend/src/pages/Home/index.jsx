import React, {
    // 
} from "react";
import {
    Heading,
    HStack,
    VStack,
    Text,
    Box,
} from "@chakra-ui/react";

import { Drawer } from "../../components/Buttons/Drawer";
import { Popover } from "../../components/Buttons/Popover";
import { Modal } from "../../components/Buttons/Modal";
const Home = () => {
    return (
        <Box p={5}>
            <Heading>
                Home
            </Heading>
            <HStack>
                
                {/* <Drawer
                    triggerButtonText="Drawer"
                    header="Drawer Example"
                >
                    <VStack>
                        <Heading>test1</Heading>
                        <Heading>test2</Heading>
                    </VStack>
                </Drawer>
                <Text>/</Text>

                <Popover
                    header="Popover Example"
                    triggerButtonText="Popover">
                    <VStack>
                        <Heading>test1</Heading>
                        <Heading>test2</Heading>
                    </VStack>
                </Popover>
                <Text>/</Text>
                <Modal
                    header="Modal Exmaple"
                    triggerButtonText="Modal"
                >
                    <VStack>
                        <Heading>test1</Heading>
                        <Heading>test2</Heading>
                    </VStack>
                </Modal> */}
            </HStack>
        </Box>
    )
}
export default Home;