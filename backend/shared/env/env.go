package env

import (
	"github.com/joho/godotenv"
	"os"
)

const EnvFilePath = "../../../.env"

func Load() {
	os.Getenv(EnvFilePath)
	_ = godotenv.Load("../../../.env")
}
