package nextcloud

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/ElHefe3/resume-api/internal/config"
)

type MultiStatus struct {
	Responses []Response `xml:"response"`
}

type Response struct {
	Href string `xml:"href"`
}

func RetrieveFilesDirectories() ([]string, error) {
	url := config.Cfg.NextcloudURL
	username := config.Cfg.NextcloudUsername
	password := config.Cfg.NextcloudPassword
    filePath := config.Cfg.NextcloudFilesDirectory

	client := &http.Client{}
	req, err := http.NewRequest("PROPFIND", url+"/remote.php/dav/files/"+username+"/"+filePath, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(username, password)
	req.Header.Set("Depth", "1")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 207 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var multistatus MultiStatus
	if err := xml.Unmarshal(body, &multistatus); err != nil {
		return nil, err
	}

	var files []string
	for _, response := range multistatus.Responses {
		files = append(files, response.Href)
	}

	return files, nil
}

func RetrieveFile(fileName string) ([]byte, string, error) {
    username := config.Cfg.NextcloudUsername
	password := config.Cfg.NextcloudPassword
    filePath := config.Cfg.NextcloudFilesDirectory
	url := config.Cfg.NextcloudURL +"/remote.php/dav/files/"+username+"/"+filePath+"/"+fileName

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	contentType := resp.Header.Get("Content-Type")

	return body, contentType, nil
}

