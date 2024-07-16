import whatsappIcon from '../assets/footer/whatsapp.svg'
import mailIcon from '../assets/footer/mail.svg'
import instagramIcon from '../assets/footer/instagram.svg'
import facebookIcon from '../assets/footer/facebook.svg'
import FaqIcon from '../assets/footer/faq-svgrepo-com.svg';

export const footerItems = [
   {
      text: '+54 9 351 111 1111',
      icon: whatsappIcon,
      isContactItem: true,
      to: 'https://wa.me/+5493511111111?text=¡Hola!%20Buenos%20días,%20tengo%20una%20consulta...',
      target:'_blank'
   },
   {
      text: 'Preguntas Frecuentes',
      icon: FaqIcon,
      isContactItem: true,
      to: '/preguntas-frecuentes'
   },
   {
      text: 'HoteleandoLATAM',
      icon: instagramIcon,
      isContactItem: false,
      to: 'https://www.instagram.com/as/',
      target:'_blank'
   },
   {
      text: 'HoteleandoLATAM',
      icon: facebookIcon,
      isContactItem: false,
      to: 'https://www.facebook.com/as/',
      target:'_blank'
   },
]