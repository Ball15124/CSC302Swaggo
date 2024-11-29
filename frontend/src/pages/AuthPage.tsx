import { LoginForm } from "../components/LoginForm.tsx";
import { RegisterForm } from "../components/RegisterForm.tsx";
import { Button } from "@chakra-ui/react";
import { useState } from "react";
import { apiService } from "../api/apiService.ts";

export const AuthPage = ({ setToken }: { setToken: (x: string) => void }) => {
    const [isLogin, setIsLogin] = useState(true); // Default to Login view

    const login = (username: string, password: string) => {
        apiService.authLoginPost({ username, password }).then((res) => {
            if (res?.data?.token) {
                setToken(res.data.token);
            } else {
                console.error("Login failed: No token received");
            }
        });
    };

    const register = (username: string, email: string, password: string) => {
        apiService.authRegisterPost({ username, email, password }).then((res) => {
            if (res?.data?.token) {
                setToken(res.data.token);
            } else {
                console.error("Login failed: No token received");
            }
        });
    }

    return (
        <>
            {isLogin ? (
                <LoginForm login={login} />
            ) : (
                <RegisterForm register={register} />
            )}
            <Button onClick={() => setIsLogin((v) => !v)}>
                Switch to {isLogin ? "Register" : "Login"}
            </Button>
        </>
    );
};
