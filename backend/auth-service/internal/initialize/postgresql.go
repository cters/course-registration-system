package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	setting "gitlab.com/dacn9315724/course-registration-ddd/backend/auth-service/pkg"
	"go.uber.org/zap"
)

// InitDB khởi tạo kết nối đến PostgreSQL và trả về SQL DB instance
func InitDB(cfg setting.Config, logger *zap.Logger) (*sql.DB, error) {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName, cfg.DBSSL)

	var err error
	DB, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Cấu hình connection pool
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(100)
	DB.SetConnMaxLifetime(time.Hour)

	logger.Info("Database connection established successfully.")
	return DB, nil
}