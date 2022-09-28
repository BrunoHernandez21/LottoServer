DROP PROCEDURE IF EXISTS `genera_orden_unico`;
delimiter $$
CREATE PROCEDURE genera_orden_unico(IN user_id int,IN card_id int) 
begin
    DECLARE time_Now datetime;
    DECLARE orden_id bigint DEFAULT 0;
    IF ( (SELECT count(cantidad) from carrito WHERE usuario_id = user_id AND activo = TRUE) > 0 ) THEN 
    set time_Now = now();
	##--crea la orden 
  	INSERT INTO ordenes ( ordenes.status, fecha_emitido, precio_total,puntos_total,usuario_id,payment_method_id,ordenes.is_suscription)
        SELECT 	"proceso",
                time_Now,
                SUM(c.total_linea),
                SUM(c.puntos_linea),
                user_id,
                card_id,
                false
        from carrito as c
        JOIN planes p ON c.plan_id = p.id 
        WHERE c.usuario_id = user_id AND c.activo = TRUE AND p.suscribcion = FALSE;
	##-- Busca su id de orden
	SET orden_id = (SELECT  id from ordenes WHERE fecha_emitido = time_Now);
    ##-- crea los items de la orden            
    INSERT INTO items_orden ( cantidad, total_linea,puntos_linea,moneda,plan_id,orden_id)
        SELECT 	c.cantidad,
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
    END IF;
        ##-- imprimo una respuesta 
    SELECT * from ordenes WHERE id = orden_id;
end
$$
delimiter ;

DROP PROCEDURE IF EXISTS `genera_orden_suscribcion`;
delimiter $$
CREATE PROCEDURE genera_orden_suscribcion(IN user_id int,IN card_id int) 
begin
    DECLARE time_Now datetime;
    DECLARE orden_id bigint DEFAULT 0;
    IF ( (SELECT count(cantidad) from carrito WHERE usuario_id = user_id AND activo = TRUE) > 0 ) THEN 
    set time_Now = now();
	##--crea la orden 
  	INSERT INTO ordenes ( ordenes.status, fecha_emitido, precio_total,puntos_total,usuario_id,payment_method_id,is_suscription)
        SELECT 	"proceso",
                time_Now,
                SUM(c.total_linea),
                SUM(c.puntos_linea),
                user_id,
                card_id,
                true
        from carrito as c
        JOIN planes p ON c.plan_id = p.id 
        WHERE c.usuario_id = user_id AND c.activo = TRUE AND p.suscribcion = true;
	##-- Busca su id de orden
	SET orden_id = (SELECT  id from ordenes WHERE fecha_emitido = time_Now);
    ##-- crea los items de la orden            
    INSERT INTO items_orden ( cantidad, total_linea,puntos_linea,moneda,plan_id,orden_id)
        SELECT 	c.cantidad,
                c.total_linea,
                c.puntos_linea,
                c.moneda,
                c.plan_id,
                orden_id 
        from carrito as c
        JOIN planes p ON c.plan_id = p.id 
        WHERE c.usuario_id = user_id AND c.activo = TRUE AND p.suscribcion = true;
	##-- limpia el carrito
	UPDATE carrito SET activo = false WHERE carrito.usuario_id = user_id;
    END IF;
        ##-- imprimo una respuesta 
    SELECT * from ordenes WHERE id = orden_id;
end
$$
delimiter ;


##-- Determinando 
DROP PROCEDURE IF EXISTS `pago_unico`;
delimiter $$
CREATE PROCEDURE pago_unico(IN iorden_id int, IN razon TEXT) 
begin
    DECLARE pcash bigint;
    DECLARE ccash bigint;
    DECLARE user_id bigint;
    DECLARE susc_id bigint;
    DECLARE resp varchar(40);
    IF ( (SELECT COUNT(id) from ordenes WHERE id = iorden_id AND is_suscription = false) != 1 ) THEN 
        SET resp = "La orden no cumple con los requisitos";
    ELSE
    ##--optenemos el usuario
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);      
	##--optenemos valores actuales de cartera
    SET ccash = (SELECT	puntos from carteras WHERE usuario_id = user_id);
	##--optenemos la suma total de pontos
    SET pcash = (SELECT puntos_total FROM ordenes WHERE id = iorden_id);
    ##--actualizamos cartera
    UPDATE carteras SET puntos = ccash + pcash WHERE usuario_id = user_id;
    ## -- crear compra
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,false);
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "pagado" WHERE ordenes.id = iorden_id;
	##--incertamos los beneficios de dias del plan
    INSERT INTO beneficios_usuario ( cobrado, usuario_id,beneficio_id)
		SELECT 	true,
        		user_id,
                bp.beneficio_id
        from items_orden i 
        JOIN beneficios_plan bp ON i.plan_id = bp.plan_id 
        WHERE i.orden_id = iorden_id;
	##--incertamos los beneficios de dias del plan
    SET susc_id = (SELECT p.id from items_orden as i INNER JOIN planes AS p ON i.plan_id = p.id WHERE p.suscribcion = true AND i.orden_id = iorden_id LIMIT 1);
    UPDATE suscripciones 
    SET monto_mensual=(SELECT precio from planes WHERE id = susc_id),
        fecha_inicio=NOW(),
        fecha_fin=DATE_ADD(NOW(), INTERVAL 1 YEAR),
        dia_corte=EXTRACT(DAY FROM NOW()),
        plan_id=susc_id,
        next_plan_id=null
      WHERE suscripciones.usuario_id = user_id AND susc_id IS NOT NULL;

	SET resp = "Proceso realizado con exito";
    END IF;
    SELECT resp;
end
$$
delimiter ;



##-- listo 
DROP PROCEDURE IF EXISTS `pago_suscribcion`;
delimiter $$
CREATE PROCEDURE pago_suscribcion(IN iorden_id int, IN razon TEXT) 
begin
    DECLARE pcash bigint;
    DECLARE ccash bigint;
    DECLARE user_id bigint;
    DECLARE susc_id bigint;
    DECLARE resp varchar(40);
    IF ( (SELECT COUNT(id) from ordenes WHERE id = iorden_id AND is_suscription = TRUE) != 1 ) THEN 
        SET resp = "La orden no cumple con los requisitos";
    ELSE
    ##--optenemos el usuario
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);      
	##--optenemos valores actuales de cartera
    SET ccash = (SELECT	puntos from carteras WHERE usuario_id = user_id);
	##--optenemos la suma total de pontos
    SET pcash = (SELECT puntos_total FROM ordenes WHERE id = iorden_id);
    ##--actualizamos cartera
    UPDATE carteras SET puntos = ccash + pcash WHERE usuario_id = user_id;
    ## -- crear compra
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,false);
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "pagado" WHERE ordenes.id = iorden_id;
	##--incertamos los beneficios de dias del plan
    INSERT INTO beneficios_usuario ( cobrado, usuario_id,beneficio_id)
		SELECT 	true,
        		user_id,
                bp.beneficio_id
        from items_orden i 
        JOIN beneficios_plan bp ON i.plan_id = bp.plan_id 
        WHERE i.orden_id = iorden_id;
	##--incertamos los beneficios de dias del plan
    SET susc_id = (SELECT p.id from items_orden as i INNER JOIN planes AS p ON i.plan_id = p.id WHERE p.suscribcion = true AND i.orden_id = iorden_id LIMIT 1);
    UPDATE suscripciones 
    SET monto_mensual=(SELECT precio from planes WHERE id = susc_id),
        fecha_inicio=NOW(),
        fecha_fin=DATE_ADD(NOW(), INTERVAL 1 YEAR),
        dia_corte=EXTRACT(DAY FROM NOW()),
        plan_id=susc_id,
        next_plan_id=null
      WHERE suscripciones.usuario_id = user_id AND susc_id IS NOT NULL;

	SET resp = "Proceso realizado con exito";
    END IF;
    SELECT resp;
end
$$
delimiter ;

##-- Listo
DROP PROCEDURE IF EXISTS `pagos_rechazado`;
delimiter $$
CREATE PROCEDURE pagos_rechazado(IN iorden_id int,IN razon TEXT) 
begin
	DECLARE user_id bigint;
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);  
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "rechazado" WHERE ordenes.id = iorden_id;
    ## -- crear compra
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,respuesta,is_error) VALUES (now(),user_id,iorden_id,razon,true);
    ## -- out
	SELECT razon;
end
$$
delimiter ;


##-- pagos_cancelado
DROP PROCEDURE IF EXISTS `pagos_cancelado`;
delimiter $$
CREATE PROCEDURE pagos_cancelado(IN iorden_id int) 
begin 
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "cancelado" WHERE ordenes.id = iorden_id;
    ## -- out
	SELECT * from ordenes WHERE id = iorden_id;
end
$$
delimiter ;

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
    SELECT "Realizado correctamente";
end
$$
delimiter ;

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
        dia_corte = null,
        next_plan_id = null
    WHERE IFNULL(s.fecha_fin = null, DATE_SUB(NOW(),INTERVAL 31 DAY)) < NOW();
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


