package dependencies

import (
	"database/sql"

	"github.com/nats-io/nats.go"
	"github.com/sirupsen/logrus"
	"github.com/teewat888/user-service/internal/config"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type Dependencies struct {
	Config *config.Config
	DB     *bun.DB
	Nats   *nats.Conn
}

func InitDependencies(cfg *config.Config) *Dependencies {
	if cfg.Env == "prod" {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	db := initDBConnection(cfg)
	nats := initNatsConnection(cfg)

	return &Dependencies{
		Config: cfg,
		DB:     db,
		Nats:   nats,
	}
}

func initDBConnection(cfg *config.Config) *bun.DB {
	sql := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(cfg.DbUrl)))
	db := bun.NewDB(sql, pgdialect.New())

	r := db.QueryRow("SELECT 1")
	if r.Err() != nil {
		panic("Cannot connect to the DB: " + r.Err().Error())
	}

	return db
}

func initNatsConnection(cfg *config.Config) *nats.Conn {
	nc, err := nats.Connect(cfg.NatsUrl)
	if err != nil {
		panic("Cannot connect to the NATS: " + err.Error())
	}

	return nc
}
