package data

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Artifact_safety struct {
	ArtifactID string `gorm:"column:artifact_id" json:"artifact_id"`
	Safety     string `gorm:"column:safety" json:"safety"`
}

type Artifact_master_phas struct {
	ID             int64            `gorm:"primary_key" json:"id"`
	Creator        string           `gorm:"column:creator" json:"creator"`
	Style          string           `gorm:"column:artifact_style_name" json:"style"`
	ExcavationDate string           `gorm:"column:date_exc" json:"date_exc"`
	TransferredBy  string           `gorm:"column:transferred_by" json:"transferred_by"`
	Width          string           `gorm:"column:width" json:"width"`
	Height         string           `gorm:"column:height" json:"height"`
	Length         string           `gorm:"column:length" json:"length"`
	Safety         *Artifact_safety
	Materials      []*Material
	Elements       []*ArtifactElement
}

type ArtifactElement struct {
	Name string `gorm:"column:artifact_element_name" json:"artifact_element_name"`
}

type Material struct{}

type CardData struct {
	db *gorm.DB
}

func NewCardData(db *gorm.DB) *CardData {
	return &CardData{db}
}

func (cd *CardData) ReadAll() ([]*Artifact_master_phas, error) {
	cards := make([]*Artifact_master_phas, 0)
	if err := cd.db.Find(&cards).Error; err != nil {
		return []*Artifact_master_phas{}, err
	}
	for i := 0; i < len(cards); i++ {
		safety := new(Artifact_safety)
		if err := cd.db.Where("artifact_id = " + string(cards[i].ID)).Error; err != nil {
			log.Println("safety error", err)
		}
		log.Println(safety)
		cards[i].Safety = safety
	}
	return cards, nil
}
