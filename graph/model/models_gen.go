// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        string  `json:"id"`
	PostID    string  `json:"postId"`
	AuthorID  string  `json:"authorId"`
	Content   string  `json:"content"`
	ParentID  *string `json:"parentId,omitempty"`
	CreatedAt string  `json:"createdAt"`
}

type Mutation struct {
}

type Post struct {
	ID            string     `json:"id"`
	Title         string     `json:"title"`
	Content       string     `json:"content"`
	Author        *User      `json:"author"`
	AllowComments bool       `json:"allowComments"`
	CreatedAt     string     `json:"createdAt"`
	Comments      []*Comment `json:"comments"`
}

type Query struct {
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
}
