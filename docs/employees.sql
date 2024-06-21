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

 Date: 21/06/2024 09:38:15
*/


-- ----------------------------
-- Table structure for employees
-- ----------------------------
DROP TABLE IF EXISTS "public"."employees";
CREATE TABLE "public"."employees" (
  "employee_id" int4 NOT NULL DEFAULT nextval('employees_employee_id_seq'::regclass),
  "employee_code" varchar(10) COLLATE "pg_catalog"."default" NOT NULL DEFAULT (to_char((CURRENT_DATE)::timestamp with time zone, 'DDMM'::text) || lpad((nextval('employee_code_seq'::regclass))::text, 3, '0'::text)),
  "employee_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "password" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "department_id" int4,
  "position_id" int4,
  "superior" int4,
  "created_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "created_by" varchar(255) COLLATE "pg_catalog"."default",
  "updated_at" timestamp(6) DEFAULT CURRENT_TIMESTAMP,
  "updated_by" varchar(255) COLLATE "pg_catalog"."default",
  "deleted_at" timestamp(6)
)
;

-- ----------------------------
-- Records of employees
-- ----------------------------
INSERT INTO "public"."employees" VALUES (7, '2006007', 'john', '$2a$10$qXGTT/HtOMX77REZq/2HKOoM3nZMIj5BQTRd8AZZ7s7.hpM2foIIO', NULL, NULL, NULL, '2024-06-20 21:45:05.645107', 'system', '2024-06-20 21:45:05.646443', NULL, NULL);
INSERT INTO "public"."employees" VALUES (9, '2006009', 'john', '$2a$10$KTDRjMl8N1G4RWeu2SBrQeKHi.R38jd1IRTQBx.sPZXxSGkOIPO2m', NULL, NULL, NULL, '2024-06-20 21:48:42.33278', 'system', '2024-06-20 21:48:42.334889', NULL, NULL);
INSERT INTO "public"."employees" VALUES (10, '2006010', 'john', '$2a$10$ymTZ.z/h4gmW/C6g7ACCT.o9kPsryn60A0q0PZrVgD3P6W107tNYi', NULL, NULL, NULL, '2024-06-20 21:51:20.092479', 'system', '2024-06-20 21:51:20.092479', 'system', NULL);
INSERT INTO "public"."employees" VALUES (11, '2006011', 'john', '$2a$10$5sV/c9I3MZxyprX5jvWPSuCAwpTgAZPCab4OMCL.Eg4EWucLK/Ulq', NULL, NULL, NULL, '2024-06-20 22:09:48.632438', 'system', '2024-06-20 22:09:48.632438', 'system', NULL);
INSERT INTO "public"."employees" VALUES (12, '2006012', 'john', '$2a$10$KMcADEocfZPOS27qqK15p.5SsRTY7nc8zd3Lap5DGSTAgRORn4VHC', NULL, NULL, NULL, '2024-06-20 22:12:48.851086', 'system', '2024-06-20 22:12:48.851086', 'system', NULL);
INSERT INTO "public"."employees" VALUES (13, '2006013', 'john', '$2a$10$nIvmFMiDecUE9huOKRHTq.Az3xAA6IAAfDVtjYfKSOvd5vE7r7CJu', NULL, NULL, NULL, '2024-06-20 22:13:50.268775', 'system', '2024-06-20 22:13:50.268775', 'system', NULL);
INSERT INTO "public"."employees" VALUES (14, '2006014', 'john', '$2a$10$RB3XDdGTmq896m1foaaQCePbLp7z9mNXcsJ/kzT/0OcHsAnnOVRk.', NULL, NULL, NULL, '2024-06-20 22:19:37.806472', 'system', '2024-06-20 22:19:37.806472', 'system', NULL);
INSERT INTO "public"."employees" VALUES (15, '2006015', 'dummy', '$2a$10$el8ts1NZTXSOINw.o9GxNeDltedTFIEY0thp8BFh1BfRKg3y.84nm', NULL, NULL, NULL, '2024-06-20 23:00:46.229875', 'john', '2024-06-20 23:00:46.229875', 'john', NULL);
INSERT INTO "public"."employees" VALUES (16, '2006016', 'dummy', '$2a$10$7Hk81lbox4nih/zkXDQBRetjNn9ePeobM0Of2xNLvffP9DbfDxBQq', NULL, NULL, NULL, '2024-06-20 23:05:00.027052', 'john', '2024-06-20 23:05:00.027052', 'john', NULL);
INSERT INTO "public"."employees" VALUES (17, '2006017', 'dummy', '$2a$10$I/Il3G4tDGo/n6Up3Ic5pOV6CA51X0.rJQkm5H8k0rqhxBH1.ets.', NULL, NULL, NULL, '2024-06-20 23:06:52.763444', 'john', '2024-06-20 23:06:52.763444', 'john', NULL);
INSERT INTO "public"."employees" VALUES (18, '2006018', 'john', '$2a$10$ohU/T5TG0apgd9b/rI5oDewtDZwHXtNpVE1rmLnOd2p7B6/awqat.', NULL, NULL, NULL, '2024-06-20 23:19:46.085694', 'system', '2024-06-20 23:19:46.085694', 'system', NULL);
INSERT INTO "public"."employees" VALUES (1, '2106001', 'John Doe', 'encrypted_password', 1, 1, 1, '2024-06-20 10:00:00', 'admin', '2024-06-20 23:31:57.091458', 'john', '2024-06-20 23:34:35.803571');
INSERT INTO "public"."employees" VALUES (19, '2006019', 'john', '$2a$10$8QLhhlf2aIr6cYrvf/n19.xitXCXDR1luNkkxhHHvbOsPL3dQxD7i', NULL, NULL, NULL, '2024-06-20 23:48:35.051447', 'system', '2024-06-20 23:48:35.051447', 'system', NULL);
INSERT INTO "public"."employees" VALUES (8, '2006008', 'john', '$2a$10$HVTCgQI5s6Vag44N/ABPPesLXFB1gm5hZtsgrTrPTchE20WQ4sCbO', 1, 3, 8, '2024-06-20 21:47:20.875502', 'system', '2024-06-21 02:11:45.443675', 'databetul', NULL);
INSERT INTO "public"."employees" VALUES (20, '2106020', 'databetul', '$2a$10$BVCw0epIvjW4oOgoUS363e9EkzGR6ZF0/3YrtY0rwxep7BrcRSCJi', 1, 3, 8, '2024-06-21 02:07:51.8802', 'john', '2024-06-21 02:12:05.935494', 'databetul', NULL);
INSERT INTO "public"."employees" VALUES (21, '2106021', 'databetull', '$2a$10$O9C90/jkCJnHDQ.HxcS8IefjQFzTyvuvocfj9tbId6helNtZdYvpW', NULL, NULL, NULL, '2024-06-21 02:13:28.79158', 'databetul', '2024-06-21 02:13:28.79158', 'databetul', NULL);
INSERT INTO "public"."employees" VALUES (22, '2106022', 'databetull', '$2a$10$u8p5epQeIsuoKzZjCODv3urO0YpNnxoCra1fqvTOjsoErGkQeesnu', 1, 1, 8, '2024-06-21 02:14:44.061718', 'databetul', '2024-06-21 02:17:11.362001', 'databetul', NULL);
INSERT INTO "public"."employees" VALUES (23, '2106023', 'atha', '$2a$10$Ant/NgTWdjveLY6.vHQwoOw4HSxlHkrpvFIaKLjXTi6YLgv0NfC5O', 2, 4, 8, '2024-06-21 02:18:36.385492', 'databetul', '2024-06-21 02:18:36.385492', 'databetul', NULL);

-- ----------------------------
-- Uniques structure for table employees
-- ----------------------------
ALTER TABLE "public"."employees" ADD CONSTRAINT "employees_employee_code_key" UNIQUE ("employee_code");

-- ----------------------------
-- Primary Key structure for table employees
-- ----------------------------
ALTER TABLE "public"."employees" ADD CONSTRAINT "employees_pkey" PRIMARY KEY ("employee_id");

-- ----------------------------
-- Foreign Keys structure for table employees
-- ----------------------------
ALTER TABLE "public"."employees" ADD CONSTRAINT "employees_department_id_fkey" FOREIGN KEY ("department_id") REFERENCES "public"."departments" ("department_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
ALTER TABLE "public"."employees" ADD CONSTRAINT "employees_position_id_fkey" FOREIGN KEY ("position_id") REFERENCES "public"."positions" ("position_id") ON DELETE NO ACTION ON UPDATE NO ACTION;
