package migrations

func createHistoryWorks() {
	timestamp := "202212042037"

	createTableWithTimestamp(
		"history_works",
		`CREATE TABLE history_works (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			company varchar(255) NOT NULL DEFAULT '',
			job_title varchar(255) NOT NULL DEFAULT '',
			description varchar(2000) NOT NULL DEFAULT '',
			from_to varchar(50) NOT NULL DEFAULT '',
			sort int(11) NOT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		timestamp,
	)
}
