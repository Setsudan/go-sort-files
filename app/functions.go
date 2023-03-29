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
 * createCategory will create a folder for a category and add the category to the all_categories.json file.
 * - category - the name of the category
 */
func CreateCategory(category string) {
	if _, err := os.Stat("images/categories/" + category); os.IsNotExist(err) {
		fmt.Println("Creating category folder")
		os.Mkdir("images/categories/"+category, 0755)
	}

	file, err := os.OpenFile("images/categories/all_categories.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	file.WriteString("\"" + category + "\",\n")

	file.Close()

}

/**
 * moveImage will move an image from the original folder to a category folder.
 * - category - the name of the category
 * - image - the path of the image
 */
func MoveImage(category string, image string) {
	if _, err := os.Stat("images/categories/" + category); os.IsNotExist(err) {
		fmt.Println("Creating category folder")
		os.Mkdir("images/categories/"+category, 0755)
	}

	os.Rename(image, "images/categories/"+category+"/"+image)
}

/**
 * deleteImage will delete an image from the original folder.
 * - image - the path of the image
 */
func deleteImage(image string) {
	os.Remove(image)
}
