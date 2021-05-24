package SexModel
import "github.com/plankiton/SexPistol/Database"

type Role struct {
    SexDB.Model
    Name    string  `json:"name,omitempty" gorm:"unique"`
    Desc    string  `json:"desc,omitempty" gorm:"unique"`
}

func (m Role) TableName() string {
    return "roles"
}
