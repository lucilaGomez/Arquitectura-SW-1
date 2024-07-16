import React, { useState } from 'react';
import axios from '../../api';
import MainLayout from '../../layouts/MainLayout/MainLayout';

const CreateAmenity = () => {
  const [name, setName] = useState('');
  const [message, setMessage] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    setMessage('');

    try {
      const response = await axios.post('/amenities', { name });
      console.log('Response:', response);
      setMessage('Amenity creada exitosamente');
      setName('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al crear amenity');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    }
  };

  return (
    <MainLayout>
      <div>
        <form onSubmit={handleSubmit}>
          <input
            type="text"
            placeholder="Nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <button type="submit">Crear Amenity</button>
        </form>
        {message && <p>{message}</p>}
      </div>
    </MainLayout>
  );
};

export default CreateAmenity;
