
/**
* Se ecnarga de inactivar los eventos que ya son caducos
*/
DROP EVENT VERIFICA_EVENTOS;

CREATE EVENT IF NOT EXISTS VERIFICA_EVENTOS
ON SCHEDULE EVERY 1 MINUTE -- AT CURRENT_TIMESTAMP + INTERVAL 1 MINUTE
STARTS CURRENT_TIMESTAMP
DO
update eventos set activo=0 where
fechahora_evento< now();