package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type Blog struct {
	Id         int    `json:"id"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	Image      string `json:"image"`
	Created_at string `json:"createdAt"`
}

func GetBlogs() ([]Blog, error) {

	rows, err := DB.Query("SELECT * FROM blogs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	blogs := make([]Blog, 0)

	for rows.Next() {
		singleBlog := Blog{}
		err = rows.Scan(&singleBlog.Id, &singleBlog.Author, &singleBlog.Title, &singleBlog.Content, &singleBlog.Image, &singleBlog.Created_at)

		if err != nil {
			return nil, err
		}

		blogs = append(blogs, singleBlog)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return blogs, nil
}
