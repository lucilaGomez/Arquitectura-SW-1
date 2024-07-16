import React, { useState } from 'react';
import {
  Paper,
  TextInput,
  PasswordInput,
  Button,
  Title,
  Text,
  Anchor,
} from '@mantine/core';
import { useNavigate } from 'react-router-dom';
import { useUserContext } from '../../../context/UserContext';
import './LoginMainPage.css';

const LoginMainPage = () => {
  const { login } = useUserContext();
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleSubmit = async () => {
    try {
      await login(email, password);
      navigate('/');
    } catch (error) {
      setError('Error al iniciar sesión');
    }
  };

  return (
    <div className="wrapper">
      <Paper className="form" radius={0} p={30}>
        <Title order={2} className="title" ta="center" mt="md" mb={50}>
          Ingresa a Hoteleando!
        </Title>
        <TextInput
          label="Email address"
          placeholder="hello@gmail.com"
          size="md"
          value={email}
          onChange={(event) => setEmail(event.currentTarget.value)}
        />
        <PasswordInput
          label="Password"
          placeholder="Your password"
          mt="md"
          size="md"
          value={password}
          onChange={(event) => setPassword(event.currentTarget.value)}
        />
        {error && <Text color="red">{error}</Text>}
        <Button fullWidth mt="xl" size="md" onClick={handleSubmit}>
          Login
        </Button>
        <Text ta="center" mt="md">
          Don't have an account?{' '}
          <Anchor href="#" fw={700} onClick={(event) => event.preventDefault()}>
            Register
          </Anchor>
        </Text>
      </Paper>
    </div>
  );
};

export default LoginMainPage;
