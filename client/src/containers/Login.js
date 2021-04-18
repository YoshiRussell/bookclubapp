import React from 'react';
import Button from 'react-bootstrap/Button';
import 'bootstrap/dist/css/bootstrap.min.css';
import '../styles/Login.css';
import { useAuth0 } from "@auth0/auth0-react";

export default function Login() {
  
    const { loginWithRedirect } = useAuth0();

    return (
        <div className="Login">
            <Button variant="dark" block size="lg" onClick={()=>loginWithRedirect()}>
            Login
            </Button>
        </div>
    );
}