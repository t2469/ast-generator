import { useContext,useState,useRef, useEffect } from "react";
import { Link } from "react-router-dom";
import { logout } from "../services/api";
import { AuthContext } from "../context/AuthContext";
import "../index.css";
import  LogoutModal  from "./LogoutModal";

function Header() {
    const { user, setUser } = useContext(AuthContext);
    const [showModal, setShowModal] = useState(false);
    const iconButtonRef = useRef<HTMLButtonElement>(null);

    const ShowModal = () => {
        setShowModal(true);
    };

    const handleLogout = async () => {
        const success = await logout();
        if (success) {
            setUser(null);
        } else {
            alert("Logout failed");
        }
    };

    return (
        <header className="fixed top-0 left-0 z-50 w-full bg-white shadow">
            <div className="container mx-auto px-4 py-3 flex items-center justify-between">
                <div className="text-xl font-bold text-gray-800">
                    <Link to="/">AST-Generator</Link>
                </div>

                <nav className="flex space-x-6 items-center">
                    <Link to="/" className="text-gray-700 hover:text-blue-500 transition-colors">
                        作成する
                    </Link>
                    {user && (
                        <Link to="/source_codes" className="text-gray-700 hover:text-blue-500 transition-colors">
                            保存されたプログラム一覧
                        </Link>
                    )}
                    <Link to="/upload" className="text-gray-700 hover:text-blue-500 transition-colors">
                        投稿
                    </Link>
                    {user ? (
                        <>
                            <button
                                 onClick={ShowModal}
                                 className="relative m-0 cursor-pointer"
                                 ref={iconButtonRef}
                            >
                            <img
                                src={user.picture}
                                alt={'${user.name}'}
                                className="w-8 h-8 rounded-full hover:brightness-75 transition duration-200"
                            />
                            </button>
                            <LogoutModal 
                                showFlag={showModal}
                                setShowModal={setShowModal}
                                handleLogout={handleLogout}
                                targetRef={iconButtonRef}
                            />
                        </>
                    ) : (
                        <Link to="/login" className="text-gray-700 hover:text-blue-500 transition-colors">
                            ログイン
                        </Link>
                    )}
                </nav>
            </div>
        </header>
    );
}
/*<button
    onClick={handleLogout}
    className="text-gray-700 hover:text-red-500 transition-colors"
>
    ログアウト
</button>*/
export default Header;
