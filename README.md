# Arquitectura de Software I - Integrador 2023
Como práctico integrador se solicita la creación de un sitio de reserva de
habitaciones para una cadena de hoteles.
1. El backend, desarrollado en Golang, que brindara todas las interfaces necesarias
para dar solución al requerimiento.
2. El frontend, desarrollado en React, representa la vista final del usuario y
consumirá los servicios desarrollados en el backend.
Para la construcción del sitio de reservas se solicitan los siguientes puntos.

- Autentificación de usuarios. Login y Permisos de Usuarios. Usuario viajante,
usuario administrador que maneja la carga de hoteles nuevos (titulo, foto,
descripción y habitaciones) y el listado de reserva de los viajantes.

- Pantalla o Interfaz de Bienvenida. (Home de la app). Listado de hoteles
disponibles en un combo (select), fecha desde - fecha hasta.

- Detalle del hotel seleccionado, disponibilidad de la reserva y posibles alternativas
a elegir desde ese listado. Los hoteles tienen el mismo tipo de pieza (para simplificar
el desarrollo) y la disponibilidad se acaba cuando se reserva la cantidad de piezas
disponibles en el período.

- Confirmación de la reserva con la necesidad de que el cliente esté registrado, si
no está registrado solicitarle el registro con sus datos personales y se confirma la
reserva.

- Página de listado de reservas, tanto del cliente como del administrador del sitio
donde pueda discriminar reservas por hotel y por día y visualice los datos de las
personas que hicieron la reservas para el caso del administrador.


Nota: No es necesario realizar algún proceso de Pagos del Mismo, se asume que el pago está aprobado y vamos por el Caso de Uso “feliz” donde se confirma la Compra o la Orden.

Extra Points:
- Posibilidad de subir fotos de los hoteles.
- Posibilidad de cargar amenities para los hoteles.

  
Condiciones de Regularidad y Exámen Final
- Para regularizar la materia se pide el desarrollo relacionado al flujo de reserva.
- Para el exámen final se solicita el trabajo completo, es decir la carga completa del
hotel con el upload de fotos y amenities.
