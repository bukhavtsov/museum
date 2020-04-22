package data

import (
	"database/sql"
	"log"

	"github.com/jinzhu/gorm"
)

type ArtifactMeasurement struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
	Length int64 `json:"length"`
}

// ArtifactMaster the main structure of artifact
type ArtifactMaster struct {
	ID                  int64                `json:"id"`
	Creator             string               `json:"creator"`
	ArtifactStyle       string               `json:"artifact_style"`
	ExcavationDate      string               `json:"date_exc"`
	TransferredBy       string               `json:"transferred_by"`
	Safety              string               `json:"safety"` // rewrite with graph
	ArtifactMeasurement *ArtifactMeasurement `json:"artifact_measurement"`
	Elements            map[string][]string  `json:"artifact_elements"`
	//Elements            []*ArtifactElement
}

// ArtifactElement some part of artifact
type ArtifactElement struct {
	ID         int64            `json:"id"`
	Name       string           `json:"name"`
	SubElement *ArtifactElement `json:"sub_element"`
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
func (cd *CardData) ReadAll() ([]*ArtifactMaster, error) {
	cards := make([]*ArtifactMaster, 0)
	// Way how to get data with relationship from db has been found, but it's not a ORM way
	// TODO: write working version without ORM way, after that rewrite to ORM
	artifactRows, err := cd.db.Raw(getArtifactsWithBasicInfo).Rows()
	if err != nil {
		log.Println(err)
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		card := getCardWithBasicInfo(artifactRows)
		card.Elements = make(map[string][]string, 0)
		cards = append(cards, card)
	}

	for _, card := range cards {
		elementsRows, err := cd.db.Raw(getArtifactElements+" WHERE ae1.artifact_id = ?", card.ID).Rows()
		if err != nil {
			log.Println(err)
		}

		for elementsRows.Next() {
			var (
				id         int64
				name       string
				parentName string
			)
			err := elementsRows.Scan(&id, &name, &parentName)
			if err != nil {
				log.Println("scan error:", err)
			}
			if parentName != "" {
				card.Elements[parentName] = append(card.Elements[parentName], name)
			}
		}
		defer elementsRows.Close()
	}

	return cards, nil
}

func getCardWithBasicInfo(artifactRows *sql.Rows) *ArtifactMaster {
	var (
		id                int64
		creator           *string
		artifactStyleName *string
		transferredBy     *string
		dateExc           *string
		height            int64
		width             int64
		length            int64
		safety            *string
	)
	err := artifactRows.Scan(&id, &creator, &artifactStyleName,
		&transferredBy, &dateExc, &height, &width, &length, &safety)
	if err != nil {
		log.Println("getCardWithBasicInfo scan error:", err)
	}
	card := new(ArtifactMaster)
	card.ID = id
	if creator != nil {
		card.Creator = *creator
	}
	if artifactStyleName != nil {
		card.ArtifactStyle = *artifactStyleName
	}
	if transferredBy != nil {
		card.TransferredBy = *transferredBy
	}
	if dateExc != nil {
		card.ExcavationDate = *dateExc
	}
	if safety != nil {
		card.Safety = *safety
	}
	card.ArtifactMeasurement = &ArtifactMeasurement{}
	card.ArtifactMeasurement.Height = height
	card.ArtifactMeasurement.Width = width
	card.ArtifactMeasurement.Length = length
	return card
}

func (cd *CardData) getArtifactParentElement(artifactID, parentElementID int64) *ArtifactElement {
	const sqlCondition = "WHERE ae1.artifact_id = ? AND ae1.artifact_parent_element_id = ?"
	subElementsRows, err := cd.db.Raw(getArtifactElements+" "+sqlCondition, artifactID, parentElementID).Rows()
	if err != nil {
		log.Println(err)
	}
	for subElementsRows.Next() {
		parent := new(ArtifactElement)
		log.Println(parent)
	}
	return nil
}
