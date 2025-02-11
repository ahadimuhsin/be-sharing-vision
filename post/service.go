package post

type PostService interface {
	Index(limit int, offset int) ([]Post, error)
	Store(input InputStorePost) (Post, error)
	SelectById(input InputPostDetail) (Post, error)
	Update(inputDetail InputPostDetail, input InputUpdatePost) (Post, error)
	Destroy(postDetail InputPostDetail) (bool, error)
}

type postService struct {
	postRepository PostRepository
}

func NewPostService(postRepository PostRepository) PostService {
	return &postService{postRepository: postRepository}
}

// Index implements PostService.
func (p *postService) Index(limit int, offset int) ([]Post, error) {
	posts, err := p.postRepository.GetAll(limit, offset)
	if err != nil {
		return posts, err
	}
	return posts, nil
}

// Store implements PostService.
func (p *postService) Store(input InputStorePost) (Post, error) {
	post := Post{}
	post.Title = input.Title
	post.Content = input.Content
	post.Category = input.Category
	post.Status = status(input.Status)

	newPost, err := p.postRepository.Store(post)

	if err != nil {
		return newPost, err
	}

	return newPost, nil
}

// SelectById implements PostService.
func (p *postService) SelectById(input InputPostDetail) (Post, error) {
	post, err := p.postRepository.SelectById(input.ID)

	if err != nil {
		return post, err
	}

	return post, nil
}

// Update implements PostService.
func (p *postService) Update(inputDetail InputPostDetail, input InputUpdatePost) (Post, error) {
	post, err := p.postRepository.SelectById(inputDetail.ID)

	if err != nil {
		return post, err
	}

	post.Title = input.Title
	post.Content = input.Content
	post.Category = input.Category
	post.Status = status(input.Status)

	updatedPost, err := p.postRepository.Update(post)

	if err != nil {
		return updatedPost, err
	}

	return updatedPost, nil
}

// Destroy implements PostService.
func (p *postService) Destroy(postDetail InputPostDetail) (bool, error) {
	_, err := p.postRepository.Destroy(postDetail)
	
	if err != nil {
        return false, err
    }
	
	return true, nil
}
