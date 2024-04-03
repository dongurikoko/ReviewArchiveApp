import { useParams,useNavigate } from 'react-router-dom';
import { useEffect,useState } from 'react';



const UpdateItem = () => {
    const params = useParams()

    const navigate = useNavigate()

    const [updateContent, setUpdateContent] = useState({
        content_id: "",
        title: "",
        before_code: "",
        after_code: "",
        review: "",
        memo: "",
        keywords: []
    })

    const getUpdateContent = async () => {
        const response = await fetch(`http://localhost:8080/lists/${params.id}`)
        const jsonResponse = await response.json()
        setUpdateContent(jsonResponse)
    }

    useEffect(() => {
        const getUpdateContent = async () => {
            const response = await fetch(`http://localhost:8080/lists/${params.id}`)
            const jsonResponse = await response.json()
            setUpdateContent(jsonResponse)
        }
        getUpdateContent()
    }, [params.id])

    const handleContentChange = (e) => {
        if (e.target.name === 'keyword') {
            // キーワードフィールドの変更時は、キーワードを配列に分割
            setUpdateContent({
                ...updateContent,
                keywords: e.target.value.split(',').map(kw => kw.trim()), // カンマで分割し、トリムする
            });
        } else {
            setUpdateContent({
                ...updateContent,
                [e.target.name]: e.target.value,
            })
        }
    }

        

    const handleSubmit = async(e) => {
        e.preventDefault();
        try{
            const response = await fetch(`http://localhost:8080/contents/${params.id}`,{
                method: "POST",
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(updateContent),
            })
            
            const jsonResponse = await response.json()
            console.log (jsonResponse)
            alert(jsonResponse.message)

            navigate("/")

        }catch(err){
            alert("編集失敗")
        }
    }
    
    return(
        <div>
            <div className = "miniTitle">コンテンツの編集</div>
            <form onSubmit={handleSubmit}>
                <input value={updateContent.title} onChange={handleContentChange} 
                type="text" name="title" placeholder="タイトル(必須)" required/>
                <textarea value={updateContent.before_code} onChange={handleContentChange}
                type="text" name="before_code" placeholder="修正前コード" rows="4"/>
                <textarea value={updateContent.after_code} onChange={handleContentChange}
                type="text" name="after_code" placeholder="修正後コード" rows="4"/>
                <textarea value={updateContent.review} onChange={handleContentChange}
                type="text" name="review" placeholder="レビュー" rows="4"/>
                <textarea value={updateContent.memo} onChange={handleContentChange}
                type="text" name="memo" placeholder="メモ" rows="4"/>
                <input value={updateContent.keywords.join(', ')} onChange={handleContentChange}
                type="text" name="keyword" placeholder="キーワード（必須） 例：文法,配列" required/>

                <button className="register-button">編集</button>
            </form>
        </div>
    )
}

export default UpdateItem;
