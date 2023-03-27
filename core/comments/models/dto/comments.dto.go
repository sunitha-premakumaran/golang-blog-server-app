package dto

import "blog-server-app/internal/blogs/models/dto"

type Reaction string

const (
	Like   Reaction = "like"
	UnLike          = "unlike"
	Clap            = "clap"
	Heart           = "heart"
)

type CreateCommentDto struct {
	Content string `json:"content"`
}

type CommentReaction struct {
	Reaction Reaction
	MadeBy   dto.User
}

type GetCommentsResultDto struct {
	Comments []GetCommentDto `json:"comments"`
}

type GetCommentDto struct {
	Content          string            `json:"content"`
	Author           dto.User          `json:"author"`
	CommentReactions []CommentReaction `json:"commentReactions"`
}

type CreateCommentResponseDto struct {
	CommentId string `json:"commentId"`
}
