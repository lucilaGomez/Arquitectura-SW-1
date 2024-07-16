import React, { useState } from 'react';
import axios from '../../api';
import MainLayout from '../../layouts/MainLayout/MainLayout';
import { Button, TextInput } from '@mantine/core';
import './CreateAmenity.css';

const CreateAmenity = () => {
  const [name, setName] = useState('');
  const [idToUpdate, setIdToUpdate] = useState('');
  const [idToDelete, setIdToDelete] = useState('');
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmitCreate = async (event) => {
    event.preventDefault();
    setMessage('');

    try {
      const response = await axios.post('/amenities', { name });
      console.log('Response:', response);
      setMessage('Amenity creado exitosamente');
      setLoading(true);
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
    } finally {
      setLoading(false);
    }
  };

  const handleSubmitUpdate = async (event) => {
    event.preventDefault();
    setMessage('');

    try {
      const response = await axios.put(`/amenities/${idToUpdate}`, { name });
      console.log('Response:', response);
      setMessage('Amenity actualizado exitosamente');
      setLoading(true);
      setName('');
      setIdToUpdate('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al actualizar amenity');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    } finally {
      setLoading(false);
    }
  };

  const handleSubmitDelete = async (event) => {
    event.preventDefault();
    setMessage('');

    try {
      const response = await axios.delete(`/amenities/${idToDelete}`);
      console.log('Response:', response);
      setMessage('Amenity eliminado exitosamente');
      setLoading(true);
      setIdToDelete('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al eliminar amenity');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <MainLayout>
      <h2 className='createAmenity-title'>Gestión de Amenities</h2>
      <form onSubmit={handleSubmitCreate}>
        <div className='createAmenity-wrapper'>
          <TextInput
            type="text"
            placeholder="Nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <Button loading={loading} type="submit">Crear Amenity</Button>
        </div>
      </form>

      <form onSubmit={handleSubmitUpdate}>
        <div className='createAmenity-wrapper'>
          <TextInput
            type="text"
            placeholder="ID del Amenity a actualizar"
            value={idToUpdate}
            onChange={(e) => setIdToUpdate(e.target.value)}
            required
          />
          <TextInput
            type="text"
            placeholder="Nuevo Nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <Button loading={loading} type="submit">Actualizar Amenity</Button>
        </div>
      </form>

      <form onSubmit={handleSubmitDelete}>
        <div className='createAmenity-wrapper'>
          <TextInput
            type="text"
            placeholder="ID del Amenity a eliminar"
            value={idToDelete}
            onChange={(e) => setIdToDelete(e.target.value)}
            required
          />
          <Button loading={loading} type="submit" color="red">Eliminar Amenity</Button>
        </div>
      </form>

      {message && <p>{message}</p>}
    </MainLayout>
  );
};

export default CreateAmenity;




/*
import React, { useState } from 'react';
import axios from '../../api';
import MainLayout from '../../layouts/MainLayout/MainLayout';
import { Button, TextInput } from '@mantine/core';
import "./CreateAmenity.css";

const CreateAmenity = () => {
  const [name, setName] = useState('');
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setMessage('');

    try {
      const response = await axios.post('/amenities', { name });
      console.log('Response:', response);
      setMessage('Amenity creada exitosamente');
      setLoading(true);
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
    } finally{
      setLoading(false);
    }
  };

  return (
    <MainLayout>
      <h2 className='createAmenity-title'>Crea los Amenities de tu hotel</h2>
        <form onSubmit={handleSubmit}>
        <div className='createAmenitie-wrapper'>
          <TextInput
            type="text"
            placeholder="Nombre"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
          <Button loading={loading} type="submit">Crear Amenity</Button>
          </div>
        </form>
        {message && <p>{message}</p>}
      
    </MainLayout>
  );
};

export default CreateAmenity;
*/