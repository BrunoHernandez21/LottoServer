##-- DROP ALL TABLES
##-- Legado
##-- DROP TABLE IF EXISTS `usuarios_roles`;
##-- DROP TABLE IF EXISTS `suscripciones`;
##-- DROP TABLE IF EXISTS `roles`;
##-- DROP TABLE IF EXISTS `results`;
##-- DROP TABLE IF EXISTS `resultado`;
##-- DROP TABLE IF EXISTS `pago`;
##-- DROP TABLE IF EXISTS `orden`;
##-- DROP TABLE IF EXISTS `hibernate_sequence`;
##-- DROP TABLE IF EXISTS `ganador`;
##-- DROP TABLE IF EXISTS `cron_task`;
##-- DROP TABLE IF EXISTS `compra`;
##-- DROP TABLE IF EXISTS `plan`;
##-- DROP TABLE IF EXISTS `carteras`;
##-- DROP TABLE IF EXISTS `apuesta_usuario`;
##-- DROP TABLE IF EXISTS `apuestas`;
##-- DROP TABLE IF EXISTS `tipo_apuesta`;
##-- DROP TABLE IF EXISTS `categoria_apuesta`;
##-- DROP TABLE IF EXISTS `videos`;
##-- DROP TABLE IF EXISTS `usuarios`;
##-- Actuales
DROP TABLE IF EXISTS `usuarios_roles`;
DROP TABLE IF EXISTS `roles`;
DROP TABLE IF EXISTS `resultado`;
DROP TABLE IF EXISTS `hibernate_sequence`;
DROP TABLE IF EXISTS `ganador`;
DROP TABLE IF EXISTS `evento_usuario`;
DROP TABLE IF EXISTS `cron_task`;
DROP TABLE IF EXISTS `eventos`;
DROP TABLE IF EXISTS `tipo_evento`;
DROP TABLE IF EXISTS `categoria_evento`;
DROP TABLE IF EXISTS `videos`;
DROP TABLE IF EXISTS `suscripciones`;
DROP TABLE IF EXISTS `beneficios_plan`;
DROP TABLE IF EXISTS `beneficios_usuario`;
DROP TABLE IF EXISTS `beneficios`;
DROP TABLE IF EXISTS `compra`;
DROP TABLE IF EXISTS `carrito`;
DROP TABLE IF EXISTS `planes`;
DROP TABLE IF EXISTS `payment_method`;
DROP TABLE IF EXISTS `carteras`;
DROP TABLE IF EXISTS `referido`;
DROP TABLE IF EXISTS `direccion`;
DROP TABLE IF EXISTS `usuarios`;

##-- Table structure for table `usuarios`
CREATE TABLE `usuarios` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `apellidom` varchar(255) DEFAULT NULL,
  `apellidop` varchar(255) DEFAULT NULL,
  `email` varchar(127) NOT NULL,
  `fecha_nacimiento` datetime(6) DEFAULT NULL,
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
(2,true,NULL,NULL,'ichimar21@gmail.com',NULL,NULL,'$2a$10$59tlZW6RvpCSnPwfKGxpR.55WwSGMQRi9Gq.2D43Nd8tZcxvQbt02',NULL,"2VksD8o@\c");
UNLOCK TABLES;

##-- Table structure for table `direccion`
CREATE TABLE `direccion` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `tipo` varchar(255) DEFAULT NULL,
  `pais` varchar(255) DEFAULT NULL,
  `ciudad` varchar(255) DEFAULT NULL,
  `calle` varchar(255) DEFAULT NULL,
  `cp` varchar(255) DEFAULT NULL,
  `numero` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vrnijaj87yt` (`user_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vrnijaj87yt` FOREIGN KEY (`user_id`) REFERENCES `usuarios` (`id`)
);

##-- Table structure for table `direccion`
CREATE TABLE `referido` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint DEFAULT NULL,
  `codigo` varchar(255) DEFAULT NULL,
  `cobrado` BOOLEAN DEFAULT false,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vr88jaj87yt` (`user_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vr88jaj87yt` FOREIGN KEY (`user_id`) REFERENCES `usuarios` (`id`)
);

##-- Table structure for table `carteras`
CREATE TABLE `carteras` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `acumulado_alto8am` int DEFAULT NULL,
  `acumulado_bajo8pm` int DEFAULT NULL,
  `aproximacion_alta00am` int DEFAULT NULL,
  `aproximacion_baja` int DEFAULT NULL,
  `oportunidades` int DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKm5oo9iahtl1p9bs4dn1ymovlb` (`usuario_id`),
  CONSTRAINT `FKm5oo9iahtl1p9bs4dn1ymovlb` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`)
);
LOCK TABLES `carteras` WRITE;
INSERT INTO `carteras` VALUES (1,0,0,0,0,0,1),(2,0,0,0,0,0,2);
UNLOCK TABLES;

##-- Table structure for table `payment_method`
CREATE TABLE `payment_method` (
  `id` bigint NOT NULL,
  `activo` BOOLEAN NOT NULL,
  `type` varchar(255) DEFAULT NULL,
  `card_number` varchar(255) DEFAULT NULL,
  `cvc` int NOT NULL,
  `default_payment` BOOLEAN NOT NULL,
  `expiry_month` int NOT NULL,
  `expiry_year` int NOT NULL,
  `holder_name` varchar(255) DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKdfsg45Tthfi8r4DF43gfd34hF` (`usuario_id`),
  CONSTRAINT `FKdfsg45Tthfi8r4DF43gfd34hF` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`)
);
LOCK TABLES `payment_method` WRITE;
UNLOCK TABLES;

##-- Table structure for table `planes`
CREATE TABLE `planes` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `acumulado_alto8am` int DEFAULT NULL,
  `acumulado_bajo8pm` int DEFAULT NULL,
  `aproximacion_alta00am` int DEFAULT NULL,
  `aproximacion_baja` int DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `oportunidades` int DEFAULT NULL,
  `precio` double DEFAULT NULL,
  `suscribcion` BOOLEAN NOT NULL DEFAULT false,
  `pago_unico` BOOLEAN NOT NULL DEFAULT false,
  PRIMARY KEY (`id`)
);
LOCK TABLES `planes` WRITE;
INSERT INTO `planes` VALUES 
(1,true,1,1,1,1,'Ordinario',4,70,true,false),
(2,true,5,5,5,5,'Promocional',20,300,true,false),
(3,true,10,10,10,10,'Platinum',40,500,true,false),
(4,true,1,1,1,1,'Ordinario',4,70,false,true),
(5,true,5,5,5,5,'Promocional',20,300,false,true),
(6,true,10,10,10,10,'Platinum',40,500,false,true);
UNLOCK TABLES;

##-- Table structure for table `carrito`
CREATE TABLE `carrito` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL default true,
  `status` varchar(255) DEFAULT NULL,
  `cantidad` int DEFAULT 1,
  `total` double DEFAULT NULL,
  `fecha_carrito` datetime(6) DEFAULT NULL,
  `plan_id` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKdfsg45Tth5345gDF43gfd34hF` (`usuario_id`),
  KEY `a8sf23R980fsdf09er234gE7R6G` (`plan_id`),
  CONSTRAINT `FKdfsg45Tth5345gDF43gfd34hF` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`),
  CONSTRAINT `a8sf23R980fsdf09er234gE7R6G` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`)
) ;
LOCK TABLES `carrito` WRITE;
UNLOCK TABLES;

##-- Table structure for table `compra`
CREATE TABLE `compra` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `fecha_pagado` datetime(6) DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `carrito_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKpsfgo6ayx335hkqudyubw5536` (`usuario_id`),
  KEY `FK1yjtle73so0wuwyiyf7utf49a` (`carrito_id`),
  CONSTRAINT `FKpsfgo6ayx335hkqudyubw5536` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`),
  CONSTRAINT `FK1yjtle73so0wuwyiyf7utf49a` FOREIGN KEY (`carrito_id`) REFERENCES `carrito` (`id`)
) ;
LOCK TABLES `compra` WRITE;
UNLOCK TABLES;

##-- Table structure for table `beneficios`
CREATE TABLE `beneficios` (
  `id` bigint AUTO_INCREMENT,
  `activo` BOOLEAN default true,
  `llave` varchar(255) DEFAULT NULL,
  `tipo` varchar(255) DEFAULT NULL,
  `moneda` varchar(127) NOT NULL,
  `valor` double DEFAULT NULL,
  `repetido` BOOLEAN NOT NULL,
  `suscripcion` BOOLEAN NOT NULL,
  `pago_individual` BOOLEAN NOT NULL,
  `referido` BOOLEAN NOT NULL,
  PRIMARY KEY (`id`)
);

##-- Table structure for table `beneficios_usuario`
CREATE TABLE `beneficios_usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `fecha_inicio` datetime(6) DEFAULT NULL,
  `fecha_fin` datetime(6) DEFAULT NULL,
  `usuario_id` bigint NOT NULL,
  `beneficio_id` bigint NOT NULL,
  `plan_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FKpsfgo6447835hkqudyubw5536` (`usuario_id`),
  KEY `FKpsfgo6986335hkqudyubw5536` (`beneficio_id`),
  KEY `FKpsfgo6097235hkqudyubw5536` (`plan_id`),
  CONSTRAINT `FKpsfgo6447835hkqudyubw5536` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`),
  CONSTRAINT `FKpsfgo6986335hkqudyubw5536` FOREIGN KEY (`beneficio_id`) REFERENCES `beneficios` (`id`),
  CONSTRAINT `FKpsfgo6097235hkqudyubw5536` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`)
);

##-- Table structure for table `beneficios_plan`
CREATE TABLE `beneficios_plan` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `beneficio_id` bigint NOT NULL,
  `plan_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `6789s0985thfi8r4DF43gfd34hF` (`beneficio_id`),
  KEY `6789sg45Tth44334DF43gfd34hF` (`plan_id`),
  CONSTRAINT `6789s0985thfi8r4DF43gfd34hF` FOREIGN KEY (`beneficio_id`) REFERENCES `beneficios` (`id`),
  CONSTRAINT `6789sg45Tth44334DF43gfd34hF` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`)
);

##-- Table structure for table `suscripciones`
CREATE TABLE `suscripciones` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL default true,
  `monto_mensual` double NOT NULL,
  `fecha_create` datetime(6) DEFAULT NULL,
  `fecha_inicio` datetime(6) DEFAULT NULL,
  `fecha_fin` datetime(6) DEFAULT NULL,
  `fecha_cobro` datetime(6) DEFAULT NULL,
  `fecha_corte` int DEFAULT NULL,
  `tipo` varchar(255) DEFAULT NULL,
  `plan_id` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK6fcib1tyjrhp8u95q3uhohqc6` (`plan_id`),
  KEY `FKh7go9iahtl5u5bs4dn1ymovlb` (`usuario_id`),
  CONSTRAINT `FK6fcib1tyjrhp8u95q3uhohqc6` FOREIGN KEY (`plan_id`) REFERENCES `planes` (`id`),
  CONSTRAINT `FKh7go9iahtl5u5bs4dn1ymovlb` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`)
);
LOCK TABLES `suscripciones` WRITE;
UNLOCK TABLES;


##--videos
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN DEFAULT '1',
  `artista` varchar(255) DEFAULT NULL,
  `canal` varchar(255) DEFAULT NULL,
  `fecha_video` datetime(6) DEFAULT NULL,
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

##-- Table structure for table `categoria_evento`
CREATE TABLE `categoria_evento` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;
LOCK TABLES `categoria_evento` WRITE;
INSERT INTO `categoria_evento` VALUES (1,'Oportunidades'),(2,'Acumulado_alto8am'),(3,'Acumulado_bajo8pm'),(4,'aproximacion_alta00am'),(5,'aproximacion_baja');
UNLOCK TABLES;

##-- Table structure for table `tipo_evento`
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
  `activo` BOOLEAN DEFAULT true,
  `fechahora_evento` datetime(6) DEFAULT NULL,
  `premio_cash` double DEFAULT NULL,
  `acumulado` double DEFAULT NULL,
  `premio_otros` varchar(255) DEFAULT NULL,
  `moneda` varchar(255) DEFAULT NULL,
  `categoria_evento_id` bigint DEFAULT NULL,
  `video_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vrnijajf0c5` (`categoria_evento_id`),
  KEY `FKs87xk1t7ytkg1xw91sntybg6m` (`video_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vrnijajf0c5` FOREIGN KEY (`categoria_evento_id`) REFERENCES `categoria_evento` (`id`),
  CONSTRAINT `FKs87xk1t7ytkg1xw91sntybg6m` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
);
LOCK TABLES `eventos` WRITE;
INSERT INTO `eventos` VALUES  
(1,false,'2022-02-14 00:00:00.000000',null,0,null,'USD',4,6),
(2,true,'2022-02-14 00:30:00.000000',5000,0,null,'USD',5,5),
(3,true,'2022-02-14 01:00:00.000000',5000,0,null,'USD',5,4),
(4,true,'2022-02-14 01:30:00.000000',5000,0,null,'USD',5,3),
(5,true,'2022-02-14 02:00:00.000000',5000,0,null,'USD',5,2),
(6,true,'2022-02-14 02:30:00.000000',5000,0,null,'USD',5,1),
(7,true,'2022-02-16 03:00:00.000000',null,0,'iPhone 13X ','USD',1,1),
(8,true,'2022-02-14 03:30:00.000000',5000,0,null,'USD',5,6),
(9,true,'2022-02-17 04:00:00.000000',5000,0,null,'USD',5,5),
(10,true,'2022-02-17 04:30:00.000000',5000,0,null,'USD',5,4),
(11,true,'2022-02-17 05:00:00.000000',5000,0,null,'USD',5,3),
(12,true,'2022-02-17 05:30:00.000000',5000,0,null,'USD',5,2),
(13,true,'2022-05-12 06:00:00.000000',null,0,'iPhone 13X','USD',1,1),
(14,true,'2022-02-22 06:30:00.000000',5000,0,null,'USD',5,6),
(15,true,'2022-02-22 07:00:00.000000',5000,0,null,'USD',5,6),
(16,true,'2022-06-22 07:30:00.000000',5000,0,null,'USD',5,6),
(17,true,'2022-06-22 08:00:00.000000',100000,0,'$100,000','USD',2,6),
(18,true,'2022-06-22 08:30:00.000000',5000,0,null,'USD',5,6),
(19,true,'2022-02-22 09:00:00.000000',5000,0,null,'USD',5,6),
(20,true,'2022-06-22 09:30:00.000000',null,0,'iPhone 13X','USD',1,6),
(21,true,'2022-02-22 10:00:00.000000',5000,0,null,'USD',5,6),
(22,true,'2022-02-22 10:30:00.000000',5000,0,null,'USD',5,6),
(23,true,'2022-02-22 11:00:00.000000',5000,0,null,'USD',5,6),
(24,true,'2022-02-22 11:30:00.000000',5000,0,null,'USD',5,6),
(25,false,'2022-02-22 12:00:00.000000',null,0,'Carrito','USD',3,6),
(26,true,'2022-02-22 12:30:00.000000',5000,0,null,'USD',5,6),
(27,true,'2022-02-22 13:00:00.000000',5000,0,null,'USD',5,6),
(28,true,'2022-02-22 13:30:00.000000',5000,0,null,'USD',5,6),
(29,true,'2022-02-22 14:00:00.000000',5000,0,null,'USD',5,6),
(30,true,'2022-02-22 14:30:00.000000',5000,0,null,'USD',5,6),
(31,true,'2022-02-22 15:00:00.000000',null,0,'iPhone 13X','USD',1,6),
(32,true,'2022-02-22 15:30:00.000000',5000,0,null,'USD',5,6),
(33,true,'2022-02-22 16:00:00.000000',5000,0,null,'USD',5,6),
(34,true,'2022-02-22 16:30:00.000000',5000,0,null,'USD',5,6),
(35,true,'2022-02-22 17:00:00.000000',5000,0,null,'USD',5,6),
(36,true,'2022-02-22 17:30:00.000000',5000,0,null,'USD',5,6),
(37,true,'2022-02-22 18:00:00.000000',null,0,'iPhone 13X','USD',1,6),
(38,true,'2022-02-22 18:30:00.000000',5000,0,null,'USD',5,6),
(39,true,'2022-02-22 19:00:00.000000',5000,0,null,'USD',5,6),
(40,true,'2022-02-22 19:30:00.000000',5000,0,null,'USD',5,6),
(41,true,'2022-02-22 20:00:00.000000',70000,0,'$70,000','USD',3,6),
(42,true,'2022-02-22 20:30:00.000000',5000,0,null,'USD',5,6),
(43,true,'2022-02-22 21:00:00.000000',5000,0,null,'USD',5,6),
(44,true,'2022-02-22 21:30:00.000000',null,0,'iPhone 13X','USD',1,6),
(45,true,'2022-02-22 22:00:00.000000',5000,0,null,'USD',5,6),
(46,true,'2022-02-22 22:30:00.000000',5000,0,null,'USD',5,6),
(47,true,'2022-02-22 23:00:00.000000',5000,0,null,'USD',5,6),
(48,true,'2022-02-22 23:30:00.000000',5000,0,null,'USD',5,6);
UNLOCK TABLES;

##-- Table structure for table `cron_task`
CREATE TABLE `cron_task` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tarea_cron` varchar(255) DEFAULT NULL,
  `evento_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKlpl54gp4rwnxqb385ujpojoum` (`evento_id`),
  CONSTRAINT `FKlpl54gp4rwnxqb385ujpojoum` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`)
) ;
LOCK TABLES `cron_task` WRITE;
INSERT INTO `cron_task` VALUES (1,'00 12 05 22 02 *',14),(2,'00 12 05 22 02 *',15),(3,'00 12 05 22 06 *',16),(4,'00 12 05 22 06 *',17),(5,'00 12 05 22 06 *',18),(6,'00 12 05 22 02 *',19),(7,'00 12 05 22 06 *',20);
UNLOCK TABLES;

##-- Table structure for table `evento_usuario`
CREATE TABLE `evento_usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `fecha` datetime(6) DEFAULT NULL,
  `views` bigint DEFAULT NULL,
  `like` bigint DEFAULT NULL,
  `comments` bigint DEFAULT NULL,
  `dislikes` bigint DEFAULT NULL,
  `saved` bigint DEFAULT NULL,
  `shared` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `evento_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKcj7cb72y8k0doxs7jajmwt0q2` (`evento_id`),
  KEY `FKlln41mrxef4w9oomu5rcbnikj` (`usuario_id`),
  CONSTRAINT `FKcj7cb72y8k0doxs7jajmwt0q2` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`),
  CONSTRAINT `FKlln41mrxef4w9oomu5rcbnikj` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`)
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
  CONSTRAINT `FKcj7cb98dfk0do0092ajmwt0q2` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`),
  CONSTRAINT `FKlln4109plf4w9ooal93cbnikj` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`),
  CONSTRAINT `FKlln41mrxeld95oom0vm2bnikj` FOREIGN KEY (`evento_usuario_id`) REFERENCES `evento_usuario` (`id`)
);
LOCK TABLES `ganador` WRITE;
UNLOCK TABLES;

##-- Table structure for table `hibernate_sequence`
CREATE TABLE `hibernate_sequence` (
  `next_val` bigint DEFAULT NULL
);
LOCK TABLES `hibernate_sequence` WRITE;
INSERT INTO `hibernate_sequence` VALUES (1);
UNLOCK TABLES;

##-- Table structure for table `resultado`
CREATE TABLE `resultado` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `hora_resultado` datetime(6) DEFAULT NULL,
  `like_count` bigint DEFAULT NULL,
  `views_count` bigint DEFAULT NULL,
  `comments_count` bigint DEFAULT NULL,
  `dislikes_count` bigint DEFAULT NULL,
  `saved_count` bigint DEFAULT NULL,
  `shared_count` bigint DEFAULT NULL,
  `video_id` varchar(255) DEFAULT NULL,
  `evento_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `16ffsg45Tthfi8r4DF43gfd34hF` (`evento_id`),
  CONSTRAINT `16ffsg45Tthfi8r4DF43gfd34hF` FOREIGN KEY (`evento_id`) REFERENCES `eventos` (`id`)
);
LOCK TABLES `resultado` WRITE;
UNLOCK TABLES;

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
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  PRIMARY KEY (`id`),
  KEY `FKisd054ko30hm3j6ljr90asype` (`user_id`),
  KEY `FKihom0uklpkfpffipxpoyf7b74` (`role_id`),
  CONSTRAINT `FKihom0uklpkfpffipxpoyf7b74` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `FKisd054ko30hm3j6ljr90asype` FOREIGN KEY (`user_id`) REFERENCES `usuarios` (`id`)
);
LOCK TABLES `usuarios_roles` WRITE;
INSERT INTO `usuarios_roles` VALUES (1,1,2),(2,2,2);
UNLOCK TABLES;

