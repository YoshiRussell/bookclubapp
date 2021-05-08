import { useEffect, useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';

const useApi = (url, options={}, method="GET", onMountCall=true) => {
    const postBool = method === "POST" ? true : false;
    const { getAccessTokenSilently } = useAuth0();
    const [state, setState] = useState({
        error: null,
        loading: true,
        data: null,
    });

    useEffect(() => {
        if(onMountCall)
            callAPI();
    },[]);
    
    const callAPI = async () => {
        try {
            const { audience, scope, ...fetchOptions } = options;
            const accessToken = await getAccessTokenSilently({ audience, scope });
            const response = await fetch(url, {
                ...fetchOptions,
                method,
                headers: {
                    Authorization: `Bearer ${accessToken}`,
                },
            });

            setState({
                data: postBool ? null : await response.json(),
                error: null,
                loading: false,
            });
        } catch (error) {
            setState({
                data: null,
                error,
                loading: false,
            });
        }
    }

    return {
        ...state,
        callAPI,
    };
};

export default useApi;