package db

import (
	"gorm.io/gorm"
)

var repo Repo

type Db *gorm.DB

type Repo struct {
	db *gorm.DB
}

func InitializeRepo(db *gorm.DB) IRepo {
	repo = Repo{
		db: db,
	}
	return repo
}

func (d Repo) Create(model any) error {
	err := d.db.Create(model).Error
	if err != nil {
		return err
	}

	return nil
}

func (d Repo) Get(model any, filter map[string]interface{}) (any, error) {
	err := d.db.Where(filter).Find(&model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (d Repo) FindOne(model any, filter map[string]interface{}) error {
	err := d.db.Where(filter).First(&model).Error
	if err != nil {
		return err
	}

	return nil
}

func (d Repo) Update(model any, filter map[string]interface{}, id int) error {
	err := d.db.Model(&model).Where("id = ?", id).Updates(filter).Error
	if err != nil {
		return err
	}

	return nil
}
