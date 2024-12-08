package openuem_utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func DownloadFile(url, filepath string, expectedHash string) error {

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return fmt.Errorf("file not found")
	}

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	// Get hash
	hash, err := GetSHA256Sum(filepath)
	if err != nil {
		return err
	}

	// Check hash
	if fmt.Sprintf("%x", hash) != expectedHash {
		return fmt.Errorf("checksum doesn't match")
	}

	return nil
}

func QueryReleasesEndpoint(url string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Second * 8,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "openuem-console")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}

	return body, nil
}
