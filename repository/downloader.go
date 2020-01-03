package repository

import (
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

type DownloadInformation struct {
	ID            string
	StartTime     time.Time
	EndTime       time.Time
	Status        string
	DownloadType  string
	Files         map[string]string
	DirectoryPath string
	URLs          []string
}

// Creates goDownload directory if it doesn't exist
func init() {

	createDirectoryIfNotExist(GLOBAL_PATH)
}

// Constructor Function for Download Information
func New(request DownloadRequest) *DownloadInformation {

	information := new(DownloadInformation)

	id := uuid.New().String()

	information.ID = id
	information.DownloadType = request.Type
	information.URLs = request.URLs
	information.Files = make(map[string]string)
	information.DirectoryPath = GLOBAL_PATH + information.ID + "/"
	information.createDirectory()

	InsertIntoDownloadCollection(information.ID, information)

	return information
}

func (information DownloadInformation) Download() {

	information.markDownloadStart()

	if strings.Compare(information.DownloadType, SERIAL) == 0 {
		information.serialDownloader()
	} else {
		go information.concurrentDownloader()
	}

}

func (information DownloadInformation) serialDownloader() {

	for _, url := range information.URLs {

		filePath := information.DirectoryPath + getFileNameFromURL(url)
		err := downloadFile(url, filePath)
		if err != nil {
			information.markDownloadEnd(FAILED)
			log.Fatal(err)
		}
		information.appendDownloadFile(url, filePath)
	}

	information.markDownloadEnd(SUCCESS)
}

func (information DownloadInformation) concurrentDownloader() {

	var wg sync.WaitGroup

	for _, url := range information.URLs {

		wg.Add(1)
		go information.concurrentDownloadHandler(url, &wg)
	}

	wg.Wait()
	information.markDownloadEnd(SUCCESS)
}

// Handles Concurrent Download Requests
func (information DownloadInformation) concurrentDownloadHandler(url string, wg *sync.WaitGroup) {

	filePath := information.DirectoryPath + getFileNameFromURL(url)

	defer wg.Done()

	err := downloadFile(url, filePath)
	if err != nil {
		information.markDownloadEnd(FAILED)
		log.Fatal(err)
	}

	information.appendDownloadFile(url, filePath)
}

func downloadFile(url string, filePath string) error {

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
