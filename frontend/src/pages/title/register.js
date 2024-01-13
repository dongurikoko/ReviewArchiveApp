import {useState} from "react";

const Register = () => {

    const [newContent, setNewContent] = useState({
        title: "",
        before_code: "",
        after_code: "",
        review: "",
        memo: "",
        keywords: [] //配列として初期化,
    })

    const handleContentChange = (e) => {
        if (e.target.name === 'keyword') {
            // キーワードフィールドの変更時は、キーワードを配列に分割
            setNewContent({
                ...newContent,
                keywords: e.target.value.split(',').map(kw => kw.trim()), // カンマで分割し、トリムする
            });
        } else {
            setNewContent({
                ...newContent,
                [e.target.name]: e.target.value,
            })
        }
    }

        

    const handleSubmit = async(e) => {
        e.preventDefault();
        try{
            const response = await fetch("http://localhost:8080/content/create",{
                method: "POST",
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(newContent),
            })
            
            const jsonResponse = await response.json()
            console.log (jsonResponse)
            alert(jsonResponse.message)

        }catch(err){
            alert("ユーザー登録失敗")
        }
    }
    return(
        <div>
            <h1>コンテンツの新規作成</h1>
            <form onSubmit={handleSubmit}>
                <input value={newContent.title} onChange={handleContentChange} 
                type="text" name="title" placeholder="タイトル" required/>
                <input value={newContent.before_code} onChange={handleContentChange}
                type="text" name="before_code" placeholder="修正前コード"/>
                <input value={newContent.after_code} onChange={handleContentChange}
                type="text" name="after_code" placeholder="修正後コード"/>
                <input value={newContent.review} onChange={handleContentChange}
                type="text" name="review" placeholder="レビュー"/>
                <input value={newContent.memo} onChange={handleContentChange}
                type="text" name="memo" placeholder="メモ"/>
                <input value={newContent.keywords.join(', ')} onChange={handleContentChange}
                type="text" name="keyword" placeholder="キーワード（カンマ区切り）" required/>

                <button>登録</button>
            </form>
        </div>
    )
}
export default Register;