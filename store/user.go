package store

import (
	"github.com/usmaarn/blogg/models"
)

func (s *Store) FindAll(users []*models.User, query map[string]any) error {
	rows, err := s.DB.Query("SELECT id, name, email, password, created_at, updated_at FROM users")
	if err != nil {
		return err
	}

	for rows.Next() {
		var user *models.User
		err = rows.Scan(&user.ID, user.Name, user.Email, user.Password, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return err
		}
		users = append(users, user)
	}
	return nil
}
