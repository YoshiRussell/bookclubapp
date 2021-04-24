import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router } from 'react-router-dom';
import './styles/index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { Auth0Provider } from '@auth0/auth0-react';

ReactDOM.render(
  <React.StrictMode>
    <Router>
        <Auth0Provider
            domain="dev-35574pmo.us.auth0.com"
            clientId="D48g8ocuBvbFz3xRhsqVnbdB8pvr2LEV"
            redirectUri={window.location.origin + '/dashboard'}
            audience="https://nillbookclub/api"
            scope="read:userbooks write:userbooks">
            <App />
        </Auth0Provider>
    </Router>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
