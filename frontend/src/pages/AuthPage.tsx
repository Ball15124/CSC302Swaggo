import {LoginForm} from "../components/LoginForm.tsx";
import {RegisterForm} from "../components/RegisterForm.tsx";
import {Button} from "@chakra-ui/react";
import {useState} from "react";

export const AuthPage = ({setToken}:{setToken:(x:string)=>void}) => {
    const [isLogin, setIsLogin] = useState(false);
    return (
        <>
            {isLogin ? <LoginForm /> : <RegisterForm />}
            <Button onClick={()=>setIsLogin(v=>!v)}>Switch to {isLogin ? "Register" : "Login"}</Button>
        </>
    );
};