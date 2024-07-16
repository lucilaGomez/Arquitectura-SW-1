import axios from '../../../api'; 
import { DatePickerInput } from '@mantine/dates';
import { useState, useEffect } from 'react';
import { Button } from '@mantine/core';
import moment from 'moment';
import "./HotelSearcher.css";
import HotelCard from "../../../components/HotelCard/HotelCard";

const HotelSearcher = () => {
    const [filters, setFilters] = useState({
        dateRangeValue: [],
    });
    const [fetchingData, setFetchingData] = useState(false);
    const [hotels, setHotels] = useState([]);
    const [userId, setUserId] = useState(null);

    useEffect(() => {
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
            <div className='hotelPicker-wrapper'>
            <DatePickerInput
                classNames={{ root: "moderationContainer-DatePicker-root" }}
                type="range"
                label="Selecciona un rango de fechas"
                value={filters.dateRangeValue}
                onChange={(value) => handleFilterChange('dateRangeValue', value)}
            />
            </div>
            <Button onClick={init} disabled={fetchingData} className="hotelSearcher-button">
                Buscar
            </Button>

            {hotels?.length > 0 && (
                <div className="hotel-list">
                    {hotels?.map(hotel => (
                        <HotelCard
                            key={hotel?.ID}
                            title={hotel?.name}
                            description={hotel?.description}
                            country={hotel?.country}
                            amenities={hotel?.amenities}
                            photo={hotel?.photos}
                            onReserve={() => handleReserve(hotel?.ID)}
                        />
                    ))}
                </div>
            )}
        </div>
    );
};

export default HotelSearcher;
