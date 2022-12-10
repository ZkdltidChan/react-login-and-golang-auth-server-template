import {
    Text,
    Image,
    HStack,
} from '@chakra-ui/react';

export default function Logo() {
    return (
        <HStack align="center" className="logo">
            <Image
                src={require("../../assets/logo48.png")}
                fallbackSrc="https://via.placeholder.com/48"
                boxSize="48px"
                alt="logo"
            />
            <Text>
                Logo
            </Text>
        </HStack>
    );
}