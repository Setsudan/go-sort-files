package routes

/* This is the route :
- `POST` `/dispatch` - handles the input field form submission and moves the image to a subfolder with the category name


The form will have the following fields :
- `image` - the image file
- `category` - the category name

 all the functions are ready in the app package "app/functions.go"
*/

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"main/go-sort-files/app"
)

// dispatch handles the form submission
func Dispatch(c *gin.Context) {
	// get the category name
	category := c.PostForm("category")
	// get the image path (string path to the file)
	imagePath := c.PostForm("image")

	// check if the category is not empty
	if category == "" {
		// return a json response
		c.JSON(200, gin.H{
			"message": "Category is empty",
		})
		return
	}

	// if category is delete then delete the image
	if category == "delete" {
		// delete the image from the original folder
		app.DeleteImage(imagePath)

		// return a json response
		c.JSON(200, gin.H{
			"message": "Image deleted",
		})

		return
	}

	fmt.Println("creating category: " + category + "")
	// create the category folder if it doesn't exist
	app.CreateCategory(category)

	fmt.Println("moving image to category folder")
	// move the image(we have it's current path) to the category folder
	app.MoveImage(category, imagePath)

	fmt.Println("image moved to category folder")
	// return a json response
	c.JSON(200, gin.H{
		"message": "Image moved",
		"wrote":   "to " + category + " folder",
	})

}
