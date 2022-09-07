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