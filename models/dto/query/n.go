package query

type N struct {
	N int `form:"n" binding:"required,numeric,min=1,max=100,default=1"`
}
