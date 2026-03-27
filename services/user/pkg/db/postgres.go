package db

import (
	"database/sql"
	"fmt"
	"learning_golang/user/internal/repository/postgres/queries"
	"logger"
)

type PostgresConfig struct {
	Host     string `env:"PG_HOST" default:"localhost"`
	Port     int    `env:"PG_PORT" default:"5432"`
	User     string `env:"PG_USER" default:"postgres"`
	Password string `env:"PG_PASSWORD" default:"postgres"`
	DBName   string `env:"PG_DB_NAME" default:"user_service"`
	SSLMode  string `env:"PG_SSLMODE" default:"disable"`
}

type PostgresDB struct {
	*sql.DB
	config      *PostgresConfig
	logger      *logger.Logger
	userQueries *queries.UserQueries
}

func NewPostgresDB(cfg *PostgresConfig, log *logger.Logger) (*PostgresDB, error) {
	connectionString := DBConnectionString(cfg)
	userQueries := queries.GetUserQueries()
	initQueries := queries.GetInitQueries()

	log.Info("Connecting to PostgreSQL")

	db, err := sql.Open("postgres", connectionString)
	defer db.Close()

	if err != nil {
		log.Fatal("Error connecting to PostgreSQL: %v", err)
		return nil, err
	}

	log.Info("Successfully connected to PostgreSQL")

	if _, err := db.Exec(initQueries.InitUserTable); err != nil {
		log.Error("Error creating table: %v", err)
	}

	return &PostgresDB{
		DB:          db,
		logger:      log,
		config:      cfg,
		userQueries: userQueries,
	}, nil
}

func DBConnectionString(cfg *PostgresConfig) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.DBName,
	)
}
