package graph

import (
	"context"
	"fmt"
	"strconv"

	"post-system/pkg/models" // Используем модели из pkg/models
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, authorID string) (*models.Post, error) {
	// Конвертация authorID из string в uint
	authorIDUint, err := strconv.ParseUint(authorID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid author ID: %v", err)
	}

	// Создание нового поста
	post := &models.Post{
		Title:    title,
		Content:  content,
		AuthorID: uint(authorIDUint),
	}

	// Сохранение в базу данных (добавь метод в свою логику)
	if err := r.DB.Create(post).Error; err != nil {
		return nil, err
	}

	return post, nil
}

// ToggleAllowComments is the resolver for the toggleAllowComments field.
func (r *mutationResolver) ToggleAllowComments(ctx context.Context, postID string) (*models.Post, error) {
	var post models.Post
	if err := r.DB.First(&post, postID).Error; err != nil {
		return nil, err
	}

	post.AllowComments = !post.AllowComments

	if err := r.DB.Save(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// AddComment is the resolver for the addComment field.
func (r *mutationResolver) AddComment(ctx context.Context, postID string, authorID string, parentID *string, content string) (*models.Comment, error) {
	authorIDUint, err := strconv.ParseUint(authorID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid author ID: %v", err)
	}

	comment := &models.Comment{
		PostID:   uint(postIDUint),
		AuthorID: uint(authorIDUint),
		Content:  content,
	}

	if parentID != nil {
		comment.ParentID = parentID
	}

	if err := r.DB.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	var posts []*models.Post
	if err := r.DB.Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post
	if err := r.DB.First(&post, id).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// NewComment is the resolver for the newComment field.
func (r *subscriptionResolver) NewComment(ctx context.Context, postID string) (<-chan *models.Comment, error) {
	comments := make(chan *models.Comment, 1)

	go func() {
		for {
			// Тут можно добавить подписку на новые комментарии
		}
	}()

	return comments, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Subscription returns SubscriptionResolver implementation.
func (r *Resolver) Subscription() SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
