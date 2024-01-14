import {useState,useEffect} from 'react'
import {Link} from 'react-router-dom'

const All = () => {
    const[allContents, setAllContents] = useState()
    const getAllcontents = async() => {
        const response = await fetch('http://localhost:8080/list/get')
        const jsonResponse = await response.json()
        setAllContents(jsonResponse)
    }

    useEffect(() => {
        const getAllcontents = async() => {
            const response = await fetch('http://localhost:8080/list/get')
            const jsonResponse = await response.json()
            setAllContents(jsonResponse)
        }
        getAllcontents()
    },[])

    return(
        <div>
           <div>
            {allContents && allContents.contents.map((content) => 
                <div key={content.content_id}>
                    <Link to ={`/contents/${content.content_id}`}>
                    <h1>{content.title}</h1>
                    </Link>
                    {content.keywords && content.keywords.map((keyword, index) => (
                    <h2 key={index}>{keyword}</h2>
                    ))}
                    </div>
            )}
            </div>
        </div>
    )
}
export default All;