package data

import (
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
		Name: "parent element",
		Children: []ArtifactElement{
			{
				Name: "child 1",
				Children: []ArtifactElement{
					{
						Name:     "sub child 1",
						Children: nil,
					},
				},
			},
			{
				Name:     "child 2",
				Children: nil,
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
	id, err := artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	idSecond, errSecond := artifactData.Add(testArtifact)
	assert.NoError(t, errSecond, fmt.Sprintf("got an error when tried to add artifact, err:%v", errSecond))
	assert.True(t, idSecond > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", idSecond))

	cleanTestDB(conn)
}

func TestReadAll(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	id, err = artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	artifacts, err := artifactData.ReadAll()
	assert.NoError(t, err)
	assert.NotEmpty(t, artifacts)
	assert.Len(t, artifacts, 2)

	assert.Equal(t, 1, artifacts[0].ID)
	assert.Equal(t, testArtifact.Creator, artifacts[0].Creator)
	assert.Equal(t, testArtifact.ArtifactMeasurement, artifacts[0].ArtifactMeasurement)
	assert.Equal(t, testArtifact.TransferredBy, artifacts[0].TransferredBy)



	assert.Equal(t, 2, artifacts[1].ID)
	assert.Equal(t, testArtifact.Creator, artifacts[1].Creator)
	assert.Equal(t, testArtifact.ArtifactMeasurement, artifacts[1].ArtifactMeasurement)
	assert.Equal(t, testArtifact.TransferredBy, artifacts[1].TransferredBy)

	cleanTestDB(conn)
}

func TestRead(t *testing.T) {
	conn, err := prepareTestDB()
	defer cleanTestDB(conn)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	artifact, err := artifactData.Read(id)
	assert.NoError(t, err)
	assert.Equal(t, id, artifact.ID)
	assert.Equal(t, testArtifact.Creator, artifact.Creator)
	assert.Equal(t, testArtifact.ArtifactMeasurement, artifact.ArtifactMeasurement)
	assert.Equal(t, testArtifact.TransferredBy, artifact.TransferredBy)
	//	assert.Equal(t, testArtifact.ExcavationDate, artifact.ExcavationDate) incorrect prefix
}

func TestUpdated(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	newArtifact := ArtifactMaster{
		ID:             0,
		Creator:        "Artsiom",
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

	cleanTestDB(conn)
}


// TODO: DELETE test fail
func TestDelete(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(testArtifact)
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
	id, err := artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	actualID, err := artifactData.insertArtifactElement(artifactData.db, id, getTestArtifactElement(1))
	assert.NoError(t, err, fmt.Sprintf("got error %+v from InsertArtifactElement when tried to insert %v", err, getTestArtifactElement(id)))
	assert.NotEqual(t, 0, actualID, "incorrect actualID, should be positive, but have got 0")

	cleanTestDB(conn)
}

func TestReadArtifactElement(t *testing.T) {
	conn, err := prepareTestDB()
	defer cleanTestDB(conn)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)
	id, err := artifactData.Add(testArtifact)
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to add artifact, err:%v", err))
	assert.True(t, id > 0, fmt.Sprintf("id less then zero, but should be higher, id: %d", id))

	actualID, err := artifactData.insertArtifactElement(artifactData.db, id, getTestArtifactElement(1))
	assert.NoError(t, err, fmt.Sprintf("got error %+v from InsertArtifactElement when tried to insert %v", err, getTestArtifactElement(1)))
	assert.NotEqual(t, 0, actualID, "incorrect actualID, should be positive, but have got 0")

	actualArtifactElements, err := artifactData.readArtifactElements(id)
	assert.NoError(t, err, fmt.Sprintf("got error %+v from readArtifactElement, with the following artifactMasterID %d", err, id))
	assert.NotEmpty(t, actualArtifactElements, "artifact element should not be nil")
}

func TestTableName(t *testing.T) {
	element := ArtifactElement{}
	actualTableName := element.TableName()
	expected := "artifact_element"
	assert.Equal(t, actualTableName, expected)
}

func TestNewArtifactData(t *testing.T) {
	actualNilData := NewArtifactData(nil)
	expectedNilData := &ArtifactData{nil}

	assert.Equal(t, actualNilData, expectedNilData)

	conn, err := prepareTestDB()
	defer cleanTestDB(conn)

	actualData := NewArtifactData(conn)
	expectedData := &ArtifactData{conn}
	assert.Equal(t, actualData, expectedData)
	assert.NoError(t, err)
}

func TestInsertTransferredByLUTIfNotExists(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)

	id, err := artifactData.insertTransferredByLUTIfNotExists(artifactData.db,"test")
	assert.NoError(t, err, "can't insert into the transferredByLUT table")
	assert.NotEqual(t, 0, id, "id should not be 0")
	assert.True(t, id > 0, "id should not be less than 1")

	idSame, errSame := artifactData.insertTransferredByLUTIfNotExists(artifactData.db, "test")
	assert.NoError(t, errSame, "can't insert the same data into the transferredByLUT table")
	assert.NotEqual(t, 0, idSame, "id should not be 0")
	assert.True(t, idSame > 0, "id should not be less than 1")

	cleanTestDB(conn)
}



func TestInsertArtifactStyleLUTIfNotExists(t *testing.T) {
	conn, err := prepareTestDB()
	assert.NoError(t, err, fmt.Sprintf("got an error when tried to prepare db, err:%v", err))

	artifactData := NewArtifactData(conn)

	id, err := artifactData.insertArtifactStyleLUTIfNotExists("test")
	assert.NoError(t, err, "can't insert into the artifact_style_lut table")
	assert.NotEqual(t, 0, id, "id should not be 0")
	assert.True(t, id > 0, "id should not be less than 1")

	idSame, errSame := artifactData.insertArtifactStyleLUTIfNotExists("test")
	assert.NoError(t, errSame, "can't insert the same data into the artifact_style_lut table")
	assert.NotEqual(t, 0, idSame, "id should not be 0")
	assert.True(t, idSame > 0, "id should not be less than 1")

	cleanTestDB(conn)
}
