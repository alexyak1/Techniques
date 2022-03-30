-- MySQL dump 10.13  Distrib 8.0.26, for Linux (x86_64)
--
-- Host: localhost    Database: techniques
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `techniques`
--

DROP TABLE IF EXISTS `techniques`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `techniques` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(128) NOT NULL,
  `belt` varchar(255) NOT NULL,
  `image_url` varchar(255) DEFAULT NULL,
  `type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `techniques`
--

LOCK TABLES `techniques` WRITE;
/*!40000 ALTER TABLE `techniques` DISABLE KEYS */;
INSERT INTO `techniques` VALUES (1,'O-soto-otoshi','yellow','https://drive.google.com/file/d/1p3JuGkd533lBv4lrzzUD4BibXy4Jjinz/view?usp=sharing','Nage-Waza'),(2,'O-uchi-gari','yellow','https://drive.google.com/file/d/1APzyii2ibDU82E_OZXilW3fl7_KZMFSr/view?usp=sharing','Nage-Waza'),(3,'O-goshi','yellow','https://drive.google.com/file/d/1V_CtDZUFY3T8Q3Zt4tKqXpBc9YV4G-j-/view?usp=sharing','Nage-Waza'),(4,'Mune-gatame','yellow','https://drive.google.com/file/d/1zoDxMR04Zi6RudTyegS_m39GHNbt49kZ/view?usp=sharing','Osaekomi-Waza'),(9,'Kuzure-kesa-gatame','yellow','https://drive.google.com/file/d/1fcfVAvjw6Ec8J3qnPvPpGVDfck2j5fn8/view?usp=sharing','Osaekomi-Waza'),(10,'Ko-soto-gari','yellow','https://drive.google.com/file/d/1zN2obD93q6xNl2bCpmDYXitb5LwysA7q/view?usp=sharing','Nage-Waza'),(11,'Ko-uchi-gari','yellow','https://drive.google.com/file/d/13Au37XbZXHxEcvqEbf_yFDhNGN9oCgiD/view?usp=sharing','Nage-Waza'),(12,'Hiza-guruma','yellow','https://drive.google.com/file/d/1_TcaJVLmIha2-q8uh3wa1xHoLxHtx3Tn/view?usp=sharing','Nage-Waza'),(13,'Eri-seoi-nage','yellow','https://drive.google.com/file/d/1sE9ACTQch9nTm_OZN6OQKPk66ZotIA4W/view?usp=sharing','Nage-Waza'),(14,'Koshi-guruma','yellow','https://drive.google.com/file/d/1TkOP0wbnN0CWQ-r8yVUzhDGSzKlD0XEm/view?usp=sharing','Nage-Waza'),(15,'Kami-shiho-gatame','yellow','https://drive.google.com/file/d/1vypc8zfwcNwAYk3CzD2QvPGf2y_8Hdtq/view?usp=sharing','Osaekomi-Waza'),(16,'Tate-shiho-gatame','yellow','https://drive.google.com/file/d/1JOAKswZmPIH2BVMKUHWClvTKsGrQgiOT/view?usp=sharing','Osaekomi-Waza'),(17,'De-ashi-harai','orange','https://drive.google.com/file/d/1DBMv0FlGz929TKMCaFTR-eiOavaGM5Rb/view?usp=sharing','Nage-Waza'),(18,'Ko-soto-gake','orange','https://drive.google.com/file/d/16j-npVYVR4pAJvJV_CKX51EqSnbBJUBz/view?usp=sharing','Nage-Waza'),(19,'Ippon-seoi-nage','orange','https://drive.google.com/file/d/1TvnUyAXTxeunJF1XflN_yW_dDl2fwRJB/view?usp=sharing','Nage-Waza'),(20,'Morote-seoi-nage','orange','https://drive.google.com/file/d/1ufSl4n22UXMNQkSqBxbudvHJZLDkwbTr/view?usp=sharing','Nage-Waza'),(21,'Tsuri-komi-goshi','orange','https://drive.google.com/file/d/1YdF5tioCh0Riz3vPiAkqy7PtdMzgt5Wr/view?usp=sharing','Nage-Waza'),(22,'Kubi-nage','orange','https://drive.google.com/file/d/10v5XIdURJXmAFAQC-Mg65RbsFfHPKLWT/view?usp=sharing','Nage-Waza'),(23,'Ushiro-kesa-gatame','orange','https://drive.google.com/file/d/10v5XIdURJXmAFAQC-Mg65RbsFfHPKLWT/view?usp=sharing','Osaekomi-Waza'),(24,'Kuzure-yoko-shiho-gatame','orange','https://drive.google.com/file/d/1FZp0THEJ47OBC-e8RPXPav5mLcDoBSce/view?usp=sharing','Osaekomi-Waza'),(25,'O-soto-gari','orange','https://drive.google.com/file/d/10I2O8sPkylJEFHx35z8G30WzSDgY3Yz_/view?usp=sharing','Nage-Waza'),(26,'Tsuri-goshi','orange','https://drive.google.com/file/d/1u9qlv3Gy34vI46O8UD9WC-AokTHNgbbP/view?usp=sharing','Nage-Waza'),(27,'Tai-otoshi','orange','https://drive.google.com/file/d/1AXM4cFEje0dQ1Z6KmfKaRRhSHw7VYKJo/view?usp=sharing','Nage-Waza'),(28,'Uchi-mata','orange','https://drive.google.com/file/d/16XWRbPj2cj4MtNY3pIuJdzChfq2twDNF/view?usp=sharing','Nage-Waza'),(29,'Harai-goshi','orange','https://drive.google.com/file/d/1zD67DXdh8AqLsdXXuuoiTcU7z0fkn3Gz/view?usp=sharing','Nage-Waza'),(30,'Yoko-shiho-gatame','orange','https://drive.google.com/file/d/1Q8dFGgVE8hRR5gsozj0L1oEuUWlL0g5C/view?usp=sharing','Osaekomi-Waza'),(31,'Kesa-gatame','orange','https://drive.google.com/file/d/1MZE17XwQBuL4o7XK5FFU-2zX_AP4am-0/view?usp=sharing','Osaekomi-Waza'),(32,'Gyaku-juji-jime','orange','https://drive.google.com/file/d/1cCYo3E84m_vOneQbuLKJcxH6wcTeaNhw/view?usp=sharing','Shime-Waza'),(33,'Nami-juji-jime','orange','https://drive.google.com/file/d/18ri-ALP525z8eCJLkmF-ZKLtpj7J9OeI/view?usp=sharing','Shime-Waza'),(34,'Kata-juji-jime','orange','https://drive.google.com/file/d/1Iv-GnvYlUJ1F3j0bJY-Mo9Aybk4tkFNO/view?usp=sharing','Shime-Waza'),(35,'Ude-garami','orange','https://drive.google.com/file/d/12fWDqXD8cyTtrhkA6PmOcV0DvzZfO7Cm/view?usp=sharing','Kansetsu-Waza'),(36,'Juji-gatame','orange','https://drive.google.com/file/d/1tiu0rL79egqINJHBa315l_zV2ntpH_FW/view?usp=sharing','Kansetsu-Waza');
/*!40000 ALTER TABLE `techniques` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-03-30 15:54:47
