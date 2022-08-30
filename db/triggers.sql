DROP TRIGGER IF EXISTS `default_user`;

delimiter $$
CREATE TRIGGER default_user AFTER INSERT ON usuarios FOR EACH ROW
begin 
	INSERT INTO carteras VALUES (null,0,0,0,NEW.id);
    INSERT INTO usuarios_roles VALUES (NEW.id,1);
    INSERT INTO propiedades_usuarios VALUES (null, 'default',null,now(),now(),NEW.id);
end
$$
delimiter ;



DROP TRIGGER IF EXISTS `beneficio_usuario`;

delimiter $$
CREATE TRIGGER default_user BEFORE INSERT ON usuarios FOR EACH ROW
begin 
	INSERT INTO carteras VALUES (null,0,0,0,0,0,NEW.id);
    INSERT INTO usuarios_roles VALUES (null,NEW.id,1);
end
$$
delimiter ;






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