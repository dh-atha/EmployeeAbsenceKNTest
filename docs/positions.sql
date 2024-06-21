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

 Date: 21/06/2024 09:38:32
*/


-- ----------------------------
-- Table structure for positions
-- ----------------------------
DROP TABLE IF EXISTS "public"."positions";
CREATE TABLE "public"."positions" (
  "position_id" int4 NOT NULL DEFAULT nextval('positions_position_id_seq'::regclass),
  "department_id" int4,
  "position_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "created_by" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_by" varchar(255) COLLATE "pg_catalog"."default",
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of positions
-- ----------------------------
INSERT INTO "public"."positions" VALUES (1, 1, 'Manager', '2024-06-20 10:00:00', 'admin', '2024-06-20 10:00:00', 'admin', NULL);
INSERT INTO "public"."positions" VALUES (2, 2, 'dummy', '2024-06-21 00:25:15.699412', 'john', '2024-06-21 00:29:16.795022', 'john', NULL);
INSERT INTO "public"."positions" VALUES (3, 2, 'Report', '2024-06-21 00:25:20.36968', 'john', '2024-06-21 00:25:20.36968', 'john', '2024-06-21 00:30:39.195624');
INSERT INTO "public"."positions" VALUES (4, 2, 'Golang Developer', '2024-06-21 02:18:07.669087', 'databetul', '2024-06-21 02:18:07.669087', 'databetul', NULL);

-- ----------------------------
-- Primary Key structure for table positions
-- ----------------------------
ALTER TABLE "public"."positions" ADD CONSTRAINT "positions_pkey" PRIMARY KEY ("position_id");

-- ----------------------------
-- Foreign Keys structure for table positions
-- ----------------------------
ALTER TABLE "public"."positions" ADD CONSTRAINT "positions_department_id_fkey" FOREIGN KEY ("department_id") REFERENCES "public"."departments" ("department_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
