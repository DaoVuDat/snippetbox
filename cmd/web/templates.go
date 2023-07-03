package main

import "github.com/DaoVuDat/snippetbox/internal/models"

// Define a templateData type to act as the holding structure for
// any dynamic data that we want to pass to our HTML templates.
// At the moment it only contains one field, but we'll add more
type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
