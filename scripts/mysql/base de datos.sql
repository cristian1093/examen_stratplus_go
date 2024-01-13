# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.39)
# Database: test
# Generation Time: 2024-01-13 03:27:05 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table bonds
# ------------------------------------------------------------

DROP TABLE IF EXISTS `bonds`;

CREATE TABLE `bonds` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `current_price` decimal(15,0) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `bonds` WRITE;
/*!40000 ALTER TABLE `bonds` DISABLE KEYS */;

INSERT INTO `bonds` (`id`, `name`, `description`, `current_price`)
VALUES
	(1,'Prueba1','bono prueba 1',12),
	(2,'Prueba2','bono prueba 2',111);

/*!40000 ALTER TABLE `bonds` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table orders
# ------------------------------------------------------------

DROP TABLE IF EXISTS `orders`;

CREATE TABLE `orders` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) DEFAULT NULL,
  `bond_id` int(11) DEFAULT NULL,
  `order_type` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL,
  `num_bonds` int(11) DEFAULT NULL,
  `price` decimal(15,4) DEFAULT NULL,
  `status` varchar(100) NOT NULL,
  `expiration` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_user_id` (`user_id`),
  KEY `fk_bond_id` (`bond_id`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
  CONSTRAINT `orders_ibfk_2` FOREIGN KEY (`bond_id`) REFERENCES `bonds` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `orders` WRITE;
/*!40000 ALTER TABLE `orders` DISABLE KEYS */;

INSERT INTO `orders` (`id`, `user_id`, `bond_id`, `order_type`, `num_bonds`, `price`, `status`, `expiration`)
VALUES
	(1,1,1,'buy',145,123.3000,'completado','2024-01-10 03:05:03'),
	(2,1,1,'buy',145,123.3000,'cancelado','2024-01-10 23:57:14'),
	(3,2,2,'buy',145,123.3000,'cancelado','2024-01-11 22:39:02');

/*!40000 ALTER TABLE `orders` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user` varchar(50) DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone` bigint(20) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `user`, `email`, `phone`, `password`)
VALUES
	(1,'Cristian','cristian@hotmail.com',5514394200,'123'),
	(2,'Armando','armando@hotmail.com',5515066338,'345'),
	(3,'Carlos B','carlos@hotmail.com',5514394211,'098'),

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



--
-- Dumping routines (PROCEDURE) for database 'test'
--
DELIMITER ;;

# Dump of PROCEDURE UpdateAndShowOrders
# ------------------------------------------------------------

/*!50003 DROP PROCEDURE IF EXISTS `UpdateAndShowOrders` */;;
/*!50003 SET SESSION SQL_MODE="ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION"*/;;
/*!50003 CREATE*/ /*!50020 DEFINER=`root`@`localhost`*/ /*!50003 PROCEDURE `UpdateAndShowOrders`(IN _user_id int(11))
BEGIN
	
	UPDATE orders SET status = "cancelado" WHERE id = id and status = "pendiente" and expiration < NOW();
	IF _user_id IS NOT NULL THEN
	SELECT o.id, o.user_id, o.bond_id, o.order_type, o.num_bonds, o.price, o.status, o.expiration, u.id as user_id, u.user, u.email, u.phone,b.id ,b.name ,b.description ,b.current_price FROM test.orders o INNER JOIN test.users u ON o.user_id = u.id INNER JOIN test.bonds b ON o.bond_id = b.id WHERE o.user_id = _user_id;
	ELSE
	SELECT o.id, o.user_id, o.bond_id, o.order_type, o.num_bonds, o.price, o.status, o.expiration, u.id as user_id, u.user, u.email, u.phone,b.id ,b.name ,b.description ,b.current_price FROM test.orders o INNER JOIN test.users u ON o.user_id = u.id INNER JOIN test.bonds b ON o.bond_id = b.id;
	END IF;
END */;;

/*!50003 SET SESSION SQL_MODE=@OLD_SQL_MODE */;;
DELIMITER ;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
