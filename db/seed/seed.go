package seed

import (
	"fmt"
	"os"

	"github.com/beslow/goblog/db"
)

var seedData = map[string]string{
	"consts": `INSERT INTO consts (id, name, description, value) VALUES
		(1, 'real_name', '', 'your name'),
		(2, 'email', '', '549174542@qq.com'),
		(3, 'profession', '', '高级服务器端工程师');`,

	"goadmin_menu": `INSERT INTO goadmin_menu (id, parent_id, ` + "`type`" + `, ` + "`order`" + `, title, icon, uri, header, plugin_name, uuid, created_at, updated_at) VALUES
		(1, 0, 1, 2, 'Admin', 'fa-tasks', '', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 1, 1, 2, 'Users', 'fa-users', '/info/manager', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(3, 1, 1, 3, 'Roles', 'fa-user', '/info/roles', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(4, 1, 1, 4, 'Permission', 'fa-ban', '/info/permission', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(5, 1, 1, 5, 'Menu', 'fa-bars', '/menu', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(6, 1, 1, 6, 'Operation log', 'fa-history', '/info/op', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(7, 0, 1, 1, 'Dashboard', 'fa-bar-chart', '/', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(8, 0, 0, 2, '资料', 'fa-angellist', '', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(9, 8, 0, 2, '履历', 'fa-fighter-jet', '/info/history_works', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(10, 8, 0, 2, '教育学习经历', 'fa-empire', '/info/history_educations', NULL, '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(11, 0, 0, 2, '博客', 'fa-adn', '/info/posts', '', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(12, 0, 0, 2, '目录', 'fa-align-justify', '/info/categories', '', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(13, 8, 0, 2, '基本设置', 'fa-certificate', '/info/consts', '', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,

	"goadmin_site": `INSERT INTO goadmin_site (id, ` + "`key`" + `, value, description, state, created_at, updated_at) VALUES
		(1, 'theme', 'adminlte', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(2, 'logger_rotate_compress', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(3, 'logger_encoder_stacktrace_key', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(4, 'session_life_time', '7200', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(5, 'bootstrap_file_path', '/Users/beslow/workspace/go-admin/bootstrap.go', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(6, 'logger_encoder_level', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(7, 'hide_visitor_user_center_entrance', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(8, 'logger_encoder_caller_key', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(9, 'site_off', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(10, 'hide_app_info_entrance', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(11, 'logger_rotate_max_age', '0', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(12, 'logger_encoder_encoding', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(13, 'animation_duration', '0.00', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(14, 'logger_encoder_level_key', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(15, 'logger_level', '0', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(16, 'color_scheme', 'skin-black', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(17, 'hide_plugin_entrance', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(18, 'allow_del_operation_log', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(19, 'app_id', 'G7XkxvITUilL', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(20, 'title', 'GoAdmin', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(21, 'error_log_path', '/Users/beslow/workspace/go-admin/logs/error.log', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(22, 'sql_log', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(23, 'logger_encoder_message_key', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(24, 'animation_delay', '0.00', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(25, 'debug', 'true', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(26, 'file_upload_engine', '{\"name\":\"local\"}', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(27, 'login_title', 'GoAdmin', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(28, 'login_logo', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(29, 'hide_tool_entrance', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(30, 'animation_type', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(31, 'operation_log_off', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(32, 'logo', 'GoAdmin', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(33, 'logger_rotate_max_backups', '0', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(34, 'logger_encoder_caller', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(35, 'custom_head_html', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(36, 'custom_foot_html', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(37, 'extra', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(38, 'domain', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(39, 'login_url', '/login', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(40, 'info_log_off', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(41, 'no_limit_login_ip', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(42, 'go_mod_file_path', '/Users/beslow/workspace/go-admin/go.mod', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(43, 'custom_403_html', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(44, 'language', 'cn', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(45, 'logger_encoder_time_key', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(46, 'logger_encoder_duration', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(47, 'prohibit_config_modification', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(48, 'exclude_theme_components', 'null', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(49, 'asset_root_path', './public/', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(50, 'info_log_path', '/Users/beslow/workspace/go-admin/logs/info.log', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(51, 'hide_config_center_entrance', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(52, 'custom_404_html', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(53, 'open_admin_api', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(54, 'access_log_path', '/Users/beslow/workspace/go-admin/logs/access.log', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(55, 'error_log_off', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(56, 'logger_encoder_time', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(57, 'footer_info', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(58, 'url_prefix', 'admin', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(59, 'access_log_off', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(60, 'logger_encoder_name_key', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(61, 'auth_user_table', 'goadmin_users', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(62, 'custom_500_html', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(63, 'mini_logo', 'GA', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(64, 'access_assets_log_off', 'false', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(65, 'logger_rotate_max_size', '0', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(66, 'asset_url', '', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(67, 'index_url', '/', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59'),
		(68, 'env', 'local', NULL, 1, '2022-12-03 12:49:59', '2022-12-03 12:49:59');`,

	"goadmin_users": `INSERT INTO goadmin_users (id, username, ` + "`password`" + `, ` + "`name`" + `, avatar, remember_token, created_at, updated_at) VALUES
		(1, 'admin', '$2a$10$LOzvKm8MNfNA1Go97pcu7eSrzARMz8qbw/Dygb10TeSHJGgAeHHXO', 'admin', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 'operator', '$2a$10$rVqkOzHjN2MdlEprRflb1eGP0oZXuSrbJLOmJagFsCd81YZm0bsh.', 'Operator', '', NULL, '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,

	"goadmin_roles": `INSERT INTO goadmin_roles (id, name, slug, created_at, updated_at) VALUES
		(1, 'Administrator', 'administrator', '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 'Operator', 'operator', '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,

	"goadmin_role_users": `INSERT INTO goadmin_role_users (role_id, user_id, created_at, updated_at) VALUES
		(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,

	"goadmin_permissions": `INSERT INTO goadmin_permissions (id, name, slug, http_method, http_path, created_at, updated_at) VALUES
		(1, 'All permission', '*', '', '*', '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 'Dashboard', 'dashboard', 'GET,PUT,POST,DELETE', '/', '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,

	"goadmin_role_permissions": `INSERT INTO goadmin_role_permissions (role_id, permission_id, created_at, updated_at) VALUES
		(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(1, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,

	"goadmin_user_permissions": `INSERT INTO goadmin_user_permissions (user_id, permission_id, created_at, updated_at) VALUES
		(1, 1, '2019-09-10 00:00:00', '2019-09-10 00:00:00'),
		(2, 2, '2019-09-10 00:00:00', '2019-09-10 00:00:00');`,
}

func Seed() {
	for table, sql := range seedData {
		seedTable(table, sql)
	}
}

func seedTable(table, insertSql string) {
	var count int64
	err := db.DB.Table(table).Count(&count).Error

	if err != nil {
		fmt.Printf("check %s fail: %#v\n", table, err)
		os.Exit(1)
	}

	if count == 0 {
		if err := db.DB.Exec(insertSql).Error; err != nil {
			fmt.Printf("seed %s fail: %#v\n", table, err)
			os.Exit(1)
		}

		fmt.Printf("seed %s finish.\n", table)
	}
}
