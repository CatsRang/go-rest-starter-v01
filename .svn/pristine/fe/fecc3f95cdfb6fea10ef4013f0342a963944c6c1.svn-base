package service

import (
	"errors"
	"log/slog"
	"sync"
	"time"

	"go_ex01/pkg/api/vo"
	"go_ex01/pkg/util"
)

type UserService struct {
	logger *slog.Logger
	users  map[int]*vo.User
	nextID int
	mu     sync.RWMutex
}

func NewUserService() *UserService {
	logger := util.GetLogger().With("component", "user_service")

	service := &UserService{
		logger: logger,
		users:  make(map[int]*vo.User),
		nextID: 1,
	}

	service.seedData()
	return service
}

func (s *UserService) seedData() {
	users := []*vo.User{
		{ID: 1, Name: "John Doe", Email: "john@go_ex01.com", CreatedAt: time.Now().Add(-24 * time.Hour)},
		{ID: 2, Name: "Jane Smith", Email: "jane@go_ex01.com", CreatedAt: time.Now().Add(-12 * time.Hour)},
	}

	for _, user := range users {
		s.users[user.ID] = user
	}
	s.nextID = 3
}

func (s *UserService) GetAllUsers() []vo.UserResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]vo.UserResponse, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, vo.UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		})
	}

	s.logger.Debug("Retrieved all users", "count", len(users))
	return users
}

func (s *UserService) GetUserByID(id int) (*vo.UserResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	user, exists := s.users[id]
	if !exists {
		s.logger.Warn("User not found", "id", id)
		return nil, errors.New("user not found")
	}

	response := &vo.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	s.logger.Debug("Retrieved user", "id", id, "name", user.Name)
	return response, nil
}

func (s *UserService) CreateUser(req vo.CreateUserRequest) (*vo.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isEmailExists(req.Email) {
		s.logger.Warn("Email already exists", "email", req.Email)
		return nil, errors.New("email already exists")
	}

	user := &vo.User{
		ID:        s.nextID,
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	s.users[user.ID] = user
	s.nextID++

	response := &vo.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	s.logger.Info("User created", "id", user.ID, "name", user.Name, "email", user.Email)
	return response, nil
}

func (s *UserService) UpdateUser(id int, req vo.UpdateUserRequest) (*vo.UserResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		s.logger.Warn("User not found for update", "id", id)
		return nil, errors.New("user not found")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		if s.isEmailExistsExcept(req.Email, id) {
			s.logger.Warn("Email already exists for another user", "email", req.Email, "id", id)
			return nil, errors.New("email already exists")
		}
		user.Email = req.Email
	}

	response := &vo.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	s.logger.Info("User updated", "id", id, "name", user.Name, "email", user.Email)
	return response, nil
}

func (s *UserService) DeleteUser(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.users[id]
	if !exists {
		s.logger.Warn("User not found for deletion", "id", id)
		return errors.New("user not found")
	}

	delete(s.users, id)
	s.logger.Info("User deleted", "id", id)
	return nil
}

func (s *UserService) isEmailExists(email string) bool {
	for _, user := range s.users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (s *UserService) isEmailExistsExcept(email string, exceptID int) bool {
	for _, user := range s.users {
		if user.Email == email && user.ID != exceptID {
			return true
		}
	}
	return false
}
