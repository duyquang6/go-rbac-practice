package repositories

import "todolist-facebook-chatbot/models"

// CachedRepository provides access to a Redis caching database.
type TaskRepository interface {
	Create(ormer *db.DB, task *models.Task) error
}