DROP VIEW IF EXISTS `eventos_videos`;
DROP VIEW IF EXISTS `plan_one`;
DROP VIEW IF EXISTS `plan_suscribcion`;
DROP VIEW IF EXISTS `plan_suscripcion`;


DROP VIEW IF EXISTS `eventos_videos`;
CREATE VIEW eventos_videos AS
SELECT 
	e.id, e.fechahora_evento,e.premio_cash,e.acumulado,e.premio_otros,e.moneda,
    e.costo,e.is_views,e.is_like,e.is_comments,e.is_saved,e.is_shared,e.is_dislikes,
    v.id as vid_id, v.artista,v.canal,v.fecha_video,v.video_id,v.thumblary,v.titulo,v.url_video,v.genero,v.proveedor
FROM eventos as e
INNER JOIN videos as v ON e.video_id = v.id 
WHERE e.activo=true and date(now())=date(fechahora_evento)
ORDER BY e.fechahora_evento;

DROP VIEW IF EXISTS `plan_one`;
CREATE VIEW plan_one AS
SELECT * FROM planes WHERE planes.suscribcion = FALSE;

DROP VIEW IF EXISTS `plan_suscripcion`;
CREATE VIEW plan_suscripcion AS
SELECT * FROM planes WHERE planes.suscribcion = TRUE;

DROP VIEW IF EXISTS `pagos_orden`;
CREATE VIEW pagos_orden AS
SELECT p.id, p.respuesta,p.fecha_pagado,p.is_error, o.status, o.fecha_emitido,  o.precio_total, o.puntos_total, o.usuario_id, o.payment_method_id
FROM pagos as p
LEFT JOIN ordenes as o ON p.orden_id = o.id;