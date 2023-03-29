package web

/*
 Here will be the gin router and server.
*/

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

/**
 func router will be the main router for the web server. It is the function
 that will be called on the main.go file.

 routes :
 	- `GET` `/` - serves the main web page with an image and an input field
	- `POST` `/dispatch` - handles the input field form submission and moves the image to a subfolder with the category name
	- `GET` `/images/categories` - give a list of all the categories
	- `GET` `/images/categories/:category` - give a list of all the images in a category
	- `GET` `/images/categories/:category/:image` - serves an image from a category
**/

// each route will have a function on their own file
// so for now we comment the lines where the functions are called

func router() {
	// create a new router
	router := gin.Default()

	// serve the main page
	router.GET("/", index)

	// handle the form submission
	// router.POST("/dispatch", dispatch)

	// serve the categories
	// router.GET("/images/categories", categories)

	// serve the images from a category
	// router.GET("/images/categories/:category", images)

	// serve an image
	// router.GET("/images/categories/:category/:image", image)

	// start the server
	fmt.Println("Server started on port 3000")
	router.Run(":3000")
}

// index serves the main page
// the html file is on app/web/html/index.html
func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
