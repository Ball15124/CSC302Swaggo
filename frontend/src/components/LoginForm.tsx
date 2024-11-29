import {Button, Input, Text} from "@chakra-ui/react";
import {useState} from "react";


export const LoginForm =({ login }: { login: (username: string, password: string) => void }) => {
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");

    const handleLogin = () => {
        if (username && password) {
            login(username, password);
        } else {
            alert("Please enter both username and password");
        }
    };

    return (
            <form style={{
                display: "flex",
                flexDirection: "column",
                gap: "1rem",
                padding: "1rem",
                border: "1px solid black",
                borderRadius: "0.5rem"
            }}>
                <Text textAlign={"center"} fontSize={"2xl"} >Login</Text>
                <Input placeholder={"username"} onChange={(e) => setUsername(e.target.value)} />
                <Input placeholder={"password"} type={"password"} onChange={(e) => setPassword(e.target.value)} />
                <Button onClick={handleLogin}>Login</Button>
            </form>
    );
};