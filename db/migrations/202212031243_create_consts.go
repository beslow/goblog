package migrations

func createConsts() {
	timestamp := "202212031243"

	createTableWithTimestamp(
		"consts",
		`CREATE TABLE consts (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			name varchar(100) NOT NULL DEFAULT '' COMMENT '常量名称',
			description varchar(500) DEFAULT NULL COMMENT '常量描述',
			value varchar(1000) DEFAULT NULL COMMENT '常量值',
			PRIMARY KEY (id)
	  		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;`,
		timestamp,
	)
}
