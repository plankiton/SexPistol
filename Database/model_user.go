package SexDatabase

import (
    "time"
)

type User struct {
    Model

    Email      string    `json:"email,omitempty" gorm:"unique,default:null"`
    Name       string    `json:"name,omitempty" gorm:"index"`
    Born       time.Time `json:"born_date,omitempty" gorm:"index"`
    Genre      string    `json:"genre,omitempty" gorm:"default:'male'"`
    PassHash   string    `json:"-"`
}

func (m User) TableName() string {
    return "users"
}

func (model *User) CheckPass(s string) bool {
    byteHash := []byte(model.PassHash)
    err := CheckPass(byteHash, s)
    if err != nil {
        return false
    }
    return true
}

func (model *User) SetPass(s string) (string, error) {
    hash, err := ToPassHash(s)
    if err != nil {
        return "", nil
    }

    model.PassHash = hash
    return model.PassHash, nil
}
