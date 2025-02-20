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
            <GoogleOAuthProvider clientId="59621130068-tnl0dbd1qbhj41uf09ippou0klbd5l7l.apps.googleusercontent.com">
                <ProsessToProvider />
            </GoogleOAuthProvider>
        </div>
    );
};

export default LoginPage;