DROP TRIGGER IF EXISTS `default_user`;

delimiter $$
CREATE TRIGGER default_user AFTER INSERT ON usuarios FOR EACH ROW
begin 
	INSERT INTO carteras(id,puntos,saldo_mxn,saldo_usd,usuario_id) VALUES (null,0,0,0,NEW.id);
    INSERT INTO usuarios_roles(user_id,role_id) VALUES (NEW.id,1);
    INSERT INTO propiedades_usuarios(id,nivel_acceso,custom_attributes,fecha_inicio,fecha_fin,usuario_id) VALUES (null, '3',null,now(),now(),NEW.id);
    INSERT INTO suscripciones(monto_mensual,suscripciones.usuario_id) VALUES (0,NEW.id);
end
$$
delimiter ;



DROP TRIGGER IF EXISTS `beneficio_core`;

delimiter $$
CREATE TRIGGER beneficio_core AFTER INSERT ON beneficios_usuario FOR EACH ROW
begin 
DECLARE caseType varchar(40);
DECLARE accesLevel varchar(40);
 SET caseType = (SELECT  tipo from beneficios WHERE id = NEW.beneficio_id);    
    CASE caseType
      WHEN "ACCES_LEVEL" THEN 
        SET accesLevel = (SELECT b.acces_id from beneficios_usuario as bu LEFT JOIN beneficios b ON bu.beneficio_id = b.id WHERE bu.id = NEW.id); 
        
        IF ( CAST(accesLevel AS UNSIGNED) < (SELECT CAST(nivel_acceso AS UNSIGNED) from propiedades_usuarios  WHERE usuario_id = NEW.usuario_id) ) THEN 
        UPDATE propiedades_usuarios 
        SET 
        	nivel_acceso = accesLevel,
        	fecha_inicio = NOW(),
            fecha_fin = DATE_ADD(now(), INTERVAL 1 MONTH)
        WHERE propiedades_usuarios.usuario_id = NEW.usuario_id;
        END IF;
      	
      WHEN "POINTS" THEN 
        UPDATE carteras 
        SET puntos = 
        	(SELECT puntos from carteras WHERE usuario_id = NEW.usuario_id) + 
            IFNULL((SELECT IFNULL(valor, 0) from beneficios WHERE id = NEW.beneficio_id),0)
            WHERE usuario_id = NEW.usuario_id;
      WHEN "CASH" THEN 
      	UPDATE carteras 
        SET saldo_mxn = 
        	(SELECT carteras.saldo_mxn from carteras WHERE usuario_id = NEW.usuario_id) + 
            IFNULL((SELECT IFNULL(valor,0) from beneficios WHERE id = NEW.beneficio_id AND moneda = "MXN" ),0)
        WHERE usuario_id = NEW.usuario_id;
        UPDATE carteras 
        SET carteras.saldo_usd = 
        	(SELECT carteras.saldo_usd from carteras WHERE usuario_id = NEW.usuario_id) + 
            IFNULL((SELECT valor from beneficios WHERE id = NEW.beneficio_id AND moneda = "USD" ),0)
            WHERE usuario_id = NEW.usuario_id;
      ELSE
        BEGIN
        END;
    END CASE;
    
end
$$
delimiter ;