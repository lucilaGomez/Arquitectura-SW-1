// src/features/Register/RegisterAuthentication/RegisterAuth.js
import React, { useState } from 'react';
import {
  TextInput,
  PasswordInput,
  Select,
  Paper,
  Title,
  Container,
  Button,
  Notification
} from '@mantine/core';
import api from '../../../api'; // Asegúrate de que la ruta sea correcta
import './RegisterAuth.css';

const RegisterAuth = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [role, setRole] = useState('');
  const [error, setError] = useState(null);
  const [success, setSuccess] = useState(null);

  const handleSubmit = async (event) => {
    event.preventDefault();
    setError(null);
    setSuccess(null);

    try {
      const response = await api.post('/signup', {
        email,
        password,
        role,
      });
      console.log('Response:', response); // Imprime la respuesta del backend en la consola
      setSuccess('Usuario registrado exitosamente');
    } catch (error) {
      console.error('Error:', error); // Imprime el error completo en la consola
      if (error.response) {
        setError(error.response.data.details || 'Error al registrar usuario');
      } else {
        setError('Error de conexión o servidor no disponible');
      }
    }
  };

  return (
    <Container size={420} my={40}>
      <Title ta="center" className="title">
        ¡Crea ahora mismo tu cuenta!
      </Title>

      <Paper withBorder shadow="md" p={30} mt={30} radius="md" component="form" onSubmit={handleSubmit}>
        {error && <Notification color="red" onClose={() => setError(null)}>{error}</Notification>}
        {success && <Notification color="green" onClose={() => setSuccess(null)}>{success}</Notification>}
        <TextInput label="Email" placeholder="you@mantine.dev" required value={email} onChange={(e) => setEmail(e.target.value)} />
        <PasswordInput label="Password" placeholder="Your password" required mt="md" value={password} onChange={(e) => setPassword(e.target.value)} />
        <Select label="Define tu rol" placeholder="Rol" data={['admin', 'user']} required mt="md" value={role} onChange={setRole} />
        <Button type="submit" fullWidth mt="xl">
          Registrarse
        </Button>
      </Paper>
    </Container>
  );
};

export default RegisterAuth;
