package app

/*
 Here we will have the functions that handles the image processing.
*/

import (
	"fmt"
	"os"
)

/**
 * initApp will create the folders that will be used by the app.
 * - images
 * - images/categories
 * - images/original
 */
func initApp() {
	if _, err := os.Stat("images"); os.IsNotExist(err) {
		fmt.Println("Creating images folder")
		os.Mkdir("images", 0755)
	}

	if _, err := os.Stat("images/categories"); os.IsNotExist(err) {
		fmt.Println("Creating categories folder")
		os.Mkdir("images/categories", 0755)
	}

	if _, err := os.Stat("images/original"); os.IsNotExist(err) {
		fmt.Println("Creating original folder")
		os.Mkdir("images/original", 0755)
	}
}

/**
 * createCategory will create a folder for a category
 * - category - the name of the category
 */
func CreateCategory(category string) {
	if _, err := os.Stat("images/categories/" + category); os.IsNotExist(err) {
		fmt.Println("Creating category folder")
		os.Mkdir("images/categories/"+category, 0755)
	}

}

/**
 * moveImage will move an image from the original folder to a category folder.
 * - category - the name of the category
 * - image - the path of the image
 */
func MoveImage(category string, imagePath string) {
	fmt.Println("Image received till functions.go")
	fmt.Println("Image from: " + imagePath)
	fmt.Println("Will go to: " + category)

	// get the image path
	pathToImage := imagePath
	fmt.Println("Okay, the image path is: " + pathToImage)

	// catory folder path
	pathToCategory := "images/categories/" + category
	fmt.Println("Okay, the category path is: " + pathToCategory)

	// get the image name
	imageName := pathToImage[16:]
	fmt.Println("Okay, the image name is: " + imageName)

	fmt.Println("Image name: " + imageName)
	fmt.Println("Path to category: " + pathToCategory)

	if _, err := os.Stat(pathToCategory + "/" + imageName); os.IsNotExist(err) {
		err := os.Rename(pathToImage, pathToCategory+"/"+imageName)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("File already exists")
	}

}

/**
 * deleteImage will delete an image from the original folder.
 * - image - the path of the image
 */
func DeleteImage(image string) {
	os.Remove(image)
}

/**
 * getImages will return a list of all the images(path to image) in the original folder.
 */
func GetImages() []string {
	var images []string

	files, err := os.ReadDir("images/original")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		images = append(images, "images/original/"+file.Name())
	}

	return images
}

/**
 * getCategories will return a list of all the categories in the categories folder.
 */
func GetCategories() []string {
	var categories []string

	files, err := os.ReadDir("images/categories")
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		categories = append(categories, file.Name())
	}

	return categories
}
