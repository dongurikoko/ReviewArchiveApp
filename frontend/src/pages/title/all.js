import {useState,useEffect} from 'react'
import {Link} from 'react-router-dom'


const All = () => {
    const[allContents, setAllContents] = useState()
    const [searchTerm, setSearchTerm] = useState('')

    const getAllcontents = async() => {
        const response = await fetch(`http://localhost:8080/list/search?keyword=${searchTerm}`)
        const jsonResponse = await response.json()
        setAllContents(jsonResponse)
    }

    useEffect(() => {
        const getAllcontents = async() => {
            const response = await fetch(`http://localhost:8080/list/search?keyword=${searchTerm}`)
            const jsonResponse = await response.json()
            setAllContents(jsonResponse)
        }
        getAllcontents()
    },[searchTerm])

    return(
        <div>
            <input
                type="text"
                value={searchTerm}
                onChange={event => setSearchTerm(event.target.value)}
            />
            
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