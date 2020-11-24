package models

import (
	"errors"

	uuid "github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // postgres
)

//Book ...
type Book struct {
	UUID        uuid.UUID `gorm:"primary_key" json:"uuid"`
	Title       string    `gorm:"size:255;not null;" json:"title"`
	Author      string    `gorm:"size:255;not null;" json:"author"`
	ImageURL    string    `gorm:"size:255;not null;" json:"image_url"`
	Description string    `gorm:"size:255;not null;" json:"description"`
}

//BeforeCreate ...
func (b *Book) BeforeCreate() {
	newuuid := uuid.New()
	b.UUID = newuuid
}

//SaveBook ...
func (b *Book) SaveBook(db *gorm.DB) (*Book, error) {
	var err error
	err = db.Debug().Create(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, err
}

//FindAllBooks ...
func (b *Book) FindAllBooks(db *gorm.DB) (*[]Book, error) {
	var err error
	books := []Book{}
	err = db.Debug().Model(&Book{}).Limit(100).Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, err
}

//FindBookByID ...
func (b *Book) FindBookByID(db *gorm.DB, id uuid.UUID) (*Book, error) {
	var err error
	err = db.Debug().Model(Book{}).Where("uuid = ?", id).Take(&b).Error
	if err != nil {
		return &Book{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Book{}, errors.New("Book Not Found")
	}
	return b, err
}

//UpdateBook ...
func (b *Book) UpdateBook(db *gorm.DB, id uuid.UUID) (*Book, error) {
	var err error
	db = db.Debug().Model(&Book{}).Where("uuid = ?", id).Take(&Book{}).Update(b)
	if db.Error != nil {
		return &Book{}, db.Error
	}
	err = db.Debug().Model(&Book{}).Where("uuid = ?", id).Take(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

//DeleteABook ...
func (b *Book) DeleteABook(db *gorm.DB, id uuid.UUID) (int64, error) {

	db = db.Debug().Model(&Book{}).Where("uuid = ?", id).Take(&Book{}).Delete(&Book{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
