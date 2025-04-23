package config

import "os"

type DBConfig struct {
	IsCloudSQL bool
	ProjectID  string
	Region     string
	Instance   string
	User       string
	Password   string
	DBName     string
	Host       string
	Port       string
}

func NewDBConfig() *DBConfig {
	useCloudSQL := os.Getenv("USE_CLOUD_SQL") == "true"

	if useCloudSQL {
		return &DBConfig{
			IsCloudSQL: true,
			ProjectID:  os.Getenv("GCP_PROJECT_ID"),
			Region:     os.Getenv("DB_REGION"),
			Instance:   os.Getenv("DB_INSTANCE"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
		}
	} else {
		return &DBConfig{
			IsCloudSQL: false,
			Host:       os.Getenv("DB_HOST"),
			Port:       os.Getenv("DB_PORT"),
			User:       os.Getenv("DB_USER"),
			Password:   os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
		}
	}
}

func IsMigrateUp() bool {
	return os.Getenv("MIGRATE_UP") == "true"
}
func IsMigrateDown() bool {
	return os.Getenv("MIGRATE_DOWN") == "true"
}
