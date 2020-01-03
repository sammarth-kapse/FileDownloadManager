package repository

import (
	"github.com/google/uuid"
	"testing"
)

func TestGetDownloadInformationByID(t *testing.T) {

	nonInsertedID := uuid.New().String()

	insertedID := uuid.New().String()
	downloadInformation := new(DownloadInformation)
	InsertIntoDownloadCollection(insertedID, downloadInformation)

	tests := []struct {
		inputID          string
		expectedPresence bool
	}{
		{insertedID, true},
		{nonInsertedID, false},
	}

	for _, test := range tests {
		if _, isPresent := GetDownloadInformationByID(test.inputID); isPresent != test.expectedPresence {
			t.Error("Input ID:", test.inputID, ", Expected Presence:", test.expectedPresence, ", Presence Obtained:", isPresent)
		}
	}
}

func TestGetFileNameFromURL(t *testing.T) {

	tests := []struct {
		inputURL, expectedFileName string
	}{
		{"http://somePath/fileName.jpg", "fileName.jpg"},
		{"http://somePath/fileName", "fileName"},
		{"http://somePath/someMorePath/fileName", "fileName"},
	}

	for _, test := range tests {
		if fileName := getFileNameFromURL(test.inputURL); fileName != test.expectedFileName {
			t.Error("Input URL:", test.inputURL, ", Expected Filename:", test.expectedFileName, ", Filename Obtained:", fileName)
		}
	}
}
