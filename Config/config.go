package Config

import (
	"fmt"
	"os"
	_ "github.com/go-sql-driver/mysql" // import the MySQL driver, which registers itself with `database/sql`
	"database/sql"
)

// DB is a variable of type `*sql.DB` that will store the database connection
var DB *sql.DB

// DBConfig represents database configuration
type DBConfig struct {
	Host     string // MySQL host address
	Port     int    // MySQL port number
	User     string // MySQL user name
	DBName   string // MySQL database name
	Password string // MySQL user password
}

// BuildDBConfig returns a pointer to a `DBConfig` struct with configuration details retrieved from environment variables
func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),     // get MySQL host address from environment variable
		Port:     3306,                        // default MySQL port number
		DBName:   os.Getenv("DB_NAME"),        // get MySQL database name from environment variable
		User:     os.Getenv("MYSQL_USER"),     // get MySQL user name from environment variable
		Password: os.Getenv("MYSQL_PASSWORD"), // get MySQL user password from environment variable
	}
	return &dbConfig // return a pointer to the `DBConfig` struct
}

// DBUri returns a string that represents the database connection URI, based on the provided `DBConfig` struct
func DBUri(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s", // MySQL connection string format
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
