import React, { useEffect, useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';
import LogoutButton from '../components/LogoutButton';
import useApi from '../hooks/useApi';

export default function Profile() {
    const { isAuthenticated } = useAuth0();
    const [userMetadata, setUserMetadata] = useState(null);
    const { error, loading, data } = useApi(
        'http://localhost:8080/mydashboard',
        {
            audience: 'https://nillbookclub/api',
            scope: 'read:userbooks write:userbooks',
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
                    {userMetadata.username}
                    {userMetadata.pageNumber}
                    <LogoutButton/>
                </div>
            </div>
        )
    );
}
