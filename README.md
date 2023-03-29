# Image Dispatcher Web App

The Image Dispatcher Web App is a Go web app that allows users to categorize and move images from one folder to another. The app uses the Gin web framework to serve a web page with an image and an input field that allows the user to categorize the image.

## Features
The Image Dispatcher Web App includes the following features:

- Displays one image at a time
- Input field with autocomplete to create a category for the image
- Creates a subfolder with the category name and moves the image to that folder
- Supports multiple image file types, including PNG, JPG, and GIF

## Architecture
The Image Dispatcher Web App is built using the following components:

- Gin web framework for serving web pages and handling HTTP requests
- fsnotify package for watching the file system for changes
- Goroutine to handle multiple requests simultaneously
- os.Rename function to move images to subfolders

## Endpoints
The Image Dispatcher Web App includes the following endpoints:

- `GET` `/` - serves the main web page with an image and an input field
- `POST` `/dispatch` - handles the input field form submission and moves the image to a subfolder with the category name
