import React from 'react';
import { Route, Switch } from 'react-router-dom';
import Home from './containers/Home';
import Login from './containers/Login';
import Profile from './containers/Profile';

export default function Routes() {
  return (
    <Switch>
      <Route exact path="/">
        <Home />
      </Route>
      <Route path="/login">
          <Login />
      </Route>
      <Route path="/profile">
          <Profile />
      </Route>
    </Switch>
  );
}