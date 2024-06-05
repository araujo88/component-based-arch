package userdao

import (
	"fmt"
	"sync"
)

// User represents a user entity.
type User struct {
	ID    int
	Name  string
	Email string
}

// UserDAO handles data access for User entities.
type UserDAO struct {
	users map[int]*User
	lock  sync.RWMutex
}

// NewUserDAO creates a new instance of UserDAO.
func NewUserDAO() *UserDAO {
	return &UserDAO{users: make(map[int]*User)}
}

// GetUser retrieves a user by ID.
func (dao *UserDAO) GetUser(id int) (*User, bool) {
	dao.lock.RLock()
	defer dao.lock.RUnlock()
	user, exists := dao.users[id]
	return user, exists
}

// CreateUser creates a new user.
func (dao *UserDAO) CreateUser(user *User) {
	dao.lock.Lock()
	defer dao.lock.Unlock()
	dao.users[user.ID] = user
}

// UpdateUser updates an existing user.
func (dao *UserDAO) UpdateUser(user *User) error {
	dao.lock.Lock()
	defer dao.lock.Unlock()
	if _, exists := dao.users[user.ID]; exists {
		dao.users[user.ID] = user
		return nil
	}
	return fmt.Errorf("user not found")
}

// DeleteUser deletes a user by ID.
func (dao *UserDAO) DeleteUser(id int) {
	dao.lock.Lock()
	defer dao.lock.Unlock()
	delete(dao.users, id)
}
