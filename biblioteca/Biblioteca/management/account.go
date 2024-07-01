package management

import (
	"time"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Number   string    `gorm:"type:varchar(100);not null"`
	Open     time.Time `gorm:"type:date;not null"`
	State    string    `gorm:"type:varchar(50);not null"`
	Email    string    `gorm:"type:varchar(100);unique"`
	Password string    `gorm:"type:varchar(250);not null"`
	PersonID uint      // Esta será la clave foránea que conecta Account con Person
	Person   Person    `gorm:"foreignKey:PersonID"` // Relación uno a uno con Person
}

// Función para crear un nuevo usuario en la base de datos
func CreateAccount(db *gorm.DB, account *Account) error {
	result := db.Create(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Setter methods
func (c *Account) SetNumber(number string) {
	c.Number = number
}

func (c *Account) SetOpen(open time.Time) {
	c.Open = open
}

func (c *Account) SetState(state string) {
	c.State = state
}

// Getter methods
func (c *Account) GetNumber() string {
	return c.Number
}

func (c *Account) GetOpen() time.Time {
	return c.Open
}

func (c *Account) GetState() string {
	return c.State
}
