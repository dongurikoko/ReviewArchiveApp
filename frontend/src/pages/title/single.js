import { useParams } from 'react-router-dom';
import { useEffect,useState } from 'react';



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
        const response = await fetch(`http://localhost:8080/list/get/${params.id}`)
        const jsonResponse = await response.json()
        setSingleContent(jsonResponse)
    }

    useEffect(() => {
        const getSingleContent = async () => {
            const response = await fetch(`http://localhost:8080/list/get/${params.id}`)
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
                <div class="code-box">
                    <span class="code-box-title">修正前コード:</span>
                    <pre><code className="code-font">{singleContent.before_code.split('  ').join('\n')}</code></pre>
                </div>
                ) : null}
                <pre><code className="code-font">{singleContent.after_code.split('  ').join('\n')}</code></pre>
                <h2>{singleContent.review}</h2>
                <h2>{singleContent.memo}</h2>
                {singleContent.keywords && singleContent.keywords.map((keyword, index) => (
                    <h2 key={index}>{keyword}</h2>
                ))}
                </>
            )}
        </div>
    )
}

export default Single;