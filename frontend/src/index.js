import React from 'react';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';
import { createRoot } from 'react-dom/client';
import { theme } from './theme/mantine';
import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import { MantineProvider } from '@mantine/core';
import { UserProvider } from './context/UserContext'; // Importa el UserProvider

const container = document.getElementById('root');
const root = createRoot(container);
root.render(
  <MantineProvider theme={theme} defaultColorScheme="dark">
    <UserProvider> {/* Envuelve tu aplicación con UserProvider */}
      <React.StrictMode>
        <App />
      </React.StrictMode>
    </UserProvider>
  </MantineProvider>
);

reportWebVitals();
