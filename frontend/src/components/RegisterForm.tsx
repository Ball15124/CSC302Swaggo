import {Button, Input, Text} from "@chakra-ui/react";

export const RegisterForm = () => {
    return (
            <form style={{
                display: "flex",
                flexDirection: "column",
                gap: "1rem",
                padding: "1rem",
                border: "1px solid black",
                borderRadius: "0.5rem"
            }}>
                <Text textAlign={"center"} fontSize={"2xl"} >Register</Text>
                <Input placeholder={"username"} />
                <Input placeholder={"email"} />
                <Input placeholder={"password"} type={"password"} />
                <Input placeholder={"re-enter password"} type={"password"} />
                <Button>Login</Button>
            </form>
    );
};