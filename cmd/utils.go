package cmd

import (
    "log"
    "net/http"
    "os"
)


func createDir(dir string) error {
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func checkLink(link string, method string) bool {
    req, err := http.NewRequest(method, link, nil)
    if err != nil {
        log.Printf("Failed to create request: %v\n", err)
        return false
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        log.Printf("Failed to send request: %v\n", err)
        return false
    }
    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
        log.Printf("Link %s is accessible\n", link)
        return true
    } else {
        log.Printf("Link %s is not accessible, status code: %d\n", link, resp.StatusCode)
        return false
    }
}
