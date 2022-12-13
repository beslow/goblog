package migrations

func createPosts() {
	timestamp := "202212091727"

	createTableWithTimestamp(
		"posts",
		`CREATE TABLE posts (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			title varchar(255) NOT NULL,
			body text NOT NULL,
			visit_count int(11) NOT NULL DEFAULT '0',
			updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		timestamp,
	)
}
