package database

import (

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type Article struct {
	Id      int64
	Title   string
	Summary string
	Body    string
}

func CreateConnection(port string, user string, password string, database string) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     port,
		User:     user,
		Password: password,
		Database: database,
	})
	return db
}

func CreateTable(db *pg.DB) error {
	// _, er := db.Exec("DROP TABLE public.articles;")
	err := db.Model((*Article)(nil)).CreateTable(&orm.CreateTableOptions{})
	return err
}

func InsertArticle(db *pg.DB, article *Article) error {
	_, err := db.Model(article).Insert()
	GetArticle(db, 0)
	return err
}

func GetArticles(db *pg.DB) (error, []Article) {
	var articles []Article
	err := db.Model(&articles).Select()
	if err != nil {
		return err, nil
	}
	return nil, articles
}

func GetArticle(db *pg.DB, Id int) (error, Article) {
	var article Article
	err := db.Model(&article).Where("id = ?", Id).Select()
	if err != nil {
		return err, article
	}
	return nil, article
}

func AddArticle(db *pg.DB, title string, summary string, bodyText string) error {
	article := &Article{
		Title:   title,
		Summary: summary,
		Body:    bodyText,
	}
	err := InsertArticle(db, article)
	if err != nil {
		return err
	}
	return nil
}
