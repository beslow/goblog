package migrations

func createHistoryEducations() {
	timestamp := "202212052104"

	createTableWithTimestamp(
		"history_educations",
		`CREATE TABLE history_educations (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			course varchar(100) NOT NULL DEFAULT '',
			city varchar(30) NOT NULL DEFAULT '',
			description varchar(1000) NOT NULL DEFAULT '',
			from_to varchar(20) NOT NULL DEFAULT '',
			sort int(11) NOT NULL DEFAULT '0',
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			PRIMARY KEY (id)
	  		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		timestamp,
	)
}
