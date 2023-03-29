/*
 This will be the route :
- `GET` `/images/get` - give a list of all the images in the original folder

We will use the `app.GetImages()` function to get the images.

The function will return a `[]string` with the images names.

We will use the `c.JSON()` function to send the response.

The response will be a json with the following format :
```json
{
	"images": [
		"/path/to/image1.jpg",
		"/path/to/image2.jpg",
		"/path/to/image3.jpg"
	]
}
```
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"main/go-sort-files/app"
)

// getImages will return a json with the images names
func GetImages(c *gin.Context) {
	// get the images
	images := app.GetImages()
	// send the response
	c.JSON(200, gin.H{
		"images": images,
	})
}
