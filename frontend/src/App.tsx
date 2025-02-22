import { BrowserRouter as Router } from 'react-router-dom';
import './App.css';
import Header from './components/Header.tsx';
import { AuthProvider } from './context/AuthProvider';
import AnimApp from './AnimationApp.tsx';

function App() {
    return (
        <>
            <AuthProvider>
                <Router basename="/">
                    <Header />
                    <div className="pt-16">
                        <AnimApp />
                    </div>
                </Router>
            </AuthProvider>
        </>
    );
}

export default App;
