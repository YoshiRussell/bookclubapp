import React, { useEffect, useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import LogoutButton from '../components/LogoutButton';
import useApi from '../hooks/useApi';
import AddBookByISBN from '../components/AddBookByISBN';

export default function Dashboard() {
    const audience = 'https://nillbookclub/api';
    const scope = 'read:userbooks write:userbooks';

    const { isAuthenticated } = useAuth0();
    const [userMetadata, setUserMetadata] = useState(null);
    const { error, loading, data } = useApi(
        'http://localhost:8080/mydashboard',
        {
            audience, 
            scope
        }
    );

    useEffect(() => {
        setUserMetadata(data);
    }, [data])

    if (loading) {
        return <div>Loading your profile...</div>
    }

    if (error) {
        return <div>Error loading your profile: {error.message}</div>
    }

    return (
        isAuthenticated && userMetadata && (
            <div>
                <div>
                    <AddBookByISBN audience={audience} scope={scope}/>
                    <LogoutButton/>
                    <h1>{userMetadata.username}</h1>
                    <h1>{userMetadata.pageNumber}</h1>
                    {userMetadata.books.map(book => 
                        <li key={book.Isbn}>{book.Title}</li>
                    )}
                </div>
            </div>
        )
    );
}
