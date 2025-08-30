DROP TABLE IF EXISTS "goadmin_menu";

CREATE TABLE IF NOT EXISTS "goadmin_menu" (
                                              `id` integer PRIMARY KEY autoincrement,
                                              `parent_id` INT NOT NULL DEFAULT '0',
                                              `order` INT NOT NULL DEFAULT '0',
                                              `type` INT NOT NULL DEFAULT '0',
                                              `title` CHAR(50) COLLATE NOCASE NOT NULL,
    `icon` CHAR(50) COLLATE NOCASE NOT NULL,
    `uri` CHAR(50) COLLATE NOCASE DEFAULT NULL,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP,
    header CHAR(150) DEFAULT NULL, plugin_name CHAR(150) NOT NULL DEFAULT '',
    uuid CHAR(100) COLLATE NOCASE DEFAULT NULL
    );

/*!------------------------------------------------------------- */

DROP TABLE IF EXISTS "goadmin_operation_log";

CREATE TABLE IF NOT EXISTS "goadmin_operation_log" (
                                                       `id` integer PRIMARY KEY autoincrement,
                                                       `user_id` INT NOT NULL,
                                                       `path` CHAR(255) COLLATE NOCASE NOT NULL,
    `method` CHAR(10) COLLATE NOCASE NOT NULL,
    `ip` CHAR(15) COLLATE NOCASE NOT NULL,
    `input` text COLLATE NOCASE NOT NULL,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
    );


/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_permissions";

CREATE TABLE IF NOT EXISTS "goadmin_permissions" (
                                                     `id` integer PRIMARY KEY autoincrement,
                                                     `name` CHAR(50) COLLATE NOCASE NOT NULL,
    `slug` CHAR(50) COLLATE NOCASE NOT NULL,
    `http_method` CHAR(255) COLLATE NOCASE DEFAULT NULL,
    `http_path` text COLLATE NOCASE,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
    );





/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_role_menu";

CREATE TABLE IF NOT EXISTS "goadmin_role_menu" (
                                                   `role_id` INT NOT NULL,
                                                   `menu_id` INT NOT NULL,
                                                   `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
                                                   `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
);


/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_role_permissions";

CREATE TABLE IF NOT EXISTS "goadmin_role_permissions" (
                                                          `role_id` INT NOT NULL,
                                                          `permission_id` INT NOT NULL,
                                                          `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
                                                          `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
);

/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_role_users";

CREATE TABLE IF NOT EXISTS "goadmin_role_users" (
                                                    `role_id` INT NOT NULL,
                                                    `user_id` INT NOT NULL,
                                                    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
                                                    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
);

/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_roles";

CREATE TABLE IF NOT EXISTS "goadmin_roles" (
                                               `id` integer PRIMARY KEY autoincrement,
                                               `name` CHAR(50) COLLATE NOCASE NOT NULL,
    `slug` CHAR(50) COLLATE NOCASE NOT NULL,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
    );


/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_session";

CREATE TABLE IF NOT EXISTS "goadmin_session" (
                                                 `id` integer PRIMARY KEY autoincrement,
                                                 `sid` CHAR(50) DEFAULT NULL,
    `values` CHAR(3000) DEFAULT NULL,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
    );

/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_site";

CREATE TABLE IF NOT EXISTS "goadmin_site" (
                                              `id` integer PRIMARY KEY autoincrement,
                                              `key` CHAR(100) COLLATE NOCASE NOT NULL,
    `value` text COLLATE NOCASE NOT NULL,
    `state` INT NOT NULL DEFAULT '0',
    `description` CHAR(3000) COLLATE NOCASE,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
    );

/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_user_permissions";

CREATE TABLE IF NOT EXISTS "goadmin_user_permissions" (
                                                          `user_id` INT NOT NULL,
                                                          `permission_id` INT NOT NULL,
                                                          `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
                                                          `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
);


/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "profile";

CREATE TABLE "profile" (
                         `id` INTEGER PRIMARY KEY AUTOINCREMENT,
                         `uuid` TEXT,
                         `photos` TEXT,
                         `resume` TEXT,
                         `resume_size` INTEGER NOT NULL DEFAULT 0,
                         `finish_state` INTEGER NOT NULL DEFAULT 0,
                         `finish_progress` INTEGER NOT NULL DEFAULT 0,
                         `pass` INTEGER NOT NULL DEFAULT 0,
                         `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT INTO `profile` (`id`, `uuid`, `photos`, `resume`, `resume_size`, `finish_state`, `finish_progress`, `pass`, `created_at`, `updated_at`)
VALUES
    (1,'eeYtHtUtQg8U7zpCNiigVVhnToj','http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png,http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png,http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png','http://yinyanghu.github.io/files/clrs_prev.pdf',13242389,0,30,0,'2020-05-15 08:29:44','2020-05-15 08:29:44'),
    (2,'AxKvrvCaZpT3zsTsmrueFuLZFg9','http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png,http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png,http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png','http://yinyanghu.github.io/files/clrs_prev.pdf',232322233,1,60,1,'2020-05-15 08:30:51','2020-05-15 08:30:51'),
    (3,'QAwrQgEfqGs7qCUNpWGmoaEP3yF','http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png,http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png,http://quick.go-admin.cn/demo/assets/dist/img/gopher_avatar.png','http://yinyanghu.github.io/files/clrs_prev.pdf',232323,2,80,1,'2020-05-15 08:31:21','2020-05-15 08:31:21');

/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "users";

CREATE TABLE IF NOT EXISTS "users" (
                                     `id` INTEGER PRIMARY KEY AUTOINCREMENT,
                                     `name` TEXT,
                                     `gender` INTEGER,
                                     `city` TEXT,
                                     `ip` TEXT,
                                     `phone` TEXT,
                                     `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                     `updated_at` TIMESTAMP
);

/*!------------------------------------------------------------- */
DROP TABLE IF EXISTS "goadmin_users";

CREATE TABLE IF NOT EXISTS "goadmin_users" (
                                               `id` integer PRIMARY KEY autoincrement,
                                               `username` CHAR(190) COLLATE NOCASE NOT NULL,
    `password` CHAR(80) COLLATE NOCASE NOT NULL DEFAULT '',
    `name` CHAR(255) COLLATE NOCASE NOT NULL,
    `avatar` CHAR(255) COLLATE NOCASE DEFAULT NULL,
    `remember_token` CHAR(100) COLLATE NOCASE DEFAULT NULL,
    `created_at` TIMESTAMP default CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP default CURRENT_TIMESTAMP
    );
/*!------------------------------------------------------------- */
INSERT INTO `goadmin_menu` (`id`, `parent_id`, `type`, `order`, `title`, `icon`, `uri`, `plugin_name`, `header`, `created_at`, `updated_at`)
VALUES
    (1,0,1,1,'Admin','fa-tasks','','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,1,1,1,'Users','fa-users','/info/manager','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (3,1,1,2,'Roles','fa-user','/info/roles','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (4,1,1,3,'Permission','fa-ban','/info/permission','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (5,1,1,4,'Menu','fa-bars','/menu','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (6,1,1,5,'Operation log','fa-history','/info/op','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (7,0,1,2,'Dashboard','fa-bar-chart','/','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');

INSERT INTO `goadmin_permissions` (`id`, `name`, `slug`, `http_method`, `http_path`, `created_at`, `updated_at`)
VALUES
    (1,'All permission','*','','*','2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,'Dashboard','dashboard','GET,PUT,POST,DELETE','/','2019-09-10 00:00:00','2019-09-10 00:00:00');


INSERT INTO `goadmin_role_menu` (`role_id`, `menu_id`, `created_at`, `updated_at`)
VALUES
    (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (1,7,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,7,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (1,8,'2019-09-11 10:20:55','2019-09-11 10:20:55'),
    (2,8,'2019-09-11 10:20:55','2019-09-11 10:20:55');


INSERT INTO `goadmin_role_permissions` (`role_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
    (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (1,2,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,2,'2019-09-10 00:00:00','2019-09-10 00:00:00');


INSERT INTO `goadmin_role_users` (`role_id`, `user_id`, `created_at`, `updated_at`)
VALUES
    (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,2,'2019-09-10 00:00:00','2019-09-10 00:00:00');



INSERT INTO `goadmin_roles` (`id`, `name`, `slug`, `created_at`, `updated_at`)
VALUES
    (1,'Administrator','administrator','2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,'Operator','operator','2019-09-10 00:00:00','2019-09-10 00:00:00');


INSERT INTO `goadmin_user_permissions` (`user_id`, `permission_id`, `created_at`, `updated_at`)
VALUES
    (1,1,'2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,2,'2019-09-10 00:00:00','2019-09-10 00:00:00');


INSERT INTO `goadmin_users` (`id`, `username`, `password`, `name`, `avatar`, `remember_token`, `created_at`, `updated_at`)
VALUES
    (1,'admin','$2a$10$U3F/NSaf2kaVbyXTBp7ppOn0jZFyRqXRnYXB.AMioCjXl3Ciaj4oy','admin','','tlNcBVK9AvfYH7WEnwB1RKvocJu8FfRy4um3DJtwdHuJy0dwFsLOgAc0xUfh','2019-09-10 00:00:00','2019-09-10 00:00:00'),
    (2,'operator','$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.','Operator','',NULL,'2019-09-10 00:00:00','2019-09-10 00:00:00');


INSERT INTO `users` (`id`, `name`, `gender`, `city`, `ip`, `phone`, `created_at`, `updated_at`)
VALUES
    (3133,'voluptatum',0,'West Dorrischester','130.24.131.165','(291)462-9','2008-10-12 07:44:28','2003-10-16 23:40:41'),
    (3134,'quaerat',0,'Mantefurt','18.206.108.141','1-170-439-','2017-01-05 23:01:17','2006-10-09 16:31:23'),
    (3135,'quibusdam',0,'Altafurt','89.162.78.57','017-065-51','1988-11-08 14:53:14','2007-03-26 20:18:35'),
    (3136,'molestias',0,'East Jadontown','131.25.27.144','+92(7)3113','2014-01-23 15:56:15','1986-06-19 20:37:54'),
    (3137,'incidunt',0,'Angelville','90.255.113.150','1-881-209-','1989-02-17 23:59:30','1970-12-05 20:00:04'),
    (3138,'exercitationem',0,'Mrazport','112.152.108.62','(101)591-1','1995-04-18 15:32:08','1989-07-06 17:23:48'),
    (3139,'cumque',0,'South Carsonborough','56.70.126.83','687-792-49','2004-09-09 12:22:21','1994-05-17 16:53:50'),
    (3140,'ab',0,'New Abigaylemouth','180.66.161.219','121.009.26','1993-07-16 16:40:39','1985-04-27 19:02:24'),
    (3141,'numquam',0,'Port Polly','118.115.157.126','764.875.85','1998-11-04 17:36:16','2003-06-16 00:32:30'),
    (3142,'ratione',0,'East Madelynn','124.144.175.243','446.459.77','1980-10-31 12:09:14','2000-08-28 21:10:47'),
    (3143,'repellat',0,'Lake Aliza','69.66.247.238','1-514-720-','1981-07-11 13:57:15','1982-11-16 19:31:11'),
    (3144,'unde',0,'Claudechester','80.187.230.130','371-412-97','1973-01-22 17:32:51','1985-10-16 07:15:04'),
    (3145,'dolores',0,'East Candida','89.169.15.90','591.507.13','1991-05-05 21:02:27','1985-10-09 18:49:14'),
    (3146,'laudantium',0,'Harrisstad','51.29.100.162','668-521-48','1981-09-12 04:20:41','1994-05-09 03:32:30'),
    (3147,'iure',0,'Kingbury','99.13.130.67','(670)383-5','1996-10-03 14:10:37','1993-04-25 20:38:23'),
    (3148,'numquam',0,'Sanfordville','89.174.176.217','015-350-08','2010-07-15 20:25:56','1990-04-21 13:27:30'),
    (3149,'alias',0,'New Jacquelynmouth','176.202.145.52','670.430.97','2000-06-07 07:57:30','2015-06-06 08:57:47'),
    (3150,'expedita',0,'Lake Hilbert','96.21.195.51','(534)858-3','2012-11-07 10:02:02','2002-04-08 21:41:02'),
    (3151,'quis',0,'Lake Neal','89.152.227.200','+07(9)3192','1990-10-22 15:41:12','2013-06-22 09:51:23'),
    (3152,'id',0,'Port Laurence','45.24.206.89','270-153-13','2013-03-28 06:34:44','2012-12-25 08:49:40'),
    (3153,'ea',0,'Cummingsmouth','119.31.3.235','628-176-55','2008-12-25 21:07:18','1987-03-04 14:45:37'),
    (3154,'sapiente',0,'West Joaquin','203.137.34.242','034.848.48','2010-03-10 04:23:48','1974-02-27 01:52:51'),
    (3155,'blanditiis',0,'Port Logan','247.71.235.180','1-354-533-','2010-03-12 00:22:42','2007-08-22 08:52:34'),
    (3156,'laborum',0,'North Odie','184.185.248.33','(349)149-5','1993-12-23 09:54:44','1990-11-07 05:09:54'),
    (3157,'sit',0,'Port Brook','254.154.238.177','1-028-949-','1997-11-18 14:26:34','1992-07-22 16:48:00'),
    (3158,'assumenda',0,'East Mackenzie','204.158.130.66','(160)960-2','1981-10-11 07:31:53','1984-01-23 14:16:56'),
    (3159,'quisquam',0,'Bashirianburgh','153.180.29.168','371-305-81','2009-02-15 00:27:53','2005-06-29 12:23:04'),
    (3160,'sed',0,'New Donny','118.254.120.78','1-685-501-','1980-05-23 09:06:18','1974-09-02 22:40:54'),
    (3161,'ut',0,'South Krisville','5.98.220.210','639.996.18','1977-08-06 01:42:41','1994-10-30 13:59:51'),
    (3162,'cumque',0,'North Emmanuel','39.91.220.233','(168)144-2','1974-03-22 07:17:19','1994-08-04 06:33:55'),
    (3163,'dolorem',0,'New Noel','180.17.202.66','+06(8)8657','1978-02-06 09:30:14','2005-03-31 17:13:48'),
    (3164,'sunt',0,'Strosinchester','41.78.165.138','(749)823-5','1985-12-25 15:40:37','2003-01-25 09:21:52'),
    (3165,'rerum',0,'West Leatha','51.234.111.252','(455)407-9','2017-02-06 23:59:34','2004-02-10 12:17:41'),
    (3166,'aut',0,'West Nels','114.157.77.105','570-533-99','1983-09-23 12:47:53','2004-03-08 06:57:28'),
    (3167,'eos',0,'Janland','190.148.202.0','267-198-73','1974-08-12 15:23:43','1984-06-17 08:04:48'),
    (3168,'omnis',0,'Nitzscheburgh','94.253.229.243','056.325.16','2007-02-19 03:02:05','1977-11-24 10:53:20'),
    (3169,'quae',0,'Gustfurt','11.246.22.247','972-910-38','1985-02-01 23:31:23','1989-08-17 05:57:51'),
    (3170,'facere',0,'Heathcoteburgh','104.70.84.237','234.144.01','1978-05-28 17:06:56','2009-06-23 00:34:41'),
    (3171,'eligendi',0,'New Candice','185.172.87.32','1-990-977-','1994-01-24 17:04:22','1990-03-14 20:10:48'),
    (3172,'autem',0,'New Martymouth','218.85.247.31','0213109440','1992-08-23 07:34:29','1991-05-23 02:26:46'),
    (3173,'velit',0,'North Chadd','152.4.77.210','1-977-024-','2018-05-12 19:40:30','1996-06-24 00:43:49'),
    (3174,'voluptatem',0,'Emardland','12.36.43.98','1-504-175-','2018-08-26 09:42:55','1984-11-20 08:11:22'),
    (3175,'sunt',0,'West Myrna','212.40.77.247','0142702071','1987-01-26 19:14:35','1998-09-02 19:16:30'),
    (3176,'minus',0,'Gennaroton','220.16.78.177','661.478.49','1999-11-02 21:00:18','2018-07-26 06:23:23'),
    (3177,'eum',0,'Mervinmouth','39.68.16.110','514-868-17','2016-07-18 15:08:34','1988-04-11 18:14:53'),
    (3178,'et',0,'Port Natalieton','29.13.164.161','+20(9)5254','2000-04-22 21:43:23','1986-06-13 03:01:07'),
    (3179,'ipsam',0,'North Rettaview','65.174.146.78','+19(5)5639','1993-02-07 05:41:54','1991-04-26 01:48:31'),
    (3180,'adipisci',0,'Conroyfurt','49.211.166.251','376.870.27','2000-04-04 06:18:45','1999-08-12 18:49:47'),
    (3181,'qui',0,'East Terryburgh','201.237.10.151','(682)499-4','1974-12-27 16:39:14','1978-03-27 05:36:35'),
    (3182,'aliquam',0,'North Darion','123.110.173.222','(386)536-3','1989-07-26 23:33:27','2007-03-10 19:38:38'),
    (3183,'blanditiis',0,'Shanelshire','161.173.251.134','693.628.85','2011-04-19 13:27:17','2000-06-27 07:32:56'),
    (3184,'provident',0,'South Amie','157.222.23.146','(941)805-3','1983-05-21 22:14:39','1974-08-13 07:59:23'),
    (3185,'et',0,'New Ryleeville','60.133.158.102','238.375.46','1981-05-23 07:26:35','2018-05-16 09:31:50'),
    (3186,'rem',0,'Krajcikview','254.68.144.153','113.589.44','2012-02-21 21:23:56','2017-11-11 18:05:11'),
    (3187,'officia',0,'Port Cruz','5.23.112.130','481.895.97','2002-06-26 19:07:28','1970-12-27 13:47:14'),
    (3188,'sit',0,'New Brooklyntown','75.170.89.171','+21(7)5136','1995-07-01 21:01:08','2005-02-24 14:01:38'),
    (3189,'error',0,'Lutherfort','78.141.190.96','+25(3)1840','1984-01-20 15:11:23','2016-09-18 14:55:22'),
    (3190,'est',0,'Linneamouth','117.182.29.64','+32(6)9290','1998-10-13 05:51:01','1993-04-20 07:28:18'),
    (3191,'molestias',0,'Miashire','160.3.34.22','408.583.99','1978-11-12 14:23:51','1989-05-01 09:34:51'),
    (3192,'dolorem',0,'Nathaven','30.142.214.7','1-015-611-','1993-02-25 17:21:00','2006-08-03 07:06:04'),
    (3193,'voluptatem',0,'Clintonberg','155.218.87.29','+36(1)8939','1998-09-14 21:54:13','1977-11-23 07:50:55'),
    (3194,'cum',0,'Billiemouth','184.218.96.17','602-921-61','1988-08-13 12:19:16','2001-08-02 09:43:17'),
    (3195,'et',0,'West Trevor','134.173.224.149','1-654-522-','2017-02-08 08:50:11','1994-02-25 02:45:06'),
    (3196,'voluptate',0,'Port Muhammad','101.162.1.247','1-469-725-','2007-05-31 21:55:19','1983-02-02 07:16:01'),
    (3197,'ut',0,'North Tomasa','34.50.218.169','1-254-394-','1998-12-01 09:58:54','1975-02-07 01:59:00'),
    (3198,'sed',0,'Hillaryport','120.212.199.52','1-926-573-','2016-07-29 19:03:12','2008-09-10 13:04:44'),
    (3199,'magni',0,'Sebastianshire','177.94.144.118','627-932-50','1973-11-10 20:24:51','2017-10-08 23:29:45'),
    (3200,'aliquid',0,'South Ellis','224.211.29.87','(493)204-8','1975-03-25 03:58:01','1990-07-28 11:14:30'),
    (3201,'omnis',0,'Anabelstad','92.249.88.62','(978)197-9','2009-01-01 18:00:34','2014-10-16 11:54:54'),
    (3202,'numquam',0,'South Ronmouth','169.106.168.28','392-241-92','1978-12-18 04:09:47','1992-08-30 11:20:23'),
    (3203,'dolores',0,'Lake Benjamin','15.117.69.132','251.485.93','2018-08-28 16:39:31','2011-07-06 04:53:46'),
    (3204,'deserunt',0,'Emersonmouth','66.51.232.79','1-447-734-','2002-11-30 12:44:53','1975-05-16 18:37:48'),
    (3205,'ut',0,'Spinkamouth','244.43.102.248','1-000-780-','2014-04-07 13:41:12','1989-05-30 18:40:16'),
    (3206,'qui',0,'South Kristopherhaven','246.92.42.67','1-484-350-','2006-06-23 16:52:48','2003-12-02 16:52:52'),
    (3207,'culpa',0,'Port Kaitlin','137.7.62.226','(970)422-8','1979-11-03 20:39:44','2008-09-19 21:25:57'),
    (3208,'assumenda',0,'Brittanystad','112.235.56.1','1-140-282-','1992-03-29 10:34:11','1987-10-13 16:36:42'),
    (3209,'tempora',0,'West Madisenbury','252.216.155.148','0364798286','1977-07-19 04:38:56','1990-02-12 06:56:49'),
    (3210,'voluptas',0,'Lake Eugenestad','159.105.105.253','(687)489-4','1983-02-04 14:30:46','2015-11-11 17:30:31'),
    (3211,'animi',0,'Gradymouth','234.240.244.211','0732322239','2010-05-25 02:06:51','1994-01-31 02:15:34'),
    (3212,'sunt',0,'Port Catharine','170.237.211.55','262.806.58','1975-10-13 05:49:59','1973-09-14 21:04:22'),
    (3213,'aspernatur',0,'North Paolo','67.211.135.138','(620)698-7','1990-12-12 18:58:51','2018-06-06 04:54:08'),
    (3214,'blanditiis',0,'Adamouth','157.101.158.128','110-851-48','1997-03-01 04:50:56','1990-05-13 20:20:15'),
    (3215,'aut',0,'Maribelburgh','214.169.56.47','219-930-75','1993-06-16 00:21:23','1982-01-22 04:54:59'),
    (3216,'voluptatibus',0,'Lake Tiannaberg','123.67.227.100','397.747.22','2014-01-27 12:53:19','2006-06-26 14:45:55'),
    (3217,'architecto',0,'Beverlystad','104.35.135.41','1-856-949-','2004-11-10 07:30:30','2000-04-30 05:34:03'),
    (3218,'perspiciatis',0,'New Dejahfurt','248.150.222.9','754.322.73','1993-09-14 09:32:03','2002-01-09 16:50:17'),
    (3219,'repellat',0,'New Humberto','103.130.113.37','866.129.48','1985-05-01 06:16:30','1996-11-03 06:40:02'),
    (3220,'consequuntur',0,'Friesenmouth','119.247.58.87','955-141-06','1976-03-20 14:24:07','1998-11-24 08:35:53'),
    (3221,'voluptas',0,'Ornfort','120.132.216.182','(589)854-2','1988-04-18 11:20:40','1992-02-26 18:57:32'),
    (3222,'laudantium',0,'South Darrell','94.253.191.196','343-005-11','1980-05-23 23:38:00','2000-02-20 03:24:03'),
    (3223,'ut',0,'Kaiafort','212.3.22.162','1-552-220-','1970-05-19 08:31:38','1998-11-01 09:55:05'),
    (3224,'rerum',0,'Quitzonview','19.97.140.221','348.900.22','2003-12-24 17:54:42','1970-02-03 17:36:52'),
    (3225,'qui',0,'New Marcelle','48.29.124.108','1-524-833-','1976-09-27 18:40:44','1993-03-05 23:57:54'),
    (3226,'autem',0,'Lake Carolineborough','243.186.200.110','485-310-14','1998-06-23 22:09:28','1987-05-01 07:29:39'),
    (3227,'quas',0,'Marshallmouth','188.233.129.117','1-200-909-','2011-07-18 19:31:07','1993-12-17 23:25:30'),
    (3228,'ab',0,'Lake Staceyhaven','23.48.214.152','0788802278','2003-04-03 18:06:42','2007-12-19 05:02:45'),
    (3229,'ipsa',0,'Kreigerborough','40.12.123.51','+77(3)1104','1997-12-07 21:03:32','2017-05-22 09:25:19'),
    (3230,'sit',0,'New Idellatown','124.93.199.84','682.372.57','1985-09-21 20:35:22','1981-12-01 18:28:11'),
    (3231,'quis',0,'Doyleburgh','115.32.241.190','530-811-53','1985-08-03 02:36:05','2007-08-05 11:37:00'),
    (3232,'explicabo',0,'Carrollland','239.130.251.156','606-189-58','1975-05-11 20:23:26','2004-08-23 10:18:05'),
    (3233,'quas',0,'Dorotheamouth','65.227.95.202','1-120-101-','1980-08-28 22:15:44','1977-05-04 14:41:36'),
    (3234,'tenetur',0,'New Alia','199.166.74.233','1-572-838-','1998-05-01 13:15:57','1970-11-11 23:53:42'),
    (3235,'quia',0,'Lloydchester','218.151.232.131','373-827-78','1990-06-26 20:36:08','1974-01-23 10:11:30'),
    (3236,'at',0,'Port Lavina','142.20.99.100','1-143-840-','2013-08-16 10:26:08','1992-08-12 10:21:22'),
    (3237,'molestias',0,'Lake Abigail','94.251.47.14','1-447-916-','1989-01-06 00:47:25','2017-03-27 18:34:24'),
    (3238,'sed',0,'East Osvaldo','216.134.64.85','042-252-27','2005-05-18 01:08:48','1993-03-08 21:51:34'),
    (3239,'culpa',0,'West Isai','72.31.156.203','163.371.96','2012-03-03 13:45:16','1973-06-12 11:43:50');