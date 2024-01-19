import { useParams,useNavigate } from 'react-router-dom';
import { useEffect,useState } from 'react';



const Delete = () => {
    const params = useParams()

    const navigate = useNavigate()

    const [deleteContent, setDeleteContent] = useState({
        content_id: "",
        title: "",
        before_code: "",
        after_code: "",
        review: "",
        memo: "",
        keywords: []
    })

    const getDeleteContent = async () => {
        const response = await fetch(`http://localhost:8080/list/get/${params.id}`)
        const jsonResponse = await response.json()
        setDeleteContent(jsonResponse)
    }

    useEffect(() => {
        const getDeleteContent = async () => {
            const response = await fetch(`http://localhost:8080/list/get/${params.id}`)
            const jsonResponse = await response.json()
            setDeleteContent(jsonResponse)
        }
        getDeleteContent()
    }, [params.id])

    const handleContentChange = (e) => {
        if (e.target.name === 'keyword') {
            // キーワードフィールドの変更時は、キーワードを配列に分割
            setDeleteContent({
                ...deleteContent,
                keywords: e.target.value.split(',').map(kw => kw.trim()), // カンマで分割し、トリムする
            });
        } else {
            setDeleteContent({
                ...deleteContent,
                [e.target.name]: e.target.value,
            })
        }
    }

        

    const handleSubmit = async(e) => {
        e.preventDefault();
        if(window.confirm("本当に削除しますか？")){
            try{
            const response = await fetch(`http://localhost:8080/content/delete/${params.id}`,{
                method: "DELETE",
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json",
                },
            })
            
            const jsonResponse = await response.json()
            console.log (jsonResponse)
            alert(jsonResponse.message)

            navigate("/")

            }catch(err){
                alert("削除失敗")
            }
        }
    }
    
    return(
        <div> 
            <div className = "miniTitle">コンテンツの削除</div>
            <form onSubmit={handleSubmit}>
            {deleteContent && ( 
                <>   
                <div className="miniTitle">{deleteContent.title}</div>
                {deleteContent && deleteContent.before_code ? (
                <div className="code-box">
                    <span className="code-box-title1">修正前コード</span>
                    <pre><code className="code-font">{deleteContent.before_code.split('  ').join('\n')}</code></pre>
                </div>
                ) : null}
                {deleteContent && deleteContent.after_code ? (
                    <div className="code-box">
                    <span className="code-box-title2">修正後コード</span>
                    <pre><code className="code-font">{deleteContent.after_code.split('  ').join('\n')}</code></pre>
                </div>
                ) : null}
                {deleteContent && deleteContent.review ? (
                    <div class="review-box">
                    <div class="review-box-title">REVIEW</div>
                   {deleteContent.review}</div>
                ) : null}
                {deleteContent && deleteContent.memo ? (
                    <div className="memo-box">
                    ---- MEMO ---- <br/><br/>
                    {deleteContent.memo}
                    </div>
                ) : null}
                <div className="keyword-container">
                キーワード: 
                {deleteContent.keywords && deleteContent.keywords.map((keyword, index) => (
                    <h2 key={index} className="keyword">#{keyword}</h2>
                ))}
                </div>
                </>
            )}
            <button class="delete-button">削除</button>
            </form>
        </div>
    )
}

export default Delete;