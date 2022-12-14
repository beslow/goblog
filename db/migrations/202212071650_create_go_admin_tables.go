package migrations

import (
	"github.com/beslow/goblog/initialize"
	"github.com/beslow/goblog/models"
)

func createGoAdminTables() {
	timestamp := "202212071650"

	var migration models.Migration
	initialize.DB.Where("timestamp = ?", timestamp).Find(&migration)

	if migration.Timestamp == "" {
		createTable(
			"goadmin_menu",
			`CREATE TABLE goadmin_menu (
				id int(10) unsigned NOT NULL AUTO_INCREMENT,
				parent_id int(11) unsigned NOT NULL DEFAULT '0',
				type tinyint(4) unsigned NOT NULL DEFAULT '0',
				`+"`order`"+` int(11) unsigned NOT NULL DEFAULT '0',
				title varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
				icon varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
				uri varchar(3000) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
				header varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				plugin_name varchar(150) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
				uuid varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		)

		createTable(
			"goadmin_operation_log",
			`CREATE TABLE goadmin_operation_log (
				id int(10) unsigned NOT NULL AUTO_INCREMENT,
				user_id int(11) unsigned NOT NULL,
				path varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
				method varchar(10) COLLATE utf8mb4_unicode_ci NOT NULL,
				ip varchar(15) COLLATE utf8mb4_unicode_ci NOT NULL,
				input text COLLATE utf8mb4_unicode_ci NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id),
				KEY admin_operation_log_user_id_index (user_id)
				) ENGINE=InnoDB AUTO_INCREMENT=112 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_permissions",
			`CREATE TABLE goadmin_permissions (
				id int(10) unsigned NOT NULL AUTO_INCREMENT,
				name varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
				slug varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
				http_method varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				http_path text COLLATE utf8mb4_unicode_ci NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id),
				UNIQUE KEY admin_permissions_name_unique (name)
			    ) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_role_menu",
			`CREATE TABLE goadmin_role_menu (
				role_id int(11) unsigned NOT NULL,
				menu_id int(11) unsigned NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				KEY admin_role_menu_role_id_menu_id_index (role_id,menu_id)
			    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_role_permissions",
			`CREATE TABLE goadmin_role_permissions (
				role_id int(11) unsigned NOT NULL,
				permission_id int(11) unsigned NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				UNIQUE KEY admin_role_permissions (role_id,permission_id)
			    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_role_users",
			`CREATE TABLE goadmin_role_users (
				role_id int(11) unsigned NOT NULL,
				user_id int(11) unsigned NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				UNIQUE KEY admin_user_roles (role_id,user_id)
			    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_roles",
			`CREATE TABLE goadmin_roles (
				id int(10) unsigned NOT NULL AUTO_INCREMENT,
				name varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
				slug varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id),
				UNIQUE KEY admin_roles_name_unique (name)
			    ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_session",
			`CREATE TABLE goadmin_session (
				id int(11) unsigned NOT NULL AUTO_INCREMENT,
				sid varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
				`+"`values`"+` varchar(3000) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
		    	) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4;`,
		)

		createTable(
			"goadmin_site",
			`CREATE TABLE goadmin_site (
				id int(11) unsigned NOT NULL AUTO_INCREMENT,
				`+"`key`"+` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				value longtext COLLATE utf8mb4_unicode_ci,
				description varchar(3000) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				state tinyint(3) unsigned NOT NULL DEFAULT '0',
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id)
			    ) ENGINE=InnoDB AUTO_INCREMENT=69 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_user_permissions",
			`CREATE TABLE goadmin_user_permissions (
				user_id int(11) unsigned NOT NULL,
				permission_id int(11) unsigned NOT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				UNIQUE KEY admin_user_permissions (user_id,permission_id)
			    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		createTable(
			"goadmin_users",
			`CREATE TABLE goadmin_users (
				id int(10) unsigned NOT NULL AUTO_INCREMENT,
				username varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
				password varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
				name varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
				avatar varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				remember_token varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
				created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
				PRIMARY KEY (id),
				UNIQUE KEY admin_users_username_unique (username)
			    ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;`,
		)

		initialize.DB.Create(&models.Migration{Timestamp: timestamp})
	}
}
