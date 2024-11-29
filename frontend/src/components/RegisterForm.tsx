import { Button, Input, Text } from "@chakra-ui/react";
import { useState } from "react";

export const RegisterForm = ({register,}: { register: (username: string, email: string, password: string) => void;
}) => {
    const [username, setUsername] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [rePassword, setRePassword] = useState("");

    const handleRegister = () => {
        if (!username || !email || !password || !rePassword) {
            alert("All fields are required.");
            return;
        }

        if (password !== rePassword) {
            alert("Passwords do not match.");
            return;
        }

        // Call the passed `register` function
        register(username, email, password);
    };

    return (
        <form
            style={{
                display: "flex",
                flexDirection: "column",
                gap: "1rem",
                padding: "1rem",
                border: "1px solid black",
                borderRadius: "0.5rem",
            }}
        >
            <Text textAlign={"center"} fontSize={"2xl"}>
                Register
            </Text>
            <Input
                placeholder={"Username"}
                value={username}
                onChange={(e) => setUsername(e.target.value)}
            />
            <Input
                placeholder={"Email"}
                type="email"
                value={email}
                onChange={(e) => setEmail(e.target.value)}
            />
            <Input
                placeholder={"Password"}
                type={"password"}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
            />
            <Input
                placeholder={"Re-enter Password"}
                type={"password"}
                value={rePassword}
                onChange={(e) => setRePassword(e.target.value)}
            />
            <Button onClick={handleRegister}>Register</Button>
        </form>
    );
};