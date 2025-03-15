const BASE_URL = import.meta.env.VITE_API_URL

const handleLogin = () => {
    window.location.href = `${BASE_URL}/auth/google/login`;
};

const LoginPage: React.FC = () => {
    return (
        <div className="login-container">
            <button onClick={handleLogin}>Googleでログイン</button>
        </div>
    );
};

export default LoginPage;