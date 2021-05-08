import { useState, useEffect } from 'react';
import useApi from '../hooks/useApi';

const AddBookByISBNButton = (props) => {
    
    const [isbnDisplay, setIsbnDisplay] = useState("");
    const { error, callAPI } = useApi(
        `http://localhost:8080/mydashboard?isbn=${isbnDisplay}`, 
        { audience: props.audience, scope: props.scope }, 
        "POST", false
    );

    const handleTextDisplay = (event) => {
        setIsbnDisplay(event.target.value);
    }

    useEffect(() => {
        console.log("error calling api: ", error);
    }, [error]);

    return (
        <div>
            <input type="text" name="isbn" value={isbnDisplay} onChange={handleTextDisplay}/>
            <button onClick={() => callAPI()}>
                POST BOOK
            </button>
        </div>
    )
}

export default AddBookByISBNButton;