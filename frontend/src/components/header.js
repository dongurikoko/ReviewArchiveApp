import { Link } from "react-router-dom"
import headerImage from '../images/icon2.png';

const Header = ({ searchTerm, setSearchTerm }) => {
    return (
        <header className="header">
            <div><Link to="/"><img src={headerImage} alt="header"/></Link></div>
            <nav className="nav-container">
                <ul className="navList">
                    <li className="navLink"><Link to="/content/new">レビュー新規作成</Link></li>
                    <li><input type="text" value={searchTerm} onChange={event => setSearchTerm(event.target.value)}
                    placeholder="🔍  キーワード検索"/></li>
                </ul>
            </nav>
        </header>
    )
}

export default Header