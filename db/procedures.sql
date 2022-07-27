DROP PROCEDURE IF EXISTS `pagos_realizado`;
delimiter $$
CREATE PROCEDURE pagos_realizado(IN iorden_id int, IN stripe_key varchar(60)) 
begin
    DECLARE pcash bigint;
    DECLARE ccash bigint;
    DECLARE user_id bigint;
    DECLARE resp varchar(40);
    IF ( (SELECT COUNT(id) from ordenes WHERE id = iorden_id) = 0 ) THEN 
        SET resp = "no exite la orden";
    ELSE
    ##--optenemos el usuario
    SET user_id = (SELECT  usuario_id from ordenes WHERE id = iorden_id);      
	##--optenemos valores actuales de cartera
    SET ccash = (SELECT	cash from carteras WHERE usuario_id = user_id);
	##--optenemos la suma total de pontos
    SET pcash = (SELECT  SUM(p.cash*i.cantidad) FROM items_orden i JOIN planes p ON i.plan_id = p.id WHERE i.orden_id = iorden_id);
    ##--actualizamos cartera
    UPDATE carteras SET cash = ccash + pcash WHERE usuario_id = user_id;
    ## crear compra
    INSERT INTO pagos (fecha_pagado,usuario_id,orden_id,stripe_id) VALUES (now(),user_id,iorden_id,stripe_key);
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "cobrado" WHERE ordenes.id = iorden_id;
    ##--incertamos los beneficios de dias del plan
    INSERT INTO beneficios_usuario ( activo, fecha_inicio, fecha_fin,usuario_id,beneficio_id,plan_id)
		SELECT 	true,
        		now(),
                DATE_ADD(now(), INTERVAL b.valor DAY),
                user_id,
                b.id,
                i.plan_id 
        from items_orden i 
        JOIN beneficios_plan bp ON i.plan_id = bp.plan_id 
    	JOIN beneficios b ON b.id = bp.beneficio_id 
        WHERE b.tipo = "DIAS" AND i.orden_id = iorden_id;
    ##--incertamos los beneficios de dias del plan
    INSERT INTO beneficios_usuario ( activo, fecha_inicio, fecha_fin,usuario_id,beneficio_id,plan_id)
		SELECT 	true,
        		now(),
                DATE_ADD(now(), INTERVAL b.valor DAY),
                user_id,
                b.id,
                i.plan_id 
        from items_orden i 
        JOIN beneficios_plan bp ON i.plan_id = bp.plan_id 
    	JOIN beneficios b ON b.id = bp.beneficio_id 
        WHERE b.tipo = "CASH" AND i.orden_id = iorden_id;
	SET resp = "Proceso realizado con exito";
    END IF;
    SELECT resp;
end
$$
delimiter ;



DROP PROCEDURE IF EXISTS `pagos_rechazado`;
delimiter $$
CREATE PROCEDURE pagos_rechazado(IN iorden_id int) 
begin
    ##--actualizamos la orden
    UPDATE ordenes SET	ordenes.status = "rechazado" WHERE ordenes.id = iorden_id;
	SELECT * from ordenes WHERE id = iorden_id;
end
$$
delimiter ;


DROP PROCEDURE IF EXISTS `genera_orden`;
delimiter $$
CREATE PROCEDURE genera_orden(IN user_id int) 
begin
    DECLARE time_Now datetime;
    DECLARE orden_id bigint;
    DECLARE resp varchar(40);
    set time_Now = now();
    IF ((SELECT count(id) from carrito WHERE carrito.usuario_id = user_id AND activo = TRUE)=0) THEN 
        SET resp = "carrito vacio";
    ELSE
        INSERT INTO ordenes ( ordenes.status, fecha_emitido, total,iva,descuento,total_iva,usuario_id)
            SELECT 	"proceso",
                    time_Now,
                    SUM(total_linea),
                    10 as iva,
                    10 as descuento,
                    (SUM(total_linea)*.90)*1.1,
                    user_id 
            from carrito 
            WHERE carrito.usuario_id = user_id AND activo = TRUE;

        SET orden_id = (SELECT  id from ordenes WHERE fecha_emitido = time_Now);
                
        INSERT INTO items_orden ( cantidad, total_linea, precio_unitario,descuento,plan_id,orden_id)
            SELECT 	cantidad,
                    total_linea,
                    precio_unitario,
                    descuento,
                    plan_id,
                    orden_id 
            from carrito 
            WHERE carrito.usuario_id = user_id AND activo = TRUE;

        UPDATE carrito SET activo = false WHERE carrito.usuario_id = user_id;
        SET resp = "finalizado correctamente";
    END IF;
    SELECT resp;
end
$$
delimiter ;
