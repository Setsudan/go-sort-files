/*
 This will be the route :
- `GET` `/categories/get` - give a list of all the categories in the all_categories.json file

We will use the `app.GetCategories()` function to get the categories.

The function will return a `[]string` with the categories names.

We will use the `c.JSON()` function to send the response.

The response will be a json with the following format :
```json
[
	"category1",
	"category2",
	"category3"
]
```
*/

package routes

import (
	"github.com/gin-gonic/gin"
	"main/go-sort-files/app"
)

// getCategories will return a json with the categories names
func GetCategories(c *gin.Context) {
	// get the categories
	categories := app.GetCategories()
	// send the response
	c.JSON(200, categories)
}
