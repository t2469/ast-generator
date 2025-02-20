import { GoogleOAuthProvider, useGoogleLogin } from '@react-oauth/google';
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './login.css'


// Googleログイン処理
function ProsessToProvider() {
  const login = useGoogleLogin({
    onSuccess: (codeResponse) => console.log(codeResponse),
  });

  return (
    <button onClick={ login }>Google Login</button>
  );
};

export default function LoginUserForm() {

    return(
        <div className="googlefield">
            <GoogleOAuthProvider clientId="775189385683-85h2mrb9ualv0l86f6v4rk4ct8qdk451.apps.googleusercontent.com">
                <ProsessToProvider />
            </GoogleOAuthProvider>
        </div>
    );
};

createRoot(document.getElementById('main')!).render(
    <StrictMode>
        <LoginUserForm />
    </StrictMode>,
  )  