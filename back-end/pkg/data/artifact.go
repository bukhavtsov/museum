package data

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type ArtifactMeasurement struct {
	Height int `json:"height"`
	Width  int `json:"width"`
	Length int `json:"length"`
}

// ArtifactMaster the main structure of artifact

// add material
type ArtifactMaster struct {
	ID                  int                  `json:"id"`
	Creator             string               `json:"creator"`
	ArtifactStyle       string               `json:"artifact_style"`
	ExcavationDate      string               `json:"date_exc"`
	TransferredBy       string               `json:"transferred_by"`
	ArtifactMeasurement *ArtifactMeasurement `json:"artifact_measurement"`
}

// ArtifactData gets connection to database
type ArtifactData struct {
	db *gorm.DB
}

// NewArtifactData creates new instance
func NewArtifactData(db *gorm.DB) *ArtifactData {
	return &ArtifactData{db}
}

// ReadAll return all artifacts from database
func (a *ArtifactData) ReadAll() ([]*ArtifactMaster, error) {
	var artifacts []*ArtifactMaster
	artifactRows, err := a.db.Raw(getArtifactsWithBasicInfoQuery).Rows()
	if err != nil {
		return nil, err
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		artifact, err := getArtifactWithBasicInfo(artifactRows)
		if err != nil {
			log.Println("got an error from getArtifactWithBasicInfo method, err is:", err)
		} else {
			artifacts = append(artifacts, artifact)
		}
	}
	return artifacts, nil
}

func getArtifactWithBasicInfo(artifactRows *sql.Rows) (*ArtifactMaster, error) {
	var (
		id                int
		creator           *string
		artifactStyleName *string
		transferredBy     *string
		dateExc           *string
		height            int
		width             int
		length            int
	)
	err := artifactRows.Scan(&id, &creator, &artifactStyleName,
		&transferredBy, &dateExc, &height, &width, &length)
	if err != nil {
		return nil, fmt.Errorf("getArtifactWithBasicInfo scan error: %w", err)
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
	artifact.ArtifactMeasurement = &ArtifactMeasurement{}
	artifact.ArtifactMeasurement.Height = height
	artifact.ArtifactMeasurement.Width = width
	artifact.ArtifactMeasurement.Length = length
	return artifact, nil
}
func (a *ArtifactData) Add(artifactMaster *ArtifactMaster) (int, error) {
	// first of all need to insert data to tables to which we have a foreign key
	insertedTransferredById, err := a.insertTransferredBy(artifactMaster.TransferredBy)
	if err != nil {
		return -1, err
	}
	fmt.Println("transferredByID:", insertedTransferredById)

	insertedArtifactMasterID, err := a.insertArtifactMaster(
		artifactMaster.Creator,
		artifactMaster.ExcavationDate,
		insertedTransferredById,
	)
	if err != nil {
		return -1, err
	}
	fmt.Println("insertedArtifactMasterID:", insertedArtifactMasterID)

	// can return nil pointer
	insertedMeasurementID, err := a.insertMeasurement(insertedArtifactMasterID, artifactMaster.ArtifactMeasurement)
	if err != nil {
		return -1, err
	}
	fmt.Println("insertedMeasurementID:", insertedMeasurementID)

	insertedStyleLUTID, err := a.insertStyleLUT(artifactMaster.ArtifactStyle)
	if err != nil {
		return -1, err
	}
	fmt.Println("insertedStyleLUTID:", insertedStyleLUTID)
	_, err = a.insertStyle(insertedArtifactMasterID, insertedStyleLUTID)
	if err != nil {
		return -1, err
	}

	return artifactMaster.ID, nil
}

func (a *ArtifactData) Update(id int, newArtifactMaster *ArtifactMaster) () {

}

func (a *ArtifactData) insertTransferredBy(transferredBy string) (insertedTransferredById int, err error) {
	result := a.db.Exec(insertTransferredBy, transferredBy)
	if result.Error != nil {
		return -1, err
	}
	transferredByRows, err := a.db.Raw(selectTransferredBy, transferredBy).Rows()
	if err != nil {
		return -1, err
	}
	for transferredByRows.Next() {
		err := transferredByRows.Scan(&insertedTransferredById)
		if err != nil {
			return -1, err
		}
		return insertedTransferredById, nil
	}
	return insertedTransferredById, nil
}

func (a *ArtifactData) insertStyleLUT(style string) (insertedStyleLUTID int, err error) {
	result := a.db.Exec(insertArtifactStyleLUT, style)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows := a.db.Raw(selectArtifactStyleLUT, style).Row()
	err = artifactStyleRows.Scan(&insertedStyleLUTID)
	if err != nil {
		return -1, err
	}
	return insertedStyleLUTID, nil
}

func (a *ArtifactData) insertStyle(artifactID, styleLUTID int) (insertedStyleID int, err error) {
	result := a.db.Exec(insertArtifactStyle, artifactID, styleLUTID)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := a.db.Raw(selectArtifactStyle, artifactID, styleLUTID).Rows()
	if err != nil {
		return -1, err
	}
	for artifactStyleRows.Next() {
		err := artifactStyleRows.Scan(&insertedStyleID)
		if err != nil {
			return -1, err
		}
		return insertedStyleID, nil
	}
	return insertedStyleID, nil
}

func (a *ArtifactData) insertArtifactMaster(creator string, excavationDate string, insertedTransferredById int) (insertedArtifactMasterID int, err error) {
	result := a.db.Exec(insertArtifactMaster, creator, excavationDate, insertedTransferredById)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := a.db.Raw(
		selectArtifactMaster,
		creator,
		excavationDate,
		insertedTransferredById).Rows()
	if err != nil {
		return -1, err
	}
	for artifactStyleRows.Next() {
		err := artifactStyleRows.Scan(&insertedArtifactMasterID)
		if err != nil {
			return -1, err
		}
		return insertedArtifactMasterID, nil
	}
	return insertedArtifactMasterID, nil
}

func (a *ArtifactData) insertMeasurement(artifactID int, artifactMeasurement *ArtifactMeasurement) (insertedMeasurement int, err error) {
	result := a.db.Exec(
		insertMeasurement,
		artifactID,
		artifactMeasurement.Length,
		artifactMeasurement.Height,
		artifactMeasurement.Width,
	)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := a.db.Raw(
		selectArtifactMeasurement,
		artifactID,
		artifactMeasurement.Length,
		artifactMeasurement.Height,
		artifactMeasurement.Width).Rows()
	if err != nil {
		return -1, err
	}
	for artifactStyleRows.Next() {
		err := artifactStyleRows.Scan(&insertedMeasurement)
		if err != nil {
			return -1, err
		}
		return insertedMeasurement, nil
	}
	return insertedMeasurement, nil
}
