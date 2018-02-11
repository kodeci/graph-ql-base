package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"ichabod/db"
	"ichabod/forms"
	"time"

	"github.com/gosimple/slug"
)

//Environment ...
type Environment struct {
	ID            int64           `db:"id, primarykey, autoincrement" json:"id"`
	ApplicationID int64           `db:"application_id" json:"application_id"`
	Title         string          `db:"title" json:"title"`
	Slug          string          `db:"slug" json:"slug"`
	Values        json.RawMessage `db:"values" json:"values"`
	UpdatedAt     time.Time       `db:"updated_at" json:"updated_at"`
	CreatedAt     time.Time       `db:"created_at" json:"created_at"`
}

//EnvironmentModel ...
type EnvironmentModel struct{}

//Create ...
func (m EnvironmentModel) Create(applicationID int64, form forms.EnvironmentCreateForm) (environment Environment, err error) {
	getDb := db.GetDB()

	applicationModel := new(ApplicationModel)

	_, err = applicationModel.One(applicationID)

	if err == nil {
		query := "INSERT INTO public.environments(application_id, title, slug, values) VALUES ($1, $2, $3, $4::jsonb) RETURNING id"
		res, err := getDb.Prepare(query)

		emptyJSON, _ := json.Marshal("")
		fmt.Println(string(emptyJSON))

		var environmentID int
		var envSlug string
		envSlug = form.Slug
		if envSlug == "" {
			envSlug = slug.Make(form.Title)
		}

		err = res.QueryRow(applicationID, form.Title, envSlug, string(emptyJSON)).Scan(&environmentID)

		if err == nil {
			err = getDb.SelectOne(&environment, "SELECT id, title, slug, values, updated_at, created_at FROM public.environments WHERE id=$1 LIMIT 1", environmentID)

			if err == nil {
				return environment, nil
			}
		}
	} else {
		return environment, errors.New("Application Not Found")
	}

	return environment, errors.New("Environment Not Created")
}

//One ...
// func (m EnvironmentModel) One(userID, id int64) (article Article, err error) {
// 	err = db.GetDB().SelectOne(&article, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 AND a.id=$2 LIMIT 1", userID, id)
// 	return article, err
// }

// //All ...
// func (m EnvironmentModel) All(userID int64) (articles []Article, err error) {
// 	_, err = db.GetDB().Select(&articles, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC", userID)
// 	return articles, err
// }

// //Update ...
// func (m EnvironmentModel) Update(userID int64, id int64, form forms.ArticleForm) (err error) {
// 	_, err = m.One(userID, id)

// 	if err != nil {
// 		return errors.New("Article not found")
// 	}

// 	_, err = db.GetDB().Exec("UPDATE public.article SET title=$1, content=$2, updated_at=$3 WHERE id=$4", form.Title, form.Content, time.Now().Unix(), id)

// 	return err
// }

// //Delete ...
// func (m EnvironmentModel) Delete(userID, id int64) (err error) {
// 	_, err = m.One(userID, id)

// 	if err != nil {
// 		return errors.New("Article not found")
// 	}

// 	_, err = db.GetDB().Exec("DELETE FROM public.article WHERE id=$1", id)

// 	return err
// }
