import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './header.css'

function App() {
    return (
        <>
            <header>
                <ul>
                    <li><a href="index.html">作成する</a></li>
                    <li><a href="search.html">探す</a></li>
                    <li><a href="upload.html">投稿</a></li>
                    <li><a href="login.html">ログイン</a></li>
                </ul>
            </header>
        </>
    )
}
  
  export default App

createRoot(document.getElementById('astheader')!).render(
    <StrictMode>
        <App />
    </StrictMode>,
  )  