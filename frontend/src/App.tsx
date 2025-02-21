import { BrowserRouter, Routes, Route } from 'react-router-dom';
import './App.css'
import ASTTree from './pages/ASTPage.tsx'
import LoginPage from './pages/LoginPage'
import Header from './components/Header.tsx'
import SearchPage from './pages/SearchPage.tsx'
import UploadPage from './pages/UploadPage.tsx'
import { AuthProvider } from './context/AuthProvider'
function App() {
    return (
        <>
            <AuthProvider>
                <BrowserRouter>
                    <Header />
                    <Routes>
                        <Route path="/" element={<ASTTree />} />
                        <Route path="/search" element={<SearchPage />} />
                        <Route path="/upload" element={<UploadPage />} />
                        <Route path="/login" element={<LoginPage />} />
                    </Routes>
                </BrowserRouter>
            </AuthProvider>
        </>

    )
}

export default App
