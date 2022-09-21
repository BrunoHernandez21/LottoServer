  DROP TABLE IF EXISTS `usuarios_roles`;
  DROP TABLE IF EXISTS `roles`;
  DROP TABLE IF EXISTS `resultado`;
  DROP TABLE IF EXISTS `ganador`;
  DROP TABLE IF EXISTS `evento_usuario`;
  DROP TABLE IF EXISTS `eventos`;
  DROP TABLE IF EXISTS `tipo_evento`;
  DROP TABLE IF EXISTS `categoria_evento`;
  DROP TABLE IF EXISTS `videos_estadisticas`;
  DROP TABLE IF EXISTS `videos`;
  DROP TABLE IF EXISTS `suscripciones`;
  DROP TABLE IF EXISTS `beneficios_plan`;
  DROP TABLE IF EXISTS `beneficios_usuario`;
  DROP TABLE IF EXISTS `beneficios_referido`;
  DROP TABLE IF EXISTS `beneficios`;
  DROP TABLE IF EXISTS `items_orden`;
  DROP TABLE IF EXISTS `pagos`;
  DROP TABLE IF EXISTS `ordenes`;
  DROP TABLE IF EXISTS `carrito`;
  DROP TABLE IF EXISTS `planes`;
  DROP TABLE IF EXISTS `payment_method`;
  DROP TABLE IF EXISTS `carteras`;
  DROP TABLE IF EXISTS `referido`;
  DROP TABLE IF EXISTS `direccion`;
  DROP TABLE IF EXISTS `propiedades_usuarios`;
  DROP TABLE IF EXISTS `acces`;
  DROP TABLE IF EXISTS `level`;
  DROP TABLE IF EXISTS `usuarios`;  

##-- Table structure for table `usuarios`
CREATE TABLE `usuarios` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL default true,
  `apellidom` varchar(255) DEFAULT NULL,
  `apellidop` varchar(255) DEFAULT NULL,
  `email` varchar(127) NOT NULL,
  `fecha_nacimiento` DATETIME DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `telefono` varchar(255) DEFAULT NULL,
  `codigo_referido` varchar(60) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_kfsp0s1tflm1cwlj8idhqsad0` (`email`),
  UNIQUE KEY `UK_kfsp0tl78lm1cwlj8idhqsad0` (`codigo_referido`)
);
LOCK TABLES `usuarios` WRITE;
INSERT INTO `usuarios` VALUES 
(1,true,NULL,NULL,'mezagg@gmail.com',NULL,NULL,'$2a$10$59tlZW6RvpCSnPwfKGxpR.55WwSGMQRi9Gq.2D43Nd8tZcxvQbt02',NULL,"2Pk6D80@&c"),
(2,true,NULL,NULL,'ichimar21@gmail.com',NULL,NULL,'93ac34edd54c8c52ea15b6026e50435bf2e6a91f',NULL,"2VksD8o@\c");
UNLOCK TABLES;
##-- Table structure for table `level`
CREATE TABLE `level` (
  `level` varchar(64) Not null,
  UNIQUE KEY `UK_levelAccesApp0001` (`level`)
);
LOCK TABLES `level` WRITE;
INSERT INTO `level` VALUES ('0'),('1'),('2'),('3');
UNLOCK TABLES;

##-- Table structure for table `acces`
CREATE TABLE `acces` (
  `level_id` varchar(64) Not null,
  `description` varchar(64) Not null,
  CONSTRAINT `UK_accesLevelrelacion0001` FOREIGN KEY (`level_id`) REFERENCES `level` (`level`) ON DELETE CASCADE
);
LOCK TABLES `acces` WRITE;
INSERT INTO `acces` VALUES 
('0','estadisticas'),
('0','estadisticas_long'),
('1','estadisticas'),
('2','nada'),
('3','nada');
UNLOCK TABLES;

##-- Table structure for table `usuarios`
CREATE TABLE `propiedades_usuarios` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nivel_acceso` varchar(60) DEFAULT '3',
  `custom_attributes` varchar(1024) DEFAULT null,
  `fecha_inicio` DATETIME DEFAULT NULL,
  `fecha_fin` DATETIME DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK_key_for_property_user` (`usuario_id`),
  KEY `UK_usuario_nivel00001` (`nivel_acceso`),
  CONSTRAINT `UK_usuario_nivel00001` FOREIGN KEY (`nivel_acceso`) REFERENCES `level` (`level`) ON DELETE CASCADE,
  CONSTRAINT `FK_key_for_property_user` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `direccion`
CREATE TABLE `direccion` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `usuario_id` bigint DEFAULT NULL,
  `tipo` varchar(255) DEFAULT NULL,
  `pais` varchar(255) DEFAULT NULL,
  `ciudad` varchar(255) DEFAULT NULL,
  `calle` varchar(255) DEFAULT NULL,
  `cp` varchar(255) DEFAULT NULL,
  `numero` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vrnijaj87yt` (`usuario_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vrnijaj87yt` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `direccion`
CREATE TABLE `referido` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `usuario_id` bigint DEFAULT NULL,
  `codigo` varchar(255) DEFAULT NULL,
  `cobrado` BOOLEAN DEFAULT false,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vr88jaj87yt` (`usuario_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vr88jaj87yt` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `carteras`
CREATE TABLE `carteras` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `puntos` int NOT NULL DEFAULT 0,
  `saldo_mxn` int NOT NULL DEFAULT 0,
  `saldo_usd` int NOT NULL DEFAULT 0,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKm5oo9iahtl1p9bs4dn1ymovlb` (`usuario_id`),
  CONSTRAINT `FKm5oo9iahtl1p9bs4dn1ymovlb` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);
LOCK TABLES `carteras` WRITE;
INSERT INTO `carteras` VALUES (1,0,0,0,1),(2,0,0,0,2);
UNLOCK TABLES;

##-- Table structure for table `payment_method`
CREATE TABLE `payment_method` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN  NOT NULL default true ,
  `type` varchar(255) DEFAULT NULL,
  `sub_type` varchar(255) DEFAULT NULL,
  `card_number` varchar(255) DEFAULT NULL,
  `cvc` int NOT NULL,
  `default_payment` BOOLEAN NOT NULL,
  `expiry_month` int NOT NULL,
  `expiry_year` int NOT NULL,
  `holder_name` varchar(255) DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKdfsg45Tthfi8r4DF43gfd34hF` (`usuario_id`),
  CONSTRAINT `FKdfsg45Tthfi8r4DF43gfd34hF` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `planes`
CREATE TABLE `planes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN   NOT NULL default true,
  `titulo` varchar(60) DEFAULT NULL,
  `descripcion` varchar(256) DEFAULT NULL,
  `pre_puntos` int DEFAULT NULL,
  `puntos` int DEFAULT NULL,
  `pre_precio` double DEFAULT 0,
  `precio` double NOT NULL DEFAULT 0,
  `moneda` varchar(8),
  `suscribcion` BOOLEAN NOT NULL DEFAULT false,
  PRIMARY KEY (`id`)
);
LOCK TABLES `planes` WRITE;
INSERT INTO `planes` VALUES 
(1,true,'Ordinario'   ,'Algo Sencillo'   ,null,4 ,null,70 ,'MXN',false),
(2,true,'Promocional' ,'Un buen comienzo',null,20,300,250 ,'MXN',false),
(3,true,'Platinum'    ,'Para Pro players',40  ,50,null,500,'MXN',false),
(4,true,'Ordinario'   ,'Algo Sencillo'   ,null,4 ,null,70 ,'MXN',true),
(5,true,'Promocional' ,'Un buen comienzo',null,20,300,250 ,'MXN',true),
(6,true,'Platinum'    ,'Para Pro players',40  ,50,null,500,'MXN',true);
UNLOCK TABLES;

##-- Table structure for table `carrito`
CREATE TABLE `carrito` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN  NOT NULL default true,
  `cantidad` int NOT NULL DEFAULT 1,
  `precio_unitario` float DEFAULT NULL,
  `total_linea` double NOT NULL DEFAULT 0,
  `puntos_unitario` float DEFAULT NULL,
  `puntos_linea` double NOT NULL DEFAULT 0,
  `moneda` varchar(10),
  `fecha_carrito` DATETIME DEFAULT NULL,
  `plan_id` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKdfsg45Tth5345gDF43gfd34hF` (`usuario_id`),
  KEY `a8sf23R980fsdf09er234gE7R6G` (`plan_id`),
  CONSTRAINT `FKdfsg45Tth5345gDF43gfd34hF` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE,
  CONSTRAINT `a8sf23R980fsdf09er234gE7R6G` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`) ON DELETE CASCADE
) ;

##-- Table structure for table `orden`
CREATE TABLE `ordenes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `status` varchar(255) DEFAULT NULL,        ## --  proceso, rechazado, pagado, cancelado,
  `fecha_emitido` DATETIME DEFAULT NULL,
  `precio_total` double NOT NULL DEFAULT 0,
  `puntos_total` double NOT NULL DEFAULT 0,
  `usuario_id` bigint DEFAULT NULL,
  `payment_method_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `a8sf99SK80fsdff02l34gE7R6G` (`usuario_id`),
  KEY `a8sf99SK80fsdff02l34gCR76G` (`payment_method_id`),
  CONSTRAINT `a8sf99SK80fsdff02l34gE7R6G` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE,
  CONSTRAINT `a8sf99SK80fsdff02l34gCR76G` FOREIGN KEY (`payment_method_id`) REFERENCES `payment_method` (`id`) ON DELETE CASCADE
) ;
##-- Table structure for table `items_orden`
CREATE TABLE `items_orden` (
  `cantidad` int NOT NULL DEFAULT 1,
  `total_linea` double NOT NULL DEFAULT 0,
  `puntos_linea` double NOT NULL DEFAULT 0,
  `moneda` varchar(10),
  `plan_id` bigint DEFAULT NULL,
  `orden_id` bigint DEFAULT NULL,
  KEY `FKpsPlankey0000000000000004` (`plan_id`),
  KEY `FK1yOrdenKey000000000000004` (`orden_id`),
  CONSTRAINT `FKpsPlankey0000000000000004` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FK1yOrdenKey000000000000004` FOREIGN KEY (`orden_id`) REFERENCES `ordenes` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `pago`
CREATE TABLE `pagos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `is_error` BOOLEAN  NOT NULL default false,
  `respuesta` TEXT DEFAULT NULL,
  `fecha_pagado` DATETIME DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `orden_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKdfsg45Tth5666pDF43gfd34hF` (`usuario_id`),
  KEY `a8sf99SK80fsdf09er234gE7R6G` (`orden_id`),
  CONSTRAINT `FKdfsg45Tth5666pDF43gfd34hF` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE,
  CONSTRAINT `a8sf99SK80fsdf09er234gE7R6G` FOREIGN KEY (`orden_id`) REFERENCES `ordenes` (`id`) ON DELETE CASCADE
) ;

##-- Table structure for table `beneficios`
CREATE TABLE `beneficios` (
  `id` bigint AUTO_INCREMENT,
  `llave` varchar(255) DEFAULT NULL, ##-- titulo
  `tipo` varchar(255) DEFAULT NULL,
  `moneda` varchar(16) DEFAULT NULL,
  `valor` float DEFAULT NULL,
  `dias` int DEFAULT NULL,
  `acces_id` varchar(64) DEFAULT NULL,
  `max_get` int DEFAULT NULL,
  PRIMARY KEY (`id`)
);

LOCK TABLES `beneficios` WRITE;
INSERT INTO `beneficios` VALUES 
(1,"7 dias de estadisticas" ,"ACCES_LEVEL",NULL ,null  ,7   ,'0' ,null),
(2,"30 dias de estadisticas","ACCES_LEVEL",NULL ,null  ,30  ,'0' ,null),
(3,"31 dias de estadisticas","ACCES_LEVEL",NULL ,null  ,31  ,'0' ,null),
(4,"Puntos extra"           ,"POINTS"     ,NULL ,4     ,null,null,null),
(5,"Puntos extra"           ,"POINTS"     ,NULL ,20    ,null,null,null),
(6,"Dinero Gratis"          ,"CASH"       ,"MXN",10    ,null,null,null),
(7,"Dinero Gratis"          ,"CASH"       ,"USD",10    ,null,null,null);
UNLOCK TABLES;

##-- Table structure for table `beneficios_referido`
CREATE TABLE `beneficios_referido` (
  `activo` BOOLEAN NOT NULL default false,
  `beneficio_id` bigint NOT NULL,
  KEY `FK_beneficio_referido0001` (`beneficio_id`),
  CONSTRAINT `FK_beneficio_referido0001` FOREIGN KEY (`beneficio_id`) REFERENCES `beneficios` (`id`) ON DELETE CASCADE
);
LOCK TABLES `beneficios_referido` WRITE;
INSERT INTO `beneficios_referido` VALUES (true,3),(true,6);
UNLOCK TABLES;

##-- Table structure for table `beneficios_usuario`
CREATE TABLE `beneficios_usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `cobrado` bigint NOT NULL default false,
  `usuario_id` bigint NOT NULL,
  `beneficio_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FKpsfgo6447835hkqudyubw5536` (`usuario_id`),
  KEY `FKpsfgo6986335hkqudyubw5536` (`beneficio_id`),
  CONSTRAINT `FKpsfgo6447835hkqudyubw5536` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKpsfgo6986335hkqudyubw5536` FOREIGN KEY (`beneficio_id`) REFERENCES `beneficios` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `beneficios_plan`
CREATE TABLE `beneficios_plan` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL default false,
  `beneficio_id` bigint NOT NULL,
  `plan_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `6789s0985thfi8r4DF43gfd34hF` (`beneficio_id`),
  KEY `6789sg45Tth44334DF43gfd34hF` (`plan_id`),
  CONSTRAINT `6789s0985thfi8r4DF43gfd34hF` FOREIGN KEY (`beneficio_id`) REFERENCES `beneficios` (`id`) ON DELETE CASCADE,
  CONSTRAINT `6789sg45Tth44334DF43gfd34hF` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`) ON DELETE CASCADE
);
LOCK TABLES `beneficios_plan` WRITE;
INSERT INTO `beneficios_plan` VALUES 
(1,true,1,1),
(2,true,1,2),
(3,true,1,3),
(4,true,2,2),
(5,true,3,3),
(6,true,3,4),
(7,true,3,5),
(8,true,3,6),
(9,true,1,4),
(10,true,1,5),
(11,true,1,6);
UNLOCK TABLES;

##-- Table structure for table `suscripciones`
CREATE TABLE `suscripciones` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `monto_mensual` float NOT NULL,
  `fecha_inicio` DATETIME DEFAULT NULL,
  `fecha_fin` DATETIME DEFAULT NULL,
  `dia_corte` int DEFAULT NULL,
  `plan_id` bigint DEFAULT NULL,
  `next_plan_id` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK6fcib1tyjrhp8u95q3uhohqc6` (`plan_id`),
  KEY `FK6fcib1tyjrhp8u95q3uhohqc7` (`plan_id`),
  KEY `FKh7go9iahtl5u5bs4dn1ymovlb` (`usuario_id`),
  CONSTRAINT `FK6fcib1tyjrhp8u95q3uhohqc6` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FK6fcib1tyjrhp8u95q3uhohqc7` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKh7go9iahtl5u5bs4dn1ymovlb` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);

##--videos
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN  DEFAULT true,
  `artista` varchar(255) DEFAULT NULL,
  `canal` varchar(255) DEFAULT NULL,
  `fecha_video` DATETIME DEFAULT NULL,
  `video_id` varchar(255) DEFAULT NULL,
  `thumblary` varchar(255) DEFAULT NULL,
  `titulo` varchar(255) DEFAULT NULL,
  `url_video` varchar(255) DEFAULT NULL,
  `genero` varchar(255) DEFAULT NULL,
  `proveedor` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);
LOCK TABLES `videos` WRITE;
INSERT INTO `videos` VALUES 
(1,1,'BANDA MS','Lizos Music','2020-10-20 00:00:00.000000','RgxM1Rv7hCs','https://i.ytimg.com/vi/RgxM1Rv7hCs/default.jpg','CERRANDO CICLOS (VIDEO OFICIAL)','https://www.youtube.com/watch?v=RgxM1Rv7hCs','Pop','Youtube'),
(2,1,'BANDA MS','Lizos Music','2020-02-28 00:00:00.000000','ova8TGDNvCo','https://i.ytimg.com/vi/ova8TGDNvCo/default.jpg','QUIÉN PIERDE MÁS (VIDEO OFICIAL)','https://www.youtube.com/watch?v=ova8TGDNvCo','Salsa','Youtube'),
(3,1,'BANDA MS','Lizos Music','2018-02-13 00:00:00.000000','2mf1Os7dAJI','https://i.ytimg.com/vi/2mf1Os7dAJI/default.jpg','SI CRUZAS LA PUERTA (LETRA)','https://www.youtube.com/watch?v=2mf1Os7dAJI','Cumbia','Youtube'),
(4,1,'Christian Nodal','Christian Nodal','2019-12-05 00:00:00.000000','Ax3psz01Q8o','https://i.ytimg.com/vi/Ax3psz01Q8o/default.jpg','Si Te Falta Alguien (Video Oficial)','https://www.youtube.com/watch?v=Ax3psz01Q8o','Banda','Youtube'),
(5,1,'Carin Leon','TAMARINDOREKORDSZ','2021-12-12 00:00:00.000000','8Bznc5tTQ9M','https://i.ytimg.com/vi/8Bznc5tTQ9M/default.jpg','Si Una Vez','https://www.youtube.com/watch?v=8Bznc5tTQ9M','Salsa','Youtube'),
(6,1,'Grupo Firme','Grupo Firme','2020-07-24 00:00:00.000000','g2BzGJnNvEw','https://i.ytimg.com/vi/g2BzGJnNvEw/default.jpg',' Ya Superame','https://www.youtube.com/watch?v=g2BzGJnNvEw','Banda','Youtube');
UNLOCK TABLES;

##--videos
CREATE TABLE `videos_estadisticas` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `video_id` bigint NOT NULL,
  `fecha` DATETIME DEFAULT NULL,  
  `views_count` bigint DEFAULT NULL,
  `like_count` bigint DEFAULT NULL,
  `comments_count` bigint DEFAULT NULL,
  `dislikes_count` bigint DEFAULT NULL,
  `saved_count` bigint DEFAULT NULL,
  `shared_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKs87xk1t7ytkg1xw9doomybg6m` (`video_id`),
  CONSTRAINT `FKs87xk1t7ytkg1xw9doomybg6m` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
);

##-- Table structure for table `categoria_evento`
CREATE TABLE `categoria_evento` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  `costo` INT NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`)
);
LOCK TABLES `categoria_evento` WRITE;
INSERT INTO `categoria_evento` VALUES (1,'Oportunidades',1),(2,'Acumulado_alto8am',3),(3,'Acumulado_bajo8pm',1),(4,'aproximacion_alta00am',2),(5,'aproximacion_baja',1);
UNLOCK TABLES;

##-- Table structure for table `tipo_evento`
##-- Esta tabla es meramente informativa
CREATE TABLE `tipo_evento` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);
LOCK TABLES `tipo_evento` WRITE;
INSERT INTO `tipo_evento` VALUES 
(1,'Views'),
(2,'Like'),
(3,'Comments'),
(4,'Dislikes'),
(5,'Saved'),
(6,'Shared');
UNLOCK TABLES;

##-- Table structure for table `eventos`
CREATE TABLE `eventos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN  DEFAULT true,
  `fechahora_evento` DATETIME DEFAULT NULL,
  `premio_cash` float DEFAULT NULL,
  `acumulado` float DEFAULT NULL,
  `premio_otros` varchar(255) DEFAULT NULL,
  `moneda` varchar(255) DEFAULT NULL,
  `categoria_evento_id` bigint DEFAULT NULL,
  `video_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vrnijajf0c5` (`categoria_evento_id`),
  KEY `FKs87xk1t7ytkg1xw91sntybg6m` (`video_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vrnijajf0c5` FOREIGN KEY (`categoria_evento_id`) REFERENCES `categoria_evento` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKs87xk1t7ytkg1xw91sntybg6m` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`) ON DELETE CASCADE
);
LOCK TABLES `eventos` WRITE;
INSERT INTO `eventos` VALUES  
(1,false,'2022-06-14 00:00:00.000000',null,0,null,'USD',4,6),
(2,true,'2022-06-14 00:30:00.000000',5000,0,null,'USD',5,5),
(3,true,'2022-06-14 01:00:00.000000',5000,0,null,'USD',5,4),
(4,true,'2022-06-14 01:30:00.000000',5000,0,null,'USD',5,3),
(5,true,'2022-06-14 02:00:00.000000',5000,0,null,'USD',5,2),
(6,true,'2022-06-14 02:30:00.000000',5000,0,null,'USD',5,1),
(7,true,'2022-06-14 03:00:00.000000',null,0,'iPhone 13X ','USD',1,1),
(8,true,'2022-06-14 03:30:00.000000',5000,0,null,'USD',5,6),
(9,true,'2022-06-14 04:00:00.000000',5000,0,null,'USD',5,5),
(10,true,'2022-06-14 04:30:00.000000',5000,0,null,'USD',5,4),
(11,true,'2022-06-14 05:00:00.000000',5000,0,null,'USD',5,3),
(12,true,'2022-06-14 05:30:00.000000',5000,0,null,'USD',5,2),
(13,true,'2022-06-14 06:00:00.000000',null,0,'iPhone 13X','USD',1,1),
(14,true,'2022-06-14 06:30:00.000000',5000,0,null,'USD',5,6),
(15,true,'2022-06-14 07:00:00.000000',5000,0,null,'USD',5,6),
(16,true,'2022-06-14 07:30:00.000000',5000,0,null,'USD',5,6),
(17,true,'2022-06-14 08:00:00.000000',100000,0,'$100,000','USD',2,6),
(18,true,'2022-06-14 08:30:00.000000',5000,0,null,'USD',5,6),
(19,true,'2022-06-14 09:00:00.000000',5000,0,null,'USD',5,6),
(20,true,'2022-06-14 09:30:00.000000',null,0,'iPhone 13X','USD',1,6),
(21,true,'2022-06-14 10:00:00.000000',5000,0,null,'USD',5,6),
(22,true,'2022-06-14 10:30:00.000000',5000,0,null,'USD',5,6),
(23,true,'2022-06-14 11:00:00.000000',5000,0,null,'USD',5,6),
(24,true,'2022-06-14 11:30:00.000000',5000,0,null,'USD',5,6),
(25,false,'2022-06-14 12:00:00.000000',null,0,'Carrito','USD',3,6),
(26,true,'2022-06-14 12:30:00.000000',5000,0,null,'USD',5,6),
(27,true,'2022-06-14 13:00:00.000000',5000,0,null,'USD',5,6),
(28,true,'2022-06-14 13:30:00.000000',5000,0,null,'USD',5,6),
(29,true,'2022-06-14 14:00:00.000000',5000,0,null,'USD',5,6),
(30,true,'2022-06-14 14:30:00.000000',5000,0,null,'USD',5,6),
(31,true,'2022-06-14 15:00:00.000000',null,0,'iPhone 13X','USD',1,6),
(32,true,'2022-06-14 15:30:00.000000',5000,0,null,'USD',5,6),
(33,true,'2022-06-14 16:00:00.000000',5000,0,null,'USD',5,6),
(34,true,'2022-06-14 16:30:00.000000',5000,0,null,'USD',5,6),
(35,true,'2022-06-14 17:00:00.000000',5000,0,null,'USD',5,6),
(36,true,'2022-06-14 17:30:00.000000',5000,0,null,'USD',5,6),
(37,true,'2022-06-14 18:00:00.000000',null,0,'iPhone 13X','USD',1,6),
(38,true,'2022-06-14 18:30:00.000000',5000,0,null,'USD',5,6),
(39,true,'2022-06-14 19:00:00.000000',5000,0,null,'USD',5,6),
(40,true,'2022-06-14 19:30:00.000000',5000,0,null,'USD',5,6),
(41,true,'2022-06-14 20:00:00.000000',70000,0,'$70,000','USD',3,6),
(42,true,'2022-06-14 20:30:00.000000',5000,0,null,'USD',5,6),
(43,true,'2022-06-14 21:00:00.000000',5000,0,null,'USD',5,6),
(44,true,'2022-06-14 21:30:00.000000',null,0,'iPhone 13X','USD',1,6),
(45,true,'2022-06-14 22:00:00.000000',5000,0,null,'USD',5,6),
(46,true,'2022-06-14 22:30:00.000000',5000,0,null,'USD',5,6),
(47,true,'2022-06-14 23:00:00.000000',5000,0,null,'USD',5,6),
(48,true,'2022-06-14 23:30:00.000000',5000,0,null,'USD',5,6);
UNLOCK TABLES;

##-- Table structure for table `evento_usuario`
CREATE TABLE `evento_usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN  NOT NULL,
  `fecha` DATETIME DEFAULT NULL,
  `views_count` bigint DEFAULT NULL,
  `like_count` bigint DEFAULT NULL,
  `comments_count` bigint DEFAULT NULL,
  `dislikes_count` bigint DEFAULT NULL,
  `saved_count` bigint DEFAULT NULL,
  `shared_count` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `evento_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKcj7cb72y8k0doxs7jajmwt0q2` (`evento_id`),
  KEY `FKlln41mrxef4w9oomu5rcbnikj` (`usuario_id`),
  CONSTRAINT `FKcj7cb72y8k0doxs7jajmwt0q2` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKlln41mrxef4w9oomu5rcbnikj` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `ganador`
CREATE TABLE `ganador` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `cantidad` bigint DEFAULT NULL,
  `concepto` varchar(255) DEFAULT NULL,
  `evento_id` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `evento_usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKcj7cb98dfk0do0092ajmwt0q2` (`evento_id`),
  KEY `FKlln4109plf4w9ooal93cbnikj` (`usuario_id`),
  KEY `FKlln41mrxeld95oom0vm2bnikj` (`evento_usuario_id`),
  CONSTRAINT `FKcj7cb98dfk0do0092ajmwt0q2` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKlln4109plf4w9ooal93cbnikj` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKlln41mrxeld95oom0vm2bnikj` FOREIGN KEY (`evento_usuario_id`) REFERENCES `evento_usuario` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `resultado`
CREATE TABLE `resultado` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `hora_resultado` DATETIME DEFAULT NULL,
  `views_count` bigint DEFAULT NULL,
  `like_count` bigint DEFAULT NULL,
  `comments_count` bigint DEFAULT NULL,
  `dislikes_count` bigint DEFAULT NULL,
  `saved_count` bigint DEFAULT NULL,
  `shared_count` bigint DEFAULT NULL,
  `video_id` varchar(255) DEFAULT NULL,
  `evento_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `16ffsg45Tthfi8r4DF43gfd34hF` (`evento_id`),
  CONSTRAINT `16ffsg45Tthfi8r4DF43gfd34hF` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`) ON DELETE CASCADE
);

##-- Table structure for table `roles`
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_ldv0v52e0udsh2h1rs0r0gw1n` (`nombre`)
) ;
LOCK TABLES `roles` WRITE;
INSERT INTO `roles` VALUES (2,'ROLE_ADMIN'),(1,'ROLE_USER');
UNLOCK TABLES;

##-- Table structure for table `usuarios_roles`
CREATE TABLE `usuarios_roles` (
  `user_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  KEY `FKisd054ko30hm3j6ljr90asype` (`user_id`),
  KEY `FKihom0uklpkfpffipxpoyf7b74` (`role_id`),
  CONSTRAINT `FKihom0uklpkfpffipxpoyf7b74` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`) ON DELETE CASCADE,
  CONSTRAINT `FKisd054ko30hm3j6ljr90asype` FOREIGN KEY (`user_id`) REFERENCES `usuarios` (`id`) ON DELETE CASCADE
);
LOCK TABLES `usuarios_roles` WRITE;
INSERT INTO `usuarios_roles` VALUES (1,2),(2,2);
UNLOCK TABLES;