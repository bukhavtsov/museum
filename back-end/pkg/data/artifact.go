package data

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

// Measurement is artifact parameters
type ArtifactMeasurement struct {
	Height int64 `json:"height"`
	Width  int64 `json:"width"`
	Length int64 `json:"length"`
}

// ArtifactElement is part of artifact
type ArtifactElement struct {
	ID            int64              `json:"id"`
	ArtifactID    int64              `json:"artifact_id"`
	Name          string             `json:"name"`
	ParentID      int64              `json:"parent_id"`
	ChildElements []*ArtifactElement `json:"child_elements"`
}

// ArtifactMaterial deeply description of artifact elements
type ArtifactMaterial struct {
	ID             int64               `json:"id"`
	ArtifactID     int64               `json:"artifact_id"`
	Quantity       int64               `json:"quantity"`
	Composition    int64               `json:"composition"`
	MaterialType   string              `json:"material_type"`
	ParentID       int64               `json:"parent_id"`
	ChildMaterials []*ArtifactMaterial `json:"child_materials"`
}

type ArtifactObjectGroup struct {
	ID               int64                  `json:"id"`
	ArtifactID       int64                  `json:"artifact_id"`
	ParentID         int64                  `json:"parent_id"`
	Name             string                 `json:"object_group_name"`
	ChildObjectGroup []*ArtifactObjectGroup `json:"child_object_group"`
}

type ArtifactPreservation struct {
	ID                int64                   `json:"id"`
	ArtifactID        int64                   `json:"artifact_id"`
	ParentID          int64                   `json:"parent_id"`
	Name              string                  `json:"preservation"`
	ChildPreservation []*ArtifactPreservation `json:"child_preservation"`
}

// ArtifactMaster the main structure of artifact
type ArtifactMaster struct {
	ID             int64                   `json:"id"`
	Creator        string                  `json:"creator"`
	ArtifactStyle  string                  `json:"artifact_style"`
	ExcavationDate string                  `json:"date_exc"`
	TransferredBy  string                  `json:"transferred_by"`
	Measurement    *ArtifactMeasurement    `json:"artifact_measurement"`
	Elements       []*ArtifactElement      `json:"artifact_elements"`
	Materials      []*ArtifactMaterial     `json:"artifact_materials"`
	ObjectGroup    []*ArtifactObjectGroup  `json:"artifact_object_group"`
	Preservation   []*ArtifactPreservation `json:"artifact_preservation"`
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
func (ad *ArtifactData) ReadAll() ([]*ArtifactMaster, error) {
	artifacts := make([]*ArtifactMaster, 0)
	// Way how to get data with relationship from db has been found, but it's not a ORM way
	// TODO: write working version without ORM way, after that rewrite to ORM
	artifactRows, err := ad.db.Raw(getArtifactsWithBasicInfoQuery).Rows()
	if err != nil {
		log.Println(err)
	}
	defer artifactRows.Close()
	for artifactRows.Next() {
		artifact := getArtifactWithBasicInfo(artifactRows)
		err := ad.initArtifactElements(artifact)
		if err != nil {
			log.Println(err)
		}
		err = ad.initArtifactObjectGroup(artifact)
		if err != nil {
			log.Println(err)
		}
		err = ad.initArtifactPreservation(artifact)
		if err != nil {
			log.Println(err)
		}
		err = ad.initArtifactMaterials(artifact)
		if err != nil {
			log.Println(err)
		}
		artifacts = append(artifacts, artifact)
	}
	return artifacts, nil
}

func (ad *ArtifactData) Create(artifact *ArtifactMaster) (int64, error) {
	err := ad.db.Exec(insertTransferredByLUTQuery, artifact.TransferredBy).Error
	var transferredID int64
	if err != nil {
		log.Println("transferredBy field didn't added, err:", err)
	}
	transferredID, err = ad.getTransferredLUTID(artifact.TransferredBy)
	if err != nil {
		log.Println("transferredBy field not found in db err:", err)
	}
	err = ad.db.Exec(insertArtifactMasterQuery,
		1, // default museum is Vetka
		1, // default excavation region is Vetka
		1, // default reg_confidence_id
		artifact.Creator,
		artifact.ExcavationDate,
		nil,           // default hist culture id
		nil,           // desc is nil by default
		nil,           // translation is nil by default
		nil,           // age min nil
		nil,           // age max nil
		nil,           // artifact info photo nil
		nil,           // artifact photo nil
		transferredID, // transferred_by_id
	).Error
	if err != nil {
		log.Println(err)
	}

	artifactRows, err := ad.db.Raw("SELECT MAX(id) FROM artifact_master").Rows()
	if err != nil {
		log.Println("artifactRows err:", err)
		return -1, err
	}

	defer artifactRows.Close()
	var artifactID int64
	artifactRows.Next()
	err = artifactRows.Scan(&artifactID)
	if err != nil {
		log.Println("artifactRows.Scan err:", err)
		return -1, err
	}

	if artifact.Measurement != nil {
		err = ad.db.Exec(
			artifactMeasurementQuery,
			artifactID,
			artifact.Measurement.Length,
			artifact.Measurement.Height,
			artifact.Measurement.Width,
		).Error
		if err != nil {
			log.Println("artifactMeasurement field didn't added, err:", err)
		}
	}

	for _, object := range artifact.ObjectGroup {
		err = ad.db.Exec(insertObjectGroupQuery, object.Name, artifactID, object.ParentID).Error
		if err != nil {
			log.Printf("insertObjectGroupQuery err: %s\n", err)
		}
		for _, childObject := range object.ChildObjectGroup {
			err = ad.db.Exec(insertObjectGroupQuery, childObject.Name, artifactID, childObject.ParentID).Error
			if err != nil {
				log.Printf("childObject insertObjectGroupQuery err: %s\n", err)
			}
		}
	}

	return artifactID, nil
}

func (ad *ArtifactData) getTransferredLUTID(transferredBy string) (int64, error) {
	artifactRows, err := ad.db.Raw(getTransferredByLUTQuery, transferredBy).Rows()
	if err != nil {
		return -1, err
	}
	defer artifactRows.Close()
	var transferredID int64
	for artifactRows.Next() {
		err := artifactRows.Scan(&transferredID)
		if err != nil {
			return -1, err
		}
	}
	return transferredID, nil
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
	artifact.Measurement = &ArtifactMeasurement{}
	artifact.Measurement.Height = height
	artifact.Measurement.Width = width
	artifact.Measurement.Length = length
	return artifact
}

func (ad *ArtifactData) initArtifactElements(artifact *ArtifactMaster) error {
	artifact.Elements = make([]*ArtifactElement, 0)
	elementsRows, err := ad.db.Raw(getArtifactElementByIDQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("elementsRows.cd.db.Raw err: %s", err)
	}
	defer elementsRows.Close()
	for elementsRows.Next() {
		var (
			id               int64
			artifactId       int64
			childElementName string
			parentElementID  sql.NullInt64
		)
		err := elementsRows.Scan(&id, &artifactId, &childElementName, &parentElementID)
		if err != nil {
			return fmt.Errorf("elementsRows.Scan err: %s", err)
		}
		childElements := make([]*ArtifactElement, 0)
		childElements, err = ad.getArtifactChildElements(artifact.ID, id)
		element := &ArtifactElement{
			ID:            id,
			ArtifactID:    artifactId,
			Name:          childElementName,
			ParentID:      parentElementID.Int64,
			ChildElements: childElements,
		}
		if len(element.ChildElements) > 0 {
			artifact.Elements = append(artifact.Elements, element)
		}
	}
	return nil
}

func (ad *ArtifactData) getArtifactChildElements(artifactID, parentID int64) ([]*ArtifactElement, error) {
	childElements := make([]*ArtifactElement, 0)
	childElementsRows, err := ad.db.Raw(getArtifactChildElementQuery, artifactID, parentID).Rows()
	if err != nil {
		return nil, fmt.Errorf("childElementsRows.cd.db.Raw.err: %s", err)
	}
	defer childElementsRows.Close()
	for childElementsRows.Next() {
		var (
			id               int64
			childArtifactID  int64
			childElementName string
			parentElementID  sql.NullInt64
		)
		err := childElementsRows.Scan(&id, &childArtifactID, &childElementName, &parentElementID)
		if err != nil {
			return nil, fmt.Errorf("childElementsRows.Scan err: %s", err)
		}

		childElement := &ArtifactElement{
			ID:         id,
			ArtifactID: childArtifactID,
			Name:       childElementName,
			ParentID:   parentElementID.Int64,
		}
		childElements = append(childElements, childElement)
	}
	return childElements, nil
}

func (ad *ArtifactData) initArtifactObjectGroup(artifact *ArtifactMaster) error {
	artifact.ObjectGroup = make([]*ArtifactObjectGroup, 0)
	objectGroupRows, err := ad.db.Raw(getArtifactObjectGroupByIDQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("objectGroupRows.cd.db.Raw err: %s", err)
	}
	defer objectGroupRows.Close()
	for objectGroupRows.Next() {
		var (
			id              int64
			artifactId      int64
			objectGroupName string
			parentElementID sql.NullInt64
		)
		err := objectGroupRows.Scan(&id, &artifactId, &objectGroupName, &parentElementID)
		if err != nil {
			return fmt.Errorf("objectGroupRows.Scan err: %s", err)
		}
		childObjectGroup := make([]*ArtifactObjectGroup, 0)
		childObjectGroup, err = ad.getArtifactChildObjectGroup(artifact.ID, id)
		objectGroup := &ArtifactObjectGroup{
			ID:               id,
			ArtifactID:       artifactId,
			Name:             objectGroupName,
			ParentID:         parentElementID.Int64,
			ChildObjectGroup: childObjectGroup,
		}
		if len(objectGroup.ChildObjectGroup) > 0 {
			artifact.ObjectGroup = append(artifact.ObjectGroup, objectGroup)
		}
	}
	return nil
}

func (ad *ArtifactData) initArtifactPreservation(artifact *ArtifactMaster) error {
	artifact.Preservation = make([]*ArtifactPreservation, 0)
	preservationRows, err := ad.db.Raw(getArtifactPreservationByIDQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("preservationRows.cd.db.Raw err: %s", err)
	}
	defer preservationRows.Close()
	for preservationRows.Next() {
		var (
			id                    int64
			artifactId            int64
			childPreservationName string
			parentPreservationID  sql.NullInt64
		)
		err := preservationRows.Scan(&id, &artifactId, &childPreservationName, &parentPreservationID)
		if err != nil {
			return fmt.Errorf("preservationRows.Scan err: %s", err)
		}
		childPreservation := make([]*ArtifactPreservation, 0)
		childPreservation, err = ad.getArtifactChildPreservation(artifact.ID, id)
		element := &ArtifactPreservation{
			ID:                id,
			ArtifactID:        artifactId,
			Name:              childPreservationName,
			ParentID:          parentPreservationID.Int64,
			ChildPreservation: childPreservation,
		}
		if len(element.ChildPreservation) > 0 {
			artifact.Preservation = append(artifact.Preservation, element)
		}
	}
	return nil
}

func (ad *ArtifactData) initArtifactMaterials(artifact *ArtifactMaster) error {
	artifact.Materials = make([]*ArtifactMaterial, 0)
	materialsRows, err := ad.db.Raw(getArtifactMaterialsByIDQuery, artifact.ID).Rows()
	if err != nil {
		return fmt.Errorf("materialsRows.cd.db.Raw err: %s", err)
	}
	defer materialsRows.Close()
	for materialsRows.Next() {
		var (
			id               int64
			artifactId       int64
			quantity         int64
			composition      sql.NullInt64
			materialType     string
			parentMaterialID sql.NullInt64
		)
		err := materialsRows.Scan(
			&id,
			&artifactId,
			&quantity,
			&composition,
			&materialType,
			&parentMaterialID,
		)
		if err != nil {
			return fmt.Errorf("materialsRows.Scan err: %s", err)
		}
		childMaterials := make([]*ArtifactMaterial, 0)
		childMaterials, err = ad.getArtifactChildMaterials(artifact.ID, id)
		element := &ArtifactMaterial{
			ID:             id,
			ArtifactID:     artifactId,
			Quantity:       quantity,
			Composition:    composition.Int64,
			MaterialType:   materialType,
			ParentID:       parentMaterialID.Int64,
			ChildMaterials: childMaterials,
		}
		if len(element.ChildMaterials) > 0 {
			artifact.Materials = append(artifact.Materials, element)
		}
	}
	return nil
}

func (ad *ArtifactData) getArtifactChildMaterials(artifactID, parentID int64) ([]*ArtifactMaterial, error) {
	childMaterials := make([]*ArtifactMaterial, 0)
	childMaterialsRows, err := ad.db.Raw(getArtifactChildMaterialsQuery, artifactID, parentID).Rows()
	if err != nil {
		return nil, fmt.Errorf("childMaterialsRows.cd.db.Raw.err: %s", err)
	}
	defer childMaterialsRows.Close()
	for childMaterialsRows.Next() {
		var (
			id               int64
			artifactId       int64
			quantity         int64
			composition      sql.NullInt64
			materialType     string
			parentMaterialID sql.NullInt64
		)
		err := childMaterialsRows.Scan(
			&id,
			&artifactId,
			&quantity,
			&composition,
			&materialType,
			&parentMaterialID,
		)
		if err != nil {
			return nil, fmt.Errorf("childMaterialsRows.Scan err: %s", err)
		}

		childMaterial := &ArtifactMaterial{
			ID:           id,
			ArtifactID:   artifactID,
			Quantity:     quantity,
			Composition:  composition.Int64,
			MaterialType: materialType,
			ParentID:     parentMaterialID.Int64,
		}
		childMaterials = append(childMaterials, childMaterial)
	}
	return childMaterials, nil
}

func (ad *ArtifactData) getArtifactChildObjectGroup(artifactID, parentID int64) ([]*ArtifactObjectGroup, error) {
	childObjectGroup := make([]*ArtifactObjectGroup, 0)
	childObjectGroupRows, err := ad.db.Raw(getArtifactChildObjectGroupQuery, artifactID, parentID).Rows()
	if err != nil {
		return nil, fmt.Errorf("childObjectGroupRows.cd.db.Raw.err: %s", err)
	}
	defer childObjectGroupRows.Close()
	for childObjectGroupRows.Next() {
		var (
			id              int64
			childArtifactID int64
			objectGroupName string
			parentElementID sql.NullInt64
		)
		err := childObjectGroupRows.Scan(&id, &childArtifactID, &objectGroupName, &parentElementID)
		if err != nil {
			return nil, fmt.Errorf("childObjectGroupRows.Scan err: %s", err)
		}

		childElement := &ArtifactObjectGroup{
			ID:         id,
			ArtifactID: artifactID,
			Name:       objectGroupName,
			ParentID:   parentElementID.Int64,
		}
		childObjectGroup = append(childObjectGroup, childElement)
	}
	return childObjectGroup, nil
}

func (ad *ArtifactData) getArtifactChildPreservation(artifactID, parentID int64) ([]*ArtifactPreservation, error) {
	childPreservation := make([]*ArtifactPreservation, 0)
	childPreservationRows, err := ad.db.Raw(getArtifactChildPreservationQuery, artifactID, parentID).Rows()
	if err != nil {
		return nil, fmt.Errorf("childPreservationRows.cd.db.Raw.err: %s", err)
	}
	defer childPreservationRows.Close()
	for childPreservationRows.Next() {
		var (
			id                   int64
			childArtifactID      int64
			preservationName     string
			parentPreservationID sql.NullInt64
		)
		err := childPreservationRows.Scan(&id, &childArtifactID, &preservationName, &parentPreservationID)
		if err != nil {
			return nil, fmt.Errorf("childPreservationRows.Scan err: %s", err)
		}

		childElement := &ArtifactPreservation{
			ID:         id,
			ArtifactID: artifactID,
			Name:       preservationName,
			ParentID:   parentPreservationID.Int64,
		}
		childPreservation = append(childPreservation, childElement)
	}
	return childPreservation, nil
}
