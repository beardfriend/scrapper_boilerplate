package infra

import (
	"boilerplate/ent"
	"boilerplate/ent/migrate"
	"context"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	client *ent.Client
}

func NewSqlite(path string) *Sqlite {
	s := new(Sqlite)
	s.init(path)

	return s
}

func (s *Sqlite) init(path string) {

	client, err := ent.Open("sqlite3", fmt.Sprintf("file:%s?_fk=1", path))
	if err != nil {
		log.Fatalf("connection failed: %v", err)
	}

	s.client = client
}

func (s *Sqlite) OnDebug() *Sqlite {
	s.client = s.client.Debug()

	return s
}

func (s *Sqlite) AutoMigrate() error {
	return s.client.Schema.Create(context.Background(), migrate.WithDropColumn(true), migrate.WithForeignKeys(false), migrate.WithDropIndex(true))
}

func (s *Sqlite) GetClient() *ent.Client {
	return s.client
}

func (s *Sqlite) Close() error {
	return s.client.Close()
}
