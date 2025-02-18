import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './header.css'

function App() {
    return (
        <>
            {
                function () {
                    const list = [];
                    for (let i = 0; i < 10; i++) {
                    list.push(<li>{i}</li>);
                    }
                    return <ul>{list}</ul>;
                }()
            }
        </>
    )
}
  
  export default App

createRoot(document.getElementById('main')!).render(
    <StrictMode>
        <App />
    </StrictMode>,
  )  