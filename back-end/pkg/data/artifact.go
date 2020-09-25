package data

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

type ArtifactMeasurement struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
	Length int64 `json:"length"`
}

// ArtifactMaster the main structure of artifact

// add material
type ArtifactMaster struct {
	ID                  int64                `json:"id"`
	Creator             string               `json:"creator"`
	ArtifactStyle       string               `json:"artifact_style"`
	ExcavationDate      string               `json:"date_exc"`
	TransferredBy       string               `json:"transferred_by"`
	ArtifactMeasurement *ArtifactMeasurement `json:"artifact_measurement"`
	Elements            map[string][]string  `json:"artifact_elements"`
	ObjectGroup         map[string][]string  `json:"artifact_object_group"`
	Preservation        map[string][]string  `json:"artifact_preservation"`
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
func (cd *ArtifactData) ReadAll() ([]*ArtifactMaster, error) {
	artifacts := make([]*ArtifactMaster, 0)
	// Way how to get data with relationship from db has been found, but it's not a ORM way
	// TODO: write working version without ORM way, after that rewrite to ORM
	artifactRows, err := cd.db.Raw(getArtifactsWithBasicInfoQuery).Rows()
	if err != nil {
		log.Println(err)
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		artifact := getArtifactWithBasicInfo(artifactRows)
		err := cd.initArtifactElements(artifact)
		if err != nil {
			log.Println(err)
		}
		err = cd.initArtifactObjectGroup(artifact)
		if err != nil {
			log.Println(err)
		}
		err = cd.initArtifactPreservation(artifact)
		if err != nil {
			log.Println(err)
		}
		if artifact.ID == 0 { //FIXME: hot fix, because rows return n+1 elements
			continue
		}
		artifacts = append(artifacts, artifact)
	}
	return artifacts, nil
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
	)
	err := artifactRows.Scan(&id, &creator, &artifactStyleName,
		&transferredBy, &dateExc, &height, &width, &length)
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
	artifact.ArtifactMeasurement = &ArtifactMeasurement{}
	artifact.ArtifactMeasurement.Height = height
	artifact.ArtifactMeasurement.Width = width
	artifact.ArtifactMeasurement.Length = length
	return artifact
}

func (cd *ArtifactData) initArtifactElements(artifact *ArtifactMaster) error {
	artifact.Elements = make(map[string][]string, 0)
	elementsRows, err := cd.db.Raw(getArtifactElementByIdQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("elementsRows.cd.db.Raw err: %s", err)
	}
	defer elementsRows.Close()
	for elementsRows.Next() {
		var (
			id                int64
			childElementName  string
			parentElementName sql.NullString
		)
		err := elementsRows.Scan(&id, &childElementName, &parentElementName)
		if err != nil {
			return fmt.Errorf("elementsRows.Scan err: %s", err)
		}
		if value, _ := parentElementName.Value(); value != nil {
			artifact.Elements[parentElementName.String] = append(artifact.Elements[parentElementName.String], childElementName)
		} else {
			_, ok := artifact.Elements[childElementName]
			if !ok {
				artifact.Elements[childElementName] = make([]string, 0)
			}
		}
	}
	return nil
}

func (cd *ArtifactData) initArtifactObjectGroup(artifact *ArtifactMaster) error {
	artifact.ObjectGroup = make(map[string][]string, 0)
	objectGroupRows, err := cd.db.Raw(getArtifactObjectGroupByIdQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("objectGroupRows.cd.db.Raw err: %s", err)
	}
	defer objectGroupRows.Close()
	for objectGroupRows.Next() {
		var (
			id                int64
			childObjectGroup  string
			parentObjectGroup sql.NullString
		)
		err := objectGroupRows.Scan(&id, &childObjectGroup, &parentObjectGroup)
		if err != nil {
			return fmt.Errorf("objectGroupRows.Scan err: %s", err)
		}
		if value, _ := parentObjectGroup.Value(); value != nil {
			artifact.ObjectGroup[parentObjectGroup.String] = append(artifact.ObjectGroup[parentObjectGroup.String], childObjectGroup)
		} else {
			_, ok := artifact.ObjectGroup[childObjectGroup]
			if !ok {
				artifact.ObjectGroup[childObjectGroup] = make([]string, 0)
			}
		}
	}
	return nil
}

func (cd *ArtifactData) initArtifactPreservation(artifact *ArtifactMaster) error {
	artifact.Preservation = make(map[string][]string, 0)
	preservationRows, err := cd.db.Raw(getArtifactPreservationByIdQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("preservationRows.cd.db.Raw err: %s", err)
	}
	defer preservationRows.Close()
	for preservationRows.Next() {
		var (
			id                 int64
			childPreservation  string
			parentPreservation sql.NullString
		)
		err := preservationRows.Scan(&id, &childPreservation, &parentPreservation)
		if err != nil {
			err := preservationRows.Scan(&id, &childPreservation, &parentPreservation)
			return fmt.Errorf("preservationRows.Scan err: %s", err)
		}

		if value, _ := parentPreservation.Value(); value != nil {
			artifact.Preservation[parentPreservation.String] = append(artifact.Preservation[parentPreservation.String], childPreservation)
		} else {
			_, ok := artifact.Preservation[childPreservation]
			if !ok {
				artifact.Preservation[childPreservation] = make([]string, 0)
			}
		}
	}
	return nil
}

func (cd *ArtifactData) Add(artifactMaster *ArtifactMaster) (int64, error) {
	//TODO: investigate transaction and rollback. Actually, I make insertion to database in different tables.
	// Need to make a rollback in case if we will got a failure in the insertion data time.

	// first of all need to insert data to tables to which we have a foreign key
	insertedTransferredById, err := cd.insertTransferredBy(artifactMaster.TransferredBy)
	if err != nil {
		return -1, err
	}
	fmt.Println("transferredByID:", insertedTransferredById)
	insertedStyleLUTID, err := cd.insertStyleLUT(artifactMaster.ArtifactStyle)
	if err != nil {
		return -1, err
	}
	fmt.Println("insertedStyleLUTID:", insertedStyleLUTID)

	insertedArtifactMasterID, err := cd.insertArtifactMaster(artifactMaster, insertedTransferredById)
	if err != nil {
		return -1, err
	}
	fmt.Println("insertedArtifactMasterID:", insertedArtifactMasterID)

	// can return nil pointer
	insertedMeasurementID, err := cd.insertMeasurement(insertedArtifactMasterID, artifactMaster.ArtifactMeasurement)
	if err != nil {
		return -1, err
	}
	fmt.Println("insertedMeasurementID:", insertedMeasurementID)
	//should be after the artifactMaster initialization
	//insertedStyleID, err := cd.insertStyle(insertedArtifactMasterID, insertedStyleLUTID)
	//if err != nil {
	//	return -1, err
	//}
	//fmt.Println("insertedStyleID", insertedStyleID)

	return -1, nil
}

func (cd *ArtifactData) insertTransferredBy(transferredBy string) (insertedTransferredById int64, err error) {
	result := cd.db.Exec(insertTransferredBy, transferredBy)
	if result.Error != nil {
		return -1, err
	}
	transferredByRows, err := cd.db.Raw(selectTransferredBy, transferredBy).Rows()
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

func (cd *ArtifactData) insertStyleLUT(style string) (insertedStyleLUTID int64, err error) {
	result := cd.db.Exec(insertArtifactStyleLUT, style)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := cd.db.Raw(selectArtifactStyleLUT, style).Rows()
	if err != nil {
		return -1, err
	}
	for artifactStyleRows.Next() {
		err := artifactStyleRows.Scan(&insertedStyleLUTID)
		if err != nil {
			return -1, err
		}
		return insertedStyleLUTID, nil
	}
	return insertedStyleLUTID, nil
}

func (cd *ArtifactData) insertStyle(artifactID, styleLUTID int64) (insertedStyleID int64, err error) {
	result := cd.db.Exec(insertArtifactStyle, artifactID, styleLUTID)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := cd.db.Raw(selectArtifactStyle, artifactID, styleLUTID).Rows()
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

func (cd *ArtifactData) insertArtifactMaster(master *ArtifactMaster, insertedTransferredById int64) (insertedArtifactMasterID int64, err error) {
	result := cd.db.Exec(insertArtifactMaster, master.Creator, master.ExcavationDate, insertedTransferredById)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := cd.db.Raw(
		selectArtifactMaster,
		master.Creator,
		master.ExcavationDate,
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

func (cd *ArtifactData) insertMeasurement(artifactID int64, artifactMeasurement *ArtifactMeasurement) (insertedMeasurement int64, err error) {
	result := cd.db.Exec(
		insertMeasurement,
		artifactID,
		artifactMeasurement.Length,
		artifactMeasurement.Height,
		artifactMeasurement.Width,
	)
	if result.Error != nil {
		return -1, err
	}
	artifactStyleRows, err := cd.db.Raw(
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
