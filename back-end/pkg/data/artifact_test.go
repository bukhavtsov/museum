// +build integration

package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/bukhavtsov/museum/back-end/db"

	"github.com/stretchr/testify/require"
)

var (
	testHost     = "localhost"
	testPort     = "5432"
	testUser     = "postgres"
	testDBname   = "postgres"
	testPassword = "postgres"
	testSSLMode  = "disable"
)

// FIXME: need to initialize table after run Test
func TestAdd(t *testing.T) {
	conn := db.GetConnection(testHost, testPort, testUser, testDBname, testPassword, testSSLMode)
	defer conn.Close()

	// prepare testArtifactData
	artifactData := NewArtifactData(conn)
	actualArtifacts, err := artifactData.ReadAll()
	if err != nil {
		t.Fatal(err, "Read all actualArtifacts method doesn't work properly")
	}

	// convert testArtifactData to bytes
	actualArtifactsJsonBytes, err := json.Marshal(actualArtifacts)
	if err != nil {
		t.Fatal(err, "Got an error when try to marshal actualArtifacts")
	}

	// get absolute path to the expected json results
	expectedJsonRelativePath := "../../tests/json/get_artifacts.json"
	expectedJsonAbsPath, err := filepath.Abs(expectedJsonRelativePath)
	require.Nil(t, err, "Can't find absolute path for file %s. Err:", expectedJsonRelativePath, err)

	// get string with expected json results via absolute path
	expectedArtifactsJson, err := os.Open(expectedJsonAbsPath)
	require.Nil(t, err, "Can't open selected path: %v", err)
	expectedArtifactsJsonBytes, err := ioutil.ReadAll(expectedArtifactsJson)
	require.Nil(t, err, "Can't read info from opened path: %v", err)

	// compare actual artifact results with expected
	isEqual, err := isJSONBytesEqual(expectedArtifactsJsonBytes, actualArtifactsJsonBytes)
	require.Nil(t, err, "Got an error: %v", err)
	require.True(t, isEqual, "Jsons are not equals!")
}

// isJSONBytesEqual compares the JSON between two byte slices
func isJSONBytesEqual(a, b []byte) (bool, error) {
	var j, j2 interface{}
	if err := json.Unmarshal(a, &j); err != nil {
		return false, err
	}
	if err := json.Unmarshal(b, &j2); err != nil {
		return false, err
	}
	return reflect.DeepEqual(j2, j), nil
}
