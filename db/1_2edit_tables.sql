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

ALTER TABLE ordenes DROP FOREIGN KEY a8sf99SK80fsdff02l34gCR76G ;
-- ALTER table ordenes DROP COLUMN ordenes.payment_method_id;
ALTER TABLE ordenes ADD moneda varchar(8) NOT null DEFAULT "MXN"
ALTER TABLE item_orden ADD titulo varchar(8) null;
ALTER table ordenes ADD payment_method_id bigint;

ALTER table suscripciones ADD stripe_customer varchar(64);
ALTER table suscripciones ADD stripe_suscription varchar(64);
ALTER table suscripciones ADD stripe_paymenth varchar(64);
ALTER table planes ADD stripe_price varchar(64);
ALTER table planes ADD stripe_produc varchar(64);


ALTER table suscripciones DROP COLUMN id;
ALTER table suscripciones DROP COLUMN dia_corte;
ALTER table suscripciones DROP COLUMN next_plan_id;
ALTER table items_orden ADD COLUMN titulo varchar(64);

ALTER table pagos DROP COLUMN id;
