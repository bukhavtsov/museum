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

// ArtifactData gets connection to database
type ArtifactData struct {
	db *gorm.DB
}

// NewCardData creates new instance
func NewCardData(db *gorm.DB) *ArtifactData {
	return &ArtifactData{db}
}

// ReadAll return all cards from database
func (cd *ArtifactData) ReadAll() ([]*ArtifactMaster, error) {
	cards := make([]*ArtifactMaster, 0)
	// Way how to get data with relationship from db has been found, but it's not a ORM way
	// TODO: write working version without ORM way, after that rewrite to ORM
	artifactRows, err := cd.db.Raw(getArtifactsWithBasicInfo).Rows()
	if err != nil {
		log.Println(err)
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		card := getArtifactWithBasicInfo(artifactRows)
		err := cd.initArtifactElements(card)
		if err != nil {
			log.Println(err)
		}
		cards = append(cards, card)
	}
	return cards, nil
}

func getArtifactWithBasicInfo(artifactRows *sql.Rows) *ArtifactMaster {
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
		log.Println("getArtifactWithBasicInfo scan error:", err)
	}
	artifact := new(ArtifactMaster)
	artifact.ID = id
	if creator != nil {
		artifact.Creator = *creator
	}
	if artifactStyleName != nil {
		artifact.ArtifactStyle = *artifactStyleName
	}
	if transferredBy != nil {
		artifact.TransferredBy = *transferredBy
	}
	if dateExc != nil {
		artifact.ExcavationDate = *dateExc
	}
	if safety != nil {
		artifact.Safety = *safety
	}
	artifact.ArtifactMeasurement = &ArtifactMeasurement{}
	artifact.ArtifactMeasurement.Height = height
	artifact.ArtifactMeasurement.Width = width
	artifact.ArtifactMeasurement.Length = length
	return artifact
}

func (cd *ArtifactData) initArtifactElements(artifact *ArtifactMaster) error {
	card.Elements = make(map[string][]string, 0)
	elementsRows, err := cd.db.Raw(getArtifactElements+" WHERE ae1.artifact_id = ?", artifact.ID).Rows()
	if err != nil {
		log.Println(err)
	}
	defer elementsRows.Close()
	for elementsRows.Next() {
		var (
			id         int64
			name       string
			parentName string
		)
		err := elementsRows.Scan(&id, &name, &parentName)
		if err != nil {
			log.Println("scan error:", err)
			return 
		}
		if parentName != "" {
			artifact.Elements[parentName] = append(artifact.Elements[parentName], name)
		}
	}
}

func (cd *ArtifactData) getArtifactParentElement(artifactID, parentElementID int64) *ArtifactElement {
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
