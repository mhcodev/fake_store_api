package driver

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectToPostgresDB connects to a postgresSQL database
func ConnectToPostgresDB() *pgxpool.Pool {
	// Initialize PostgreSQL connection
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbMode := os.Getenv("DB_SSL_MODE")

	fmt.Println("dbUser: ", dbUser)
	fmt.Println("dbPass: ", dbPass)
	fmt.Println("dbHost: ", dbHost)
	fmt.Println("dbPort: ", dbPort)
	fmt.Println("dbName: ", dbName)
	fmt.Println("dbMode: ", dbMode)

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		dbMode,
	)

	config, err := pgxpool.ParseConfig(dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse connection string: %v\n", err)
		os.Exit(1)
	}

	config.MaxConns = 30

	dbpool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
	}

	var dbError error

	for i := range 5 { // Retry 5 times
		err = dbpool.Ping(context.Background())
		if err != nil {
			dbError = err
			fmt.Println("Triying to connect to DB: ", i)
		} else {
			dbError = nil
			break
		}
		time.Sleep(time.Second * 2)
	}

	if dbError != nil {
		fmt.Fprintf(os.Stderr, "Is not ready to make queries, check and run again: %v", err)
		os.Exit(1)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeCacheDescribe

	var dbConnected string
	_, err = dbpool.Exec(context.Background(), "DEALLOCATE ALL")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec query failed: %v\n", err)
		os.Exit(1)
	}

	err = dbpool.QueryRow(context.Background(), "select 'Database connected!' as text").Scan(&dbConnected)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(dbConnected)
	return dbpool
}
