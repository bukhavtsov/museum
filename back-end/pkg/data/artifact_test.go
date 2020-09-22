package data

import (
	"testing"
	"time"

	"github.com/bukhavtsov/museum/back-end/db"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	dbname   = "postgres"
	password = "postgres"
	sslmode  = "disable"
)

func getTestArtifactPreservation() map[string][]string {
	return map[string][]string{}
}
func getTestArtifactElements() map[string][]string {
	return map[string][]string{}
}
func getTestArtifactObjectGroup() map[string][]string {
	return map[string][]string{}
}

func TestAdd(t *testing.T) {
	conn := db.GetConnection(host, port, user, dbname, password, sslmode)
	defer conn.Close()
	//1. create artifact with insertion of all fields
	artifactMaster := new(ArtifactMaster)
	artifactMaster.ID = 0
	artifactMaster.ExcavationDate = time.Now().String()
	artifactMaster.Creator = "Test Creator"
	artifactMaster.ArtifactStyle = "Belarussian style"
	artifactMaster.TransferredBy = "Good Boy"
	artifactMaster.ArtifactMeasurement = &ArtifactMeasurement{Width: 10, Height: 20, Length: 30}
	artifactMaster.Preservation = getTestArtifactPreservation()
	artifactMaster.Elements = getTestArtifactElements()
	artifactMaster.ObjectGroup = getTestArtifactObjectGroup()
}
