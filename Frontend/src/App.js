import React from 'react';
import {
  ChakraProvider,
  theme,
} from '@chakra-ui/react';
// import { ColorModeSwitcher } from './ColorModeSwitcher';
import Router from './router/router';
import { BrowserRouter } from "react-router-dom";
// import { AuthProvider } from './context';
import {AuthProvider} from './hook/auth'
function App() {
  return (
    <ChakraProvider theme={theme}>
      <AuthProvider>
        {/* <ColorModeSwitcher /> */}
        <BrowserRouter>
          <Router />
        </BrowserRouter>
      </AuthProvider>
    </ChakraProvider>
  );
}

export default App;
