package management

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title           string    `gorm:"type:varchar(255);not null"`
	PublicationDate time.Time `gorm:"not null"`
	File            []byte    `gorm:"not null"` // Almacena PDF como un arreglo de bytes
	AuthorID        uint      `gorm:"not null"` // Clave foránea para la relación con Author
	Author          Author    `gorm:"foreignKey:AuthorID"`

	Barcode   string    `gorm:"type:varchar(100);not null"`
	PageCount int       `gorm:"not null"`
	Format    string    `gorm:"type:varchar(50);not null"`
	DueDate   time.Time `gorm:"type:date;not null"`
}

func GetBooks(db *gorm.DB) ([]Book, error) {
	var books []Book
	result := db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

// Método que permite la adición de un nuevo libro
func CreateBook(db *gorm.DB, book *Book) error {
	result := db.Create(&book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Métodos setter
func (b *Book) SetPublicationDate(date string) {
	b.PublicationDate, _ = time.Parse("2006-01-02", date)
}

func (b *Book) SetTitle(title string) {
	b.Title = title
}

func (b *Book) SetFile(file []byte) {
	b.File = file
}

func (b *Book) SetBarcode(barcode string) {
	b.Barcode = barcode
}

func (b *Book) SetPageCount(pageCount int) {
	b.PageCount = pageCount
}

func (b *Book) SetFormat(format string) {
	b.Format = format
}

func (b *Book) SetDueDate(date string) {
	b.DueDate, _ = time.Parse("2006-01-02", date)
}

// Métodos getter
func (b *Book) GetPublicationDate() string {
	return b.PublicationDate.Format("2006-01-02")
}

func (b *Book) GetTitle() string {
	return b.Title
}

func (b *Book) GetFile() []byte {
	return b.File
}

func (b *Book) GetBarcode() string {
	return b.Barcode
}

func (b *Book) GetPageCount() int {
	return b.PageCount
}

func (b *Book) GetFormat() string {
	return b.Format
}

func (b *Book) GetDueDate() string {
	return b.DueDate.Format("2006-01-02")
}
