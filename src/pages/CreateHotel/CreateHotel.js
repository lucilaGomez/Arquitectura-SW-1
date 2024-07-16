import React, { useState } from 'react';
import axios from '../../api';
import MainLayout from '../../layouts/MainLayout/MainLayout';

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
      console.log('Response:', response);
      setMessage('Availability creada exitosamente');
      setAvailability('');
    } catch (error) {
      console.error('Error:', error);
      if (error.response) {
        setMessage(error.response.data.details || 'Error al crear availability');
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
        <form onSubmit={handleHotelSubmit}>
          <input placeholder="Nombre" value={name} onChange={(e) => setName(e.target.value)} required />
          <input placeholder="Descripcion" value={description} onChange={(e) => setDescription(e.target.value)} required />
          <input placeholder="Direccion" value={address} onChange={(e) => setAddress(e.target.value)} required />
          <input placeholder="Ciudad" value={city} onChange={(e) => setCity(e.target.value)} required />
          <input placeholder="Pais" value={country} onChange={(e) => setCountry(e.target.value)} required />
          <input placeholder="Amenities" value={amenities} onChange={(e) => setAmenities(e.target.value)} required />
          <input placeholder="Fotos" value={photos} onChange={(e) => setPhotos(e.target.value)} required />
          <button type="submit">Crear Hotel</button>
        </form>
        <form onSubmit={handleAvailabilitySubmit}>
          <input placeholder="Hotel ID" value={hotelID} onChange={(e) => setHotelID(e.target.value)} required />
          <input placeholder="Availability" value={availability} onChange={(e) => setAvailability(e.target.value)} required />
          <button type="submit">Crear Availability</button>
        </form>
        {message && <p>{message}</p>}
      </div>
    </MainLayout>
  );
};

export default CreateHotel;
