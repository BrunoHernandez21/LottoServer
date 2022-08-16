DROP VIEW IF EXISTS `eventos_videos`;
DROP VIEW IF EXISTS `plan_one`;
DROP VIEW IF EXISTS `plan_suscribcion`;
DROP VIEW IF EXISTS `plan_suscripcion`;


DROP VIEW IF EXISTS `eventos_videos`;
CREATE VIEW eventos_videos AS
SELECT 
	 eventos.id, eventos.fechahora_evento,eventos.premio_cash,eventos.acumulado,eventos.premio_otros,eventos.moneda,eventos.categoria_evento_id,
     videos.id as vid_id, videos.artista,videos.canal,videos.fecha_video,videos.video_id,videos.thumblary,
     videos.titulo,videos.url_video,videos.genero,videos.proveedor
    FROM eventos 
    INNER JOIN videos ON eventos.video_id = videos.id 
    WHERE eventos.activo=true
    ORDER BY eventos.fechahora_evento;

DROP VIEW IF EXISTS `plan_one`;
CREATE VIEW plan_one AS
SELECT * FROM planes WHERE planes.suscribcion = FALSE;

DROP VIEW IF EXISTS `plan_suscripcion`;
CREATE VIEW plan_suscripcion AS
SELECT * FROM planes WHERE planes.suscribcion = TRUE;