import { useEffect, useState } from 'react';
import { useAuth0 } from '@auth0/auth0-react';

const useApi = (url, options ={}) => {
    const { getAccessTokenSilently } = useAuth0();
    const [state, setState] = useState({
        error: null,
        loading: true,
        data: null,
    });

    useEffect(() => {
        (async () => {
            try {
                const { audience, scope, ...fetchOptions } = options;
                const accessToken = await getAccessTokenSilently({ audience, scope });
                const response = await fetch(url, {
                    ...fetchOptions,
                    headers: {
                        Authorization: `Bearer ${accessToken}`,
                    },
                });

                setState({
                    ...state,
                    data: await response.json(),
                    error: null,
                    loading: false,
                });
            } catch (error) {
                setState({
                    ...state,
                    error,
                    loading: false,
                });
            }
        })();
    },[url]);

    return {
        ...state,
    };
};

export default useApi;