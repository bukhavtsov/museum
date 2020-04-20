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

// ArtifactMaster the main structure of artifact
type ArtifactMaster struct {
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
		err := artifactRows.Scan(&id, &creator, &artifactStyleName,
			&transferredBy, &dateExc, &height, &width, &length, &safety)
		if err != nil {
			log.Println("scan error:", err)
		}
		card := new(ArtifactMaster)
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

	for _, card := range cards {
		elementsRows, err := cd.db.Raw(getArtifactElements+" WHERE ae1.artifact_id = ?", card.ID).Rows()
		if err != nil {
			log.Println(err)
		}

		for elementsRows.Next() {
			var (
				id       int64
				name     string
				parentID int64
			)

			err := elementsRows.Scan(&id, &name, &parentID)
			if err != nil {
				log.Println("scan error:", err)
			}
			log.Println(id)
			log.Println(name)
			log.Println(parentID)
			log.Println("-------")
			if parentID == 0 { // FIXME: draft condition for testing rows operation
				sqlCondition := "WHERE ae1.artifact_id = ? AND ae1.artifact_parent_element_id = ?"
				subElementsRows, err := cd.db.Raw(getArtifactElements+" "+sqlCondition, card.ID, id).Rows()
				if err != nil {
					log.Println(err)
				}
				for subElementsRows.Next() {
					var (
						subID       int64
						subName     string
						subParentID int64
					)
					err := subElementsRows.Scan(&subID, &subName, &subParentID)
					if err != nil {
						log.Println("sub scan error:", err)
					}
					log.Println("		", subID)
					log.Println("		", subName)
					log.Println("		", subParentID)
					log.Println("-------------------")
				}
				subElementsRows.Close()
			}
		}
		defer elementsRows.Close()
	}

	return cards, nil
}

//TODO: добавить условие: если parent_id != 0, то добавить текущий подэлемент в список к корневому
// 1. ReadAll разбить на методы
// 2. Присвоить иерархию элементов объекту card
