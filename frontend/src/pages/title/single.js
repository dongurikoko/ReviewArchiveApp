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
                <h1>{singleContent.title}</h1>
                <h2>{singleContent.before_code}</h2>
                <h2>{singleContent.after_code}</h2>
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