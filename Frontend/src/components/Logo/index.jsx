import {
    Flex,
    Image
} from '@chakra-ui/react';

export default function Logo() {
    return (
        <Flex>
            <Image
                src={require("../../assets/logo48.png")}
                fallbackSrc="https://via.placeholder.com/48"
                boxSize="48px"
                alt="logo"
            />
        </Flex>
    );
}