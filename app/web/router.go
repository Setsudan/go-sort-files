package web

/*
 Here will be the gin router and server.
*/

import (
	"fmt"
	"net/http"

	"main/go-sort-files/app/web/routes"

	"github.com/gin-gonic/gin"
)

/**
 func router will be the main router for the web server. It is the function
 that will be called on the main.go file.

 routes :
 	- `GET` `/` - serves the main web page with an image and an input field
	- `POST` `/dispatch` - handles the input field form submission and moves the image to a subfolder with the category name
	- `GET` `/images/get` - give a list of all the images in the original folder
	- `GET` `/images/categories` - give a list of all the categories
	- `GET` `/images/categories/:category` - give a list of all the images in a category
	- `GET` `/images/categories/:category/:image` - serves an image from a category
**/

// each route will have a function on their own file
// so for now we comment the lines where the functions are called

func Router() {
	// create a new router
	router := gin.Default()

	// serve the main page
	router.GET("/", index)

	// handle the form submission
	router.POST("/dispatch", routes.Dispatch)

	// get all the images
	router.GET("/images/get", routes.GetImages)

	// serve all the images from the original folder
	router.StaticFS("/images/original", http.Dir("images/original"))

	// get all the categories
	router.GET("/categories/get", routes.GetCategories)

	// start the server
	fmt.Println("Server started on port 3000")
	router.Run(":3000")
}

// index serves the main page
func index(c *gin.Context) {
	// the html file is supposed to be in /templates/index.html in the project folder
	file := "templates/index.html"
	c.File(file)

	// serve the main page
	c.HTML(http.StatusOK, file, gin.H{})
}
