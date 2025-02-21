import { useEffect, useState } from 'react'

import './header.css'
import { Link } from 'react-router-dom'
import { getCurrentUser, UserInfo } from '../services/api'

function Header() {
    const [user, setUser] = useState<UserInfo | null>(null);

    useEffect(() => {
        async function fetchUser() {
            const currentUser = await getCurrentUser();
            setUser(currentUser);
        }
        fetchUser();
    }, []);
    
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
                </ul>
            </header>
        </>
    )
}

export default Header
