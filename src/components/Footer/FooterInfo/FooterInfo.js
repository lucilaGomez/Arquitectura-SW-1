import './FooterInfo.css'
import { Link } from 'react-router-dom'
import { Button } from '@mantine/core'
import { footerItems } from '../../../utils/consts';

export default function FooterInfo() {
   return (
      <div className='footer-info-container'>
         <div className='footer-logo-container'>
            <div className='footer-column-container'>
               <Button variant='outline' color={'#FFF'} component={Link} to='/contacto' w='100%' >CONTACTO</Button>
            </div>
         </div>
         <div className='footer-contact-container'>
            <div className='footer-column-container'>
               <h4 className='footer-column-title'>Contacto y Ayuda</h4>
               {footerItems.filter(item => item.isContactItem).map(i => (
                  <Link key={i.text}
                  to={i.to}
                  className='footer-item-container'
                  target={i.target ? '_blank' : undefined}>
                     < img className='footer-item-icon' src={i.icon} alt='icono' />
                     <span className='footer-item-text'>{i.text}</span>
                  </Link>
               ))}
            </div>
         </div>
         <div className='footer-socials-container'>
            <div className='footer-column-container'>
               <h4 className='footer-column-title'>Nuestras redes</h4>
               {footerItems.filter(item => !item.isContactItem).map(i => (
                  <Link key={i.text} to={i.to} target='_blank' className='footer-item-container'>
                     < img className='footer-item-icon' src={i.icon} alt='icono' />
                     <span className='footer-item-text'>{i.text}</span>
                  </Link>
               ))}
            </div>

         </div>
      </div>
   )
}
