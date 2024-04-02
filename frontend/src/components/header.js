import React from 'react';
import { Link } from "react-router-dom"
import headerImage from '../images/icon2.png';
import { useNavigate } from 'react-router-dom';
import { auth } from '../main';

const Header = ({ searchTerm, setSearchTerm }) => {
    const navigate = useNavigate();
    const handleLogout = async () => {
        try {
            await auth.signOut();
            localStorage.removeItem('jwt');
            navigate("/signin");
        } catch (err) {
            alert(err.message);
            console.error(err);
        }
    };
    return (
        <header className="header">
            <div><Link to="/"><img src={headerImage} alt="header"/></Link></div>
            <nav className="nav-container">
                <ul className="navList">
                    <li className="navLink"><Link to="/content/new">レビュー新規作成</Link></li>
                    <li><input type="text" value={searchTerm} onChange={event => setSearchTerm(event.target.value)}
                    placeholder="🔍  キーワード検索"/></li>
                    <button onClick={handleLogout}>Logout</button>
                </ul>
            </nav>
        </header>
    )
}

export default Header
