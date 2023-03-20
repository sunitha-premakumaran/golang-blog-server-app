package dto

type CreateCommentDto struct {
	Content string `json:"content"`
}

type GetCommentDto struct {
}

type CreateCommentResponseDto struct {
	CommentId string `json:"commentId"`
}
