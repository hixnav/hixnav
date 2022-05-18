package cmd

import "os"

func createDir(dir string) error {
	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		return err
	}
	return nil
}
