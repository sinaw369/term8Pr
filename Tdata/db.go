package Tdata

import (
	"database/sql"
	"github.com/GoAdminGroup/go-admin/modules/logger"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB() error {
	// Path to the SQLite database file
	dbPath := "./Tdata/database.db"

	// Create/Open SQLite database
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		//	log.Fatalf("error opening database: %v", err)
		logger.Fatalf("error opening database: %v", err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// Initialize the migration driver
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		//log.Fatalf("error creating migration driver: %v", err)
		logger.Fatalf("error creating migration driver: %v", err)
	}

	// Create migration instance with the driver
	m, err := migrate.NewWithDatabaseInstance("file://Tdata/migrations", "sqlite", driver)
	if err != nil {
		//log.Fatalf("error creating migration instance: %v", err)
		logger.Fatalf("error creating migration instance: %v", err)
	}

	// Force the migration to version 0
	if err = m.Force(0); err != nil {
		//log.Fatalf("error forcing migration version: %v", err)
		logger.Fatalf("error forcing migration version: %v", err)
	}
	// Apply all migrations
	//if err := m.Drop(); err != nil && err != migrate.ErrNoChange {
	//	log.Fatalf("error applying migrations: %v", err)
	//}
	// Apply all migrations
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		//log.Fatalf("error applying migrations: %v", err)
		logger.Fatalf("error applying migrations: %v", err)
	}

	//fmt.Println("Migrations applied successfully!")
	logger.Info(" =====> Migrations applied successfully!")

	// Additional code to interact with the migrated database can be added here

	// Query to retrieve table names
	//rows, err := db.Query("SELECT name FROM sqlite_master WHERE type='table'")
	//if err != nil {
	//	//log.Fatalf("error querying tables: %v", err)
	//	logger.Fatalf("error querying tables: %v", err)
	//}
	//defer func(rows *sql.Rows) {
	//	err = rows.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(rows)
	//
	//// Iterate over the rows and print table names
	////fmt.Println("Tables:")
	//for rows.Next() {
	//	var tableName string
	//	if err := rows.Scan(&tableName); err != nil {
	//		log.Fatalf("error scanning table name: %v", err)
	//	}
	//	fmt.Println(tableName)
	//}
	//if err := rows.Err(); err != nil {
	//	log.Fatalf("error iterating over rows: %v", err)
	//}
	return nil
}
