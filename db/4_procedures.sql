DROP PROCEDURE IF EXISTS `genera_orden`;
delimiter $$
CREATE PROCEDURE genera_orden(IN user_id int,IN card_id int) 
begin
    DECLARE time_Now datetime;
    DECLARE orden_id bigint DEFAULT 0;
    IF ( (SELECT count(cantidad) from carrito WHERE usuario_id = user_id AND activo = TRUE) > 0 ) THEN 
    set time_Now = now();
	##--crea la orden 
  	INSERT INTO ordenes ( ordenes.status, fecha_emitido, precio_total,puntos_total,usuario_id,payment_method_id)
        SELECT 	"proceso",
                time_Now,
                SUM(total_linea),
                SUM(puntos_linea),
                user_id,
                card_id
        from carrito 
        WHERE carrito.usuario_id = user_id AND activo = TRUE;
	##-- Busca su id de orden
	SET orden_id = (SELECT  id from ordenes WHERE fecha_emitido = time_Now);
    ##-- crea los items de la orden            
    INSERT INTO items_orden ( cantidad, total_linea,puntos_linea,moneda,plan_id,orden_id)
        SELECT 	cantidad,
                carrito.total_linea,
                carrito.puntos_linea,
                carrito.moneda,
                plan_id,
                orden_id 
        from carrito 
        WHERE carrito.usuario_id = user_id AND activo = TRUE;
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
    IF ( (SELECT COUNT(id) from ordenes WHERE id = iorden_id) = 0 ) THEN 
        SET resp = "no exite la orden";
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
    SELECT "Realizado correctamente";
end
$$
delimiter ;

##-- verificar_suscribciones
DROP PROCEDURE IF EXISTS `verificar_suscribciones`;
delimiter $$
CREATE PROCEDURE verificar_suscribciones() 
begin 
    SELECT "Realizado correctamente";
end
$$
delimiter ;

##-- verificar_propiedades_usuario
DROP PROCEDURE IF EXISTS `verificar_propiedades_usuario`;
delimiter $$
CREATE PROCEDURE verificar_propiedades_usuario() 
begin 
    SELECT "Realizado correctamente";
end
$$
delimiter ;