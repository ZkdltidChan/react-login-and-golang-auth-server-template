import axios from "axios";

const MOCK_URL = "https://855caced-2ea4-45bc-bc84-bcd8c0f101e5.mock.pstmn.io"

export async function get_mock_url(params) {
    const requestOptions = {
        url: `${MOCK_URL}/users`,
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
        params,
    }
    try {
        const response = await axios(requestOptions)
        if (response.status === 200) {
            return response.data
        } else {
            // TODO
            return "error"
        }
    } catch (e) {
        console.log(e)
    }
}