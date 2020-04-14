package data

import (
	"log"

	"github.com/jinzhu/gorm"
)

// ArtifactSafety discribes current quality of artifact
type ArtifactSafety struct {
	ID         string
	ArtifactID string `gorm:"column:artifact_id" json:"artifact_id"`
	Safety     string `gorm:"column:safety" json:"safety"`
}

// ArtifactMasterPhas the main structure of artifact
type ArtifactMasterPhas struct {
	ID             int64  `gorm:"primary_key" json:"id"`
	Creator        string `gorm:"column:creator" json:"creator"`
	Style          string `gorm:"column:artifact_style_name" json:"style"`
	ExcavationDate string `gorm:"column:date_exc" json:"date_exc"`
	TransferredBy  string `gorm:"column:transferred_by" json:"transferred_by"`
	Width          string `gorm:"column:width" json:"width"`
	Height         string `gorm:"column:height" json:"height"`
	Length         string `gorm:"column:length" json:"length"`
	Safety         string
	Materials      []*Material
	Elements       []*ArtifactElement
}

// ArtifactElement some part of artifact
type ArtifactElement struct {
	Name string `gorm:"column:artifact_element_name" json:"artifact_element_name"`
}

// Material describes material with additional information for spicific artifact
type Material struct{}

// CardData gets connection to database
type CardData struct {
	db *gorm.DB
}

// NewCardData creates new instance
func NewCardData(db *gorm.DB) *CardData {
	return &CardData{db}
}

// ReadAll return all cards from database
func (cd *CardData) ReadAll() ([]*ArtifactMasterPhas, error) {
	cd.db.SingularTable(true) // gives apportunity to use table with singular name
	cards := make([]*ArtifactMasterPhas, 0)
	if err := cd.db.Find(&cards).Error; err != nil {
		return []*ArtifactMasterPhas{}, err
	}
	for _, card := range cards {
		safety := new(ArtifactSafety)
		if err := cd.db.Where("artifact_id = ?", card.ID).First(&safety).Error; err != nil {
			log.Println("safety error:", err)
		}
		card.Safety = safety.Safety
	}
	return cards, nil
}
