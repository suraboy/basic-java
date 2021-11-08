package postgres

import (
	"github.com/suraboy/go-fiber-api/infrastructure"
	"gorm.io/gorm"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */

type Postgres struct {
	Connection *gorm.DB
}

func NewPostgres() *Postgres {
	conn := infrastructure.PostgresConnection()
	return &Postgres{
		Connection: conn,
	}
}
