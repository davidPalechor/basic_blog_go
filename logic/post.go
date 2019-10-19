package logic

import (
	"basic_blog_go/utils"
	"basic_blog_go/viewmodels"
	"fmt"
)

// PostLogic defines a struct that will be associated to all operations
// related with the treatment of posts
type PostLogic struct {
	DB utils.DBOrmer
}

func NewPostLogic() *PostLogic {
	return &PostLogic{
		DB: utils.NewDBWrapper(),
	}
}

// CreatePost inserts a new registry into 'post' table
func (p *PostLogic) CreatePost(obj viewmodels.PostRequest) error {
	if _, err := p.DB.Insert(&obj); err != nil {
		return fmt.Errorf("Failed to create post")
	}
	return nil
}
