import React, { useEffect, useState } from 'react';
import axios from '../../api';
import { Container, Title, List, Paper, Text, Button } from '@mantine/core';
import MainLayout from '../../layouts/MainLayout/MainLayout';
import './UserReservation.css';

const UserReservation = () => {
  const [reservations, setReservations] = useState([]);

  useEffect(() => {
    const fetchReservations = async () => {
      try {
        const response = await axios.get('/reservations/my');
        setReservations(response.data.reservations);
      } catch (error) {
        console.error('Error fetching reservations:', error);
      }
    };

    fetchReservations();
  }, []);

  const handleCancelReservation = async (id) => {
    try {
      await axios.delete(`/reservations/${id}`);
      // Actualizar la lista de reservas después de cancelar una reserva si es necesario
      const updatedReservations = reservations.filter(reservation => reservation.ID !== id);
      setReservations(updatedReservations);
    } catch (error) {
      console.error('Error cancelling reservation:', error);
      // Manejo de errores, si es necesario
    }
  };

  return (
    <MainLayout overflow="hidden">
      <Title order={2} align="center" mb="lg">
        Mis Reservas
      </Title>
      {reservations.length > 0 ? (
        <div className='userReservation-div' size="md" my="xl">
          <ul className='userReservation-div' spacing="sm" size="sm" center>
            {reservations.map(reservation => (
              <Paper classNames={{root:"userReservation-root"}} key={reservation.ID} shadow="xs" p="md" withBorder>
                <div>
                  <p className='userReservation-p'><strong>Hotel:</strong> {reservation.Hotel?.name || 'N/A'}</p>
                  <p className='userReservation-p'><strong>Check-in:</strong> {reservation.check_in ? new Date(reservation.check_in).toLocaleDateString() : 'N/A'}</p>
                  <p className='userReservation-p'><strong>Check-out:</strong> {reservation.check_out ? new Date(reservation.check_out).toLocaleDateString() : 'N/A'}</p>
                </div>
                <div className='userReservation-buttonDiv'>
                  <Button
                    fw="normal"
                    classNames={{root:"userReservation-buttons"}}
                    onClick={() => handleCancelReservation(reservation.ID)}
                  >
                    Cancelar reserva
                  </Button>
                </div>
              </Paper>
            ))}
          </ul>
        </div>
      ) : (
        <Text align="center">No tienes reservas.</Text>
      )}
    </MainLayout>
  );
};

export default UserReservation;
