package management

import (
	"time"

	"gorm.io/gorm"
)

type BookItem struct {
	gorm.Model
	Barcode   string    `gorm:"type:varchar(100);not null"`
	PageCount int       `gorm:"not null"`
	Format    string    `gorm:"type:varchar(50);not null"`
	DueDate   time.Time `gorm:"type:date;not null"`
	BookID    uint      `gorm:"not null"`
	Book      Book      `gorm:"foreignkey:BookID"`
}

// Setter methods
func (b *BookItem) SetBarcode(barcode string) {
	b.Barcode = barcode
}

func (b *BookItem) SetPageCount(pageCount int) {
	b.PageCount = pageCount
}

func (b *BookItem) SetFormat(format string) {
	b.Format = format
}

func (b *BookItem) SetDueDate(dueDate time.Time) {
	b.DueDate = dueDate
}

// Getter methods
func (b *BookItem) GetBarcode() string {
	return b.Barcode
}

func (b *BookItem) GetPageCount() int {
	return b.PageCount
}

func (b *BookItem) GetFormat() string {
	return b.Format
}

func (b *BookItem) GetDueDate() time.Time {
	return b.DueDate
}
