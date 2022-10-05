package utils

import "content/src/models"

func Contains(data map[models.MovieLight][]*models.GenreLight, element uint8) bool {
	for _, genre := range data {
		for _, genre := range genre {
			if element == genre.GenreID {
				return true
			}
		}
	}
	return false
}
