import React, { useEffect, useState } from 'react';
import axios from '../../api';
import { Container, Title, List, Paper, Text } from '@mantine/core';
import MainLayout from '../../layouts/MainLayout/MainLayout';

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

  return (
    <MainLayout overflow="hidden">
      <Container size="md" my="xl">
        <Title order={2} align="center" mb="lg">
          Mis Reservas
        </Title>
        {reservations.length > 0 ? (
          <List spacing="sm" size="sm" center>
            {reservations.map(reservation => (
              <Paper key={reservation.ID} shadow="xs" p="md" withBorder>
                <Text><strong>Hotel:</strong> {reservation.Hotel?.name || 'N/A'}</Text>
                <Text><strong>Check-in:</strong> {reservation.check_in ? new Date(reservation.check_in).toLocaleDateString() : 'N/A'}</Text>
                <Text><strong>Check-out:</strong> {reservation.check_out ? new Date(reservation.check_out).toLocaleDateString() : 'N/A'}</Text>
              </Paper>
            ))}
          </List>
        ) : (
          <Text align="center">No tienes reservas.</Text>
        )}
      </Container>
    </MainLayout>
  );
};

export default UserReservation;
