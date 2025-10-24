package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"main.go/internal/app"
)

type Storage struct {
	db *sql.DB
}

func New(StoragePath string) (*Storage, error) {
	const op = "storage.postgres.New"
	db, err := sql.Open("postgres", StoragePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func CreateCat(s Storage, name string, years_of_exp int, breed string, salary int) (string, error) {

	const op = "storage.postgres.createcat"

	_, err := s.db.Exec("INSERT INTO cats (name, years_of_exp, breed, salary) VALUES($1, $2, $3, $4)", name, years_of_exp, breed, salary)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return name, nil
}

func ListSingle_SpyCat(s Storage, name string) (app.Cats, error) {
	const op = "storage.postgres.listspycats"

	stmt, err := s.db.Prepare("SELECT name, years_of_exp, breed, salary FROM cats WHERE name = $1")
	if err != nil {
		return app.Cats{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRow(name)

	cat := app.Cats{}
	err = row.Scan(&cat.Name, &cat.YearsOfExperience, &cat.Breed, &cat.Salary)

	return cat, nil
}

func AllSpyCats(s Storage) ([]app.Cats, error) {
	const op = "storage.postgres.listallspycats"

	row, err := s.db.Query("SELECT name, years_of_exp, breed, salary FROM cats")
	if err != nil {
		return []app.Cats{}, fmt.Errorf("%s: %w", op, err)
	}

	list := []app.Cats{}
	for row.Next() {
		cats := app.Cats{}
		err = row.Scan(&cats.Name, &cats.YearsOfExperience, &cats.Breed, &cats.Salary)

		list = append(list, cats)
	}

	return list, nil
}

func DeleteCat(s Storage, name string) error {
	const op = "storage.postgres.deletecat"

	_, err := s.db.Exec("DELETE FROM cats WHERE name = $1", name)
	if err != nil {
		return fmt.Errorf("%s: %w: ", op, err)
	}

	return nil
}

func UpdateCat(s Storage, salary int, name string) error {
	const op = "storage.postgres.updatecat"

	_, err := s.db.Exec("UPDATE cats SET salary = $1 WHERE name = $2", salary, name)
	if err != nil {
		return fmt.Errorf("%s: %w: ", op, err)
	}

	return nil
}
