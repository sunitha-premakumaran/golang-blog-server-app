package dto

import (
	shared "blog-server-app/modules/shared/dto"
)

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
	MadeBy   shared.User
}

type GetCommentsResultDto struct {
	comments []GetCommentDto
}

type GetCommentDto struct {
	Content          string            `json:"content"`
	Author           shared.User       `json:"author"`
	CommentReactions []CommentReaction `json:"commentReactions"`
}

type CreateCommentResponseDto struct {
	CommentId string `json:"commentId"`
}
