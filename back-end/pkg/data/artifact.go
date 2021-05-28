package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
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
	ArtifactStyle       string               `json:"artifact_style"` // should be array
	ExcavationDate      string               `json:"date_exc"`
	TransferredBy       string               `json:"transferred_by"`
	ArtifactMeasurement *ArtifactMeasurement `json:"artifact_measurement"`
	ArtifactElements    []ArtifactElement    `json:"artifact_elements,omitempty"`
}

type ArtifactElement struct {
	Name     string            `json:"name"`
	Children []ArtifactElement `json:"children,omitempty"`
}

// ArtifactData gets connection to database
type ArtifactData struct {
	db *gorm.DB
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (ArtifactElement) TableName() string {
	return "artifact_element"
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
		return nil, fmt.Errorf("err when try to read basic info, err: %w", err)
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		artifact, err := getArtifactWithBasicInfo(artifactRows)
		if err != nil {
			return nil, fmt.Errorf("got an error from getArtifactWithBasicInfo method, err is: %w", err)
		}
		artifact.ArtifactElements, err = a.readArtifactElements(artifact.ID)
		if err != nil {
			return nil, fmt.Errorf("got an error from readArtifactElements method, err is: %w", err)
		}
		artifacts = append(artifacts, artifact)
	}
	return artifacts, nil
}

// Read return artifact by id from database
func (a *ArtifactData) Read(id int) (*ArtifactMaster, error) {
	var artifact *ArtifactMaster
	artifactRows, err := a.db.Raw(getArtifactsWithBasicInfoByIDQuery, id).Rows()
	if err != nil {
		return nil, fmt.Errorf("error when tried to execute getArtifactsWithBasicInfoByIDQuery, err: %w", err)
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		artifact, err = getArtifactWithBasicInfo(artifactRows)
		if err != nil {
			return nil, fmt.Errorf("error when tried to getArtifactWithBasicInfo, err: %w", err)
		}
	}
	return artifact, nil
}

func getArtifactWithBasicInfo(artifactRows *sql.Rows) (*ArtifactMaster, error) {
	var (
		id                *int
		creator           *string
		artifactStyleName *string
		transferredBy     *string
		dateExc           *string
		height            *int
		width             *int
		length            *int
	)
	err := artifactRows.Scan(&id, &creator, &artifactStyleName,
		&transferredBy, &dateExc, &height, &width, &length)
	if err != nil {
		return nil, fmt.Errorf("getArtifactWithBasicInfo scan error: %w", err)
	}
	artifact := new(ArtifactMaster)
	if id != nil {
		artifact.ID = *id
	}
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
	if height != nil {
		artifact.ArtifactMeasurement.Height = *height
	}
	if width != nil {
		artifact.ArtifactMeasurement.Width = *width
	}
	if length != nil {
		artifact.ArtifactMeasurement.Length = *length
	}

	return artifact, nil
}
func (a *ArtifactData) Add(artifactMaster ArtifactMaster) (int, error) {
	// first of all need to insert data to tables to which we have a foreign key

	transferredById, err := a.insertTransferredByLUTIfNotExists(artifactMaster.TransferredBy)
	if err != nil {
		return -1, fmt.Errorf("got an error when tried to call insertTransferredByLUTIfNotExists method, in Add, err: %w", err)
	}

	//artifactStyleLUTId, err := a.insertArtifactStyleLUTIfNotExists(artifactMaster.ArtifactStyle)
	//if err != nil {
	//	return -1, fmt.Errorf("got an error when tried to call insertArtifactStyleLUTIfNotExists method, in Add, err: %w", err)
	//}

	insertedArtifactMasterID, err := a.insertArtifactMaster(
		artifactMaster.Creator,
		artifactMaster.ExcavationDate,
		transferredById,
	)
	if err != nil {
		return -1, fmt.Errorf("error when tried to insertArtifactMaster, err %w", err)
	}

	if &artifactMaster.ArtifactMeasurement != nil {
		_, err = a.insertMeasurement(insertedArtifactMasterID, artifactMaster.ArtifactMeasurement)
	} else {
		log.Println("ArtifactMeasurement has not been added, because has a nil value")
	}
	if err != nil {
		return -1, fmt.Errorf("error when tried to insertMeasurement, err %w", err)
	}

	err = a.insertArtifactElements(insertedArtifactMasterID, artifactMaster.ArtifactElements)
	if err != nil {
		return -1, fmt.Errorf("error when tried to insertArtifactElements %w", err)
	}
	return insertedArtifactMasterID, nil
}

func (a *ArtifactData) insertArtifactElements(artifactMasterID int, elements []ArtifactElement) error {
	for _, element := range elements {
		_, err := a.insertArtifactElement(artifactMasterID, element)
		if err != nil {
			return fmt.Errorf("error from insertArtifactElements when tried to execute insertArtifactElement, err: %w", err)
		}
	}
	return nil
}

func (a *ArtifactData) insertArtifactElement(artifactMasterID int, element ArtifactElement) (int, error) {
	b, err := json.Marshal(&element)
	if err != nil {
		return -1, fmt.Errorf("err when tried Marshal %v, err: %w", element, err)
	}
	elementJSON := string(b)
	row := a.db.Raw(insertArtifactElement, artifactMasterID, elementJSON).Row()
	if row.Err() != nil {
		return -1, fmt.Errorf("err when tried to insert %v into the db. err: %w", element, row.Err())
	}
	var lastInsertedID int
	err = row.Scan(&lastInsertedID)
	if err != nil {
		return -1, fmt.Errorf("err when tried to scan inserted id from db. err: %w", err)
	}
	return lastInsertedID, nil

}

func (a *ArtifactData) Update(artifactMasterID int, newArtifactMaster *ArtifactMaster) error {
	newTransferredById, err := a.insertTransferredByLUTIfNotExists(newArtifactMaster.TransferredBy)
	if err != nil {
		return fmt.Errorf("got an error when tried to call insertTransferredByLUTIfNotExists method, err: %w", err)
	}
	updateArtifactMasterRow := a.db.Exec(updateArtifactMaster, newArtifactMaster.Creator, newArtifactMaster.ExcavationDate, newTransferredById, artifactMasterID)
	if updateArtifactMasterRow.Error != nil {
		return fmt.Errorf("got an error when tried to updateArtifactMaster %w", err)
	}

	updateArtifactMeasurementRow := a.db.Exec(updateArtifactMeasurement, newArtifactMaster.ArtifactMeasurement.Length, newArtifactMaster.ArtifactMeasurement.Height, newArtifactMaster.ArtifactMeasurement.Width, artifactMasterID)
	if updateArtifactMeasurementRow.Error != nil {
		return fmt.Errorf("got an error when tried to updateArtifactMeasurementRow %w", err)
	}

	//artifactStyleLUTID, err := a.insertArtifactStyleLUTIfNotExists(newArtifactMaster.ArtifactStyle)
	//if err != nil {
	//	return fmt.Errorf("got an error when tried to insertArtifactStyleLUTIfNotExists, error is : %w", err)
	//}
	//
	//updateArtifactStyleRow:= a.db.Exec(updateArtifactStyle, artifactStyleLUTID, artifactMasterID)
	//if updateArtifactStyleRow.Error != nil {
	//	return fmt.Errorf("can't execute updateArtifactStyle, got an error: %e", err)
	//}
	return nil
}

// 	check type in artifact style lut ? yes : return artifact_style_id
// 	otherwise add new type to artifact style lut, then return  artifact_style_id
func (a *ArtifactData) insertArtifactStyleLUTIfNotExists(artifactStyle string) (int, error) {
	var artifactStyleLUTExistingID int
	err := a.db.Raw(selectArtifactStyleLUT, artifactStyle).Row().Scan(&artifactStyleLUTExistingID)
	if err != nil {
		log.Printf("no artifact style lut record with artifactStyle: %s. Create new", artifactStyle)
		artifactStyleLUTInsertedID, err := a.insertStyleLUT(artifactStyle)
		if err != nil {
			return -1, fmt.Errorf("insertArtifactStyleLUTIfNotExists called insertStyleLUT err: %w", err)
		}
		return artifactStyleLUTInsertedID, nil
	}
	return artifactStyleLUTExistingID, nil
}

func (a *ArtifactData) insertTransferredByLUTIfNotExists(transferredBy string) (int, error) {
	var transferredByIDExisting int
	err := a.db.Raw(getTransferredByIdFieldByName, transferredBy).Row().Scan(&transferredByIDExisting)
	if err != nil {
		log.Printf("no transfered by with name %s. Create new", transferredBy)
		transferredById, err := a.insertTransferredByLUT(transferredBy)
		if err != nil {
			return -1, fmt.Errorf("insertTransferredByLUTIfNotExists called insertTransferredByLUT err: %w", err)
		}
		return transferredById, nil
	}
	return transferredByIDExisting, nil
}

func (a *ArtifactData) insertTransferredByLUT(transferredBy string) (insertedTransferredById int, err error) {
	result := a.db.Exec(insertTransferredBy, transferredBy)
	if result.Error != nil {
		return -1, fmt.Errorf("error when tried to exec insertTransferredBy query, err: %w", result.Error)
	}
	transferredByRow := a.db.Raw(selectTransferredBy, transferredBy).Row()
	err = transferredByRow.Scan(&insertedTransferredById)
	if err != nil {
		return -1, fmt.Errorf("error when tried to scan insertedTransferredById, err: %v", err)
	}
	return insertedTransferredById, nil
}

func (a *ArtifactData) insertStyleLUT(style string) (insertedStyleLUTID int, err error) {
	result := a.db.Exec(insertArtifactStyleLUT, style)
	if result.Error != nil {
		return -1, fmt.Errorf("error when tried to exec insertArtifactStyleLUT query, err: %v", err)
	}
	artifactStyleRows := a.db.Raw(selectArtifactStyleLUT, style).Row()
	err = artifactStyleRows.Scan(&insertedStyleLUTID)
	if err != nil {
		return -1, fmt.Errorf("error when tried to scan insertedStyleLUTID, err: %v", err)
	}
	return insertedStyleLUTID, nil
}

func (a *ArtifactData) insertStyle(artifactID, styleLUTID int) (insertedStyleID int, err error) {
	result := a.db.Exec(insertArtifactStyle, artifactID, styleLUTID)
	if result.Error != nil {
		return -1, fmt.Errorf("error when tried to exec insertArtifactStyle query, err: %v", err)
	}
	artifactStyleRows, err := a.db.Raw(selectArtifactStyle, artifactID, styleLUTID).Rows()
	if err != nil {
		return -1, err
	}
	for artifactStyleRows.Next() {
		err := artifactStyleRows.Scan(&insertedStyleID)
		if err != nil {
			return -1, fmt.Errorf("error when tried to scan insertedStyleID, err: %v", err)
		}
		return insertedStyleID, nil
	}
	return insertedStyleID, nil
}

func (a *ArtifactData) insertArtifactMaster(creator string, excavationDate string, insertedTransferredById int) (insertedArtifactMasterID int, err error) {
	result := a.db.Raw(insertArtifactMaster, creator, excavationDate, insertedTransferredById).Row()
	if result.Err() != nil {
		return -1, fmt.Errorf("error when tried to exec insertArtifactMaster query, err: %v", err)
	}
	err = result.Scan(&insertedArtifactMasterID)
	if err != nil {
		return -1, fmt.Errorf("error when tried to scan insertedArtifactMasterID, err: %v", err)
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
		return -1, fmt.Errorf("error when tried to exec insertMeasurement, err: %v", err)
	}
	artifactMeasurementRows, err := a.db.Raw(
		selectArtifactMeasurement,
		artifactID,
		artifactMeasurement.Length,
		artifactMeasurement.Height,
		artifactMeasurement.Width).Rows()
	if err != nil {
		return -1, fmt.Errorf("error when tried to Raw selectArtifactMeasurement, err: %v", err)
	}
	for artifactMeasurementRows.Next() {
		err := artifactMeasurementRows.Scan(&insertedMeasurement)
		if err != nil {
			return -1, fmt.Errorf("error when tried to scan insertedMeasurement, err: %v", err)
		}
		return insertedMeasurement, nil
	}
	return insertedMeasurement, nil
}

func (a *ArtifactData) Delete(artifactId int) error {
	resDeleteMeasurement := a.db.Exec(deleteMeasurement, artifactId)
	if resDeleteMeasurement.Error != nil {
		return fmt.Errorf("got an error when tried to execute deleteMeasurement, error is: %w", resDeleteMeasurement.Error)
	}

	resDeleteArtifactElements := a.db.Exec(deleteArtifactElements, artifactId)
	if resDeleteArtifactElements.Error != nil {
		return fmt.Errorf("err when tried to delete ArtifactElements, err: %w", resDeleteArtifactElements.Error)
	}

	resDeleteArtifactMaster := a.db.Exec(deleteArtifactMaster, artifactId)
	if resDeleteArtifactMaster.Error != nil {
		return fmt.Errorf("got an error when tried to execute resDeleteArtifactMaster, error is: %w", resDeleteArtifactMaster.Error)
	}
	return nil
}

func (a *ArtifactData) readArtifactElements(artifactId int) ([]ArtifactElement, error) {
	var artifactElements []ArtifactElement
	artifactElementsRows, err := a.db.Raw(selectArtifactElement, artifactId).Rows()
	if err != nil {
		return nil, fmt.Errorf("got an error when tried to selectArtifactElement, error is: %w", err)
	}
	for artifactElementsRows.Next() {
		var (
			artifactElementJSON string
			artifactElement     ArtifactElement
		)
		err := artifactElementsRows.Scan(&artifactElementJSON)
		if err != nil {
			return nil, fmt.Errorf("err when tried to scan artifactElementJSON, err: %w", err)
		}
		err = json.Unmarshal([]byte(artifactElementJSON), &artifactElement)
		if err != nil {
			return nil, fmt.Errorf("err when tried to Unmarshal artifactElementsJSON, err: %w", err)
		}
		artifactElements = append(artifactElements, artifactElement)
	}

	return artifactElements, nil
}
