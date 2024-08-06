-- Adminer 4.8.1 MySQL 8.0.30 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `u628725475_ecommerce` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `u628725475_ecommerce`;

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `ID` bigint unsigned NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `category` (`ID`, `NAME`) VALUES
(1,	'Website Development'),
(2,	'E-commerce Solutions'),
(3,	'Portfolio Templates'),
(4,	'Blog Templates'),
(5,	'Business Templates');

DROP TABLE IF EXISTS `detailtransaksi`;
CREATE TABLE `detailtransaksi` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trx_id` bigint unsigned DEFAULT NULL,
  `product_id` bigint unsigned DEFAULT NULL,
  `qty` bigint DEFAULT NULL,
  `total` decimal(10,2) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `idx_DETAILTRANSAKSI_trx_id` (`trx_id`),
  KEY `idx_DETAILTRANSAKSI_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `detailtransaksi` (`id`, `trx_id`, `product_id`, `qty`, `total`, `created_at`, `updated_at`) VALUES
(1,	1,	1,	1,	15000.00,	'2024-08-06 10:38:05',	'2024-08-06 10:38:05'),
(2,	2,	2,	1,	250000.00,	'2024-08-06 10:38:20',	'2024-08-06 10:38:20'),
(3,	3,	3,	1,	5000.00,	'2024-08-06 10:38:28',	'2024-08-06 10:38:28'),
(4,	4,	4,	1,	30000.00,	'2024-08-06 10:38:37',	'2024-08-06 10:38:37'),
(5,	5,	5,	1,	75000.00,	'2024-08-06 10:38:49',	'2024-08-06 10:38:49');

DROP TABLE IF EXISTS `product`;
CREATE TABLE `product` (
  `ID` bigint unsigned NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `PRICE` decimal(10,2) DEFAULT NULL,
  `SLUG` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `IMAGE_URL` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `CATEGORY_ID` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `fk_PRODUCT_category` (`CATEGORY_ID`),
  CONSTRAINT `fk_PRODUCT_category` FOREIGN KEY (`CATEGORY_ID`) REFERENCES `category` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `product` (`ID`, `NAME`, `PRICE`, `SLUG`, `IMAGE_URL`, `CATEGORY_ID`) VALUES
(1,	'Custom Website Design',	50000.00,	'custom-website-design',	NULL,	1),
(2,	'Online Store Setup',	45000.00,	'online-store-setup',	NULL,	2),
(3,	'Art Portfolio Template',	600000.00,	'art-portfolio-template',	NULL,	3),
(4,	'Personal Blog Template',	800000.00,	'personal-blog-template',	'',	4),
(5,	'Corporate Business Template',	40000.00,	'corporate-business-template',	NULL,	5);

DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile` (
  `ID` bigint unsigned NOT NULL AUTO_INCREMENT,
  `NAME` longtext COLLATE utf8mb4_general_ci,
  `GENDER` longtext COLLATE utf8mb4_general_ci,
  `CITY` longtext COLLATE utf8mb4_general_ci,
  `DATE` datetime(3) DEFAULT NULL,
  `ADDRESS` longtext COLLATE utf8mb4_general_ci,
  `PHONE` longtext COLLATE utf8mb4_general_ci,
  `user_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `fk_PROFILE_user` (`user_id`),
  CONSTRAINT `fk_PROFILE_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `profile` (`ID`, `NAME`, `GENDER`, `CITY`, `DATE`, `ADDRESS`, `PHONE`, `user_id`, `created_at`, `updated_at`) VALUES
(3,	'John Doe',	'Male',	'New York',	'2023-01-15 00:00:00.000',	'123 6th St',	'123-456-7890',	2,	'2024-08-06 17:35:09.000',	'2024-08-06 17:35:09.000'),
(4,	'Sara Smith',	'Female',	'Los Angeles',	'2023-01-20 00:00:00.000',	'789 Main St',	'234-567-8901',	3,	'2024-08-06 17:35:55.000',	'2024-08-06 17:35:55.000'),
(5,	'Alex Jones',	'Male',	'Chicago',	'2023-02-10 00:00:00.000',	'456 Elm St',	'345-678-9012',	4,	'2024-08-06 17:36:33.000',	'2024-08-06 17:36:33.000'),
(6,	'Lisa White',	'Female',	'Female',	'2023-03-05 00:00:00.000',	'123 Oak St',	'456-789-0123',	5,	'2024-08-06 17:37:01.000',	'2024-08-06 17:37:01.000'),
(7,	'Mark Brown',	'Male',	'Philadelphia',	'2023-04-16 00:00:00.000',	'789 Maple St',	'567-890-1234',	6,	'2024-08-06 17:37:25.000',	'2024-08-06 17:37:25.000');

DROP TABLE IF EXISTS `review`;
CREATE TABLE `review` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `USER_ID` bigint unsigned DEFAULT NULL,
  `PRODUCT_ID` bigint unsigned DEFAULT NULL,
  `TRX_ID` bigint unsigned DEFAULT NULL,
  `CONTENT` longtext COLLATE utf8mb4_general_ci,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_REVIEW_user` (`USER_ID`),
  KEY `fk_REVIEW_product` (`PRODUCT_ID`),
  KEY `fk_REVIEW_transaction` (`TRX_ID`),
  CONSTRAINT `fk_REVIEW_product` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `product` (`ID`),
  CONSTRAINT `fk_REVIEW_transaction` FOREIGN KEY (`TRX_ID`) REFERENCES `transaksi` (`id`),
  CONSTRAINT `fk_REVIEW_user` FOREIGN KEY (`USER_ID`) REFERENCES `user` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `review` (`id`, `USER_ID`, `PRODUCT_ID`, `TRX_ID`, `CONTENT`, `created_at`, `updated_at`) VALUES
(2,	2,	1,	1,	'Excellent custom website, highly recommended!',	'2024-08-06 16:57:12.000',	'2024-08-06 16:57:12.000'),
(3,	3,	3,	1,	'Still waiting for the project to complete.',	'2024-08-06 16:57:47.000',	'2024-08-06 16:57:47.000'),
(4,	4,	4,	4,	'Great template for artists.',	'2024-08-06 16:57:57.000',	'2024-08-06 16:57:57.000'),
(5,	5,	5,	5,	'Perfect for our business needs.',	'2024-08-06 16:58:11.000',	'2024-08-06 16:58:11.000'),
(7,	7,	1,	1,	'lorem ipsumee',	'2024-08-06 17:02:43.000',	'2024-08-06 17:02:43.000'),
(11,	7,	4,	5,	'test aja updated 23wx',	'2024-08-06 17:20:11.632',	'2024-08-06 17:34:25.394');

DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `ID` bigint unsigned NOT NULL AUTO_INCREMENT,
  `NAME` longtext COLLATE utf8mb4_general_ci,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `role` (`ID`, `NAME`, `created_at`, `updated_at`) VALUES
(1,	'Admin',	'2024-08-06 16:34:18.000',	'2024-08-06 16:34:18.000'),
(2,	'Developer',	'2024-08-06 16:34:33.000',	'2024-08-06 16:34:33.000'),
(3,	'Client',	'2024-08-06 16:34:42.000',	'2024-08-06 16:34:42.000'),
(4,	'Reviewer',	'2024-08-06 16:34:50.000',	'2024-08-06 16:34:50.000'),
(5,	'Guest',	'2024-08-06 16:34:54.000',	'2024-08-06 16:34:54.000');

DROP TABLE IF EXISTS `transaksi`;
CREATE TABLE `transaksi` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `trx_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `product_id` bigint unsigned DEFAULT NULL,
  `user_id` bigint unsigned DEFAULT NULL,
  `status` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `total` bigint DEFAULT NULL,
  `pay_date` datetime DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_TRANSAKSI_trx_id` (`trx_id`),
  KEY `idx_TRANSAKSI_product_id` (`product_id`),
  KEY `idx_TRANSAKSI_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `transaksi` (`id`, `trx_id`, `product_id`, `user_id`, `status`, `total`, `pay_date`, `created_at`, `updated_at`) VALUES
(1,	'TRX1001',	1,	1,	'Completed',	15000,	NULL,	'2024-08-06 09:52:39',	'2024-08-06 09:52:39'),
(2,	'TRX1002',	2,	3,	'Pending',	30000,	NULL,	'2024-08-06 09:53:02',	'2024-08-06 09:53:02'),
(3,	'TRX1003',	3,	3,	'Completed',	40000,	NULL,	'2024-08-06 09:53:19',	'2024-08-06 09:53:19'),
(4,	'TRX1004',	4,	4,	'Completed',	60000,	NULL,	'2024-08-06 09:53:32',	'2024-08-06 09:53:32'),
(5,	'TRX1005',	5,	5,	'Pending',	700000,	NULL,	'2024-08-06 09:53:53',	'2024-08-06 09:53:53');

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `ID` bigint unsigned NOT NULL AUTO_INCREMENT,
  `USERNAME` longtext COLLATE utf8mb4_general_ci,
  `PASSWORD` longtext COLLATE utf8mb4_general_ci,
  `EMAIL` longtext COLLATE utf8mb4_general_ci,
  `role_id` bigint unsigned DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`ID`),
  KEY `fk_USER_role` (`role_id`),
  CONSTRAINT `fk_USER_role` FOREIGN KEY (`role_id`) REFERENCES `role` (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `user` (`ID`, `USERNAME`, `PASSWORD`, `EMAIL`, `role_id`, `created_at`, `updated_at`) VALUES
(2,	'johnDoe',	'$2a$10$twW1DX0tuMHA.L5MIt6NGOp4ctSpx4Li5Yt9pRdbm37N6d40KvW3K',	'john.doe@example.com',	1,	'2024-08-06 16:34:59.687',	'2024-08-06 16:34:59.687'),
(3,	'saraSmith',	'$2a$10$SjZPC.fHVXonf2tdNd4SQuELl6yNJI87v5THH8VArz660MdYq3BSC',	'sara.smith@example.com',	1,	'2024-08-06 16:35:56.477',	'2024-08-06 16:35:56.477'),
(4,	'alexJones',	'$2a$10$ZAa3jV6Iqsa78W7gdiQtnOvtOhBdz8oK7nWqSYScgSpmBrQuCcgv2',	'alex.jones@example.com',	1,	'2024-08-06 16:36:15.092',	'2024-08-06 16:36:15.092'),
(5,	'lisaWhite',	'$2a$10$tGZ1tQKd1RkYXWAWUMNdlOfjQFkCwSN6pQCdNMEwXr94niJm4c73m',	'lisa.white@example.com',	1,	'2024-08-06 16:36:26.348',	'2024-08-06 16:36:26.348'),
(6,	'markBrown',	'$2a$10$prJ.h062Su736P7hFaDcbuub2hmkiTJwecodwmD8y3rjxUGmo2Cu6',	'mark.brown@example.com',	1,	'2024-08-06 16:36:37.102',	'2024-08-06 16:36:37.102'),
(7,	'kotone',	'$2a$10$ir3kpuEjweooT..rWQBzo.Ad7gGOeoW0KWwp5pK.TDykhq2jCE5ZO',	'kotone@test.com',	1,	'2024-08-06 16:37:11.042',	'2024-08-06 16:37:11.042');

-- 2024-08-06 10:39:08