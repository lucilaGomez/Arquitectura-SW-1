import React from 'react';
import { IconHeart } from '@tabler/icons-react';
import { Card, Image, Text, Group, Badge, Button, ActionIcon } from '@mantine/core';
import './HotelCard.css';
import HomeCarousel from '../HomeCarousel/HomeCarousel';

const HotelCard = ({ title, description, country, amenities, onReserve, photo }) => {
  console.log(photo);
  const features = amenities.map((amenity, index) => (
    <Badge variant="light" key={index}>
      {amenity.name}
    </Badge>
  ));

  const photoUrls = photo.map(photoItem => photoItem.url);

  return (
    <div className='card-container'>
      <Card withBorder radius="md" p="md" className="card">
        <Card.Section>
          <Image src={photoUrls} />
        </Card.Section>

        <Card.Section className="section" mt="md">
          <Group justify="apart">
            <Text fz="lg" fw={500}>
              {title}
            </Text>
            <Badge size="sm" variant="light">
              {country}
            </Badge>
          </Group>
          <Text fz="sm" mt="xs">
            {description}
          </Text>
        </Card.Section>

        <Card.Section className="section">
          <Text mt="md" className="label" c="dimmed">
            Amenities
          </Text>
          <Group gap={7} mt={5}>
            {features}
          </Group>
        </Card.Section>

        <Group mt="xs">
          <Button radius="md" style={{ flex: 1 }} onClick={onReserve}>
            Reservar
          </Button>
          <ActionIcon variant="default" radius="md" size={36}>
            <IconHeart className="like" stroke={1.5} />
          </ActionIcon>
        </Group>
      </Card>
    </div>
  );
}

export default HotelCard;
