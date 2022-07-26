DROP VIEW IF EXISTS `eventos_videos`;
DROP VIEW IF EXISTS `plan_one`;
DROP VIEW IF EXISTS `plan_suscribcion`;

CREATE VIEW eventos_videos AS
SELECT 
	 eventos.id,eventos.activo, eventos.fechahora_evento,eventos.premio_cash,eventos.acumulado,eventos.premio_otros,eventos.moneda,
     eventos.categoria_evento_id,
     videos.artista,videos.canal,videos.fecha_video,videos.video_id,videos.thumblary,videos.titulo,videos.url_video,videos.genero,videos.proveedor
    FROM eventos 
    INNER JOIN videos ON eventos.video_id = videos.id 
    ORDER BY eventos.fechahora_evento;

CREATE VIEW plan_one AS
SELECT * FROM planes WHERE planes.suscribcion = FALSE;

CREATE VIEW plan_suscribcion AS
SELECT * FROM planes WHERE planes.suscribcion = TRUE;