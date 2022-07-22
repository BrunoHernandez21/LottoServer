DROP PROCEDURE IF EXISTS `pagos`;
delimiter $$
CREATE PROCEDURE pagos(car_id int) 
begin 
	DECLARE calto8a int; 
    DECLARE calto0a int;
    DECLARE cbajo8p int;
    DECLARE cbajoss int;
    DECLARE copor int;
	DECLARE palto8a int; 
    DECLARE palto0a int;
    DECLARE pbajo8p int;
    DECLARE pbajoss int;
    DECLARE popor int;
    
    DECLARE lan_id int;
    DECLARE user_id int;
    DECLARE canti int;
    
    SELECT lan_id = plan_id, user_id = usuario_id, canti = cantidad 
    from carrito 
    WHERE id = car_id;  
	##--optenemos valores actuales de cartera
    SELECT 	calto8a = acumulado_alto8am, 		calto0a = aproximacion_alta00am, 
    		cbajoss = aproximacion_baja, 		cbajo8p = acumulado_bajo8pm , 
            copor = oportunidades 
    	from carteras 
        WHERE usuario_id = user_id;
	##--optenemos valores del plan
    SELECT 	palto8a = acumulado_alto8am, 		palto0a = aproximacion_alta00am, 
    		pbajoss = aproximacion_baja, 		pbajo8p = acumulado_bajo8pm , 
            popor = oportunidades 
    	from planes 
        WHERE id = lan_id;
    ##--actualizamos cartera
    UPDATE carteras 
    SET   	acumulado_alto8am		= calto8a + (palto8a*canti),
    		aproximacion_alta00am	= calto0a + (palto0a*canti),
            acumulado_bajo8pm		= cbajo8p + (pbajo8p*canti),
            aproximacion_baja		= cbajoss + (pbajoss*canti),
            oportunidades			= copor   + (popor*canti)
            WHERE usuario_id = user_id;
    ##--actualizamos cartera
    UPDATE carrito 
    SET   	activo = false,
    		carrito.status = "finalizado"
            WHERE carrito.id = car_id;
    ## crear compra
    INSERT into compra VALUES (null,now(),user_id,carrito_id);
    ##--incertamos los beneficios de dias del plan
    INSERT INTO beneficios_usuario ( activo, fecha_inicio, fecha_fin,usuario_id,beneficio_id,plan_id) 
		SELECT true,now(),DATE_ADD(now(), INTERVAL beneficios.valor DAY),user_id,beneficios.id,lan_id from beneficios 
    	WHERE id IN (SELECT id from beneficios_plan WHERE plan_id = lan_id) AND beneficios.tipo = "DIAS";
	##--optenemos valores actuales de dinero del plan
	INSERT INTO beneficios_usuario ( activo, fecha_inicio, fecha_fin,usuario_id,beneficio_id,plan_id) 
		SELECT true,now(),now(),user_id,beneficios.id,lan_id from beneficios 
    	WHERE id IN (SELECT id from beneficios_plan WHERE plan_id = lan_id) AND beneficios.tipo = "CASH";
        
    
end
$$
delimiter ;







DROP PROCEDURE IF EXISTS `pruebas`;
delimiter $$
CREATE PROCEDURE pruebas(car_id bigint) 
begin 
    
    DECLARE lan_id int;
    DECLARE user_id int;
    DECLARE canti int;
    SELECT car_id;
    SELECT id from carrito WHERE id = 2;
    SELECT lan_id = plan_id, user_id = usuario_id, canti = cantidad 
    from carrito 
    WHERE id = car_id;  
	SELECT lan_id,user_id,canti;
        
end
$$
delimiter ;