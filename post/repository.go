package post

import (
	"gorm.io/gorm"
)

type PostRepository interface {
	Store(post Post) (Post, error)
	Update(post Post) (Post, error)
	GetAll(limit int, offset int) ([]Post, error)
	SelectById(id int) (Post, error)
	Destroy(postDetail InputPostDetail) (bool, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db: db}
}

// GetAll implements PostRepository.
func (p *postRepository) GetAll(limit int, offset int) ([]Post, error) {
	var posts []Post
	
	err := p.db.Find(&posts).Error
	
	if (limit > 0 && offset > 0){
		err = p.db.Limit(limit).Offset(offset).Find(&posts).Error
	}
	
	if err != nil {
		return posts, err
	}

	return posts, nil
}

// Store implements PostRepository.
func (p *postRepository) Store(post Post) (Post, error) {
	err := p.db.Create(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil
}

// SelectById implements PostRepository.
func (p *postRepository) SelectById(id int) (Post, error) {
	var post Post

	err := p.db.First(&post, id).Error

	if err != nil {
		return post, err
	}

	return post, nil
}

// Update implements PostRepository.
func (p *postRepository) Update(post Post) (Post, error) {
	err := p.db.Save(&post).Error

	if err != nil {
		return post, err
	}

	return post, nil
}

// Destroy implements PostRepository.
func (p *postRepository) Destroy(postDetail InputPostDetail) (bool, error) {
	post := Post{
		ID: int64(postDetail.ID),
	}

	err := p.db.Delete(&post).Error

	if err!= nil {
        return false, err
    }
	return true, nil
}
