package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	MigrationsPath  string
	MigrationsTable string
	UserDB          string
	PassDB          string
	HostDB          string
	PortDB          int
	DBName          string
}

const (
	migrationsPathName  = "MIGRATIONS_PATH"
	migrationsTableName = "MIGRATIONS_TABLE"
	userDBName          = "USER_DB"
	passDBName          = "PASS_DB"
	hostDBName          = "HOST_DB"
	portDBName          = "PORT_DB"
	dbName              = "DB_NAME"
)

func MustLoad() Config {
	emptyName := ""
	defer emptyNameErr(emptyName)
	migrationsPath := os.Getenv(migrationsPathName)
	if migrationsPath == "" {
		emptyName = migrationsPathName
		return Config{}
	}
	migrationsTable := os.Getenv(migrationsTableName)
	if migrationsTable == "" {
		emptyName = migrationsTableName
		return Config{}
	}
	userDB := os.Getenv(userDBName)
	if userDB == "" {
		emptyName = userDBName
		return Config{}
	}
	passDB := os.Getenv(passDBName)
	if passDB == "" {
		emptyName = passDBName
		return Config{}
	}
	hostDB := os.Getenv(hostDBName)
	if hostDB == "" {
		emptyName = hostDBName
		return Config{}
	}
	db := os.Getenv(dbName)
	if db == "" {
		emptyName = dbName
		return Config{}
	}
	portStr := os.Getenv(portDBName)
	portDB, err := strconv.Atoi(portStr)
	if err != nil {
		emptyName = portDBName
		return Config{}
	}

	return Config{
		MigrationsPath:  migrationsPath,
		MigrationsTable: migrationsTable,
		UserDB:          userDB,
		PassDB:          passDB,
		HostDB:          hostDB,
		PortDB:          portDB,
		DBName:          db,
	}
}

func emptyNameErr(emptyName string) {
	if emptyName != "" {
		log.Fatalf("%s is not set", emptyName)
	}
}
