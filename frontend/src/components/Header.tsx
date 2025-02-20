import './header.css'
import { Link } from 'react-router-dom'

function Header() {
    return (
        <>
            <header id="header">
                <ul>
                    <li><Link to="/">作成する</Link></li>
                    <li><Link to="/search">探す</Link></li>
                    <li><Link to="/upload">投稿</Link></li>
                    <li><Link to="/login">ログイン</Link></li>
                </ul>
            </header>
        </>
    )
}

export default Header
