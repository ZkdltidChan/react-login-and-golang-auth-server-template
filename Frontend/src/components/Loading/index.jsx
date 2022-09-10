import { Spinner, Box } from '@chakra-ui/react'


function Loading({ ...rest }) {
    return (
        <Box
            {...rest}
        >
            <Spinner
                thickness='4px'
                speed='0.65s'
                emptyColor='gray.200'
                color='blue.500'
                size='xl'
            />
        </Box>
    )
}

export default Loading;