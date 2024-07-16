import './ErrorPage.css';
import ErrorGif from '../../assets/404ErrorPage/Connection_Error.gif';
import { Button } from '@mantine/core';
import { useNavigate } from 'react-router-dom';
const ErrorPage = () => {
  const navigate = useNavigate ();
  const handleButtonRedirection = () => {
    navigate ('/');
  }
  return (
    <div className='errorPage-container'>
      <div className='errorPage-father-div'>
        <div className='errorPage-iframe-div'>
          <img 
            className='errorPage404-iframe'
            src={ErrorGif}
            title='Error 404 Página no encontrada.'
          />
        </div>
        <div className='errorPage-content-div'>
          <h2 className='errorPage-title-styling'>¡UPS!</h2>
          <p className='errorPage-p-styling'>Parece que la página que estabas buscando no existe. ¡Te invitamos a volver al inicio!</p>
        </div>
      </div>
      <div className='errorPage-button-div'>
        <Button
        fw="normal"
        fullWidth
        onClick={handleButtonRedirection}
        >
          Inicio
        </Button>
      </div>
    </div>
  );
};

export default ErrorPage;