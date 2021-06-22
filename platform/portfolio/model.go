package portfolio

import (
	"github.com/jinzhu/gorm"
)

// Blog posts
type Blog struct {
	gorm.Model
	Title string
	Topic string
	Text  string
}

// User for admin portal
type User struct {
	gorm.Model
	Username         string // should be unique
	First_Name       string
	Last_Name        string
	Password         string
	Super_User_Check bool
	Email            string
}

// Projects fo r code portfolio page
type Project struct {
	gorm.Model
	Title string
	Img   []byte
	Desc  string
	Url   string
}

//------------------------------ Constructors ------------------------------

func New_Blog(title string, topic string, text string) *Blog {
	blog := &Blog{
		Title: title,
		Topic: topic,
		Text:  text,
	}
	return blog
}

func New_Project(title string, img []byte, desc string, url string) *Project {
	project := &Project{
		Title: title,
		Img:   img,
		Desc:  desc,
		Url:   url,
	}
	return project
}
