package repository

type PgConfig struct {
	PgHost string
	PgPort string
	PgUser string
	PgPass string
	PgBase string
}

var cfg PgConfig

func init() {
	cfg.PgHost = "localhost"
	cfg.PgPort = "5432"
	cfg.PgUser = "postgres"
	cfg.PgPass = "postgres"
	cfg.PgBase = "UrlDB"
}
