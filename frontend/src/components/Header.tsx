import { useContext } from 'react'

import './header.css'
import { Link } from 'react-router-dom'
import { logout } from '../services/api'
import { AuthContext } from '../context/AuthContext'

function Header() {
    const { user, setUser } = useContext(AuthContext);

    const handleLogout = async () => {
        const success = await logout();
        if (success) {
            setUser(null);
        } else {
            alert("Logout failed");
        }
    };

    return (
        <>
            <header id="header">
                <ul>
                    <li><Link to="/">作成する</Link></li>
                    <li><Link to="/search">探す</Link></li>
                    <li><Link to="/upload">投稿</Link></li>
                    <li>
                        {user ? (
                            <div>{user.name}</div>
                        ) : (
                            <Link to="/login">ログイン</Link>
                        )}
                    </li>
                    {user && (
                        <li>
                            <button onClick={handleLogout}>ログアウト</button>
                        </li>
                    )}
                </ul>
            </header>
        </>
    )
}

export default Header
