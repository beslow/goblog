package migrations

func createComments() {
	timestamp := "202212101614"

	createTableWithTimestamp(
		"comments",
		`CREATE TABLE comments (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			post_id int(11) DEFAULT NULL,
			name varchar(255) DEFAULT '',
			email varchar(255) DEFAULT NULL,
			body text NOT NULL,
			created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
		    ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		timestamp,
	)
}
