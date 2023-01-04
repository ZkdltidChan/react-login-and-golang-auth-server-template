// TODO: api提到api層定義
import axios from "axios";
import { BASE_URL, LOGIN_ERROR, LOGIN_SUCCESS, LOGOUT, REQUEST_LOGIN } from "../../constants/actions";


export const loginUser = async (dispatch, loginPayload) => {
    const requestOptions = {
        url: `${BASE_URL}/signin`,
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        data: loginPayload
    }
    try {
        dispatch({ type: REQUEST_LOGIN, payload: null })
        const response = await axios(requestOptions)
        if (response.status === 200) {
            dispatch({ type: LOGIN_SUCCESS, payload: response.data })
            localStorage.setItem('currentUser', JSON.stringify(response.data))
            return response.data
        } else {
            dispatch({ type: LOGIN_ERROR, error: response.data.error[0] })
        }
        return;
    } catch (e) {
        dispatch({ type: LOGIN_ERROR, error: e })
    }
}

export async function loginUserWithDiscord(dispatch, params) {
    const requestOptions = {
        url: `${BASE_URL}/oauth/discord/callback`,
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        },
        params,
    }
    try {
        dispatch({ type: REQUEST_LOGIN, payload: null })
        const response = await axios(requestOptions)
        if (response.status === 200) {
            dispatch({ type: LOGIN_SUCCESS, payload: response.data })
            localStorage.setItem('currentUser', JSON.stringify(response.data))
            return response.data
        } else {
            dispatch({ type: LOGIN_ERROR, error: response.data.error[0] })
        }
        return;
    } catch (e) {
        dispatch({ type: LOGIN_ERROR, error: e })
    }
}


export async function logout(dispatch) {
    dispatch({ type: LOGOUT });
    localStorage.removeItem('currentUser');
    localStorage.removeItem('token');
}