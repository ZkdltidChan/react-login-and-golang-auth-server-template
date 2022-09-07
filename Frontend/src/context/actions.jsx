// TODO: api提到api層定義
import axios from "axios";
const BASE_URL = 'http://localhost:5002/v1';

export const loginUser = async (dispatch, loginPayload) => {
    const requestOptions = {
        url: `${BASE_URL}/login`,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: loginPayload
    }
    try {
        dispatch({ type: 'REQUEST_LOGIN', payload: null })
        const response = await axios(requestOptions)
        if (response.status === 200) {
            dispatch({ type: 'LOGIN_SUCCESS', payload: response.data })
            localStorage.setItem('currentUser', JSON.stringify(response.data))
            return response.data
        } else {
            dispatch({ type: 'LOGIN_ERROR', error: response.data.error[0] })
        }
        return;
    } catch (e) {
        dispatch({ type: 'LOGIN_ERROR', error: e })
    }
}


export async function logout(dispatch) {
    dispatch({ type: 'LOGOUT' });
    localStorage.removeItem('currentUser');
    localStorage.removeItem('token');
}