import React from 'react';
import Navbar from './components/navbar/Navbar'
import './App.css';
import Routes from './Routes'

function App() {
  return (
    <div className="App">
      <Navbar/>
      <Routes/>
    </div>
  );
}

export default App;
