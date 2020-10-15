package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	// This is implicitly needed by gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type config struct {
	Vars varConfig `json:"vars"`
	DB   dbConfig  `json:"db"`
}

type varConfig struct {
	Port    string `json:"port"`
	Env     string `json:"env"`
}

type dbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func (s *server) loadConfig() config {
	f, err := os.Open(".config.json")
	if err != nil {
		s.panic("config file missing", err)
	}

	var c config
	err = json.NewDecoder(f).Decode(&c)
	if err != nil {
		s.panic("json decoder failed", err)
	}

	return c
}

func (s *server) isProd() bool {
	return s.vars.Env == "prod"
}

func (s *server) setupDB(c dbConfig) {
	const dialect = "postgres"
	connectionStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Host, c.Port, c.User, c.Password, c.DBName,
	)

	db, err := gorm.Open(dialect, connectionStr)
	if err != nil {
		s.logErr("DB connection failed", err)
	}

	db.LogMode(!s.isProd())

	s.db = db
}

func (s *server) autoMigrate() {
	err := s.db.AutoMigrate().Error
	if err != nil {
		s.logErr("Automigration problem", err)
	}
}

func (s *server) destructiveReset() {
	err := s.db.DropTableIfExists().Error
	if err != nil {
		s.logErr("DestructiveReset problem", err)
	}
	s.autoMigrate()
}
