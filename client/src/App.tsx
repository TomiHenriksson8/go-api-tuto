import { useState, useEffect } from 'react';
import { Route, Routes, useNavigate } from 'react-router-dom';
import Login from './components/Login';
import Navbar from './components/Navbar';
import Register from './components/Register';
import TodoPage from './components/TodoPage';
import Welcome from './components/Welcome';
import './index.css';

function App() {

  return (
    <div className='h-screen'>
      <Navbar />
      <Routes>
        <Route path="/" element={<TodoPage />} />
        <Route path="/register" element={<Register />} />
        <Route path="/login" element={<Login />} />
        <Route path="/welcome" element={<Welcome />} />
      </Routes>
    </div>
  );
}

export default App;
