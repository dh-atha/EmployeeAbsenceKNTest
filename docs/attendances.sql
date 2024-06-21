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

 Date: 21/06/2024 09:37:46
*/


-- ----------------------------
-- Table structure for attendances
-- ----------------------------
DROP TABLE IF EXISTS "public"."attendances";
CREATE TABLE "public"."attendances" (
  "attendance_id" int4 NOT NULL DEFAULT nextval('attendances_attendance_id_seq'::regclass),
  "employee_id" int4,
  "location_id" int4,
  "absent_in" timestamp(6),
  "absent_out" timestamp(6),
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "created_by" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_by" varchar(255) COLLATE "pg_catalog"."default",
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of attendances
-- ----------------------------
INSERT INTO "public"."attendances" VALUES (3, 1, 2, NULL, NULL, '2024-06-21 01:04:24.115084', 'john', '2024-06-21 01:04:24.115084', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (4, 14, 2, NULL, NULL, '2024-06-21 01:06:52.100294', 'john', '2024-06-21 01:06:52.100294', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (2, 1, 2, '2023-06-21 08:00:00', '2023-06-21 09:00:00', '2024-06-21 01:04:23.894269', 'john', '2024-06-21 01:22:30.96585', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (5, 14, 2, NULL, NULL, '2024-06-21 01:25:36.450989', 'john', '2024-06-21 01:25:36.450989', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (6, 14, 2, '2023-06-21 08:00:00', NULL, '2024-06-21 01:28:20.021606', 'john', '2024-06-21 01:28:20.021606', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (8, 8, 2, '2023-06-21 08:00:00', NULL, '2024-06-21 01:29:24.096177', 'john', '2024-06-21 01:29:24.096177', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (9, 14, 2, '2024-06-21 01:29:35.240455', NULL, '2024-06-21 01:29:35.240456', 'john', '2024-06-21 01:29:35.240456', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (7, 14, 2, '2023-06-21 08:00:00', '2023-06-21 09:00:00', '2024-06-21 01:28:30.525179', 'john', '2024-06-21 01:31:06.1809', 'john', NULL);
INSERT INTO "public"."attendances" VALUES (1, 1, 2, NULL, NULL, '2024-06-21 01:04:20.556165', 'john', '2024-06-21 01:04:20.556165', 'john', '2024-06-21 01:31:39.284921');
INSERT INTO "public"."attendances" VALUES (10, 20, 2, '2024-06-21 02:12:40.384946', NULL, '2024-06-21 02:12:40.384947', 'databetul', '2024-06-21 02:12:40.384947', 'databetul', NULL);
INSERT INTO "public"."attendances" VALUES (11, 20, 1, '2024-06-21 02:15:38.447616', '2024-06-21 09:15:38.447616', '2024-06-21 02:15:38.447616', 'databetul', '2024-06-21 02:16:31.25243', 'databetul', NULL);
INSERT INTO "public"."attendances" VALUES (12, 23, 1, '2024-06-21 02:20:13.688722', '2024-06-21 09:15:38.447616', '2024-06-21 02:20:13.688723', 'atha', '2024-06-21 02:20:40.815429', 'atha', NULL);

-- ----------------------------
-- Primary Key structure for table attendances
-- ----------------------------
ALTER TABLE "public"."attendances" ADD CONSTRAINT "attendances_pkey" PRIMARY KEY ("attendance_id");

-- ----------------------------
-- Foreign Keys structure for table attendances
-- ----------------------------
ALTER TABLE "public"."attendances" ADD CONSTRAINT "attendances_employee_id_fkey" FOREIGN KEY ("employee_id") REFERENCES "public"."employees" ("employee_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."attendances" ADD CONSTRAINT "attendances_location_id_fkey" FOREIGN KEY ("location_id") REFERENCES "public"."locations" ("location_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
