import { useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';

const AddBookByISBNButton = (props) => {

    const [isbnDisplay, setIsbnDisplay] = useState("");
    const { getAccessTokenSilently } = useAuth0();
    
    const handleTextDisplay = (event) => {
        setIsbnDisplay(event.target.value);
    }

    const handlePostingISBN = async () => {
        try {
            const accessToken = await getAccessTokenSilently({ audience: props.audience, scope: props.scope });
            fetch(`http://localhost:8080/mydashboard?isbn=${isbnDisplay}`,
                {
                    method: 'POST',
                    headers: {
                        Authorization: `Bearer ${accessToken}`,
                    },
                });
        } catch(error) {
            console.log(error)
        }
    }

    return (
        <div>
            <input type="text" name="isbn" value={isbnDisplay} onChange={handleTextDisplay}/>
            <button onClick={() => handlePostingISBN()}>
                POST BOOK
            </button>
        </div>
    )
}

export default AddBookByISBNButton;