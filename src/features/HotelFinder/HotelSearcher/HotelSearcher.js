import axios from '../../../api'; // Ajusta esta ruta según la ubicación de api.js en tu proyecto
import { DatePickerInput } from '@mantine/dates';
import { useState, useEffect } from 'react';
import { Button } from '@mantine/core';
import moment from 'moment';
import "./HotelSearcher.css";

const HotelSearcher = () => {
    const [filters, setFilters] = useState({
        dateRangeValue: [],
    });
    const [fetchingData, setFetchingData] = useState(false);
    const [hotels, setHotels] = useState([]);
    const [userId, setUserId] = useState(null);

    useEffect(() => {
        // Fetch the current user information, including user_id
        const fetchUser = async () => {
            try {
                const response = await axios.get('/validate');
                setUserId(response.data.user.ID);
            } catch (error) {
                console.error("Error fetching user data:", error);
            }
        };

        fetchUser();
    }, []);

    const handleFilterChange = (filterName, value) => {
        setFilters(prevFilters => ({
            ...prevFilters,
            [filterName]: value
        }));
    };

    const init = async () => {
        setFetchingData(true);
        const [startDate, endDate] = filters.dateRangeValue.map(date => date ? moment(date).format('YYYY-MM-DD') : null);

        console.log("Fetching data with filters:", { startDate, endDate });

        try {
            const response = await axios.get('/hotels-availability', {
                params: {
                    start_date: startDate,
                    end_date: endDate
                }
            });
            console.log("Response data:", response.data);
            setHotels(response.data);
        } catch (error) {
            console.error("Error fetching data:", error);
        } finally {
            setFetchingData(false);
        }
    };

    const handleReserve = async (hotelId) => {
        const [checkIn, checkOut] = filters.dateRangeValue.map(date => date ? moment(date).format('YYYY-MM-DD') : null);
        console.log("Attempting to create reservation with data:", {
            hotel_id: hotelId,
            user_id: userId,
            check_in: checkIn,
            check_out: checkOut,
            total_price: 100 // Puedes ajustar este valor según sea necesario
        });

        try {
            const response = await axios.post('/reservations', {
                hotel_id: hotelId,
                user_id: userId,
                check_in: checkIn,
                check_out: checkOut,
                total_price: 100 // Puedes ajustar este valor según sea necesario
            });
            console.log("Reservation response:", response.data);
            alert("Reserva creada exitosamente!");
        } catch (error) {
            console.error("Error creando la reserva:", error);
            alert("Error al crear la reserva");
        }
    };

    return (
        <div className="hotelSearcher-wrapper">
            <DatePickerInput
                classNames={{ root: "moderationContainer-DatePicker-root" }}
                type="range"
                label="Selecciona un rango de fechas"
                value={filters.dateRangeValue}
                onChange={(value) => handleFilterChange('dateRangeValue', value)}
            />
            <Button onClick={init} disabled={fetchingData} className="hotelSearcher-button">
                Buscar
            </Button>

            {hotels.length > 0 && (
                <div className="hotel-list">
                    {hotels.map(hotel => (
                        <div key={hotel.ID} className="hotel-item">
                            <h3>{hotel.name}</h3>
                            <p>{hotel.description}</p>
                            <p>{hotel.address}, {hotel.city}, {hotel.country}</p>
                            <p>Disponibilidad: {hotel.availability}</p>
                            <Button onClick={() => handleReserve(hotel.ID)} className="hotelSearcher-button">
                                Reservar
                            </Button>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default HotelSearcher;
