package learn_sql

// example: NewMySQL("user:pass@tcp(host:3306)/db")
func NewMySQL(dsn string) *sql.DB, error {
	return sql.Open("mysql", dsn)
}

// example: NewSQLite("/tmp/db.sqlite")
func NewSQLite(dsn string)  *sql.DB, error {
	return sql.Open("sqlite3", dns)
}