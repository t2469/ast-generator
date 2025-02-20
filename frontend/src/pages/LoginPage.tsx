import { GoogleOAuthProvider, useGoogleLogin } from '@react-oauth/google';

function ProsessToProvider() {
    const login = useGoogleLogin({
        onSuccess: (codeResponse) => console.log(codeResponse),
        flow: "auth-code",
        scope: "email profile openid",
    });

    return (
        <button onClick={login}>Googleでログイン</button>
    );
};

const LoginPage: React.FC = () => {
    return (
        <div className="login-container">
            <GoogleOAuthProvider clientId="775189385683-85h2mrb9ualv0l86f6v4rk4ct8qdk451.apps.googleusercontent.com">
                <ProsessToProvider />
            </GoogleOAuthProvider>
        </div>
    );
};

export default LoginPage;