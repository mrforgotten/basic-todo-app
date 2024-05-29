package config

import (
	"context"
	"os"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

func GetDbConfig() *pg.Options {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	opts := &pg.Options{
		Addr:     os.Getenv("PG_HOST") + ":" + os.Getenv("PG_PORT"),
		User:     os.Getenv("PG_USERNAME"),
		Password: os.Getenv("PG_PASSWORD"),
		Database: os.Getenv("PG_DATABASE"),
		OnConnect: func(ctx context.Context, cn *pg.Conn) error {
			_, err := cn.Exec("SET search_path=public")
			if err != nil {
				return err
			}
			return nil
		},
	}

	return opts
}
