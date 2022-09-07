import React, {
    // 
} from "react";
import {
    Heading,
    Text,
} from "@chakra-ui/react";
import { useParams } from "react-router-dom";
const EmbedPageWithID = () => {
    let { id } = useParams();
    return (
        <>
            <Heading>
                Embed Page With ID
            </Heading>
            <Text>
                ID = {id}
            </Text>
        </>
    )
}
export default EmbedPageWithID;