import {Provider} from './components/ui/provider'
import {Box, Theme} from "@chakra-ui/react";
import {AuthPage} from "./pages/AuthPage.tsx";
import {TodoListPage} from "./pages/TodoListPage.tsx";
import {useState} from "react";

function App() {
    const [token, setToken] = useState<string | null>(null);
  return (
    <Provider>
        <Theme appearance="light">
            <Box gap={10} flexDirection={"column"} minH={"100vh"} minW={"100vw"} alignItems={"center"} display={"flex"} justifyContent={"center"} >
                {!token ?
                    <AuthPage setToken={setToken}/> :
                <TodoListPage/>}
            </Box>
        </Theme>
    </Provider>
  )
}

export default App
