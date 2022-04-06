/*
 Navicat Premium Data Transfer

 Source Server         : local_dev
 Source Server Type    : MySQL
 Source Server Version : 80027
 Source Host           : localhost:3306
 Source Schema         : ginessential

 Target Server Type    : MySQL
 Target Server Version : 80027
 File Encoding         : 65001

 Date: 06/04/2022 22:57:20
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  `telephone` varchar(11) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (1, '2022-04-06 14:52:36', '2022-04-06 14:52:36', NULL, 'admin', '13888888888', '$2a$10$8MWYWyFoJmi7CzERzuT7GO.KuCv.yIq3uvRe/cuNYVsRakKZHWSC.');
INSERT INTO `users` VALUES (2, '2022-04-06 14:55:02', '2022-04-06 14:55:02', NULL, 'wd8f2QnFhk', '16888888888', '$2a$10$X4UIdmeYOUEoZe7cPOIwkOsCcR46EJo7biofdYFL9bVxGhgsCSDk6');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
