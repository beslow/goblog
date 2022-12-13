package migrations

func createCategories() {
	timestamp := "202212101941"

	createTableWithTimestamp(
		"categories",
		`CREATE TABLE categories (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(255) NOT NULL DEFAULT '',
			created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		timestamp,
	)
}
