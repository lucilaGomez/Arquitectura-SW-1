import React from 'react';
import { Link } from "react-router-dom";
import { useUserContext } from '../../context/UserContext';
import "./Header.css";

const Header = () => {
  const { user, logout } = useUserContext();

  return (
    <header className="header__container">
      <h1 className="header__logo"><a className="header__logo1" href="/">HoteleandoLATAM</a></h1>
      <ul className="header__list">
        <Link className="header__a__list" to="/" >Home</Link>
        <Link className="header__a__list" to="/hotel-finder" >Buscar Hotel</Link>
        {user ? (
          <>
            <Link className="header__a__list" to="/user-reservation" >Mis Reservas</Link>
            {user.Role === 'admin' && (
              <>
                <Link className="header__a__list" to="/crear-hotel" >Crear Hotel</Link>
                <Link className="header__a__list" to="/crear-amenity" >Crear Amenity</Link>
              </>
            )}
            <span className="header__user__email">{user.Email}</span>
            <button onClick={logout} className="header__logout__button">Logout</button>
          </>
        ) : (
          <>
            <Link className="header__a__list" to="/login" >Login</Link>
            <Link className="header__a__list" to="/registro" >Registro</Link>
          </>
        )}
      </ul>
    </header>
  );
};

export default Header;
