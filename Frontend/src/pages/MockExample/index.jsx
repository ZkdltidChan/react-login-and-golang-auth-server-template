import React, {
    // 
    useState,
} from "react";
import {
    useColorModeValue,
    Container,
    Flex,
    Heading,
    VStack,
    Button,
    Divider,
    HStack,
    Text,
    Select,
    Input,
    Table,
    Thead,
    Tr,
    Th,
    Tbody,
    Td,
} from "@chakra-ui/react";

import { get_mock_url } from "../../hook/api/mock_api";

const MockExample = () => {
    const [users, setUser] = useState([])
    const [page, setPage] = useState(1)
    const [size, setSize] = useState(5)
    // const { loading } = useAuthState()
    const onClick = async () => {
        setUser(await get_mock_url({
            page:page,
            size:size
        }))
        console.log(page)
        console.log(size)
        console.log(users)
    }
    const pageOnChange = (e) => {
        const value = e.target.value;
        setPage(value);
    }
    const sizeOnChange = (e) => {
        const value = e.target.value;
        setSize(value)
    }

    return (
        <>
            <Container p={3} py={5}>
                <VStack>
                    <Flex
                        shadow="xl"
                        rounded="lg"
                        p={8}
                        maxW="lg"
                        minW={{ base: "100%", md: "md" }}
                        justify="center"
                        bg={useColorModeValue('white', 'gray.700')}
                    >
                        <VStack w="100%" spacing={5}>
                            <Heading textAlign="center">
                                Mock server data
                            </Heading>
                            <HStack w="100%">
                                <Text>
                                    Page Size
                                </Text>
                                <Select placeholder='Table Size'
                                    onChange={sizeOnChange}
                                >
                                    <option value='5'>5</option>
                                    <option value='10'>10</option>
                                    <option value='15'>15</option>
                                </Select>
                            </HStack>
                            <HStack w="100%">
                                <Text>Page Index: </Text>
                                <Input
                                    // value={size}
                                    onChange={pageOnChange}
                                    placeholder='PageIndex'
                                />
                            </HStack>
                            <Button onClick={onClick}>
                                Get Mock Data
                            </Button>
                        </VStack>


                    </Flex>
                    <Flex
                        shadow="xl"
                        rounded="lg"
                        p={8}
                        maxW="lg"
                        minW={{ base: "100%", md: "md" }}
                        justify="center"
                        bg={useColorModeValue('white', 'gray.700')}
                    >
                        <Table>
                            <Thead>
                                <Tr>
                                    {/* <Th>Email</Th> */}
                                    <Th>First Name</Th>
                                    <Th isNumeric>Last Name</Th>
                                    <Th isNumeric>Create At</Th>
                                </Tr>
                            </Thead>
                            <Tbody>
                                {users && users.data ?
                                    users.data.map((user, index) => {
                                        return (
                                            <Tr key={index}>
                                                {/* <Td  w="100%" overflow="scroll">{user.email}</Td> */}
                                                <Td>{user.first_name}</Td>
                                                <Td>{user.last_name}</Td>
                                                <Td>{user.create_at}</Td>
                                            </Tr>
                                        )
                                    }) : ""
                                }
                            </Tbody>
                            
                        </Table>
                    </Flex>
                </VStack>
            </Container >
        </>
    )
}
export default MockExample;