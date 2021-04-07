import React from 'react';
import ReactDOM from 'react-dom';
import App from '../App'
import { Auth0Provider } from '@auth0/auth0-react';

ReactDOM.render(
  <Auth0Provider
    domain="cnguy.us.auth0.com"
    clientId="OPZkgMVjFR5u8apnDuyRLT0BjfAZWfId"
    redirectUri={window.location.origin}
  >
    <App />
  </Auth0Provider>,
  document.getElementById("root")
);