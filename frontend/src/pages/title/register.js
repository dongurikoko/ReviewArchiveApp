const Register = () => {

    const handleSubmit = () => {
        try{
            fetch("http://localhost:8080/content/create",{
                method: "POST",
                headers: {
                    "Accept": "application/json",
                    "Content-Type": "application/json",
                },
                body:"ダミー"
            })
        }catch(err){}
    }
    return(
        <div>
            <h1>コンテンツの新規作成</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" name="title" placeholder="タイトル" required/>
                <input type="text" name="before_code" placeholder="修正前コード"/>
                <input type="text" name="after_code" placeholder="修正後コード"/>
                <input type="text" name="review" placeholder="レビュー"/>
                <input type="text" name="memo" placeholder="メモ"/>
                <input type="text" name="keyword" placeholder="キーワード" required/>
                <button type="submit">登録</button>
            </form>
        </div>
    )
}
export default Register;