-- MariaDB dump 10.19  Distrib 10.6.5-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: hfcms_users
-- ------------------------------------------------------
-- Server version	10.6.5-MariaDB-1:10.6.5+maria~focal

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `department_users`
--

DROP TABLE IF EXISTS `department_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `department_users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `department_id` int(10) DEFAULT NULL,
  `user_id` int(10) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `department_users`
--

LOCK TABLES `department_users` WRITE;
/*!40000 ALTER TABLE `department_users` DISABLE KEYS */;
INSERT INTO `department_users` VALUES (1,4,3,NULL,0,'2022-01-13 12:03:51'),(2,5,4,NULL,0,'2022-01-13 12:03:51'),(3,5,2,NULL,0,'2022-01-13 12:03:51'),(4,5,1,NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `department_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `departments`
--

DROP TABLE IF EXISTS `departments`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `departments` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `state` tinyint(1) DEFAULT 0 COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `departments`
--

LOCK TABLES `departments` WRITE;
/*!40000 ALTER TABLE `departments` DISABLE KEYS */;
INSERT INTO `departments` VALUES (1,NULL,'dbcd','SDP','Sales Department',0,0,'2022-01-13 12:03:51'),(2,NULL,'dfgh','PRD','Public Relations Department',0,0,'2022-01-13 12:03:51'),(3,NULL,'djkl','R&D','Research and Development Department',0,0,'2022-01-13 12:03:51'),(4,NULL,'dnop','FD','Finance Department',0,0,'2022-01-13 12:03:51'),(5,NULL,'drst','F&B','Food & Beverage Department',0,0,'2022-01-13 12:03:51'),(6,1,'duvw','SDO','Sales Department Office',0,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `departments` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `permissions`
--

DROP TABLE IF EXISTS `permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `permissions` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  UNIQUE KEY `id` (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `permissions`
--

LOCK TABLES `permissions` WRITE;
/*!40000 ALTER TABLE `permissions` DISABLE KEYS */;
INSERT INTO `permissions` VALUES (1,NULL,'pabc','delete','oh no!',NULL,0,'2022-01-13 12:03:51'),(2,NULL,'pdef','update','be carefull',NULL,0,'2022-01-13 12:03:51'),(3,NULL,'pghi','create','Create what?',NULL,0,'2022-01-13 12:03:51'),(4,NULL,'pjkl','query','ok, show you',NULL,0,'2022-01-13 12:03:51'),(5,2,'pmon','article','update article',NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_permissions`
--

DROP TABLE IF EXISTS `role_permissions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_permissions` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` int(10) DEFAULT NULL,
  `permission_id` int(10) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_permissions`
--

LOCK TABLES `role_permissions` WRITE;
/*!40000 ALTER TABLE `role_permissions` DISABLE KEYS */;
INSERT INTO `role_permissions` VALUES (1,1,1,NULL,0,'2022-01-13 12:03:51'),(2,1,2,NULL,0,'2022-01-13 12:03:51'),(3,1,3,NULL,0,'2022-01-13 12:03:51'),(4,1,4,NULL,0,'2022-01-13 12:03:51'),(5,1,5,NULL,0,'2022-01-13 12:03:51'),(6,2,3,NULL,0,'2022-01-13 12:03:51'),(7,2,4,NULL,0,'2022-01-13 12:03:51'),(8,3,4,NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `role_permissions` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_usergroups`
--

DROP TABLE IF EXISTS `role_usergroups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_usergroups` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` int(10) DEFAULT NULL,
  `usergroup_id` int(10) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_usergroups`
--

LOCK TABLES `role_usergroups` WRITE;
/*!40000 ALTER TABLE `role_usergroups` DISABLE KEYS */;
INSERT INTO `role_usergroups` VALUES (1,1,1,NULL,0,'2022-01-13 12:03:51'),(2,2,2,NULL,0,'2022-01-13 12:03:51'),(3,3,3,NULL,0,'2022-01-13 12:03:51'),(4,3,4,NULL,0,'2022-01-13 12:03:51'),(5,3,5,NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `role_usergroups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `role_users`
--

DROP TABLE IF EXISTS `role_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `role_users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `role_id` int(10) DEFAULT NULL,
  `user_id` int(10) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `role_users`
--

LOCK TABLES `role_users` WRITE;
/*!40000 ALTER TABLE `role_users` DISABLE KEYS */;
INSERT INTO `role_users` VALUES (1,1,1,NULL,0,'2022-01-13 12:03:51'),(2,2,2,NULL,0,'2022-01-13 12:03:51'),(3,2,3,NULL,0,'2022-01-13 12:03:51'),(4,3,4,NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `role_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `roles`
--

DROP TABLE IF EXISTS `roles`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `roles` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `roles`
--

LOCK TABLES `roles` WRITE;
/*!40000 ALTER TABLE `roles` DISABLE KEYS */;
INSERT INTO `roles` VALUES (1,NULL,'rabc','master','You know, here\'s my spot.',NULL,0,'2022-01-13 12:03:51'),(2,NULL,'rdef','submaster','You know, master is my boss.',NULL,0,'2022-01-13 12:03:51'),(3,NULL,'rdeg','users','Ordinary people',NULL,0,'2022-01-13 12:03:51'),(4,NULL,'rdeh','guest','May be newbie here',NULL,0,'2022-01-13 12:03:51'),(5,NULL,'rghi','blocked','clear out here!',NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `roles` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usergroup_users`
--

DROP TABLE IF EXISTS `usergroup_users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usergroup_users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `usergroup_id` int(10) DEFAULT NULL,
  `user_id` int(10) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usergroup_users`
--

LOCK TABLES `usergroup_users` WRITE;
/*!40000 ALTER TABLE `usergroup_users` DISABLE KEYS */;
INSERT INTO `usergroup_users` VALUES (1,1,1,NULL,0,'2022-01-13 12:03:51'),(2,2,2,NULL,0,'2022-01-13 12:03:51'),(3,3,3,NULL,0,'2022-01-13 12:03:51'),(4,4,3,NULL,0,'2022-01-13 12:03:51'),(5,5,3,NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `usergroup_users` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usergroups`
--

DROP TABLE IF EXISTS `usergroups`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `usergroups` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `state` tinyint(1) DEFAULT NULL COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE KEY `code` (`code`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usergroups`
--

LOCK TABLES `usergroups` WRITE;
/*!40000 ALTER TABLE `usergroups` DISABLE KEYS */;
INSERT INTO `usergroups` VALUES (1,NULL,'ug001','GameMaster','My plasure',NULL,0,'2022-01-13 12:03:51'),(2,NULL,'ug100','Room Master','All rooms control',NULL,0,'2022-01-13 12:03:51'),(3,2,'ug101','Room 1','good day',NULL,0,'2022-01-13 12:03:51'),(4,2,'ug102','Room 2','good day',NULL,0,'2022-01-13 12:03:51'),(5,2,'ug103','Room 3','good day',NULL,0,'2022-01-13 12:03:51');
/*!40000 ALTER TABLE `usergroups` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `realname` varchar(255) DEFAULT NULL,
  `nickname` varchar(255) DEFAULT NULL,
  `avatar_url` varchar(255) DEFAULT NULL,
  `phone` varchar(11) DEFAULT NULL,
  `user_ip` int(4) unsigned DEFAULT NULL,
  `state` tinyint(1) DEFAULT 0 COMMENT 'user state: 0=normal, 1=disable',
  `deleted` tinyint(1) DEFAULT 0 COMMENT 'soft deleted: 0=undelete,1=deleted',
  `create_time` timestamp NOT NULL DEFAULT current_timestamp(),
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES (1,'zhangsan','*036BD541925B8EE4682D9DE696D3549DFB089957','zhangsan','zhangsansan','zhangsan.jpg','13912345678',3232235530,0,0,'2022-01-13 12:03:51','2022-01-13 12:03:51'),(2,'lisi','*29B17A67C2C3D2779E95FA08BB42A6AD25C63D55','lisi','lisisi','lisi.jpg','13912345679',3232235531,0,0,'2022-01-13 12:03:51','2022-01-13 12:03:51'),(3,'wangwu','*7FD5C4BCBD56844F22C4F9372DFE46E8BAEA79AF','wangwu','wangwuwu','wangwu.jpg','13912345670',3232235532,0,0,'2022-01-13 12:03:51','2022-01-13 12:03:51'),(4,'test','*22CBF14EBDE8814586FF12332FA2B6023A7603BB','test','testt','test.jpg','13912345671',3232235789,0,0,'2022-01-13 12:03:51','2022-01-13 12:03:51'),(5,'guest','*33CE4D2EA3F507528311D7F10073F1A6EEA433DC','guest','guestt','guest.jpg','13912345611',3232235790,0,0,'2022-01-13 12:03:51','2022-01-13 12:03:51');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-01-13 12:22:09
