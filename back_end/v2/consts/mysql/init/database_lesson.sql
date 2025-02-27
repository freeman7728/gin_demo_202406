/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80033 (8.0.33)
 Source Host           : localhost:3306
 Source Schema         : database_lesson

 Target Server Type    : MySQL
 Target Server Version : 80033 (8.0.33)
 File Encoding         : 65001

 Date: 23/06/2024 21:20:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` int UNSIGNED NOT NULL,
  `account` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqe_account`(`account` ASC) USING BTREE,
  CONSTRAINT `fk_account_employee_1` FOREIGN KEY (`id`) REFERENCES `employee` (`id`) ON DELETE CASCADE ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES (47, '2024060047', '490652');
INSERT INTO `account` VALUES (48, '2024060048', '426280');
INSERT INTO `account` VALUES (49, '2024060049', '747992');
INSERT INTO `account` VALUES (50, '2024060050', '570625');
INSERT INTO `account` VALUES (51, '2024060051', '381196');

-- ----------------------------
-- Table structure for detail
-- ----------------------------
DROP TABLE IF EXISTS `detail`;
CREATE TABLE `detail`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `product_id` int UNSIGNED NOT NULL,
  `quantity` int UNSIGNED NULL DEFAULT NULL,
  `total_price` decimal(10, 2) NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `list_id` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_detail_product_1`(`product_id` ASC) USING BTREE,
  INDEX `fk_detail_list_1`(`list_id` ASC) USING BTREE,
  CONSTRAINT `fk_detail_list_1` FOREIGN KEY (`list_id`) REFERENCES `list` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_detail_product_1` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of detail
-- ----------------------------
INSERT INTO `detail` VALUES (26, 26, 50, 4999.50, 'First detail item', 15);
INSERT INTO `detail` VALUES (27, 35, 25, 2249.75, 'Second detail item', 15);
INSERT INTO `detail` VALUES (28, 27, 10, 499.90, 'Third detail item', 15);
INSERT INTO `detail` VALUES (29, 26, 100, 9999.00, 'Fourth detail item', 16);
INSERT INTO `detail` VALUES (30, 30, 75, 9749.25, 'Fifth detail item', 19);

-- ----------------------------
-- Table structure for employee
-- ----------------------------
DROP TABLE IF EXISTS `employee`;
CREATE TABLE `employee`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `salary` decimal(10, 2) NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `deleted_at` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `level` int NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email_is_auth` int NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 52 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee
-- ----------------------------
INSERT INTO `employee` VALUES (47, '邓秀英', '93928490652', 4129.00, '老板', NULL, 0, '1357880399@qq.com', 0);
INSERT INTO `employee` VALUES (48, '曹晓明', '10468426280', 3152.00, '管理员', NULL, 1, 'cxia@icloud.com', 0);
INSERT INTO `employee` VALUES (49, '陈子韬', '04842747992', 3284.00, '新员工', NULL, 2, 'zc6@icloud.com', 0);
INSERT INTO `employee` VALUES (50, '石致远', '77604570625', 3299.00, '新员工', NULL, 2, 'szhiyuan504@icloud.com', 0);
INSERT INTO `employee` VALUES (51, '董岚', '69252381196', 2743.00, '新员工', NULL, 2, 'dolan811@hotmail.com', 0);

-- ----------------------------
-- Table structure for list
-- ----------------------------
DROP TABLE IF EXISTS `list`;
CREATE TABLE `list`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `quantity` int NULL DEFAULT NULL,
  `total_price` decimal(10, 2) NULL DEFAULT NULL,
  `time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `employee_id` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_list_employee_1`(`employee_id` ASC) USING BTREE,
  CONSTRAINT `fk_list_employee_1` FOREIGN KEY (`employee_id`) REFERENCES `employee` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 20 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of list
-- ----------------------------
INSERT INTO `list` VALUES (15, 85, 7749.15, '2024-06-23 10:00:00', 'First list item', 49);
INSERT INTO `list` VALUES (16, 100, 9999.00, '2024-06-23 11:00:00', 'Second list item', 49);
INSERT INTO `list` VALUES (17, NULL, NULL, '2024-06-23 12:00:00', 'Third list item', 49);
INSERT INTO `list` VALUES (18, NULL, NULL, '2024-06-23 13:00:00', 'Fourth list item', 50);
INSERT INTO `list` VALUES (19, 75, 9749.25, '2024-06-23 14:00:00', 'Fifth list item', 51);

-- ----------------------------
-- Table structure for producer
-- ----------------------------
DROP TABLE IF EXISTS `producer`;
CREATE TABLE `producer`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `short_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `contact_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `contact_tel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of producer
-- ----------------------------
INSERT INTO `producer` VALUES (19, 'Producer One', 'Prod1', '1234 First St - City A - EX 12345', '111-111-1111', 'contact1@example.com', 'Alice', '111-222-3333', 'First producer note');
INSERT INTO `producer` VALUES (20, 'Producer Two', 'Prod2', '5678 Second St - City B - EX 67890', '222-222-2222', 'contact2@example.com', 'Bob', '222-333-4444', 'Second producer note');
INSERT INTO `producer` VALUES (21, 'Producer Three', 'Prod3', '9101 Third St - City C - EX 10112', '333-333-3333', 'contact3@example.com', 'Charlie', '333-444-5555', 'Third producer note');

-- ----------------------------
-- Table structure for product
-- ----------------------------
DROP TABLE IF EXISTS `product`;
CREATE TABLE `product`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `price` decimal(10, 2) NOT NULL,
  `introduction` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `producer_id` int UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_product_producer_1`(`producer_id` ASC) USING BTREE,
  CONSTRAINT `fk_product_producer_1` FOREIGN KEY (`producer_id`) REFERENCES `producer` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (26, 'Product 1', 99.99, 'Introduction for Product 1', 'Note for Product 1', 19);
INSERT INTO `product` VALUES (27, 'Product 2', 49.99, 'Introduction for Product 2', 'Note for Product 2', 19);
INSERT INTO `product` VALUES (28, 'Product 3', 199.99, 'Introduction for Product 3', 'Note for Product 3', 19);
INSERT INTO `product` VALUES (29, 'Product 4', 79.99, 'Introduction for Product 4', 'Note for Product 4', 19);
INSERT INTO `product` VALUES (30, 'Product 5', 129.99, 'Introduction for Product 5', 'Note for Product 5', 21);
INSERT INTO `product` VALUES (31, 'Product 6', 149.99, 'Introduction for Product 6', 'Note for Product 6', 21);
INSERT INTO `product` VALUES (32, 'Product 7', 199.99, 'Introduction for Product 7', 'Note for Product 7', 21);
INSERT INTO `product` VALUES (33, 'Product 8', 39.99, 'Introduction for Product 8', 'Note for Product 8', 20);
INSERT INTO `product` VALUES (34, 'Product 9', 299.99, 'Introduction for Product 9', 'Note for Product 9', 20);
INSERT INTO `product` VALUES (35, 'Product 10', 89.99, 'Introduction for Product 10', 'Note for Product 10', 20);

-- ----------------------------
-- Procedure structure for CreateEmployeeWithAccount
-- ----------------------------
DROP PROCEDURE IF EXISTS `CreateEmployeeWithAccount`;
delimiter ;;
CREATE PROCEDURE `CreateEmployeeWithAccount`(IN p_name VARCHAR(255),
    IN p_tel VARCHAR(255),
    IN p_password VARCHAR(255),
    IN p_salary DECIMAL,
    IN p_note VARCHAR(255),
		IN p_email VARCHAR(255),
    IN p_level INT,
    OUT p_account VARCHAR(255))
BEGIN
    DECLARE last_id INT;
    DECLARE current_year VARCHAR(4);
    DECLARE current_month VARCHAR(2);
    DECLARE id_suffix VARCHAR(4);
    
    -- 插入新员工记录
    INSERT INTO employee (name, tel, salary, note, level,email)
    VALUES (p_name, p_tel, p_salary, p_note, p_level,p_email);
    
    -- 获取最后插入的ID
    SET last_id = LAST_INSERT_ID();
    
    -- 获取当前年份和月份
    SET current_year = YEAR(CURDATE());
    SET current_month = LPAD(MONTH(CURDATE()), 2, '0');
    
    -- 获取ID后四位
    SET id_suffix = LPAD(last_id % 10000, 4, '0');
    
    -- 生成员工账号
    SET p_account = CONCAT(current_year, current_month, id_suffix);
    
    -- 插入账号记录到account表
    INSERT INTO account (account, password, id)
    VALUES (p_account, p_password, last_id);
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table detail
-- ----------------------------
DROP TRIGGER IF EXISTS `before_insert_detail`;
delimiter ;;
CREATE TRIGGER `before_insert_detail` BEFORE INSERT ON `detail` FOR EACH ROW BEGIN
    DECLARE product_price DECIMAL(10, 2);
    
    -- 查询产品的单价
    SELECT price INTO product_price FROM `product` WHERE id = NEW.product_id;
    
    -- 计算总价
    SET NEW.total_price = NEW.quantity * product_price;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table detail
-- ----------------------------
DROP TRIGGER IF EXISTS `after_insert_detail`;
delimiter ;;
CREATE TRIGGER `after_insert_detail` AFTER INSERT ON `detail` FOR EACH ROW BEGIN
    UPDATE `list`
    SET `quantity` = (SELECT SUM(`quantity`) FROM `detail` WHERE `list_id` = NEW.list_id),
        `total_price` = (SELECT SUM(`total_price`) FROM `detail` WHERE `list_id` = NEW.list_id)
    WHERE `id` = NEW.list_id;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table detail
-- ----------------------------
DROP TRIGGER IF EXISTS `after_update_detail`;
delimiter ;;
CREATE TRIGGER `after_update_detail` AFTER UPDATE ON `detail` FOR EACH ROW BEGIN
    UPDATE `list`
    SET `quantity` = (SELECT SUM(`quantity`) FROM `detail` WHERE `list_id` = NEW.list_id),
        `total_price` = (SELECT SUM(`total_price`) FROM `detail` WHERE `list_id` = NEW.list_id)
    WHERE `id` = NEW.list_id;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table detail
-- ----------------------------
DROP TRIGGER IF EXISTS `after_delete_detail`;
delimiter ;;
CREATE TRIGGER `after_delete_detail` AFTER DELETE ON `detail` FOR EACH ROW BEGIN
    UPDATE `list`
    SET `quantity` = (SELECT SUM(`quantity`) FROM `detail` WHERE `list_id` = OLD.list_id),
        `total_price` = (SELECT SUM(`total_price`) FROM `detail` WHERE `list_id` = OLD.list_id)
    WHERE `id` = OLD.list_id;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table employee
-- ----------------------------
DROP TRIGGER IF EXISTS `check_level_before_insert`;
delimiter ;;
CREATE TRIGGER `check_level_before_insert` BEFORE INSERT ON `employee` FOR EACH ROW BEGIN
   IF NEW.level >= 3 THEN
      SET NEW.level = 2;
   END IF;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table employee
-- ----------------------------
DROP TRIGGER IF EXISTS `set_email_auth`;
delimiter ;;
CREATE TRIGGER `set_email_auth` BEFORE INSERT ON `employee` FOR EACH ROW BEGIN
   SET NEW.email_is_auth = FALSE;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table employee
-- ----------------------------
DROP TRIGGER IF EXISTS `check_level_before_update`;
delimiter ;;
CREATE TRIGGER `check_level_before_update` BEFORE UPDATE ON `employee` FOR EACH ROW BEGIN
   IF NEW.level >= 3 THEN
      SET NEW.level = 2;
   END IF;
END
;;
delimiter ;

-- ----------------------------
-- Triggers structure for table employee
-- ----------------------------
DROP TRIGGER IF EXISTS `reset_auth`;
delimiter ;;
CREATE TRIGGER `reset_auth` BEFORE UPDATE ON `employee` FOR EACH ROW BEGIN
   IF NEW.email != OLD.email THEN
      SET NEW.email_is_auth = 0;
   END IF;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
