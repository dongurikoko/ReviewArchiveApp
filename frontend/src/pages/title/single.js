import { useParams } from 'react-router-dom';
import { useEffect,useState } from 'react';
import {Link} from 'react-router-dom';



const Single = () => {
    const params = useParams()

    const [singleContent, setSingleContent] = useState({
        content_id: "",
        title: "",
        before_code: "",
        after_code: "",
        review: "",
        memo: "",
        keywords: []
    })

    const getSingleContent = async () => {
        const response = await fetch(`http://localhost:8080/lists/${params.id}`)
        const jsonResponse = await response.json()
        setSingleContent(jsonResponse)
    }

    useEffect(() => {
        const getSingleContent = async () => {
            const response = await fetch(`http://localhost:8080/lists/${params.id}`)
            const jsonResponse = await response.json()
            setSingleContent(jsonResponse)
        }
        getSingleContent()
    }, [params.id])

    
    return(
        <div> 
            {singleContent && ( 
                <>   
                <div className="miniTitle">{singleContent.title}</div>
                {singleContent && singleContent.before_code ? (
                <div className="code-box">
                    <span className="code-box-title1">修正前コード</span>
                    <pre><code className="code-font">{singleContent.before_code.split('  ').join('\n')}</code></pre>
                </div>
                ) : null}
                {singleContent && singleContent.after_code ? (
                    <div className="code-box">
                    <span className="code-box-title2">修正後コード</span>
                    <pre><code className="code-font">{singleContent.after_code.split('  ').join('\n')}</code></pre>
                </div>
                ) : null}
                {singleContent && singleContent.review ? (
                    <div class="review-box">
                    <div class="review-box-title">REVIEW</div>
                   {singleContent.review}</div>
                ) : null}
                {singleContent && singleContent.memo ? (
                    <div className="memo-box">
                    ---- MEMO ---- <br/><br/>
                    {singleContent.memo}
                    </div>
                ) : null}
                <div className="keyword-container">
                キーワード: 
                {singleContent.keywords && singleContent.keywords.map((keyword, index) => (
                    <h2 key={index} className="keyword">#{keyword}</h2>
                ))}
                </div>
                <div className='button-container12'>
                    <Link to={`/content/update/${singleContent.content_id}`} className="content-button1">編集</Link>
                    <Link to={`/content/delete/${singleContent.content_id}`} className="content-button2">削除</Link>
                </div>
                </>
            )}
        </div>
    )
}

export default Single;
