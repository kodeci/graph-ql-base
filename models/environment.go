package models

import (
	"encoding/json"
	"errors"
	"ichabod/db"
	"ichabod/forms"
	"log"
	"time"

	"github.com/gosimple/slug"
)

//Environment ...
type Environment struct {
	ID            int64           `db:"id, primarykey, autoincrement" json:"id"`
	ApplicationID int64           `db:"application_id" json:"application_id"`
	Title         string          `db:"title" json:"title"`
	Slug          string          `db:"slug" json:"slug"`
	Data          json.RawMessage `db:"data" json:"data"`
	Values        map[string]string
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
}

//EnvironmentModel ...
type EnvironmentModel struct{}

//Create ...
func (m EnvironmentModel) Create(applicationID int64, form forms.EnvironmentCreateForm) (environment Environment, err error) {
	getDb := db.GetDB()

	applicationModel := new(ApplicationModel)

	_, err = applicationModel.One(applicationID)

	if err == nil {
		query := "INSERT INTO public.environments(application_id, title, slug, data) VALUES ($1, $2, $3, $4::jsonb) RETURNING id"
		res, err := getDb.Prepare(query)

		if err != nil {
			return environment, errors.New("Error occured")
		}

		emptyJSON, _ := json.Marshal("")

		envSlug := form.Slug
		if envSlug == "" {
			envSlug = slug.Make(form.Title)
		}

		var environmentID int
		err = res.QueryRow(applicationID, form.Title, envSlug, string(emptyJSON)).Scan(&environmentID)

		if err == nil {
			err = getDb.SelectOne(&environment, "SELECT id, application_id, title, slug, data, updated_at, created_at FROM public.environments WHERE id=$1 LIMIT 1", environmentID)

			if err == nil {
				return environment, nil
			}
		}
	} else {
		return environment, errors.New("Application Not Found")
	}

	return environment, errors.New("Environment Not Created")
}

//Get ...
func (m EnvironmentModel) Get(appID int64, slug string) (environment Environment, err error) {
	err = db.GetDB().SelectOne(&environment, "SELECT id, title, slug, data, updated_at, created_at FROM public.environments WHERE application_id = $1 AND slug = $2 LIMIT 1", appID, slug)
	return environment, err
}

// //All ...
// func (m EnvironmentModel) All(userID int64) (articles []Article, err error) {
// 	_, err = db.GetDB().Select(&articles, "SELECT a.id, a.title, a.content, a.updated_at, a.created_at, json_build_object('id', u.id, 'name', u.name, 'email', u.email) AS user FROM public.article a LEFT JOIN public.user u ON a.user_id = u.id WHERE a.user_id=$1 ORDER BY a.id DESC", userID)
// 	return articles, err
// }

//Update ...
func (m EnvironmentModel) Update(appID int64, slug string, title string, key string, value string) (err error) {

	environment, err := m.Get(appID, slug)

	log.Println(environment)

	if err != nil {
		return errors.New("Environment not found")
	}

	// Set title off of incoming; if no incoming title, set with existing title
	if title == "" {
		title = environment.Title
	}

	// if title is set, just update the title; otherwise update key/value pair
	if title != "" && key == "" || value == "" {
		return errors.New("Key/Value pair missing")
	}

	json.Unmarshal(environment.Data, &environment.Values)

	if environment.Values == nil {
		environment.Values = make(map[string]string)
	}
	environment.Values[key] = value
	values, err := json.Marshal(environment.Values)

	if err != nil {
		return errors.New("key value pair failed")
	}

	_, err = db.GetDB().Exec("UPDATE public.environments SET title=$1, data=$2::json WHERE id=$3", title, string(values), environment.ID)

	return err
}

// //Delete ...
// func (m EnvironmentModel) Delete(userID, id int64) (err error) {
// 	_, err = m.One(userID, id)

// 	if err != nil {
// 		return errors.New("Article not found")
// 	}

// 	_, err = db.GetDB().Exec("DELETE FROM public.article WHERE id=$1", id)

// 	return err
// }
