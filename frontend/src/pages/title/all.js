import {useState,useEffect} from 'react'
import {Link} from 'react-router-dom'
import Header from '../../components/header'
import { useAuthContext } from '../../context/AuthContext'
import { useNavigate } from "react-router-dom";

const All = ({ searchTerm, setSearchTerm }) => {
    const[allContents, setAllContents] = useState()
    const navigate = useNavigate()

    const info = useAuthContext()
    
    // JWTトークンを使用してバックエンドからコンテンツを取得する関数
    const getAllcontents = async () => {
        // localStorageからJWTトークンを取得
        const jwt = localStorage.getItem('jwt');
        const response = await fetch(`http://localhost:8080/lists/search?keyword=${searchTerm}`, {
            method: 'GET', // HTTPメソッド
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${jwt}` // AuthorizationヘッダーにJWTトークンを設定
            },
        });
        const jsonResponse = await response.json();
        setAllContents(jsonResponse);
    }
    
    

    useEffect(() => {
        if (!info.user) {
            navigate("/signin");
            return; // ユーザーがいない場合はここで処理を中断
        }
        getAllcontents()
    }, [info.user, searchTerm, navigate])
    
    return(
        <div>   
            <div className="btn-container">
            {allContents && allContents.contents && allContents.contents.map((content) => 
                <div key={content.content_id}>
                    <Link to ={`/contents/${content.content_id}`} className="btn btn-border-shadow btn-border-shadow--color"><span>
                    <h1 style={{ marginTop: '10px',marginBottom: '10px'}}>{content.title}</h1>
                    <h3>---- キーワード ----</h3>
                    {content.keywords && content.keywords.map((keyword, index) => (
                    <h3 key={index}>{keyword}</h3>
                    ))}
                    </span>
                    </Link>
                </div>
            )}
            </div>
        </div>
    )
}

export default All;
