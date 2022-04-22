package db

import (
	"fmt"
	"github.com/nightlord189/so5hw/internal/config"
	"github.com/nightlord189/so5hw/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io/ioutil"
)

var ErrRecordNotFound = gorm.ErrRecordNotFound

type Manager struct {
	Config      *config.Config
	DB          *gorm.DB
	truncateSQL string
	fillDataSQL string
}

func InitDb(cfg *config.Config) (*Manager, error) {
	dbInstance, err := connectDb(cfg)
	if err != nil {
		return nil, err
	}
	dbManager := Manager{
		DB:     dbInstance,
		Config: cfg,
	}
	err = dbManager.loadFixtures()
	return &dbManager, err
}

func (d *Manager) loadFixtures() error {
	content, err := ioutil.ReadFile("configs/truncate.sql")
	if err != nil {
		return err
	}
	d.truncateSQL = string(content)
	content, err = ioutil.ReadFile("configs/fixture.sql")
	if err != nil {
		return err
	}
	d.fillDataSQL = string(content)
	return nil
}

func connectDb(cfg *config.Config) (*gorm.DB, error) {
	address := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Name, cfg.DB.Password)
	dbLogger := logger.Discard
	if cfg.DB.Log {
		dbLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.Open(address), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("connect err: %w", err)
	}
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)
	if cfg.DB.Migrate {
		err = autoMigrate(db)
		if err != nil {
			return nil, fmt.Errorf("migrate err: %w", err)
		}
	}
	return db, err
}

type sqlMigration struct {
	SQL        string
	CheckError bool
}

var sqlMigrations = []sqlMigration{
	{
		SQL:        "ALTER TABLE public.image ADD CONSTRAINT image_fk FOREIGN KEY (product_id) REFERENCES public.product(id) ON DELETE CASCADE ON UPDATE CASCADE;",
		CheckError: false,
	},
}

//AutoMigrate - применить миграции
func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		model.MerchandiserDB{},
		model.CustomerDB{},
		model.ImageDB{},
		model.ProductDB{},
	)
	if err != nil {
		return err
	}
	return applyMigrations(db, sqlMigrations)
}

func applyMigrations(db *gorm.DB, migrations []sqlMigration) error {
	for _, migration := range sqlMigrations {
		err := db.Exec(migration.SQL).Error
		if err != nil && migration.CheckError {
			return err
		}
	}
	return nil
}
