package migrations

import (
	context "context"
	sql "database/sql"

	goose "github.com/pressly/goose/v3"

	db_pool "app/db"
	models "app/src/models"
)

func init() {
	goose.AddMigrationContext(UpInitialSchema, DownInitialSchema)
}

func UpInitialSchema(ctx context.Context, tx *sql.Tx) error {
	gormDb, err := db_pool.GetTxDB(tx)
	if err != nil {
		return err
	}

	return gormDb.WithContext(ctx).AutoMigrate(
		&models.User{},
		&models.Follower{},
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
		&models.Follower{},
		&models.User{},
	)
}
