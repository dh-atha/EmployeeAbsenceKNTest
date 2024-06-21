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

 Date: 21/06/2024 09:38:24
*/


-- ----------------------------
-- Table structure for locations
-- ----------------------------
DROP TABLE IF EXISTS "public"."locations";
CREATE TABLE "public"."locations" (
  "location_id" int4 NOT NULL DEFAULT nextval('locations_location_id_seq'::regclass),
  "location_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "created_by" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_by" varchar(255) COLLATE "pg_catalog"."default",
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of locations
-- ----------------------------
INSERT INTO "public"."locations" VALUES (1, 'Head Office', '2024-06-20 10:00:00', 'admin', '2024-06-20 10:00:00', 'admin', NULL);
INSERT INTO "public"."locations" VALUES (2, 'coworking', '2024-06-21 00:38:17.535281', 'john', '2024-06-21 00:39:57.184012', 'john', NULL);
INSERT INTO "public"."locations" VALUES (3, 'HQ', '2024-06-21 00:38:46.464408', 'john', '2024-06-21 00:38:46.464408', 'john', '2024-06-21 00:40:31.228101');

-- ----------------------------
-- Primary Key structure for table locations
-- ----------------------------
ALTER TABLE "public"."locations" ADD CONSTRAINT "locations_pkey" PRIMARY KEY ("location_id");
