const handleLogin = () => {
    window.location.href = "http://localhost:8080/auth/google/login";
};

const LoginPage: React.FC = () => {
    return (
        <div className="login-container">
            <button onClick={handleLogin}>Googleでログイン</button>
        </div>
    );
};

export default LoginPage;