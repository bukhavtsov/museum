package data

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestReadAll(t *testing.T) {
	assert.Equal(t, true, true, "test should works")
}

func TestGetArtifactWithBasicInfo(t *testing.T) {
	sdb, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("faild to create mock DB: %v", err)
	}
	defer sdb.Close()

	db, err := gorm.Open("postgres", sdb)
	if err != nil {
		t.Fatalf("faild to create gorm DB: %v", err)
	}
	defer db.Close()

}
