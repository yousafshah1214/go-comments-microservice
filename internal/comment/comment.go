package comment

import "gorm.io/gorm"

type Service struct {
	DB *gorm.DB
}

type Comment struct {
	gorm.Model
	Slug   string
	Body   string
	Author string
}

type ServiceInterface interface {
	GetComment(id uint) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	CreateComment(comment Comment) (Comment, error)
	UpdateComment(id uint, newComment Comment) (Comment, error)
	DeleteComment(id uint) error
	GetAllComments() ([]Comment, error)
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}

func (s *Service) GetComment(id uint) (Comment, error) {
	var comment Comment

	result := s.DB.First(&comment, id)
	if result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment

	result := s.DB.Find(&comments).Where("slug = ?", slug)
	if result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

func (s *Service) GetAllComments() ([]Comment, error) {
	var comments []Comment

	result := s.DB.Find(&comments)
	if result.Error != nil {
		return []Comment{}, result.Error
	}

	return comments, nil
}

func (s *Service) CreateComment(comment Comment) (Comment, error) {
	result := s.DB.Save(&comment)
	if result.Error != nil {
		return Comment{}, result.Error
	}

	return comment, nil
}

func (s *Service) UpdateComment(id uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(id)

	if err != nil {
		return Comment{}, err
	}

	result := s.DB.Model(&comment).Updates(newComment)
	if result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

func (s *Service) DeleteComment(id uint) (Comment, error) {
	var comment Comment

	result := s.DB.Delete(&comment, id)

	if result.Error != nil {
		return Comment{}, nil
	}

	return comment, nil
}
