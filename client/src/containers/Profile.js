import React from 'react';
import { useAuth0 } from '@auth0/auth0-react';


export default function Profile() {
    const { user, isAuthenticated, isLoading } = useAuth0();

    if(isLoading) {
        return <div>Loading your profile...</div>
    }

    return (
        isAuthenticated && (
            <div>
                {user.name}
                {user.email}
            </div>
        )
    );
}