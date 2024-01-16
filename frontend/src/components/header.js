import { Link } from "react-router-dom"
import headerImage from '../images/headerImage2.png';

const Header = () => {
    return (
        <header className="header">
            <div><Link to="/"><img src={headerImage} alt="header"/></Link></div>
            <nav>
                <ul className="navList">
                    <li className="navLink"><Link to="/content/new">レビュー新規作成</Link></li>
                </ul>
            </nav>
        </header>
    )
}

export default Header