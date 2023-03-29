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
	// get the image path
	image, _ := c.FormFile("image")

	// create the category folder if it doesn't exist
	app.CreateCategory(category)

	// move the image to the category folder
	err := c.SaveUploadedFile(image, "images/categories/"+category+"/"+image.Filename)
	if err != nil {
		fmt.Println(err)
	}

}
