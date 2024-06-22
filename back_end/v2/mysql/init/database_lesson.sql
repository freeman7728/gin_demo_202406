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

 Date: 22/06/2024 16:58:10
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
INSERT INTO `account` VALUES (8, '2024060008', 'password123');
INSERT INTO `account` VALUES (24, '2024060024', '265823');
INSERT INTO `account` VALUES (29, '2024060029', '299223');
INSERT INTO `account` VALUES (31, '2024060031', '123123');

-- ----------------------------
-- Table structure for auth
-- ----------------------------
DROP TABLE IF EXISTS `auth`;
CREATE TABLE `auth`  (
  `id` int UNSIGNED NOT NULL,
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `level` int NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `unique_token`(`token` ASC) USING BTREE,
  CONSTRAINT `fk_auth_employee_1` FOREIGN KEY (`id`) REFERENCES `employee` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth
-- ----------------------------

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
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of detail
-- ----------------------------
INSERT INTO `detail` VALUES (6, 6, 0, 15153.60, '', 3);
INSERT INTO `detail` VALUES (7, 6, 123, 15153.60, '12312', 3);
INSERT INTO `detail` VALUES (8, 6, 123, 15153.60, '12312', 3);
INSERT INTO `detail` VALUES (9, 6, 123, 15153.60, '12312', 3);
INSERT INTO `detail` VALUES (10, 6, 123, 15153.60, '12312', 3);
INSERT INTO `detail` VALUES (11, 6, 123, 15153.60, '12312', 3);
INSERT INTO `detail` VALUES (14, 6, 123, 15153.60, '12312', 6);
INSERT INTO `detail` VALUES (15, 6, 1000, 15153.60, '', 6);
INSERT INTO `detail` VALUES (16, 6, 123, 15153.60, '12312', 6);
INSERT INTO `detail` VALUES (17, 6, 123, 15153.60, '12312', 6);
INSERT INTO `detail` VALUES (18, 6, 123, 15153.60, '12312', 6);
INSERT INTO `detail` VALUES (19, 6, 123, 15153.60, '12312', 6);
INSERT INTO `detail` VALUES (24, 21, 10, 199.90, 'First detail item', 9);

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
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uniqe_tel`(`tel` ASC) USING BTREE COMMENT '手机号唯一'
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of employee
-- ----------------------------
INSERT INTO `employee` VALUES (8, 'John Doe', '123456789', 5000.00, 'New employee', NULL, 1);
INSERT INTO `employee` VALUES (12, '123', '123123', 123.33, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (14, '123', '1236123', 123.33, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (16, '123', '12361123', 123.33, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (17, '123', '1123', 123.33, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (18, '123', '123323', 123.33, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (24, '123', '1265823', 123.00, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (29, '123', '111299223', 123.00, '1231', NULL, 1231);
INSERT INTO `employee` VALUES (31, 'name', '123345213', 12312.00, '1231', NULL, 2);

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
) ENGINE = InnoDB AUTO_INCREMENT = 15 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of list
-- ----------------------------
INSERT INTO `list` VALUES (3, 615, 90921.60, '2023-01-01T12:00:00Z', 'This is a note for List 1.', 14);
INSERT INTO `list` VALUES (4, 123, 123.34, '123123', '12312', 8);
INSERT INTO `list` VALUES (5, 123, 234.00, '123', '23', 18);
INSERT INTO `list` VALUES (6, 1615, 90921.60, '123123', '12312', 8);
INSERT INTO `list` VALUES (7, 123, 123.34, '123123', '12312', 8);
INSERT INTO `list` VALUES (8, 123, 123.34, '123123', '12312', 8);
INSERT INTO `list` VALUES (9, 10, 199.90, '123123', '12312', 8);
INSERT INTO `list` VALUES (14, 20, 299.99, '2023-01-02T15:00:00Z', 'This is a note for List 2.', 14);

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
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of producer
-- ----------------------------
INSERT INTO `producer` VALUES (3, 'Producer Tw', 'Prod2', '5678 Second St, City B, EX 67890', '222-222-2222', 'contact2@example.com', 'Bob', '222-333-4444', 'Second producer note.');
INSERT INTO `producer` VALUES (6, 'Producer Three', 'Prod3', '9101 Third St, City C, EX 10112', '333-333-3333', 'contact3@example.com', 'Charlie', '333-444-5555', 'Third producer note.');
INSERT INTO `producer` VALUES (7, 'Producer One', 'Prod1', '1234 First St, City A, EX 12345', '111-111-1111', 'contact1@example.com', 'Alice', '111-222-3333', 'First producer note.');
INSERT INTO `producer` VALUES (8, 'Producer Two', 'Prod2', '5678 Second St, City B, EX 67890', '222-222-2222', 'contact2@example.com', 'Bob', '222-333-4444', 'Second producer note.');
INSERT INTO `producer` VALUES (9, 'Producer Three', 'Prod3', '9101 Third St, City C, EX 10112', '333-333-3333', 'contact3@example.com', 'Charlie', '333-444-5555', 'Third producer note.');
INSERT INTO `producer` VALUES (10, 'Producer One', 'Prod1', '1234 First St, City A, EX 12345', '111-111-1111', 'contact1@example.com', 'Alice', '111-222-3333', 'First producer note.');
INSERT INTO `producer` VALUES (11, 'Producer Two', 'Prod2', '5678 Second St, City B, EX 67890', '222-222-2222', 'contact2@example.com', 'Bob', '222-333-4444', 'Second producer note.');
INSERT INTO `producer` VALUES (12, 'Producer Three', 'Prod3', '9101 Third St, City C, EX 10112', '333-333-3333', 'contact3@example.com', 'Charlie', '333-444-5555', 'Third producer note.');
INSERT INTO `producer` VALUES (13, 'Producer One', 'Prod1', '1234 First St, City A, EX 12345', '111-111-1111', 'contact1@example.com', 'Alice', '111-222-3333', 'First producer note.');
INSERT INTO `producer` VALUES (14, 'Producer Two', 'Prod2', '5678 Second St, City B, EX 67890', '222-222-2222', 'contact2@example.com', 'Bob', '222-333-4444', 'Second producer note.');
INSERT INTO `producer` VALUES (15, 'Producer Three', 'Prod3', '9101 Third St, City C, EX 10112', '333-333-3333', 'contact3@example.com', 'Charlie', '333-444-5555', 'Third producer note.');
INSERT INTO `producer` VALUES (16, 'Producer One', 'Prod1', '1234 First St, City A, EX 12345', '111-111-1111', 'contact1@example.com', 'Alice', '111-222-3333', 'First producer note.');
INSERT INTO `producer` VALUES (17, 'Producer Two', 'Prod2', '5678 Second St, City B, EX 67890', '222-222-2222', 'contact2@example.com', 'Bob', '222-333-4444', 'Second producer note.');
INSERT INTO `producer` VALUES (18, 'Producer Three', 'Prod3', '9101 Third St, City C, EX 10112', '333-333-3333', 'contact3@example.com', 'Charlie', '333-444-5555', 'Third producer note.');

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
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of product
-- ----------------------------
INSERT INTO `product` VALUES (6, '1231', 123.20, '1231sda', '123as', 3);
INSERT INTO `product` VALUES (7, 'Product 1', 19.99, 'This is the introduction for Product 1.', 'Note for Product 1', 9);
INSERT INTO `product` VALUES (8, '1231', 123.20, '1231sda', '123as', 3);
INSERT INTO `product` VALUES (9, 'Product 1', 19.99, 'This is the introduction for Product 1.', 'Note for Product 1', 9);
INSERT INTO `product` VALUES (14, 'Product 1', 19.99, 'This is the introduction for Product 1.', 'Note for Product 1', 3);
INSERT INTO `product` VALUES (15, 'Product 2', 29.99, 'This is the introduction for Product 2.', 'Note for Product 2', 3);
INSERT INTO `product` VALUES (16, 'Product 1', 19.99, 'This is the introduction for Product 1.', 'Note for Product 1', 3);
INSERT INTO `product` VALUES (18, 'Product 1', 19.99, 'This is the introduction for Product 1.', 'Note for Product 1', 3);
INSERT INTO `product` VALUES (21, 'Product 1', 19.99, 'This is the introduction for Product 1.', 'Note for Product 1', 3);

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
    IN p_level INT,
    OUT p_account VARCHAR(255))
BEGIN
    DECLARE last_id INT;
    DECLARE current_year VARCHAR(4);
    DECLARE current_month VARCHAR(2);
    DECLARE id_suffix VARCHAR(4);
    
    -- 插入新员工记录
    INSERT INTO employee (name, tel, salary, note, level)
    VALUES (p_name, p_tel, p_salary, p_note, p_level);
    
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
DROP TRIGGER IF EXISTS `check_level_before_update`;
delimiter ;;
CREATE TRIGGER `check_level_before_update` BEFORE UPDATE ON `employee` FOR EACH ROW BEGIN
   IF NEW.level >= 3 THEN
      SET NEW.level = 2;
   END IF;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
