package testsetup

import (
	"context"
	"os"

	"gorm.io/gorm"
)

var TEST_CONNECTION_STRING = os.Getenv("TEST_DB_CONNECTION_STRING")

const DB_SCHEME = "dronetask"

type DB struct {
	client *gorm.DB
}

func (db *DB) Init(client *gorm.DB) {
	db.client = client
}

func (db *DB) Get(ctx context.Context, obj interface{}, result interface{}) {
	db.client.WithContext(ctx).Where(obj).Last(&result)
}
