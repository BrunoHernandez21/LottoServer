DROP TABLE IF EXISTS `apuesta_usuario`;
CREATE TABLE `apuesta_usuario` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `cantidad` int NOT NULL DEFAULT 1,
  `fecha` datetime(6) DEFAULT NULL,
  `comentarios` bigint DEFAULT NULL,
  `likes` bigint DEFAULT NULL,
  `vistas` bigint DEFAULT NULL,
  `Dislikes` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `apuesta_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKcj7cb72y8k0doxs7jajmwt0q2` (`apuesta_id`),
  KEY `FKlln41mrxef4w9oomu5rcbnikj` (`usuario_id`),
  CONSTRAINT `FKcj7cb72y8k0doxs7jajmwt0q2` FOREIGN KEY (`apuesta_id`) REFERENCES `apuestas` (`id`),
  CONSTRAINT `FKlln41mrxef4w9oomu5rcbnikj` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`)
);

##-- Table structure for table `apuestas`
DROP TABLE IF EXISTS `apuestas`;
CREATE TABLE `apuestas` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN DEFAULT true,
  `fechahoraapuesta` datetime(6) DEFAULT NULL,
  `precio` double DEFAULT NULL,
  `acumulado` double DEFAULT NULL,
  `premio` varchar(255) DEFAULT NULL,
  `categoria_apuesta_id` bigint DEFAULT NULL,
  `tipo_apuesta_id` bigint DEFAULT NULL,
  `video_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKk6e2a82e9uvkc8vrnijajf0c5` (`categoria_apuesta_id`),
  KEY `FKpofbiappg24mbidd4ytolkrsl` (`tipo_apuesta_id`),
  KEY `FKs87xk1t7ytkg1xw91sntybg6m` (`video_id`),
  CONSTRAINT `FKk6e2a82e9uvkc8vrnijajf0c5` FOREIGN KEY (`categoria_apuesta_id`) REFERENCES `categoria_apuesta` (`id`),
  CONSTRAINT `FKpofbiappg24mbidd4ytolkrsl` FOREIGN KEY (`tipo_apuesta_id`) REFERENCES `tipo_apuesta` (`id`),
  CONSTRAINT `FKs87xk1t7ytkg1xw91sntybg6m` FOREIGN KEY (`video_id`) REFERENCES `videos` (`id`)
) ;
LOCK TABLES `apuestas` WRITE;
INSERT INTO `apuestas` VALUES  
(1,false,'2022-02-14 00:00:00.000000',0,0,'',4,1,6),
(2,true,'2022-02-14 00:30:00.000000',5000,0,'$5,000',5,2,5),
(3,true,'2022-02-14 01:00:00.000000',5000,0,'$5,000',5,3,4),
(4,true,'2022-02-14 01:30:00.000000',5000,0,'$5,000',5,4,3),
(5,true,'2022-02-14 02:00:00.000000',5000,0,'$5,000',5,1,2),
(6,true,'2022-02-14 02:30:00.000000',5000,0,'$5,000',5,2,1),
(7,true,'2022-02-16 03:00:00.000000',0,0,'iPhone 13X ',1,3,1),
(8,true,'2022-02-14 03:30:00.000000',5000,0,'5,000',5,4,6),
(9,true,'2022-02-17 04:00:00.000000',5000,0,'$5,000',5,1,5),
(10,true,'2022-02-17 04:30:00.000000',5000,0,'$5,000',5,2,4),
(11,true,'2022-02-17 05:00:00.000000',5000,0,'$5,000',5,3,3),
(12,true,'2022-02-17 05:30:00.000000',5000,0,'$5,000',5,4,2),
(13,true,'2022-05-12 06:00:00.000000',0,0,'iPhone 13X',1,1,1),
(14,true,'2022-02-22 06:30:00.000000',5000,0,'$5,000',5,2,6),
(15,true,'2022-02-22 07:00:00.000000',5000,0,'$5,000',5,3,6),
(16,true,'2022-06-22 07:30:00.000000',5000,0,'$5,000',5,4,6),
(17,true,'2022-06-22 08:00:00.000000',100000,0,'$100,000',2,1,6),
(18,true,'2022-06-22 08:30:00.000000',5000,0,'$5,000',5,2,6),
(19,true,'2022-02-22 09:00:00.000000',5000,0,'$5,000',5,3,6),
(20,true,'2022-06-22 09:30:00.000000',0,0,'iPhone 13X',1,4,6),
(21,true,'2022-02-22 10:00:00.000000',5000,0,'$5,000',5,1,6),
(22,true,'2022-02-22 10:30:00.000000',5000,0,'$5,000',5,2,6),
(23,true,'2022-02-22 11:00:00.000000',5000,0,'$5,000',5,3,6),
(24,true,'2022-02-22 11:30:00.000000',5000,0,'$5,000',5,4,6),
(25,false,'2022-02-22 12:00:00.000000',0,0,'',3,1,6),
(26,true,'2022-02-22 12:30:00.000000',5000,0,'$5,000',5,2,6),
(27,true,'2022-02-22 13:00:00.000000',5000,0,'$5,000',5,3,6),
(28,true,'2022-02-22 13:30:00.000000',5000,0,'$5,000',5,4,6),
(29,true,'2022-02-22 14:00:00.000000',5000,0,'$5,000',5,1,6),
(30,true,'2022-02-22 14:30:00.000000',5000,0,'$5,000',5,2,6),
(31,true,'2022-02-22 15:00:00.000000',0,0,'iPhone 13X',1,3,6),
(32,true,'2022-02-22 15:30:00.000000',5000,0,'$5,000',5,4,6),
(33,true,'2022-02-22 16:00:00.000000',5000,0,'$5,000',5,1,6),
(34,true,'2022-02-22 16:30:00.000000',5000,0,'$5,000',5,2,6),
(35,true,'2022-02-22 17:00:00.000000',5000,0,'$5,000',5,3,6),
(36,true,'2022-02-22 17:30:00.000000',5000,0,'$5,000',5,4,6),
(37,true,'2022-02-22 18:00:00.000000',0,0,'iPhone 13X',1,1,6),
(38,true,'2022-02-22 18:30:00.000000',5000,0,'$5,000',5,2,6),
(39,true,'2022-02-22 19:00:00.000000',5000,0,'$5,000',5,3,6),
(40,true,'2022-02-22 19:30:00.000000',5000,0,'$5,000',5,4,6),
(41,true,'2022-02-22 20:00:00.000000',70000,0,'$70,000',3,1,6),
(42,true,'2022-02-22 20:30:00.000000',5000,0,'$5,000',5,2,6),
(43,true,'2022-02-22 21:00:00.000000',5000,0,'$5,000',5,3,6),
(44,true,'2022-02-22 21:30:00.000000',0,0,'iPhone 13X',1,4,6),
(45,true,'2022-02-22 22:00:00.000000',5000,0,'$5,000',5,1,6),
(46,true,'2022-02-22 22:30:00.000000',5000,0,'$5,000',5,2,6),
(47,true,'2022-02-22 23:00:00.000000',5000,0,'$5,000',5,3,6),
(48,true,'2022-02-22 23:30:00.000000',5000,0,'$5,000',5,4,6);
UNLOCK TABLES;

##-- Table structure for table `carteras`
DROP TABLE IF EXISTS `carteras`;
CREATE TABLE `carteras` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `acumulado_alto8am` int DEFAULT NULL,
  `acumulado_bajo8pm` int DEFAULT NULL,
  `aproximacion_alta00am` int DEFAULT NULL,
  `aproximacion_baja` int DEFAULT NULL,
  `oportunidades` int DEFAULT NULL,
  `id_usuario` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKm5oo9iahtl1p9bs4dn1ymovlb` (`id_usuario`),
  CONSTRAINT `FKm5oo9iahtl1p9bs4dn1ymovlb` FOREIGN KEY (`id_usuario`) REFERENCES `usuarios` (`id`)
);

##-- Table structure for table `categoria_apuesta`
DROP TABLE IF EXISTS `categoria_apuesta`;
CREATE TABLE `categoria_apuesta` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;
LOCK TABLES `categoria_apuesta` WRITE;
INSERT INTO `categoria_apuesta` VALUES (1,'Oportunidades'),(2,'Acumulado_alto8am'),(3,'Acumulado_bajo8pm'),(4,'aproximacion_alta00am'),(5,'aproximacion_baja');
UNLOCK TABLES;



##-- Table structure for table `compra`
DROP TABLE IF EXISTS `compra`;
CREATE TABLE `compra` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `cantidad` int DEFAULT 1,
  `amount` double DEFAULT NULL,
  `fecha_compra` datetime(6) DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  `plan_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKpsfgo6ayx335hkqudyubw5536` (`usuario_id`),
  KEY `FK1yjtle73so0wuwyiyf7utf49a` (`plan_id`),
  CONSTRAINT `FKpsfgo6ayx335hkqudyubw5536` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`),
  CONSTRAINT `FK1yjtle73so0wuwyiyf7utf49a` FOREIGN KEY (`plan_id`) REFERENCES `plan` (`id`)
) ;
LOCK TABLES `compra` WRITE;
UNLOCK TABLES;

##-- Table structure for table `cron_task`
DROP TABLE IF EXISTS `cron_task`;
CREATE TABLE `cron_task` (
  `id` int NOT NULL AUTO_INCREMENT,
  `tarea_cron` varchar(255) DEFAULT NULL,
  `apuesta_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKlpl54gp4rwnxqb385ujpojoum` (`apuesta_id`),
  CONSTRAINT `FKlpl54gp4rwnxqb385ujpojoum` FOREIGN KEY (`apuesta_id`) REFERENCES `apuestas` (`id`)
) ;
LOCK TABLES `cron_task` WRITE;
INSERT INTO `cron_task` VALUES (1,'00 12 05 22 02 *',14),(2,'00 12 05 22 02 *',15),(3,'00 12 05 22 06 *',16),(4,'00 12 05 22 06 *',17),(5,'00 12 05 22 06 *',18),(6,'00 12 05 22 02 *',19),(7,'00 12 05 22 06 *',20);
UNLOCK TABLES;

##-- Table structure for table `ganador`
DROP TABLE IF EXISTS `ganador`;
CREATE TABLE `ganador` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `cantidad` bigint DEFAULT NULL,
  `concepto` varchar(255) DEFAULT NULL,
  `id_apuesta` int DEFAULT NULL,
  `id_ganador` int DEFAULT NULL,
  `id_usuario` int DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;
LOCK TABLES `ganador` WRITE;
UNLOCK TABLES;

##-- Table structure for table `hibernate_sequence`
DROP TABLE IF EXISTS `hibernate_sequence`;
CREATE TABLE `hibernate_sequence` (
  `next_val` bigint DEFAULT NULL
) ;
LOCK TABLES `hibernate_sequence` WRITE;
INSERT INTO `hibernate_sequence` VALUES (1);
UNLOCK TABLES;


##-- Table structure for table `orden`
DROP TABLE IF EXISTS `orden`;
CREATE TABLE `orden` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activa` BOOLEAN NOT NULL,
  `cantidad` int DEFAULT 1,
  `amount` double DEFAULT NULL,
  `fecha_orden` datetime(6) DEFAULT NULL,
  `id_charges` varchar(255) DEFAULT NULL,
  `orden_status` varchar(255) DEFAULT NULL,
  `id_plan` bigint DEFAULT NULL,
  `usuario_id` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FKdfsg45Tth5345gDF43gfd34hF` (`usuario_id`),
  KEY `a8sf23R980fsdf09er234gE7R6G` (`id_plan`),
  CONSTRAINT `FKdfsg45Tth5345gDF43gfd34hF` FOREIGN KEY (`usuario_id`) REFERENCES `usuarios` (`id`),
  CONSTRAINT `a8sf23R980fsdf09er234gE7R6G` FOREIGN KEY (`id_plan`) REFERENCES `plan` (`id`)
) ;
LOCK TABLES `orden` WRITE;
UNLOCK TABLES;

##-- Table structure for table `pago`
DROP TABLE IF EXISTS `pago`;
CREATE TABLE `pago` (
  `id` bigint NOT NULL,
  `card_number` varchar(255) DEFAULT NULL,
  `cvc` int NOT NULL,
  `default_payment` bit(1) NOT NULL,
  `expiry_month` int NOT NULL,
  `expiry_year` int NOT NULL,
  `holder_name` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
);
LOCK TABLES `pago` WRITE;
UNLOCK TABLES;

##-- Table structure for table `plan`
DROP TABLE IF EXISTS `plan`;
CREATE TABLE `plan` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `acumulado_alto8am` int DEFAULT NULL,
  `acumulado_bajo8pm` int DEFAULT NULL,
  `aproximacion_alta00am` int DEFAULT NULL,
  `aproximacion_baja` int DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `oportunidades` int DEFAULT NULL,
  `precio` double DEFAULT NULL,
  PRIMARY KEY (`id`)
);
LOCK TABLES `plan` WRITE;
INSERT INTO `plan` VALUES 
(1,true,0,0,0,0,'Ordinario',1,50),
(2,true,0,0,0,1,'Promocional',5,250),
(3,true,1,1,8,8,'Platinum',36,1630);
UNLOCK TABLES;

##-- Table structure for table `resultado`
DROP TABLE IF EXISTS `resultado`;
CREATE TABLE `resultado` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `comment_count` bigint DEFAULT NULL,
  `fechahoraapuesta` datetime(6) DEFAULT NULL,
  `id_apuesta` int DEFAULT NULL,
  `id_video` varchar(255) DEFAULT NULL,
  `like_count` bigint DEFAULT NULL,
  `view_count` bigint DEFAULT NULL,
  PRIMARY KEY (`id`)
);
LOCK TABLES `resultado` WRITE;
UNLOCK TABLES;

##-- Table structure for table `results`
DROP TABLE IF EXISTS `results`;
CREATE TABLE `results` (
  `idresults` int NOT NULL AUTO_INCREMENT COMMENT 'favorite_count, id_apuesta ,id_video, like_count, view_count',
  `comment_count` varchar(45) DEFAULT NULL,
  `favorite_count` varchar(45) DEFAULT NULL,
  `id_apuesta` varchar(45) DEFAULT NULL,
  `id_video` varchar(45) DEFAULT NULL,
  `like_count` varchar(45) DEFAULT NULL,
  `view_count` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`idresults`)
);
LOCK TABLES `results` WRITE;
INSERT INTO `results` VALUES (1,'4889','0','13','RgxM1Rv7hCs','265751','64903112'),(2,'4889','0','13','RgxM1Rv7hCs','265752','64903378'),(4,'4889','0','13','RgxM1Rv7hCs','265753','64905674'),(5,'4889','0','13','RgxM1Rv7hCs','265795','64920311'),(6,'16898','0','16','g2BzGJnNvEw','1590356','263016531'),(7,'16898','0','17','g2BzGJnNvEw','1590356','263016531'),(8,'16898','0','20','g2BzGJnNvEw','1590356','263016531'),(9,'16898','0','18','g2BzGJnNvEw','1590356','263016531');
UNLOCK TABLES;

##-- Table structure for table `roles`
DROP TABLE IF EXISTS `roles`;
CREATE TABLE `roles` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_ldv0v52e0udsh2h1rs0r0gw1n` (`nombre`)
) ;
LOCK TABLES `roles` WRITE;
INSERT INTO `roles` VALUES (2,'ROLE_ADMIN'),(1,'ROLE_USER');

UNLOCK TABLES;

##-- Table structure for table `suscripciones`
DROP TABLE IF EXISTS `suscripciones`;
CREATE TABLE `suscripciones` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN NOT NULL,
  `acumulado_alto8am` int DEFAULT NULL,
  `acumulado_bajo8pm` int DEFAULT NULL,
  `aproximacion_alta00am` int DEFAULT NULL,
  `aproximacion_baja` int DEFAULT NULL,
  `fecha_inicio` datetime(6) DEFAULT NULL,
  `oportunidades` int DEFAULT NULL,
  `id_plan` bigint DEFAULT NULL,
  `id_usuario` bigint DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `FK6fcib1tyjrhp8u95q3uhohqc6` (`id_plan`),
  KEY `FKh7go9iahtl5u5bs4dn1ymovlb` (`id_usuario`),
  CONSTRAINT `FK6fcib1tyjrhp8u95q3uhohqc6` FOREIGN KEY (`id_plan`) REFERENCES `plan` (`id`),
  CONSTRAINT `FKh7go9iahtl5u5bs4dn1ymovlb` FOREIGN KEY (`id_usuario`) REFERENCES `usuarios` (`id`)
);
LOCK TABLES `suscripciones` WRITE;
INSERT INTO `suscripciones` VALUES 
(1,true,1,1,8,8,'2022-04-29 14:36:45.465000',36,3,1);
UNLOCK TABLES;

##-- Table structure for table `tipo_apuesta`
DROP TABLE IF EXISTS `tipo_apuesta`;
CREATE TABLE `tipo_apuesta` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `nombre` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ;
LOCK TABLES `tipo_apuesta` WRITE;
INSERT INTO `tipo_apuesta` VALUES (1,'Views'),(2,'Like'),(3,'comments'),(4,'Dislikes');
UNLOCK TABLES;

##-- Table structure for table `usuarios`
DROP TABLE IF EXISTS `usuarios`;
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
  PRIMARY KEY (`id`),
  UNIQUE KEY `UK_kfsp0s1tflm1cwlj8idhqsad0` (`email`)
);
LOCK TABLES `usuarios` WRITE;
INSERT INTO `usuarios` VALUES (1,true,NULL,NULL,'mezagg@gmail.com',NULL,NULL,'$2a$10$59tlZW6RvpCSnPwfKGxpR.55WwSGMQRi9Gq.2D43Nd8tZcxvQbt02',NULL);
UNLOCK TABLES;

##-- Table structure for table `usuarios_roles`
DROP TABLE IF EXISTS `usuarios_roles`;
CREATE TABLE `usuarios_roles` (
  `user_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  UNIQUE KEY `FKisd054ko30hm3j6ljr90asype` (`user_id`),
  KEY `FKihom0uklpkfpffipxpoyf7b74` (`role_id`),
  CONSTRAINT `FKihom0uklpkfpffipxpoyf7b74` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`),
  CONSTRAINT `FKisd054ko30hm3j6ljr90asype` FOREIGN KEY (`user_id`) REFERENCES `usuarios` (`id`)
);
LOCK TABLES `usuarios_roles` WRITE;
INSERT INTO `usuarios_roles` VALUES (1,2),(2,2);
UNLOCK TABLES;

##--videos
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `activo` BOOLEAN DEFAULT '1',
  `artista` varchar(255) DEFAULT NULL,
  `canal` varchar(255) DEFAULT NULL,
  `fecha_video` datetime(6) DEFAULT NULL,
  `id_video` varchar(255) DEFAULT NULL,
  `thumblary` varchar(255) DEFAULT NULL,
  `titulo` varchar(255) DEFAULT NULL,
  `url_video` varchar(255) DEFAULT NULL,
  `genero` varchar(255) DEFAULT NULL,
  
  PRIMARY KEY (`id`)
);
LOCK TABLES `videos` WRITE;
INSERT INTO `videos` VALUES 
(1,1,'BANDA MS','Lizos Music','2020-10-20 00:00:00.000000','RgxM1Rv7hCs','https://i.ytimg.com/vi/RgxM1Rv7hCs/default.jpg','CERRANDO CICLOS (VIDEO OFICIAL)','https://www.youtube.com/watch?v=RgxM1Rv7hCs','Pop'),
(2,1,'BANDA MS','Lizos Music','2020-02-28 00:00:00.000000','ova8TGDNvCo','https://i.ytimg.com/vi/ova8TGDNvCo/default.jpg','QUIÉN PIERDE MÁS (VIDEO OFICIAL)','https://www.youtube.com/watch?v=ova8TGDNvCo','Salsa'),
(3,1,'BANDA MS','Lizos Music','2018-02-13 00:00:00.000000','2mf1Os7dAJI','https://i.ytimg.com/vi/2mf1Os7dAJI/default.jpg','SI CRUZAS LA PUERTA (LETRA)','https://www.youtube.com/watch?v=2mf1Os7dAJI','Cumbia'),
(4,1,'Christian Nodal','Christian Nodal','2019-12-05 00:00:00.000000','Ax3psz01Q8o','https://i.ytimg.com/vi/Ax3psz01Q8o/default.jpg','Si Te Falta Alguien (Video Oficial)','https://www.youtube.com/watch?v=Ax3psz01Q8o','Banda'),
(5,1,'Carin Leon','TAMARINDOREKORDSZ','2021-12-12 00:00:00.000000','8Bznc5tTQ9M','https://i.ytimg.com/vi/8Bznc5tTQ9M/default.jpg','Si Una Vez','https://www.youtube.com/watch?v=8Bznc5tTQ9M','Salsa'),
(6,1,'Grupo Firme','Grupo Firme','2020-07-24 00:00:00.000000','g2BzGJnNvEw','https://i.ytimg.com/vi/g2BzGJnNvEw/default.jpg',' Ya Superame','https://www.youtube.com/watch?v=g2BzGJnNvEw','Banda');
UNLOCK TABLES;