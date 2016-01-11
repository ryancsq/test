/*
Navicat MySQL Data Transfer

Source Server         : 500wan
Source Server Version : 50616
Source Host           : rds3bhb1ed059c58i02wo.mysql.rds.aliyuncs.com:3306
Source Database       : new_ha

Target Server Type    : MYSQL
Target Server Version : 50616
File Encoding         : 65001

Date: 2016-01-05 11:00:03
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `pk_pan_map`
-- ----------------------------
DROP TABLE IF EXISTS `pk_pan_map`;
CREATE TABLE `pk_pan_map` (
  `pan_id` int(11) NOT NULL AUTO_INCREMENT,
  `pan_desc` varchar(50) NOT NULL,
  `pan_value` float(8,2) NOT NULL,
  PRIMARY KEY (`pan_id`),
  KEY `desc` (`pan_desc`)
) ENGINE=InnoDB AUTO_INCREMENT=64 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pk_pan_map
-- ----------------------------
INSERT INTO `pk_pan_map` VALUES ('1', '受半球/一球', '0.75');
INSERT INTO `pk_pan_map` VALUES ('2', '受半球', '0.50');
INSERT INTO `pk_pan_map` VALUES ('3', '受一球', '1.00');
INSERT INTO `pk_pan_map` VALUES ('4', '受球半/两球', '1.75');
INSERT INTO `pk_pan_map` VALUES ('5', '一球', '-1.00');
INSERT INTO `pk_pan_map` VALUES ('6', '受球半', '1.50');
INSERT INTO `pk_pan_map` VALUES ('7', '平手', '0.00');
INSERT INTO `pk_pan_map` VALUES ('8', '平手/半球', '-0.25');
INSERT INTO `pk_pan_map` VALUES ('9', '球半', '-1.50');
INSERT INTO `pk_pan_map` VALUES ('10', '受平手/半球', '0.25');
INSERT INTO `pk_pan_map` VALUES ('11', '半球', '-0.50');
INSERT INTO `pk_pan_map` VALUES ('12', '受两球半', '2.50');
INSERT INTO `pk_pan_map` VALUES ('13', '受三球', '3.00');
INSERT INTO `pk_pan_map` VALUES ('14', '受四球', '4.00');
INSERT INTO `pk_pan_map` VALUES ('15', '受两球半/三球', '2.75');
INSERT INTO `pk_pan_map` VALUES ('16', '受两球/两球半', '2.25');
INSERT INTO `pk_pan_map` VALUES ('17', '受三球半', '3.50');
INSERT INTO `pk_pan_map` VALUES ('18', '半球/一球', '-0.75');
INSERT INTO `pk_pan_map` VALUES ('19', '受一球/球半', '1.25');
INSERT INTO `pk_pan_map` VALUES ('20', '受两球', '2.00');
INSERT INTO `pk_pan_map` VALUES ('21', '球半/两球', '-1.75');
INSERT INTO `pk_pan_map` VALUES ('22', '一球/球半', '-1.25');
INSERT INTO `pk_pan_map` VALUES ('23', '两球', '-2.00');
INSERT INTO `pk_pan_map` VALUES ('24', '两球半', '-2.50');
INSERT INTO `pk_pan_map` VALUES ('25', '两球半/三球', '-2.75');
INSERT INTO `pk_pan_map` VALUES ('26', '两球/两球半', '-2.25');
INSERT INTO `pk_pan_map` VALUES ('27', '三球/三球半', '-3.25');
INSERT INTO `pk_pan_map` VALUES ('28', '三球', '-3.00');
INSERT INTO `pk_pan_map` VALUES ('29', '三球半', '-3.50');
INSERT INTO `pk_pan_map` VALUES ('30', '四球半', '-4.50');
INSERT INTO `pk_pan_map` VALUES ('31', '四球', '-4.00');
INSERT INTO `pk_pan_map` VALUES ('32', '三球半/四球', '-3.75');
INSERT INTO `pk_pan_map` VALUES ('33', '受三球半/四球', '3.75');
INSERT INTO `pk_pan_map` VALUES ('34', '受四球/四球半', '4.25');
INSERT INTO `pk_pan_map` VALUES ('35', '受三球/三球半', '3.25');
INSERT INTO `pk_pan_map` VALUES ('36', '受四球半', '4.50');
INSERT INTO `pk_pan_map` VALUES ('37', '五球', '-5.00');
INSERT INTO `pk_pan_map` VALUES ('38', '四球半/五球', '-4.75');
INSERT INTO `pk_pan_map` VALUES ('39', '七球半/八球', '-7.75');
INSERT INTO `pk_pan_map` VALUES ('40', '七球半', '-7.50');
INSERT INTO `pk_pan_map` VALUES ('41', '八球', '-8.00');
INSERT INTO `pk_pan_map` VALUES ('42', '七球', '-7.00');
INSERT INTO `pk_pan_map` VALUES ('43', '六球半', '-6.50');
INSERT INTO `pk_pan_map` VALUES ('44', '五球半', '-5.50');
INSERT INTO `pk_pan_map` VALUES ('45', '五球/五球半', '-5.25');
INSERT INTO `pk_pan_map` VALUES ('46', '五球半/六球', '-6.75');
INSERT INTO `pk_pan_map` VALUES ('47', '六球', '-6.00');
INSERT INTO `pk_pan_map` VALUES ('48', '四球/四球半', '-4.25');
INSERT INTO `pk_pan_map` VALUES ('49', '八球半', '-8.50');
INSERT INTO `pk_pan_map` VALUES ('50', '受五球半', '5.50');
INSERT INTO `pk_pan_map` VALUES ('51', '受五球', '5.00');
INSERT INTO `pk_pan_map` VALUES ('52', '受四球半/五球', '4.75');
INSERT INTO `pk_pan_map` VALUES ('53', '受五球/五球半', '5.25');
INSERT INTO `pk_pan_map` VALUES ('54', '受五球半/六球', '5.75');
INSERT INTO `pk_pan_map` VALUES ('55', '受六球', '6.00');
INSERT INTO `pk_pan_map` VALUES ('56', '受六球半', '6.50');
INSERT INTO `pk_pan_map` VALUES ('57', '受七球半', '7.50');
INSERT INTO `pk_pan_map` VALUES ('58', '受六球/六球半', '6.25');
INSERT INTO `pk_pan_map` VALUES ('59', '受六球半/七球', '6.75');
INSERT INTO `pk_pan_map` VALUES ('60', '受七球', '7.00');
INSERT INTO `pk_pan_map` VALUES ('61', '六球/六球半', '-6.25');
INSERT INTO `pk_pan_map` VALUES ('62', '受九球', '9.00');
INSERT INTO `pk_pan_map` VALUES ('63', '九球', '-9.00');
