import { BrowserRouter as Router} from 'react-router-dom';
import './App.css'
import Header from './components/Header.tsx'
import { AuthProvider } from './context/AuthProvider'
import AnimApp from './AnimationApp.tsx'

function App() {
    return (
        <>
            <AuthProvider>
                <Router basename="/">
                    <Header />
                    <AnimApp />
                </Router>
            </AuthProvider>
        </>

    )
}

export default App
