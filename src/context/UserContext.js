import React, { createContext, useContext, useState, useEffect } from 'react';
import axios from '../api';

const UserContext = createContext();

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    // Obtener el usuario actual
    axios.get('/auth/me')
      .then(response => {
        setUser(response.data.user);
      })
      .catch(error => {
        console.error('Failed to fetch user:', error);
      });
  }, []);

  const login = async (email, password) => {
    try {
      const response = await axios.post('/login', { email, password });
      setUser(response.data.user);
    } catch (error) {
      console.error('Failed to login:', error);
      throw error;
    }
  };

  const logout = async () => {
    try {
      await axios.post('/auth/logout');
      setUser(null);
    } catch (error) {
      console.error('Failed to logout:', error);
    }
  };

  return (
    <UserContext.Provider value={{ user, login, logout }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUserContext = () => {
  return useContext(UserContext);
};
