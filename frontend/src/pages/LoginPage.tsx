import React, { useState } from 'react';
import { GoogleOAuthProvider, useGoogleLogin } from '@react-oauth/google';

function ProsessToProvider() {
    const login = useGoogleLogin({
        onSuccess: (codeResponse) => console.log(codeResponse),
        flow: "auth-code",
        scope: "email profile openid",
    });
  
    return (
      <button onClick={ login }>Googleでログイン</button>
    );
  };

const LoginPage: React.FC = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        console.log('Login attempt:', email, password);
    };

    return (
        <div className="login-container">
            <h2>ログイン</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label htmlFor="email">メールアドレス:</label>
                    <input
                        type="email"
                        id="email"
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                        required
                    />
                </div>
                <div className="form-group">
                    <label htmlFor="password">パスワード:</label>
                    <input
                        type="password"
                        id="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">ログイン</button>
            </form>
            <GoogleOAuthProvider clientId="775189385683-85h2mrb9ualv0l86f6v4rk4ct8qdk451.apps.googleusercontent.com">
                <ProsessToProvider />
            </GoogleOAuthProvider>
        </div>
    );
};

export default LoginPage;