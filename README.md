# Image Dispatcher Web App

The Image Dispatcher Web App is a Go web app that allows users to categorize and move images from one folder to another. The app uses the Gin web framework to serve a web page with an image and an input field that allows the user to categorize the image.

## Features
The Image Dispatcher Web App includes the following features:

- Displays one image at a time
- Input field to create a category for the image
- Creates a subfolder with the category name and moves the image to that folder
- Supports multiple image file types, including PNG, JPG, and GIF
- Support Videos


## Endpoints
The Image Dispatcher Web App includes the following endpoints:

- `GET` `/` - serves the main web page with an image and an input field
- `POST` `/dispatch` - handles the input field form submission and moves the image to a subfolder with the category name
