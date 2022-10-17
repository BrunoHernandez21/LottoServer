
autorrellenado de video mediante vidID al crear evento  
pedir CCV de la tarjeta y no guardarlo en el BD                            // espera
terminar de separar plan y suscripcion                                     // espera

 ## --  proceso, rechazado, pagado, cancelado

 si Guardar la tarjeta
 ccv guardar ASH1
 pedir id y ccv
Agregar el costo del evento a la peticion                                       // espera

terminar de separar plan y suscripcion                                          // espera
Reintentar comprar una orden y cancelarla                                       // espera
websocket                                                                       // espera
agregar la respuesta directamente a la lista existente                          // espera

"type": "invalid_request_error"