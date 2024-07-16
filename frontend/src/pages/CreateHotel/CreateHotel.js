import React, { useState } from 'react';
import axios from '../../api';
import MainLayout from '../../layouts/MainLayout/MainLayout';
import { Button, TextInput } from '@mantine/core';
import './CreateHotel.css';

const CreateHotel = () => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [address, setAddress] = useState('');
  const [city, setCity] = useState('');
  const [country, setCountry] = useState('');
  const [amenities, setAmenities] = useState('');
  const [photos, setPhotos] = useState('');
  const [hotelID, setHotelID] = useState('');
  const [availability, setAvailability] = useState('');
  const [availabilityID, setAvailabilityID] = useState('');
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const handleHotelSubmit = async (event) => {
    event.preventDefault();
    setMessage('');

    const hotelData = {
      name,
      description,
      address,
      city,
      country,
      amenities: amenities.split(','),
      photos: photos.split(',')
    };

    try {
      const response = await axios.post('/hotels', hotelData);
      console.log('Response:', response);
      setMessage('Hotel creado exitosamente');
      setHotelID(response.data.ID);
      setName('');
      setDescription('');
      setAddress('');
      setCity('');
      setCountry('');
      setAmenities('');
      setPhotos('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al crear hotel');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    }
  };

  const handleHotelUpdate = async (event) => {
    event.preventDefault();
    setMessage('');

    const hotelData = {
        id: hotelID,
        name,
        description,
        address,
        city,
        country,
        amenities: amenities.split(','),
        photos: photos.split(',')
    };

    try {
        const response = await axios.put(`/hotels/${hotelID}`, hotelData);
        setMessage('Hotel actualizado exitosamente');
        
        // Actualiza los estados con los datos actualizados
        setName(response.data.Name);
        setDescription(response.data.Description);
        setAddress(response.data.Address);
        setCity(response.data.City);
        setCountry(response.data.Country);
        // Actualiza amenities y photos si son necesarios
        // ...

    } catch (error) {
        console.error('Error:', error);
        if (error.response) {
            setMessage(error.response.data.details || 'Error al actualizar hotel');
        } else if (error.request) {
            setMessage('No response from server');
        } else {
            setMessage('Error en la solicitud');
        }
    }
};


  const handleHotelDelete = async () => {
    setMessage('');

    try {
      await axios.delete(`/hotels/${hotelID}`);
      setMessage('Hotel eliminado exitosamente');
      setHotelID('');
      setName('');
      setDescription('');
      setAddress('');
      setCity('');
      setCountry('');
      setAmenities('');
      setPhotos('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al eliminar hotel');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    }
  };

  const handleAvailabilitySubmit = async (event) => {
    event.preventDefault();
    setMessage('');

    const availabilityData = {
      hotel_id: hotelID,
      available: parseInt(availability, 10)
    };

    try {
      const response = await axios.post('/availability', availabilityData);
      setMessage('Availability creada exitosamente');
      setAvailabilityID(response.data.ID);
      setAvailability('');
      setLoading(true);
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al crear availability');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    } finally {
      setLoading(false);
    }
  };

  const handleAvailabilityUpdate = async (event) => {
    event.preventDefault();
    setMessage('');

    const availabilityData = {
      hotel_id: hotelID,
      available: parseInt(availability, 10)
    };

    try {
      await axios.put(`/availability/${availabilityID}`, availabilityData);
      setMessage('Availability actualizada exitosamente');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al actualizar availability');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    }
  };

  return (
    <MainLayout>
      <div className='createHotel-father'>
        <div className='createHotel-wrapper'>
          <h2 className='hotel-title'>Crea tu hotel</h2>
          <form className='createHotel-firstForm' onSubmit={handleHotelSubmit}>
            <div className='createHotel-firstForm-div'>
              <TextInput placeholder="Nombre" value={name} onChange={(e) => setName(e.target.value)} required />
              <TextInput placeholder="Descripcion" value={description} onChange={(e) => setDescription(e.target.value)} required />
              <TextInput placeholder="Direccion" value={address} onChange={(e) => setAddress(e.target.value)} required />
              <TextInput placeholder="Ciudad" value={city} onChange={(e) => setCity(e.target.value)} required />
              <TextInput placeholder="Pais" value={country} onChange={(e) => setCountry(e.target.value)} required />
              <TextInput placeholder="Amenities" value={amenities} onChange={(e) => setAmenities(e.target.value)} required />
              <TextInput placeholder="Fotos" value={photos} onChange={(e) => setPhotos(e.target.value)} required />
              <Button type="submit">Crear Hotel</Button>
            </div>
          </form>
          <h2 className='hotel-title'>Actualizar hotel</h2>
          <form className='createHotel-firstForm' onSubmit={handleHotelUpdate}>
            <div className='createHotel-firstForm-div'>
              <TextInput placeholder="ID del Hotel" value={hotelID} onChange={(e) => setHotelID(e.target.value)} required />
              <TextInput placeholder="Nombre" value={name} onChange={(e) => setName(e.target.value)} required />
              <TextInput placeholder="Descripcion" value={description} onChange={(e) => setDescription(e.target.value)} required />
              <TextInput placeholder="Direccion" value={address} onChange={(e) => setAddress(e.target.value)} required />
              <TextInput placeholder="Ciudad" value={city} onChange={(e) => setCity(e.target.value)} required />
              <TextInput placeholder="Pais" value={country} onChange={(e) => setCountry(e.target.value)} required />
              <TextInput placeholder="Amenities" value={amenities} onChange={(e) => setAmenities(e.target.value)} required />
              <TextInput placeholder="Fotos" value={photos} onChange={(e) => setPhotos(e.target.value)} required />
              <Button type="submit">Actualizar Hotel</Button>
            </div>
          </form>
          <div className="createHotel-buttons">
            <TextInput placeholder="ID del Hotel" value={hotelID} onChange={(e) => setHotelID(e.target.value)} required />
            <Button onClick={handleHotelDelete} color="red">Eliminar Hotel</Button>
          </div>
          <h2 className='hotel-title'>Define la disponibilidad de tu hotel</h2>
          <form className='createHotel-secondForm' onSubmit={handleAvailabilitySubmit}>
            <div className='createHotel-firstForm-div'>
              <TextInput placeholder="Hotel ID" value={hotelID} onChange={(e) => setHotelID(e.target.value)} required />
              <TextInput placeholder="Availability" value={availability} onChange={(e) => setAvailability(e.target.value)} required />
              <Button loading={loading} type="submit">Crear Availability</Button>
            </div>
          </form>
          {message && <p className='createHotel-messagetext'>{message}</p>}
        </div>
      </div>
    </MainLayout>
  );
};

export default CreateHotel;



/*import React, { useState } from 'react';
import axios from '../../api';
import MainLayout from '../../layouts/MainLayout/MainLayout';
import "./CreateHotel.css";
import { Button, TextInput } from '@mantine/core';

const CreateHotel = () => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [address, setAddress] = useState('');
  const [city, setCity] = useState('');
  const [country, setCountry] = useState('');
  const [amenities, setAmenities] = useState('');
  const [photos, setPhotos] = useState('');
  const [hotelID, setHotelID] = useState('');
  const [availability, setAvailability] = useState('');
  const [message, setMessage] = useState('');
  const [loading, setLoading] = useState(false);

  const handleHotelSubmit = async (event) => {
    event.preventDefault();
    setMessage('');

    const hotelData = {
      name,
      description,
      address,
      city,
      country,
      amenities: amenities.split(','),
      photos: photos.split(',')
    };

    try {
      const response = await axios.post('/hotels', hotelData);
      console.log('Response:', response);
      setMessage('Hotel creado exitosamente');
      setHotelID(response.data.ID);
      setName('');
      setDescription('');
      setAddress('');
      setCity('');
      setCountry('');
      setAmenities('');
      setPhotos('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al crear hotel');
      } else if (error.request) {
        setMessage('No response from server');
      } else {
        setMessage('Error en la solicitud');
      }
    }
  };

  const handleAvailabilitySubmit = async (event) => {
    event.preventDefault();
    setMessage('');

    const availabilityData = {
      hotel_id: hotelID,
      available: parseInt(availability, 10)
    };

    try {
      const response = await axios.post('/availability', availabilityData);
      setMessage('Availability creada exitosamente');
      setAvailability('');
      setLoading(true);
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al crear availability');
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
      <div className='createHotel-father'>
      <div className='createHotel-wrapper'>
        <h2 className='hotel-title'>Crea tu hotel</h2>
        <form className='createHotel-firstForm' onSubmit={handleHotelSubmit}>
          <div className='createHotel-firstForm-div'>
          <TextInput placeholder="Nombre" value={name} onChange={(e) => setName(e.target.value)} required />
          <TextInput placeholder="Descripcion" value={description} onChange={(e) => setDescription(e.target.value)} required />
          <TextInput placeholder="Direccion" value={address} onChange={(e) => setAddress(e.target.value)} required />
          <TextInput placeholder="Ciudad" value={city} onChange={(e) => setCity(e.target.value)} required />
          <TextInput placeholder="Pais" value={country} onChange={(e) => setCountry(e.target.value)} required />
          <TextInput placeholder="Amenities" value={amenities} onChange={(e) => setAmenities(e.target.value)} required />
          <TextInput placeholder="Fotos" value={photos} onChange={(e) => setPhotos(e.target.value)} required />
          <Button type="submit">Crear Hotel</Button>
          </div>
        </form>
        <h2 className='hotel-title'>Define la disponibilidad de tu hotel</h2>
        <form className='createHotel-secondForm' onSubmit={handleAvailabilitySubmit}>
        <div className='createHotel-firstForm-div'>
          <TextInput placeholder="Hotel ID" value={hotelID} onChange={(e) => setHotelID(e.target.value)} required />
          <TextInput placeholder="Availability" value={availability} onChange={(e) => setAvailability(e.target.value)} required />
          <Button loading={loading} type="submit">Crear Availability</Button>
          </div>
        </form>
        {message && <p className='createHotel-messagetext'>{message}</p>}
      </div>
      </div>
    </MainLayout>
  );
};

export default CreateHotel;
*/