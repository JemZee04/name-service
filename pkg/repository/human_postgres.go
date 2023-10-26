package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"name-service/model"
	"strings"
)

type HumanPostgres struct {
	db *sqlx.DB
}

func NewHumanPostgres(db *sqlx.DB) *HumanPostgres {
	return &HumanPostgres{db: db}
}

func (r *HumanPostgres) Create(human model.Human) (int, error) {
	var id int
	query := fmt.Sprintf(
		"insert into %s (name, surname, patronymic, age, gender, nationality) values ($1, $2, $3, $4, $5, $6) returning id",
		humanTable,
	)

	logrus.Infof("createQuery: %s", query)

	row := r.db.QueryRow(query, human.Name, human.Surname, human.Patronymic, human.Age, human.Gender, human.Nationality)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *HumanPostgres) GetAll(filter model.FilterHuman, pageSize int) ([]model.Human, error) {
	var humans []model.Human

	query := fmt.Sprintf(
		`select * from %s 
         where position($1 in name)>0 
         and position($2 in surname)>0 
         and position($3 in patronymic)>0 
         and position($4 in gender)>0 
         and position($5 in nationality)>0
         and $6 <= age and age <= $7
         limit $8 offset $9;`, humanTable,
	)

	logrus.Infof("getQuery: %s", query)

	err := r.db.Select(
		&humans, query, filter.Name, filter.Surname, filter.Patronymic,
		filter.Gender, filter.Nationality, filter.MinAge, filter.MaxAge, pageSize, filter.Page,
	)

	return humans, err
}

func (r *HumanPostgres) Delete(id int) error {
	query := fmt.Sprintf("delete from %s where $1 = id", humanTable)

	logrus.Infof("deleteQuery: %s", query)

	_, err := r.db.Exec(query, id)

	return err
}

func (r *HumanPostgres) Update(id int, input model.UpdateHumanInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("name=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}
	if input.Surname != nil {
		setValues = append(setValues, fmt.Sprintf("surname=$%d", argId))
		args = append(args, *input.Surname)
		argId++
	}
	if input.Patronymic != nil {
		setValues = append(setValues, fmt.Sprintf("patronymic=$%d", argId))
		args = append(args, *input.Patronymic)
		argId++
	}
	if input.Age != nil {
		setValues = append(setValues, fmt.Sprintf("age=$%d", argId))
		args = append(args, *input.Age)
		argId++
	}
	if input.Gender != nil {
		setValues = append(setValues, fmt.Sprintf("gender=$%d", argId))
		args = append(args, *input.Gender)
		argId++
	}
	if input.Nationality != nil {
		setValues = append(setValues, fmt.Sprintf("nationality=$%d", argId))
		args = append(args, *input.Nationality)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("update %s set %s where id = $%d", humanTable, setQuery, argId)

	args = append(args, id)

	logrus.Infof("updateQuery: %s", query)
	logrus.Infof("args: %s", args)

	_, err := r.db.Exec(query, args...)
	return err
}
