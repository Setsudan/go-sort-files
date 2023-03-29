package main

/*
- Gin web framework for serving web pages and handling HTTP requests
- fsnotify package for watching the file system for changes
- Goroutine to handle multiple requests simultaneously
- os.Rename function to move images to subfolders
- os.MkdirAll function to create subfolders
- os.Remove function to delete images from the original folder
*/

/// Main functions
/*
 Here we will only create a new router and start the server.
 So we only need gin, net/http, and fmt packages.
*/

import (
	"fmt"

	"main/go-sort-files/app/web"
)

func main() {
	fmt.Println("Starting server...")
	web.Router()
}
