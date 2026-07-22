package migrations

import (
	db_pool "app/db"
	models "app/src/models"
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	// Clean, standard registration
	goose.AddMigrationContext(UpInitialSchema, DownInitialSchema)
}

func UpInitialSchema(ctx context.Context, tx *sql.Tx) error {
	gormDb, err := db_pool.GetTxDB(tx)
	if err != nil {
		return err
	}

	return gormDb.WithContext(ctx).AutoMigrate(
		&models.User{},
		&models.Followers{},
		&models.OfficeDetails{},
		&models.ResidentialDetails{},
	)
}

func DownInitialSchema(ctx context.Context, tx *sql.Tx) error {
	gormDB, err := db_pool.GetTxDB(tx)
	if err != nil {
		return err
	}

	return gormDB.WithContext(ctx).Migrator().DropTable(
		&models.ResidentialDetails{},
		&models.OfficeDetails{},
		&models.Followers{},
		&models.User{},
	)
}
