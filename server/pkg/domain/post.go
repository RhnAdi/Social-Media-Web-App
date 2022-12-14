package domain

import (
	"github.com/RhnAdi/Gomle/internal/auth"
	"github.com/RhnAdi/Gomle/pkg/dto"
	"github.com/RhnAdi/Gomle/pkg/models"
)

type Post struct {
	ID      string         `json:"id"`
	Content string         `json:"content"`
	Files   []models.Image `json:"files"`
}

type PostService interface {
	FindAll() ([]models.Post, error)
	Find(id string) (models.Post, error)
	FindMyPost(claim auth.JWTClaim) ([]dto.MyPostResponse, error)
	Create(claim auth.JWTClaim, post Post) (models.Post, error)
	Update(claim auth.JWTClaim, post Post) (models.Post, error)
	Delete(claim auth.JWTClaim, post Post) (models.Post, error)
	FollowingPosts(claim auth.JWTClaim) ([]models.Post, error)
	AddComment(claim auth.JWTClaim, post_id string, comment dto.CommentRequest) (models.Comment, error)
}

type PostDB interface {
	FindAll() ([]models.Post, error)
	Find(models.Post) (models.Post, error)
	FindMyPost(id string) ([]models.Post, error)
	Create(models.Post) (models.Post, error)
	Update(models.Post) (models.Post, error)
	Delete(models.Post) (models.Post, error)
	FollowingPosts(userId string) ([]models.Post, error)
	AddComment(comment models.Comment) (models.Comment, error)
}
