package content

type CreateSetupContent struct {
	UserId   int64  `json:"user_id" validate:"required"`
	UserSlug string `json:"user_slug" validate:"user_slug"`
}
