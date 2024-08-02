/*
 Navicat Premium Data Transfer

 Source Server         : ecommerce - bootcamp
 Source Server Type    : MySQL
 Source Server Version : 101108 (10.11.8-MariaDB-cll-lve)
 Source Host           : 151.106.117.51:3306
 Source Schema         : u628725475_ecommerce

 Target Server Type    : MySQL
 Target Server Version : 101108 (10.11.8-MariaDB-cll-lve)
 File Encoding         : 65001

 Date: 02/08/2024 10:29:37
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for CATEGORY
-- ----------------------------
DROP TABLE IF EXISTS `CATEGORY`;
CREATE TABLE `CATEGORY`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of CATEGORY
-- ----------------------------
INSERT INTO `CATEGORY` VALUES (1, 'Website Development');
INSERT INTO `CATEGORY` VALUES (2, 'E-commerce Solutions');
INSERT INTO `CATEGORY` VALUES (3, 'Portfolio Templates');
INSERT INTO `CATEGORY` VALUES (4, 'Blog Templates');
INSERT INTO `CATEGORY` VALUES (5, 'Business Templates');

-- ----------------------------
-- Table structure for DETAILTRANSAKSI
-- ----------------------------
DROP TABLE IF EXISTS `DETAILTRANSAKSI`;
CREATE TABLE `DETAILTRANSAKSI`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `TRX_ID` int NULL DEFAULT NULL,
  `PRODUCT_ID` int NULL DEFAULT NULL,
  `QTY` int NULL DEFAULT NULL,
  `TOTAL` decimal(10, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `TRX_ID`(`TRX_ID` ASC) USING BTREE,
  INDEX `PRODUCT_ID`(`PRODUCT_ID` ASC) USING BTREE,
  CONSTRAINT `DETAILTRANSAKSI_ibfk_1` FOREIGN KEY (`TRX_ID`) REFERENCES `TRANSAKSI` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `DETAILTRANSAKSI_ibfk_2` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `PRODUCT` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of DETAILTRANSAKSI
-- ----------------------------
INSERT INTO `DETAILTRANSAKSI` VALUES (1, 1, 1, 1, 1500.00);
INSERT INTO `DETAILTRANSAKSI` VALUES (2, 2, 2, 1, 2500.00);
INSERT INTO `DETAILTRANSAKSI` VALUES (3, 3, 3, 1, 50.00);
INSERT INTO `DETAILTRANSAKSI` VALUES (4, 4, 4, 1, 25.00);
INSERT INTO `DETAILTRANSAKSI` VALUES (5, 5, 5, 1, 75.00);

-- ----------------------------
-- Table structure for PRODUCT
-- ----------------------------
DROP TABLE IF EXISTS `PRODUCT`;
CREATE TABLE `PRODUCT`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `PRICE` decimal(10, 2) NOT NULL,
  `CATEGORY_ID` int NULL DEFAULT NULL,
  `SLUG` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `IMAGE_URL` int NULL DEFAULT NULL,
  `STATUS` int NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `CATEGORY_ID`(`CATEGORY_ID` ASC) USING BTREE,
  CONSTRAINT `PRODUCT_ibfk_1` FOREIGN KEY (`CATEGORY_ID`) REFERENCES `CATEGORY` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of PRODUCT
-- ----------------------------
INSERT INTO `PRODUCT` VALUES (1, 'Custom Website Design', 1500.00, 1, 'custom-website-design', NULL, NULL);
INSERT INTO `PRODUCT` VALUES (2, 'Online Store Setup', 2500.00, 2, 'online-store-setup', NULL, NULL);
INSERT INTO `PRODUCT` VALUES (3, 'Art Portfolio Template', 50.00, 3, 'art-portfolio-template', NULL, NULL);
INSERT INTO `PRODUCT` VALUES (4, 'Personal Blog Template', 25.00, 4, 'personal-blog-template', NULL, NULL);
INSERT INTO `PRODUCT` VALUES (5, 'Corporate Business Template', 75.00, 5, 'corporate-business-template', NULL, NULL);

-- ----------------------------
-- Table structure for PROFILE
-- ----------------------------
DROP TABLE IF EXISTS `PROFILE`;
CREATE TABLE `PROFILE`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `GENDER` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `CITY` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `DATE` date NULL DEFAULT NULL,
  `ADDRESS` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `PHONE` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `USER_ID` int NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `USER_ID`(`USER_ID` ASC) USING BTREE,
  CONSTRAINT `PROFILE_ibfk_1` FOREIGN KEY (`USER_ID`) REFERENCES `USER` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of PROFILE
-- ----------------------------
INSERT INTO `PROFILE` VALUES (1, 'John Doe', 'Male', 'New York', '2023-01-15', '123 6th St', '123-456-7890', 1);
INSERT INTO `PROFILE` VALUES (2, 'Sara Smith', 'Female', 'Los Angeles', '2023-01-20', '789 Main St', '234-567-8901', 2);
INSERT INTO `PROFILE` VALUES (3, 'Alex Jones', 'Male', 'Chicago', '2023-02-10', '456 Elm St', '345-678-9012', 3);
INSERT INTO `PROFILE` VALUES (4, 'Lisa White', 'Female', 'Houston', '2023-03-05', '123 Oak St', '456-789-0123', 4);
INSERT INTO `PROFILE` VALUES (5, 'Mark Brown', 'Male', 'Philadelphia', '2023-04-16', '789 Maple St', '567-890-1234', 5);

-- ----------------------------
-- Table structure for REVIEW
-- ----------------------------
DROP TABLE IF EXISTS `REVIEW`;
CREATE TABLE `REVIEW`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `PRODUCT_ID` int NULL DEFAULT NULL,
  `USER_ID` int NULL DEFAULT NULL,
  `TRX_ID` int NULL DEFAULT NULL,
  `CONTENT` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `PRODUCT_ID`(`PRODUCT_ID` ASC) USING BTREE,
  INDEX `USER_ID`(`USER_ID` ASC) USING BTREE,
  INDEX `TRX_ID`(`TRX_ID` ASC) USING BTREE,
  CONSTRAINT `REVIEW_ibfk_1` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `PRODUCT` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `REVIEW_ibfk_2` FOREIGN KEY (`USER_ID`) REFERENCES `USER` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `REVIEW_ibfk_3` FOREIGN KEY (`TRX_ID`) REFERENCES `TRANSAKSI` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of REVIEW
-- ----------------------------
INSERT INTO `REVIEW` VALUES (1, 1, 1, 1, 'Excellent custom website, highly recommended!');
INSERT INTO `REVIEW` VALUES (2, 2, 2, 2, 'Still waiting for the project to complete.');
INSERT INTO `REVIEW` VALUES (3, 3, 3, 3, 'Great template for artists.');
INSERT INTO `REVIEW` VALUES (4, 4, 4, 4, 'Simple and clean blog template.');
INSERT INTO `REVIEW` VALUES (5, 5, 5, 5, 'Perfect for our business needs.');

-- ----------------------------
-- Table structure for ROLE
-- ----------------------------
DROP TABLE IF EXISTS `ROLE`;
CREATE TABLE `ROLE`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `NAME` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ROLE
-- ----------------------------
INSERT INTO `ROLE` VALUES (1, 'Admin');
INSERT INTO `ROLE` VALUES (2, 'Developer');
INSERT INTO `ROLE` VALUES (3, 'Client');
INSERT INTO `ROLE` VALUES (4, 'Reviewer');
INSERT INTO `ROLE` VALUES (5, 'Guest');

-- ----------------------------
-- Table structure for TRANSAKSI
-- ----------------------------
DROP TABLE IF EXISTS `TRANSAKSI`;
CREATE TABLE `TRANSAKSI`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `TRX_ID` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `PRODUCT_ID` int NULL DEFAULT NULL,
  `USER_ID` int NULL DEFAULT NULL,
  `STATUS` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `TOTAL` decimal(10, 2) NULL DEFAULT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  INDEX `PRODUCT_ID`(`PRODUCT_ID` ASC) USING BTREE,
  INDEX `USER_ID`(`USER_ID` ASC) USING BTREE,
  CONSTRAINT `TRANSAKSI_ibfk_1` FOREIGN KEY (`PRODUCT_ID`) REFERENCES `PRODUCT` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `TRANSAKSI_ibfk_2` FOREIGN KEY (`USER_ID`) REFERENCES `USER` (`ID`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of TRANSAKSI
-- ----------------------------
INSERT INTO `TRANSAKSI` VALUES (1, 'TRX1001', 1, 1, 'Completed', 1500.00);
INSERT INTO `TRANSAKSI` VALUES (2, 'TRX1002', 2, 2, 'Pending', 2500.00);
INSERT INTO `TRANSAKSI` VALUES (3, 'TRX1003', 3, 3, 'Completed', 50.00);
INSERT INTO `TRANSAKSI` VALUES (4, 'TRX1004', 4, 4, 'Completed', 25.00);
INSERT INTO `TRANSAKSI` VALUES (5, 'TRX1005', 5, 5, 'Pending', 75.00);

-- ----------------------------
-- Table structure for USER
-- ----------------------------
DROP TABLE IF EXISTS `USER`;
CREATE TABLE `USER`  (
  `ID` int NOT NULL AUTO_INCREMENT,
  `USERNAME` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `PASSWORD` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `EMAIL` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`ID`) USING BTREE,
  UNIQUE INDEX `EMAIL`(`EMAIL` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of USER
-- ----------------------------
INSERT INTO `USER` VALUES (1, 'johnDoe', 'john123', 'john.doe@example.com');
INSERT INTO `USER` VALUES (2, 'saraSmith', 'sara123', 'sara.smith@example.com');
INSERT INTO `USER` VALUES (3, 'alexJones', 'alex123', 'alex.jones@example.com');
INSERT INTO `USER` VALUES (4, 'lisaWhite', 'lisa123', 'lisa.white@example.com');
INSERT INTO `USER` VALUES (5, 'markBrown', 'mark123', 'mark.brown@example.com');

SET FOREIGN_KEY_CHECKS = 1;
