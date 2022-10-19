##-------------------------------------- Proces de Pasarela de pagos
##----------- Pagos unicos
##-- generar Orden
DROP PROCEDURE IF EXISTS `genera_orden`;
delimiter $$
CREATE PROCEDURE genera_orden(IN user_id int) 
genorden:begin
DECLARE time_Now datetime;
DECLARE orden_id bigint DEFAULT 0;
IF  (SELECT count(id) from carrito WHERE usuario_id = user_id AND activo = TRUE) = 0  THEN
    SELECT * from ordenes WHERE id = null;
	LEAVE genorden;
END IF;
set time_Now = now();
##--crea la orden 
INSERT INTO ordenes ( ordenes.status, fecha_emitido, precio_total,puntos_total,usuario_id,is_suscription,moneda)
	SELECT 	"proceso",
		time_Now,
        SUM(c.total_linea),
        SUM(c.puntos_linea),
        user_id,
        false,
        "MXN"
	from carrito as c
    JOIN planes p ON c.plan_id = p.id 
    WHERE c.usuario_id = user_id AND c.activo = TRUE AND p.suscribcion = FALSE;
##-- Busca su id de orden
SET orden_id = (SELECT  id from ordenes WHERE fecha_emitido = time_Now AND usuario_id = user_id);
##-- crea los items de la orden            
INSERT INTO items_orden ( cantidad,titulo,total_linea,puntos_linea,moneda,plan_id,orden_id)
	SELECT 	c.cantidad,
    	p.titulo,
        c.total_linea,
        c.puntos_linea,
        c.moneda,
        c.plan_id,
        orden_id
	from carrito as c
    JOIN planes p ON c.plan_id = p.id 
    WHERE c.usuario_id = user_id AND c.activo = TRUE AND p.suscribcion = FALSE;
##-- limpia el carrito
UPDATE carrito SET activo = false WHERE carrito.usuario_id = user_id;
SELECT * from ordenes WHERE id = orden_id;
end
$$
delimiter ;
## -- Se pago Satisfctoriamente
DROP PROCEDURE IF EXISTS `orden_pagada`;
delimiter $$
CREATE PROCEDURE orden_pagada(IN iorden_id int, IN razon TEXT) 
ordenpag:begin
DECLARE pcash bigint;
DECLARE ccash bigint;
DECLARE user_id bigint;
DECLARE susc_id bigint;
DECLARE resp varchar(40);
IF ( (SELECT COUNT(id) from ordenes WHERE id = iorden_id AND is_suscription = false) != 1 ) THEN
	##--TODO Hospital
	SELECT "Orden corrupta";
    LEAVE ordenpag;
END IF;
    ##--obtenemos el usuario
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);      
	##--obtenemos valores actuales de cartera
    SET ccash = (SELECT	puntos from carteras WHERE usuario_id = user_id);
	##--obtenemos la suma total de pontos
    SET pcash = (SELECT puntos_total FROM ordenes WHERE id = iorden_id);
    ##--actualizamos cartera
    UPDATE carteras SET puntos = ccash + pcash WHERE usuario_id = user_id;
    ## -- crear compra
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,false);
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "pagado" WHERE ordenes.id = iorden_id;
	##--insertamos los beneficios del plan
    INSERT INTO beneficios_usuario ( cobrado, usuario_id,beneficio_id)
		SELECT 	true,
        		user_id,
                bp.beneficio_id
        from items_orden i 
        JOIN beneficios_plan bp ON i.plan_id = bp.plan_id 
        WHERE i.orden_id = iorden_id; 
    
    SELECT "Proceso realizado con exito";
end
$$
delimiter ;
## -- Error en el pago
DROP PROCEDURE IF EXISTS `orden_rechazada`;
delimiter $$
CREATE PROCEDURE orden_rechazada(IN iorden_id int,IN razon TEXT) 
begin
	DECLARE user_id bigint;
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);  
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "rechazado" WHERE ordenes.id = iorden_id;
    ## -- crear compra fallida
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,true);
    ## -- out
	SELECT razon;
end
$$
delimiter ; 
##----------- Subscripciones
##-- Generar Subscripcion  (orden)
DDROP PROCEDURE IF EXISTS `orden_subscripcion`;
delimiter $$
CREATE PROCEDURE orden_subscripcion(IN user_id int,IN plan_id int) 
ordensubs:begin
DECLARE time_Now datetime;
DECLARE orden_id bigint DEFAULT 0;
IF  (SELECT count(id) from planes WHERE planes.id = plan_id AND activo = TRUE AND suscribcion = true) = 0  THEN
    LEAVE ordensubs;
END IF;

set time_Now = now();
INSERT INTO ordenes ( ordenes.status, fecha_emitido, precio_total,puntos_total,usuario_id,is_suscription,moneda)
	SELECT 	"proceso",
		time_Now,
                p.precio,
                p.puntos,
                user_id,
                true,
                p.moneda
        from planes p 
        WHERE p.id = plan_id;
	SET orden_id = (SELECT  id from ordenes WHERE fecha_emitido = time_Now AND usuario_id = user_id);
    INSERT INTO items_orden ( cantidad,titulo,total_linea,puntos_linea,moneda,plan_id,orden_id)
	SELECT 	1,
    	p.titulo,
        p.precio,
        p.puntos,
        p.moneda,
        p.id,
        orden_id
	from planes p
    WHERE p.activo = TRUE AND p.suscribcion = TRUE AND p.id = plan_id ;        
    SELECT * FROM ordenes WHERE fecha_emitido = time_Now AND usuario_id = user_id;
end
$$
delimiter ;
##-- Suscribcion Pagada (stack/mes)
DROP PROCEDURE IF EXISTS `suscribcion_aceptada`;
delimiter $$
CREATE PROCEDURE suscribcion_aceptada(IN iorden_id int, IN razon TEXT,IN str_suscript varchar(64)) 
susacept:begin
    DECLARE pcash bigint;
    DECLARE ccash bigint;
    DECLARE user_id bigint;
    DECLARE plann_id bigint;
    IF  (SELECT COUNT(id) from ordenes WHERE id = iorden_id AND is_suscription = true) != 1  THEN 
    	SELECT "Orden corrupta";
        LEAVE susacept;
    END IF;
    ##--optenemos el usuario
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);                                    
	##--optenemos valores actuales de cartera
    SET ccash = (SELECT	puntos from carteras WHERE usuario_id = user_id);
	##--optenemos la suma total de puntos
    SET pcash = (SELECT puntos_total FROM ordenes WHERE id = iorden_id);
    ##--actualizamos cartera
    UPDATE carteras SET puntos = ccash + pcash WHERE usuario_id = user_id;
    ## -- crear compra
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,false);
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "suscrito" WHERE ordenes.id = iorden_id;
	##--incertamos los beneficios de dias del plan
    INSERT INTO beneficios_usuario ( cobrado, usuario_id,beneficio_id)
		SELECT 	true,
        		user_id,
                bp.beneficio_id
        from items_orden i 
        JOIN beneficios_plan bp ON i.plan_id = bp.plan_id 
        WHERE i.orden_id = iorden_id;
	##--actualizamos la suscribcion
    SET plann_id = (SELECT p.id from items_orden as i INNER JOIN planes AS p ON i.plan_id = p.id WHERE i.orden_id = iorden_id);
    UPDATE suscripciones sus
    SET monto_mensual=(SELECT precio from planes WHERE id = plann_id),
        fecha_inicio=NOW(),
        fecha_fin=DATE_ADD(NOW(), INTERVAL 1 MONTH),
        plan_id=plann_id,
        stripe_suscription = str_suscript
      WHERE sus.usuario_id = user_id AND plann_id IS NOT NULL;
    SELECT "Proceso realizado con exito";
end
$$
delimiter ;
##-- Suscribcion RECHADAZADA (Retira los beneficios)
DROP PROCEDURE IF EXISTS `suscribcion_rechazada`;
delimiter $$
CREATE PROCEDURE suscribcion_rechazada(IN iorden_id int, IN razon TEXT) 
begin
    DECLARE user_id bigint;
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);  
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "finalizado" WHERE ordenes.id = iorden_id;
    ## -- crear compra fallida
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,true);
    ## -- out
	SELECT razon;
end
$$
delimiter ;
##--------------------------------------  Flow Control
##----------- Proceso que requieren calcular
##-- generar_ganador
DROP PROCEDURE IF EXISTS `generar_ganador`;
delimiter $$
CREATE PROCEDURE generar_ganador() 
begin 
    INSERT INTO ganador ( cantidad, cantidad_acumulada, concepto, evento_id, usuario_id, evento_usuario_id)
	(SELECT 	e.premio_cash,
     			e.acumulado,
        		e.premio_otros,
     			e.id,
                eu.usuario_id,
                eu.id
        from eventos as e 
       	LEFT JOIN evento_usuario as eu ON e.id = eu.evento_id   
     	LEFT JOIN videos_estadisticas ve ON e.video_id = ve.video_id  
        WHERE 	e.fechahora_evento > DATE_SUB(NOW(),INTERVAL 10 MINUTE ) AND
     			MINUTE(ve.fecha) = MINUTE(NOW()) AND
     			(eu.views_count = ve.views_count OR
                 eu.like_count = ve.views_count OR  
                 eu.shared_count = ve.views_count OR  
                 eu.comments_count = ve.views_count OR  
                 eu.saved_count = ve.views_count OR 
                 eu.dislikes_count = ve.views_count));
    UPDATE evento_usuario as eu 
        LEFT JOIN eventos as e ON e.id = eu.evento_id 
        SET eu.activo = false
        WHERE e.fechahora_evento < DATE_ADD(NOW(), INTERVAL 3 MINUTE) AND eu.activo = true;
    UPDATE eventos as e  
            SET e.activo = false
            WHERE e.fechahora_evento < DATE_ADD(NOW(), INTERVAL 3 MINUTE);
    UPDATE videos SET activo = false WHERE id NOT IN (SELECT video_id FROM eventos WHERE activo = true);
    SELECT "Realizado correctamente";
end
$$
delimiter ;
##----------- Procesos de verificacion
##-- verificar_propiedades_usuario
DROP PROCEDURE IF EXISTS `verificar_tablas_usuario`;
delimiter $$
CREATE PROCEDURE verificar_tablas_usuario(IN user_id bigint) 
begin
    IF (SELECT COUNT(pu.id) from propiedades_usuarios pu WHERE pu.usuario_id = user_id) = 0  THEN 
        INSERT INTO propiedades_usuarios(id,nivel_acceso,custom_attributes,fecha_inicio,fecha_fin,usuario_id) 
        VALUES (null,'3',null,now(),now(),user_id);
    END IF;
   	IF (SELECT COUNT(s.id) from suscripciones s WHERE s.usuario_id = user_id) = 0  THEN 
    	INSERT INTO suscripciones(monto_mensual,suscripciones.usuario_id) VALUES (0,user_id);
    END IF;
    IF (SELECT COUNT(c.id) from carteras c WHERE c.usuario_id = user_id) = 0  THEN 
        INSERT INTO carteras(puntos,saldo_mxn,saldo_usd,usuario_id) VALUES (0,0,0,user_id);
    END IF;
    IF (SELECT COUNT(ur.id) from usuarios_roles ur WHERE ur.user_id = user_id) = 0  THEN 
        INSERT INTO usuarios_roles(user_id,role_id) VALUES (user_id,1);
    END IF;
    SELECT "Realizado correctamente";
end
$$
delimiter ;

##-- verificar_propiedades_usuario
DROP PROCEDURE IF EXISTS `verificar_propiedades_usuario`;
delimiter $$
CREATE PROCEDURE verificar_propiedades_usuario() 
begin 
	UPDATE propiedades_usuarios as pu
    LEFT JOIN suscripciones as s ON pu.usuario_id = s.usuario_id  
    SET	pu.nivel_acceso 	=  
    	IF(
            s.plan_id != null,
           (SELECT b.acces_id from beneficios_plan as bp LEFT JOIN beneficios as b ON bp.beneficio_id = b.id WHERE bp.plan_id = s.plan_id), 
           "3"
          ),
    	pu.fecha_inicio	=  IF(s.plan_id != null, s.fecha_inicio, null),
        pu.fecha_fin 		=  IF(s.plan_id != null, s.fecha_fin, null)
    WHERE IFNULL(pu.fecha_fin = null, DATE_SUB(NOW(),INTERVAL 31 DAY)) < NOW();
    SELECT "Realizado correctamente";
end
$$
delimiter ;




##-------------------------------------- Funciones legacy
##----------- Propiedades usuario
##-- verificar_suscribciones
DROP PROCEDURE IF EXISTS `verificar_suscribciones`;
delimiter $$
CREATE PROCEDURE verificar_suscribciones() 
begin 
    UPDATE suscripciones as s
    SET	monto_mensual = 0,
    	fecha_inicio = null,
        fecha_fin = null,
        plan_id = null,
    WHERE IFNULL(s.fecha_fin = null, DATE_SUB(NOW(),INTERVAL 32 DAY)) < NOW();
    SELECT "Realizado correctamente";
end
$$
delimiter ;






DROP PROCEDURE IF EXISTS `prueba`;
DELIMITER $$
CREATE PROCEDURE prueba()
BEGIN
DECLARE EXIT HANDLER FOR SQLEXCEPTION, SQLWARNING
    BEGIN
        
	    SELECT "Error en clase";
        ROLLBACK;
    END;
	START TRANSACTION;
		myclass:BEGIN 
        IF true THEN 
        	SELECT * from pagos WHERE null;
        	SELECT "mi clase";
        	LEAVE myclass; 
        END IF; 
        SELECT "Esto ya no";
	END;
END
$$
delimiter ;