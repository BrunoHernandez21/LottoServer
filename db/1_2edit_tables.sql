ALTER TABLE pagos ADD is_error BOOLEAN  NOT NULL default false;
ALTER TABLE pagos ADD respuesta TEXT DEFAULT NULL;
ALTER TABLE pagos DROP stripe_id;

ALTER TABLE suscripciones ADD next_plan_id bigint DEFAULT NULL;
ALTER TABLE suscripciones DROP fecha_creado;
ALTER TABLE suscripciones DROP fecha_cobro;
                