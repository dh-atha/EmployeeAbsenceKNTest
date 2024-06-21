/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : PostgreSQL
 Source Server Version : 150003 (150003)
 Source Host           : localhost:5432
 Source Catalog        : absence
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 150003 (150003)
 File Encoding         : 65001

 Date: 21/06/2024 09:37:58
*/


-- ----------------------------
-- Table structure for departments
-- ----------------------------
DROP TABLE IF EXISTS "public"."departments";
CREATE TABLE "public"."departments" (
  "department_id" int4 NOT NULL DEFAULT nextval('departments_department_id_seq'::regclass),
  "department_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "created_by" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_by" varchar(255) COLLATE "pg_catalog"."default",
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of departments
-- ----------------------------
INSERT INTO "public"."departments" VALUES (1, 'Human Resources', '2024-06-20 10:00:00', 'admin', '2024-06-20 10:00:00', 'admin', NULL);
INSERT INTO "public"."departments" VALUES (2, 'Backend', '2024-06-21 00:06:19.512653', 'john', '2024-06-21 00:12:18.70502', 'john', NULL);
INSERT INTO "public"."departments" VALUES (3, 'IT', '2024-06-21 00:09:08.348731', 'john', '2024-06-21 00:09:08.348731', 'john', '2024-06-21 00:12:57.826993');

-- ----------------------------
-- Primary Key structure for table departments
-- ----------------------------
ALTER TABLE "public"."departments" ADD CONSTRAINT "departments_pkey" PRIMARY KEY ("department_id");
