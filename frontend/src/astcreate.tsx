import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'

function App() {
    return (
        <>
            <textarea name="codeinput" placeholder="コードを入力してください"/>
            <div id="arrow"></div>
            <textarea id="astreturn" name="astreturn" disabled/>
        </>
    )
}
  
  export default App

createRoot(document.getElementById('astfield')!).render(
    <StrictMode>
        <App />
    </StrictMode>,
  )  