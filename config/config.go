package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Postgres PostgresConfig
	Server   ServerConfig
	Mongo    MongoDBConfig
	Redis    RedisConfig
}

type PostgresConfig struct {
	PDB_NAME     string
	PDB_PORT     string
	PDB_PASSWORD string
	PDB_USER     string
	PDB_HOST     string
}

type RedisConfig struct {
	RDB_ADDRESS  string
	RDB_PASSWORD string
}

type ServerConfig struct {
	USER_SERVICE string
	USER_ROUTER  string
}

type MongoDBConfig struct {
	MDB_ADDRESS string
	MDB_NAME    string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error while loading .env file: %v", err)
	}

	return &Config{
		Postgres: PostgresConfig{
			PDB_HOST:     cast.ToString(coalesce("PDB_HOST", "localhost")),
			PDB_PORT:     cast.ToString(coalesce("PDB_PORT", "5432")),
			PDB_USER:     cast.ToString(coalesce("PDB_USER", "postgres")),
			PDB_NAME:     cast.ToString(coalesce("PDB_NAME", "postgres")),
			PDB_PASSWORD: cast.ToString(coalesce("PDB_PASSWORD", "3333")),
		},
		Server: ServerConfig{
			USER_SERVICE: cast.ToString(coalesce("USER_SERVICE", ":1234")),
			USER_ROUTER:  cast.ToString(coalesce("USER_ROUTER", ":1234")),
		},
		Redis: RedisConfig{
			RDB_ADDRESS:  cast.ToString(coalesce("RDB_ADDRESS", "localhost:6379")),
			RDB_PASSWORD: cast.ToString(coalesce("RDB_PASSWORD", "")),
		},
		Mongo: MongoDBConfig{
			MDB_ADDRESS: cast.ToString(coalesce("MDB_ADDRESS", "mongodb://localhost:27017")),
			MDB_NAME:    cast.ToString(coalesce("MDB_NAME", "test")),
		},
	}
}

func coalesce(key string, value interface{}) interface{} {
	val, exist := os.LookupEnv(key)
	if exist {
		return val
	}
	return value
}
