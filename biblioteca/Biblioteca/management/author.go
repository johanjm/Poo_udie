package management

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	PersonID  uint   `gorm:"uniqueIndex;not null"` // Foreign key for Person
	Biography string `gorm:"type:varchar(100);null"`
	Person    Person `gorm:"foreignKey:PersonID"`
	Books     []Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // One-to-many relationship with Book
}

// Constructor that allows the addition of a new author
func CreateAuthor(db *gorm.DB, author *Author) error {
	result := db.Create(&author)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateAuthor(db *gorm.DB, id int, author *Author) error {
	// Buscar el autor por ID

	fmt.Println("show me ", id)
	var existingAuthor Author
	if err := db.Preload("Person").First(&existingAuthor, id).Error; err != nil {
		return err
	}

	// Actualizar los campos del autor
	existingAuthor.Biography = author.Biography // Actualizar el campo Biography si es necesario

	// Actualizar los campos de la persona del autor
	existingAuthor.Person.FullName = author.Person.FullName
	existingAuthor.Person.Birthdate = author.Person.Birthdate
	existingAuthor.Person.Nationality = author.Person.Nationality
	existingAuthor.Person.Phone = author.Person.Phone

	// Guardar los cambios en la base de datos
	if err := db.Save(&existingAuthor).Error; err != nil {
		return err
	}

	// if db.Statement.Changed("Name", "Admin") { // if Name or Role changed
	// 	tx.Statement.SetColumn("Age", 18)
	//   }

	return nil
}

func GetAuthorAndBooks(db *gorm.DB, id uint) (*Author, error) {
	var author Author
	result := db.Preload("Books").Preload("Person").First(&author, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &author, nil
}

func GetAuthors(db *gorm.DB) ([]Author, error) {
	var authors []Author
	result := db.Preload("Person").Preload("Books").Find(&authors)
	if result.Error != nil {
		return nil, result.Error
	}
	return authors, nil
}
func GetAuthorById(db *gorm.DB, id uint) (*Author, error) {
	var author Author
	result := db.Preload("Person").First(&author, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &author, nil
}

func DeleteAuthor(db *gorm.DB, id int) error {
	result := db.Delete(&Author{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Display the books belonging to each author
func (a *Author) ShowBooks(db *gorm.DB) error {
	var author Author
	if err := db.Preload("Books").Preload("Person").First(&author, a.ID).Error; err != nil {
		return err
	}

	fmt.Printf("Books by %s:\n", author.Person.FullName)
	for _, book := range author.Books {
		fmt.Printf("- %s (Publication Date: %s)\n", book.Title, book.PublicationDate.Format("2006-01-02"))
	}
	return nil
}

// Setter methods for Author
func (a *Author) SetFullName(fullName string) {
	a.Person.FullName = fullName
}

func (a *Author) SetPhone(phone string) {
	a.Person.Phone = phone
}

func (a *Author) SetBirthdate(birthdate time.Time) {
	a.Person.Birthdate = birthdate
}

func (a *Author) SetEthnicity(ethnicity string) {
	a.Person.Ethnicity = ethnicity
}

func (a *Author) SetNationality(nationality string) {
	a.Person.Nationality = nationality
}

func (a *Author) SetCivilStatus(civilStatus string) {
	a.Person.CivilStatus = civilStatus
}

// Getter methods for Author
func (a *Author) GetFullName() string {
	return a.Person.FullName
}

func (a *Author) GetPhone() string {
	return a.Person.Phone
}

func (a *Author) GetBirthdate() time.Time {
	return a.Person.Birthdate
}

func (a *Author) GetEthnicity() string {
	return a.Person.Ethnicity
}

func (a *Author) GetNationality() string {
	return a.Person.Nationality
}

func (a *Author) GetCivilStatus() string {
	return a.Person.CivilStatus
}
