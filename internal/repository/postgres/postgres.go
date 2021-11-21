package postgres

import (
	"gorm.io/gorm"
)

/**
 * Created by boy.sirichai on 8/11/2021 AD
 */

type Postgres struct {
	Connection *gorm.DB
}

func NewPostgres(conn *gorm.DB) *Postgres {
	return &Postgres{
		Connection: conn,
	}
}
