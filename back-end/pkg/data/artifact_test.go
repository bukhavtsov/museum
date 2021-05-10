package data

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/bukhavtsov/museum/back-end/db"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")

	testArtifact = ArtifactMaster{
		ID:             0,
		Creator:        "TestExample",
		ArtifactStyle:  "1323123123123",
		ExcavationDate: "1979-02-17",
		TransferredBy:  "qeqwe",
		ArtifactMeasurement: &ArtifactMeasurement{
			Height: 10,
			Width:  20,
			Length: 30,
		},
	}
)

func getTestArtifactElement(artifactMasterID int) ArtifactElement {
	return ArtifactElement{
		ArtifactID: artifactMasterID,
		Name:       "parent element",
		Children: []ArtifactElement{
			{
				ArtifactID: artifactMasterID,
				Name:       "child 1",
				Children: []ArtifactElement{
					{
						ArtifactID: artifactMasterID,
						Name:       "sub child 1",
						Children:   nil,
					},
				},
			},
			{
				ArtifactID: artifactMasterID,
				Name:       "child 2",
				Children:   nil,
			},
		},
	}
}

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "1001"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "postgres"
	}
	if password == "" {
		password = "postgres"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func prepareTestDB() (*gorm.DB, error) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	path := filepath.Join("../../../db/scripts/init-tables.sql")

	sqlBytes, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		return nil, fmt.Errorf("got an error when read script, err:%v", ioErr)
	}
	sql := string(sqlBytes)
	conn.Exec(sql)
	return conn, nil
}

func cleanTestDB(conn *gorm.DB) {
	defer conn.Close()
	conn.Exec("DROP SCHEMA public CASCADE;")
	conn.Exec("CREATE SCHEMA public;")
	fmt.Println("clean db")
}

func TestCreate(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	cleanTestDB(conn)
}

func TestReadAll(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	artifacts, err := artifactData.ReadAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, artifacts)
	assert.Len(t, artifacts, 1)

	cleanTestDB(conn)
}

func TestRead(t *testing.T) {
	conn, err := prepareTestDB()
	defer cleanTestDB(conn)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	artifact, err := artifactData.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, id, artifact.ID)
	assert.Equal(t, testArtifact.Creator, artifact.Creator)
	assert.Equal(t, testArtifact.ArtifactMeasurement, artifact.ArtifactMeasurement)
	assert.Equal(t, testArtifact.TransferredBy, artifact.TransferredBy)
	//	assert.Equal(t, testArtifact.ExcavationDate, artifact.ExcavationDate) incorrect prefix
	//	assert.Equal(t, testArtifact.ArtifactStyle, artifact.ArtifactStyle) not implemented
}

func TestUpdated(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	newArtifact := ArtifactMaster{
		ID:             0,
		Creator:        "Artsiom",
		ArtifactStyle:  "OldSchool",
		ExcavationDate: "1972-02-17",
		TransferredBy:  "Some One",
		ArtifactMeasurement: &ArtifactMeasurement{
			Height: 20,
			Width:  40,
			Length: 50,
		},
	}
	err = artifactData.Update(id, &newArtifact)
	assert.NoError(t, err)

	artifact, err := artifactData.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, id, artifact.ID)
	assert.Equal(t, newArtifact.Creator, artifact.Creator)
	assert.Equal(t, newArtifact.ArtifactMeasurement, artifact.ArtifactMeasurement)
	assert.Equal(t, newArtifact.TransferredBy, artifact.TransferredBy)
	//	assert.Equal(t, testArtifact.ExcavationDate, artifact.ExcavationDate) incorrect prefix
	//	assert.Equal(t, testArtifact.ArtifactStyle, artifact.ArtifactStyle) not implemented

	cleanTestDB(conn)
}

func TestDelete(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	err = artifactData.Delete(id)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to delete artifact, err:%v", err))

	artifacts, err := artifactData.ReadAll()
	assert.NoError(t, err)
	assert.Empty(t, artifacts)

	cleanTestDB(conn)
}

func TestInsertArtifactElement(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	actualID, err := artifactData.InsertArtifactElement(getTestArtifactElement(id))
	assert.NoError(t, err, fmt.Sprintf("got error %+v from InsertArtifactElement when tried to insert %v", err, getTestArtifactElement(id)))
	assert.NotEqual(t, -1, actualID, "incorrect actualID, should be positive, but have got -1")

	cleanTestDB(conn)
}

func TestReadArtifactElement(t *testing.T) {
	conn, err := prepareTestDB()
	defer cleanTestDB(conn)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(&testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))
	testArtifactElement := getTestArtifactElement(id)
	actualID, err := artifactData.InsertArtifactElement(testArtifactElement)
	assert.NoError(t, err, fmt.Sprintf("got error %+v from InsertArtifactElement when tried to insert %v", err, testArtifactElement))
	assert.NotEqual(t, -1, actualID, "incorrect actualID, should be positive, but have got -1")

	actualArtifactElements, err := artifactData.readArtifactElements(id)
	assert.NoError(t, err, fmt.Sprintf("got error %+v from readArtifactElement, with the following artifactMasterID %d", err, id))
	assert.NotEmpty(t, actualArtifactElements, "artifact element should not be nil")
	assert.Equal(t, testArtifactElement.Name, actualArtifactElements[0].Name)
	assert.Equal(t, testArtifactElement.ArtifactID, actualArtifactElements[0].ArtifactID)
	assert.NotEmpty(t, actualArtifactElements[0].Children)

	assert.Equal(t, testArtifactElement.Children[0].Name, actualArtifactElements[0].Children[0].Name)
	assert.Equal(t, testArtifactElement.Children[0].ArtifactID, actualArtifactElements[0].Children[0].ArtifactID)
	assert.NotEmpty(t, actualArtifactElements[0].Children[0].Children)

	assert.Equal(t, testArtifactElement.Children[0].Children[0].Name, actualArtifactElements[0].Children[0].Children[0].Name)
	assert.Equal(t, testArtifactElement.Children[0].Children[0].ArtifactID, actualArtifactElements[0].Children[0].Children[0].ArtifactID)
	assert.Empty(t, actualArtifactElements[0].Children[0].Children[0].Children)

	assert.Equal(t, testArtifactElement.Children[1].Name, actualArtifactElements[0].Children[1].Name)
	assert.Equal(t, testArtifactElement.Children[1].ArtifactID, actualArtifactElements[0].Children[1].ArtifactID)
	assert.Empty(t, actualArtifactElements[0].Children[1].Children)
}

func TestFilterChildren(t *testing.T) {

	testArtifactElement := []ArtifactElement{
		{
			ID:         1,
			ArtifactID: 1,
			Name:       "parent element",
		},
		{
			ID:         2,
			ArtifactID: 1,
			Name:       "child 1",
			ParentID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
		{
			ID:         3,
			ArtifactID: 1,
			Name:       "sub child 1",
			ParentID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
		},
		{
			ID:         4,
			ArtifactID: 1,
			Name:       "child 2",
			ParentID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
	}

	other, children, err := filterChildren(testArtifactElement, 1)
	assert.NoErrorf(t, err, "error from filterChildren %w", err)
	assert.Len(t, children, 2)
	assert.Len(t, other, 2)
	assert.Equal(t, 1, other[0].ID)
	assert.Equal(t, 1, other[0].ArtifactID)
	assert.Equal(t, "parent element", other[0].Name)
	assert.Equal(t, sql.NullInt64{}, other[0].ParentID)
	assert.Empty(t, other[0].Children)
	assert.Equal(t, 3, other[1].ID)
	assert.Equal(t, 1, other[1].ArtifactID)
	assert.Equal(t, "sub child 1", other[1].Name)
	assert.Equal(t, sql.NullInt64{Int64: 2, Valid: true}, other[1].ParentID)
	assert.Empty(t, other[1].Children)

	assert.Equal(t, 2, children[0].ID)
	assert.Equal(t, 1, children[0].ArtifactID)
	assert.Equal(t, "child 1", children[0].Name)
	assert.Equal(t, sql.NullInt64{Int64: 1, Valid: true}, children[0].ParentID)
	assert.Empty(t, children[0].Children)
	assert.Equal(t, 4, children[1].ID)
	assert.Equal(t, 1, children[1].ArtifactID)
	assert.Equal(t, "child 2", children[1].Name)
	assert.Equal(t, sql.NullInt64{Int64: 1, Valid: true}, children[1].ParentID)
	assert.Empty(t, children[1].Children)
}

func TestGroupArtifactElements(t *testing.T) {
	testArtifactElement := []ArtifactElement{
		{
			ID:         1,
			ArtifactID: 1,
			Name:       "parent element",
		},
		{
			ID:         2,
			ArtifactID: 1,
			Name:       "child 1",
			ParentID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
		{
			ID:         3,
			ArtifactID: 1,
			Name:       "sub child 1",
			ParentID: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
		},
		{
			ID:         4,
			ArtifactID: 1,
			Name:       "child 2",
			ParentID: sql.NullInt64{
				Int64: 1,
				Valid: true,
			},
		},
	}

	other, err := groupArtifactElements(testArtifactElement)
	fmt.Println(err)
	fmt.Println(other)
	intB, _ := json.MarshalIndent(&other, "", "\t")
	fmt.Println(string(intB))
}
