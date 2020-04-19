package data

import (
	"log"

	"github.com/jinzhu/gorm"
)

type ArtifactMeasurement struct {
	Height int64
	Width  int64
	Length int64
}

// ArtifactMasterPhas the main structure of artifact
type ArtifactMasterPhas struct {
	ID                  int64                `json:"id"`
	Creator             string               `json:"creator"`
	ArtifactStyle       string               `json:"artifact_style"`
	ExcavationDate      string               `json:"date_exc"`
	TransferredBy       string               `json:"transferred_by"`
	Safety              string               `json:"safety"` // rewrite with graph
	ArtifactMeasurement *ArtifactMeasurement `json:"artifact_measurement"`
	Materials           []*Material
	Elements            []*ArtifactElement
}

// ArtifactElement some part of artifact
type ArtifactElement struct {
	Name string `gorm:"column:artifact_element_name" json:"artifact_element_name"`
}

// Material describes material with additional information for specific artifact
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
	cards := make([]*ArtifactMasterPhas, 0)
	// Way how to get data with relationship from db has been found, but it's not a ORM way
	// TODO: write working version without ORM way, after that rewrite to ORM
	rows, err := cd.db.Raw(getBasicArtifactInfo).Rows()
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var (
			id                int64
			creator           *string
			artifactStyleName *string
			transferredBy     string
			dateExc           string
			height            int64
			width             int64
			length            int64
			safety            string
		)
		err := rows.Scan(&id, &creator, &artifactStyleName,
			&transferredBy, &dateExc, &height, &width, &length, &safety)
		if err != nil {
			log.Println("scan error:", err)
		}
		card := new(ArtifactMasterPhas)
		card.ID = id
		if creator != nil {
			card.Creator = *creator
		}
		if artifactStyleName != nil {
			card.ArtifactStyle = *artifactStyleName
		}

		card.TransferredBy = transferredBy
		card.ExcavationDate = dateExc
		card.ArtifactMeasurement = &ArtifactMeasurement{
			Height: height,
			Width:  width,
			Length: length,
		}
		card.Safety = safety
		cards = append(cards, card)
	}
	return cards, nil
}
