import './App.css';
import { AuthScreen } from "./auth/Auth";
import { ChakraProvider } from '@chakra-ui/react'

function App() {
  return (
      <ChakraProvider>
          <div className="App">
              <header className="App-header">
                  <AuthScreen></AuthScreen>
              </header>
          </div>
      </ChakraProvider>
  );
}

export default App;