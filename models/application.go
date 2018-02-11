package models

import (
	"errors"
	"ichabod/db"
	"ichabod/forms"
	"log"
	"time"
)

//Application ...
type Application struct {
	ID        int       `db:"id, primarykey, autoincrement" json:"id"`
	Title     string    `db:"title" json:"title"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

//ApplicationModel ...
type ApplicationModel struct{}

//Create ...
func (m ApplicationModel) Create(form forms.ApplicationCreateForm) (application Application, err error) {
	getDb := db.GetDB()

	query := "INSERT INTO public.applications(title) VALUES($1) RETURNING id"

	res, err := getDb.Prepare(query)

	var applicationID int

	err = res.QueryRow(form.Title).Scan(&applicationID)

	if err == nil {
		err = getDb.SelectOne(&application, "SELECT id, title, updated_at, created_at FROM public.applications WHERE id=$1 LIMIT 1", applicationID)
		log.Println(err)

		if err == nil {
			return application, nil
		}
	}

	return application, errors.New("Not Created")
}

//One ...
func (m ApplicationModel) One(id int64) (application Application, err error) {
	err = db.GetDB().SelectOne(&application, "SELECT id, title, created_at, updated_at FROM applications WHERE id = $1", id)
	return application, err
}
