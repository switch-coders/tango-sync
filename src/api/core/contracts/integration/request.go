package integration

type Request struct {
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"email"`
	TangoKey string `form:"tango_key" binding:"required"`
}
