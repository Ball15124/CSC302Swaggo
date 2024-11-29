import {Button, Input, Text} from "@chakra-ui/react";

export const LoginForm = () => {
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
                <Input placeholder={"username"} />
                <Input placeholder={"password"} type={"password"} />
                <Button>Login</Button>
            </form>
    );
};