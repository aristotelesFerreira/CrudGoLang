package models

import (
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nickname string `gorm:"size255;not null;unique" json:"nickname"`
	Email    string `gorm:"size:100;not null; unique" json:"email"`
	Password string `gorm:"size:100;not null" json:"password"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)

	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	// u.CreatedAt = time.Now()
	// u.UpdatedAt = time.Now()
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	err := db.Debug().Create(&u).Error

	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	users := []User{}
	err := db.Debug().Model(&User{}).Limit(100).Find(&users).Error

	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	err := u.BeforeSave()

	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password": u.Password,
			"nickname": u.Nickname,
			"email":    u.Email,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}

	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error

	if err != nil {
		return &User{}, err
	}

	return u, nil
}

func (u *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
