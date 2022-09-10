import React, { useContext, createContext, useReducer } from "react";
import { AuthReducer, initialState } from './reducer'

const AuthStateContext = createContext(null)
const AuthDispatchContext = createContext(null)

export function useAuthState() {
    const context = useContext(AuthStateContext);
    if (context === undefined) {
        throw new Error("useAuthState is only available within AuthProvider.")
    }

    return context;
}

export function useAuthDispatch() {
    const context = useContext(AuthDispatchContext);
    if (context === undefined) {
        throw new Error("useAuthDispatch is only available within AuthProvider.")
    }
    return context;
}

export const AuthProvider = ({ children }) => {
    
    const [user, dispatch] = useReducer(AuthReducer, initialState)

    return (
        <AuthStateContext.Provider value={user}>
            <AuthDispatchContext.Provider value={dispatch}>
                {children}
            </AuthDispatchContext.Provider>
        </AuthStateContext.Provider>
    )
}
