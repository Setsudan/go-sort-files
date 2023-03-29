package routes

/*
 - `GET` `/download` - compress the categories folder into a zip and download it
*/

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

// DownloadCategories will compress the categories folder into a zip and download it
func DownloadCategories(c *gin.Context) {
	// get the categories folder
	categoriesFolder := "images/categories"

	// get the categories folder
	originalFolder := "images/original"

	// create a new zip file
	zipFile, err := os.Create("images/categories.zip")
	if err != nil {
		fmt.Println(err)
	}

	// create a new zip writer
	zipWriter := zip.NewWriter(zipFile)

	// defer the closing of the zip file
	defer zipWriter.Close()

	// add the categories folder to the zip file
	addFolderToZip(categoriesFolder, zipWriter)

	// add the original folder to the zip file
	addFolderToZip(originalFolder, zipWriter)

	// close the zip file
	zipWriter.Close()

	// download the zip file
	c.File("images/categories.zip")
}

// addFolderToZip will add a folder to a zip file
func addFolderToZip(folder string, zipWriter *zip.Writer) {

	// get the files in the folder
	files, err := filepath.Glob(folder + "/*")
	if err != nil {
		fmt.Println(err)
	}

	// loop through the files
	for _, file := range files {

		// get the file name
		fileName := strings.Split(file, "/")[2]

		// open the file
		fileToZip, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
		}

		// get the file info
		info, err := fileToZip.Stat()
		if err != nil {
			fmt.Println(err)
		}

		// create a new file header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			fmt.Println(err)
		}

		// set the name of the file
		header.Name = fileName

		// create a new file in the zip file
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			fmt.Println(err)
		}

		// copy the file to the zip file
		_, err = io.Copy(writer, fileToZip)
		if err != nil {
			fmt.Println(err)
		}

		// close the file
		fileToZip.Close()
	}
}
