ALTER TABLE pagos ADD is_error BOOLEAN  NOT NULL default false;
ALTER TABLE pagos ADD respuesta TEXT DEFAULT NULL;
ALTER TABLE pagos DROP stripe_id;

ALTER TABLE suscripciones ADD next_plan_id bigint DEFAULT NULL;
ALTER TABLE suscripciones DROP fecha_creado;
ALTER TABLE suscripciones DROP fecha_cobro;
                
UPDATE ordenes SET ordenes.status = "proceso" WHERE ordenes.status = "pendiente";
UPDATE ordenes SET ordenes.status = "pagado" WHERE ordenes.status = "pagada";
UPDATE ordenes SET ordenes.status = "pagado" WHERE ordenes.status = "cobrado";

ALTER TABLE ganador ADD cantidad_acumulada float DEFAULT NULL;
ALTER TABLE ordenes ADD is_suscription bool DEFAULT false NOT NULL;
DROP PROCEDURE IF EXISTS `genera_orden`;
DROP PROCEDURE IF EXISTS `pagos_rechazado`;


ALTER TABLE eventos DROP FOREIGN KEY FKk6e2a82e9uvkc8vrnijajf0c5;
ALTER TABLE eventos DROP categoria_evento_id;
ALTER TABLE eventos ADD costo INT NOT NULL default 0;
ALTER TABLE eventos ADD is_views BOOL NOT NULL default false;
ALTER TABLE eventos ADD is_like BOOL NOT NULL default false;
ALTER TABLE eventos ADD is_comments BOOL NOT NULL default false;
ALTER TABLE eventos ADD is_saved BOOL NOT NULL default false;
ALTER TABLE eventos ADD is_shared BOOL NOT NULL default false;
ALTER TABLE eventos ADD is_dislikes BOOL NOT NULL default false;
DROP TABLE IF EXISTS `categoria_evento`;

ALTER TABLE payment_method DROP cvc;
ALTER TABLE payment_method ADD cvc varchar(256) NOT NULL;