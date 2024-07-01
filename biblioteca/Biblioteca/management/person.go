package management

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	CINIC       string    `gorm:"type:varchar(100); null"`
	FullName    string    `gorm:"type:varchar(100);not null"`
	Phone       string    `gorm:"type:varchar(20);null"`
	Birthdate   time.Time `gorm:"type:date;null"`
	Ethnicity   string    `gorm:"type:varchar(50);null"`
	Nationality string    `gorm:"type:varchar(50);null"`
	CivilStatus string    `gorm:"type:varchar(50);null"`
}

// Setter methods
func (p *Person) SetCINIC(ciNic string) {
	p.CINIC = ciNic
}

func (p *Person) SetFullName(fullName string) {
	p.FullName = fullName
}

func (p *Person) SetPhone(phone string) {
	p.Phone = phone
}

func (p *Person) SetBirthdate(birthdate time.Time) {
	p.Birthdate = birthdate
}

func (p *Person) SetEthnicity(ethnicity string) {
	p.Ethnicity = ethnicity
}

func (p *Person) SetNationality(nationality string) {
	p.Nationality = nationality
}

func (p *Person) SetCivilStatus(civilStatus string) {
	p.CivilStatus = civilStatus
}

// Getter methods
func (p *Person) GetCINIC() string {
	return p.CINIC
}

func (p *Person) GetFullName() string {
	return p.FullName
}

func (p *Person) GetPhone() string {
	return p.Phone
}

func (p *Person) GetBirthdate() time.Time {
	return p.Birthdate
}

func (p *Person) GetEthnicity() string {
	return p.Ethnicity
}

func (p *Person) GetNationality() string {
	return p.Nationality
}

func (p *Person) GetCivilStatus() string {
	return p.CivilStatus
}
