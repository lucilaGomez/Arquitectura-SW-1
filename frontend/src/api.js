import axios from 'axios';

const instance = axios.create({
  baseURL: 'http://localhost:3000', // Asegúrate de que esto sea correcto
  withCredentials: true // Asegúrate de que las cookies se envíen con cada solicitud
});

export default instance;
