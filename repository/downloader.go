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

func New(request DownloadRequest) *DownloadInformation {

	information := new(DownloadInformation)

	id := uuid.New().String()

	information.ID = id
	information.DownloadType = request.Type
	information.URLs = request.URLs
	information.Files = make(map[string]string)
	information.DirectoryPath = GLOBAL_PATH + information.ID + "/"
	information.createDirectory()

	insertIntoDownloadCollection(information.ID, information)

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

	for _, v := range information.URLs {

		filePath := information.DirectoryPath + getFileName(v)
		err := downloadFile(v, filePath)
		if err != nil {
			information.markDownloadEnd(FAILED)
			log.Fatal(err)
		}
		information.appendDownloadFile(v, filePath)
	}

	information.markDownloadEnd(SUCCESS)
}

func (information DownloadInformation) concurrentDownloader() {

	var wg sync.WaitGroup

	for _, v := range information.URLs {

		wg.Add(1)
		go information.concurrentDownloadHandler(v, &wg)
	}

	wg.Wait()
	information.markDownloadEnd(SUCCESS)
}

func (information DownloadInformation) concurrentDownloadHandler(url string, wg *sync.WaitGroup) {

	filePath := information.DirectoryPath + getFileName(url)

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
