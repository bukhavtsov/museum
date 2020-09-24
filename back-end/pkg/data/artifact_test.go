package data

import (
	"testing"
	"time"

	"github.com/bukhavtsov/museum/back-end/db"
)

var (
	testHost     = "localhost"
	testPort     = "5432"
	testUser     = "postgres"
	testDBname   = "postgres"
	testPassword = "postgres"
	testSSLMode  = "disable"
)

func getTestArtifactPreservation() map[string][]string {
	return map[string][]string{
		"main side is good":                {"sub side test1 have bad quality", "sub side test2 with good quality"},
		"sub side test1 have bad quality":  {"sub side test3 with good quality"},
		"sub side test3 with good quality": {"sub side test4 is bad", "sub side test5 is normal", "sub side test6 it's ok"},
	}
}

func getTestArtifactElements() map[string][]string {
	return map[string][]string{
		"element 1": {"sub element 1 of element 1", "sub element 2 of element 2", "sub element 3 of element 3"},
		"element 2": {"sub element 2 of element 1", "sub element 2 of element 2", "sub element 3 of element 3"},
	}
}
func getTestArtifactObjectGroup() map[string][]string {
	return map[string][]string{}
}

func TestAdd(t *testing.T) {
	conn := db.GetConnection(testHost, testPort, testUser, testDBname, testPassword, testSSLMode)
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
