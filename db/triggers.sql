DROP TRIGGER IF EXISTS `default_user`;

delimiter $$
CREATE TRIGGER default_user AFTER INSERT ON usuarios FOR EACH ROW
begin 
	INSERT INTO carteras VALUES (null,0,0,0,0,0,NEW.id);
    INSERT INTO usuarios_roles VALUES (null,NEW.id,1);
end
$$
delimiter ;


