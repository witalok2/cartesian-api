package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/cartesian-api/utils/json"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"go.elastic.co/apm/module/apmsql"
	_ "go.elastic.co/apm/module/apmsql/mysql"
	_ "go.elastic.co/apm/module/apmsql/pq"
	_ "go.elastic.co/apm/module/apmsql/sqlite3"
)

type dbConfig struct {
	FilePath     string
	PathToDBFile string `json:"path_to_db_file"`
	Type         string `json:"type"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	DataBase     string `json:"name"`
	SSLMode      string `json:"ssl_mode"`

	MaxLifeTime       int `json:"max_life_time"`
	MaxOpenConnection int `json:"max_open_connection"`
	MaxIdleConnection int `json:"max_idle_connection"`

	Context *context.Context
}

var connections = map[string]*sqlx.DB{}
var connectionsCtx = map[string]*sql.DB{}

var Vconfig = dbConfig{}

func get(filePath string) (*sqlx.DB, error) {
	if db, found := connections[filePath]; found {
		return db, nil
	}
	return nil, fmt.Errorf(`error database not found. File path: %s`, filePath)
}

func getCtx(filePath string) (*sql.DB, error) {
	if db, found := connectionsCtx[filePath]; found {
		return db, nil
	}
	return nil, fmt.Errorf(`error database not found. File path: %s`, filePath)
}

func connect(config dbConfig) (*sqlx.DB, error) {
	var (
		db  *sqlx.DB
		err error
	)

	switch config.Type {
	case "postgres":
		db, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s port=%d password=%s host=%s dbname=%s sslmode=%s",
			config.User,
			config.Port,
			config.Password,
			config.Host,
			config.DataBase,
			config.SSLMode,
		))

	case "mysql":
		db, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DataBase,
		))

	case "sqlite3":
		db, err = sqlx.Connect("sqlite3", config.PathToDBFile)

	default:
		return nil, errors.New("error database type is not supported")
	}

	if err != nil {
		return nil, fmt.Errorf(`error connecting to database of type "%s" because of: %s`, config.Type, err.Error())
	}

	maxLifeTime := time.Duration(config.MaxLifeTime)

	db.SetMaxOpenConns(config.MaxOpenConnection)
	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetConnMaxLifetime(maxLifeTime * time.Minute)

	connections[config.FilePath] = db
	return db, nil
}

func connectCtx(config dbConfig) (*sql.DB, error) {
	var (
		db  *sql.DB
		err error
	)

	switch config.Type {
	case "mysql":
		if config.Context == nil {
			return nil, err
		}

		db, err = apmsql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&multiStatements=true",
			config.User,
			config.Password,
			config.Host,
			config.Port,
			config.DataBase,
		))
		if err != nil {
			return nil, err
		}
		if err = db.PingContext(*config.Context); err != nil {
			db.Close()
			return nil, err
		}
	default:
		return nil, errors.New("error database type is not supported")
	}

	if err != nil {
		return nil, fmt.Errorf(`error connecting to database of type "%s" because of: %s`, config.Type, err.Error())
	}

	maxLifeTime := time.Duration(config.MaxLifeTime)

	db.SetMaxOpenConns(config.MaxOpenConnection)
	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetConnMaxLifetime(maxLifeTime * time.Minute)

	connectionsCtx[config.FilePath] = db
	return db, nil
}

// GetByFile Create a database connection through
// the path of a file
func GetByFile(filePath string) (*sqlx.DB, error) {

	if db, err := get(filePath); err == nil {
		return db, nil
	}

	var (
		config dbConfig
		err    error
	)

	if err = json.UnmarshalFile(filePath, &config); err != nil {
		return nil, err
	}

	config.FilePath = filePath
	Vconfig = config

	return connect(config)
}

// GetByFile Create a database connection through
// the path of a file
func GetByFileCtx(ctx *context.Context, filePath string) (*sql.DB, error) {

	if db, err := getCtx(filePath); err == nil {
		return db, nil
	}

	var (
		config dbConfig
		err    error
	)

	if err = json.UnmarshalFile(filePath, &config); err != nil {
		return nil, err
	}

	config.FilePath = filePath
	config.Context = ctx
	Vconfig = config

	return connectCtx(config)

}

// MustGetByFile Create a database connection through
// the path of a file and generates a panic in case of error
func MustGetByFile(filePath string) *sqlx.DB {
	db, err := GetByFile(filePath)
	if err != nil {
		panic(err)
	}
	return db
}

// MustGetByFile Create a database connection through
// the path of a file and generates a panic in case of error
func MustGetByFileCtx(ctx context.Context, filePath string) *sql.DB {
	db, err := GetByFileCtx(&ctx, filePath)
	if err != nil {
		panic(err)
	}
	return db
}