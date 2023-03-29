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
	"net/http"

	"github.com/gin-gonic/gin"

	"main/go-sort-files/app"
)

// dispatch handles the form submission
func dispatch(c *gin.Context) {
	// get the image file
	image, err := c.FormFile("image")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	// get the category name
	category := c.PostForm("category")

	// create the category folder
	app.CreateCategory(category)

	// move the image to the category folder
	// the image will be moved to the original folder
	// and then moved to the category folder
	// so we can use the original folder to delete the images
	// if we want to
	app.MoveImage(category, image.Filename)

	// redirect to the main page
	c.Redirect(http.StatusMovedPermanently, "/")
}
