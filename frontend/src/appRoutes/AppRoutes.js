import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "../pages/Home/Home";
import HotelFinder from "../pages/HotelFinder/HotelFinder";
import Login from "../pages/Login/Login";
import UserReservation from "../pages/UserReservation/UserReservation";
import HotelDetails from "../pages/HotelDetails/HotelDetails";
import Register from "../pages/Register/Register";
import Error404Page from "../pages/Error404Page/Error404Page";
import ReservationsView from "../pages/ReservationsView/ReservationsView";
import CreateHotel from "../pages/CreateHotel/CreateHotel";
import CreateAmenity from "../pages/CreateAmenity/CreateAmenity";

const AppRoutes = () => {
  return (
    <BrowserRouter>
      <Routes>
          <Route exact path="/" element={<Home />}/>
          <Route path="hotel-finder" element={<HotelFinder />} />
          <Route path="login" element={<Login />} />
          <Route path="user-reservation" element={<UserReservation />} />
          <Route path="hotel-details" element={<HotelDetails />} />
          <Route path="registro" element={<Register />} />
          <Route path="reservas" element={<ReservationsView />} />
          <Route path="crear-hotel" element={<CreateHotel />} />
          <Route path="crear-amenity" element={<CreateAmenity />} />
          <Route path="*" element={<Error404Page />} />
      </Routes>
    </BrowserRouter>
  )
}
export default AppRoutes