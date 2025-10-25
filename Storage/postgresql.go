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

func CreateCat(s Storage, Cat *app.Cats) (string, error) {

	const op = "storage.postgres.createcat"
	var id int
	err := s.db.QueryRow("INSERT INTO cats (name, years_of_exp, breed, salary) VALUES($1, $2, $3, $4) RETURNING id", Cat.Name, Cat.YearsOfExperience, Cat.Breed, Cat.Salary).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	Cat.Id = id

	return Cat.Name, nil
}

func Get_ListSingle_SpyCat(s Storage, name string) (app.Cats, error) {
	const op = "storage.postgres.listspycats"

	stmt, err := s.db.Prepare("SELECT * FROM cats WHERE name = $1")
	if err != nil {
		return app.Cats{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRow(name)

	cat := app.Cats{}
	err = row.Scan(&cat.Id, &cat.Name, &cat.YearsOfExperience, &cat.Breed, &cat.Salary)

	return cat, nil
}

func Show_AllSpyCats(s Storage) ([]app.Cats, error) {
	const op = "storage.postgres.listallspycats"

	row, err := s.db.Query("SELECT * FROM cats")
	if err != nil {
		return []app.Cats{}, fmt.Errorf("%s: %w", op, err)
	}

	list := []app.Cats{}
	for row.Next() {
		cats := app.Cats{}
		err = row.Scan(&cats.Id, &cats.Name, &cats.YearsOfExperience, &cats.Breed, &cats.Salary)

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

func CreateTarget(s Storage, app *app.Target) error {
	const op = "storage.postgres.createtarget"
	var id int
	err := s.db.QueryRow("INSERT INTO target (name, country, notes, complete_state) VALUES ($1, $2, $3, $4) RETURNING id", app.Target_name, app.Country, app.Notes, app.CompleteState_target).Scan(&id)
	if err != nil {

		return fmt.Errorf("%s: %w: ", op, err)
	}

	app.TargetId = id
	fmt.Println(app.TargetId)
	return nil
}

func CreateMission(s Storage, mission *app.Missions) error {
	const op = "storage.postgres.createmission"

	var id int
	err := s.db.QueryRow("INSERT INTO missions (cat_id, target_id, complete_state) VALUES ($1, $2, $3) RETURNING id", mission.Cat_id, mission.Target_id, mission.CompleteState).Scan(&id)
	if err != nil {
		return fmt.Errorf("%s: %w: ", op, err)
	}
	mission.MissionsId = id
	return nil
}

func DeleteMission(s Storage, mission *app.Missions) error {
	const op = "storage.postgres.deletemission"
	fmt.Println(mission.MissionsId)

	_, err := s.db.Exec("DELETE FROM missions WHERE id = $1 ", mission.MissionsId)
	if err != nil {
		return fmt.Errorf("%s: %w: ", op, err)
	}

	return nil
}

func UpdateMission(s Storage, mission *app.Missions) error {
	const op = "storage.postgres.updatemission"

	_, err := s.db.Exec("UPDATE missions SET complete_state = $1 WHERE id = $2  ", mission.CompleteState, mission.MissionsId)
	if err != nil {
		return fmt.Errorf("%s: %w: ", op, err)
	}

	return nil
}

func Show_AllMissions(s Storage) ([]app.Missions, error) {
	const op = "storage.postgres.allmissions"

	row, err := s.db.Query("SELECT * FROM missions")
	if err != nil {
		return []app.Missions{}, fmt.Errorf("%s: %w", op, err)
	}

	list := []app.Missions{}
	for row.Next() {
		mission := app.Missions{}
		err = row.Scan(&mission.MissionsId, &mission.Cat_id, &mission.Target_id, &mission.CompleteState)

		list = append(list, mission)
	}

	return list, nil
}

func Get_SingleMission(s Storage, mission app.Missions) (app.Missions, error) {
	const op = "storage.postgres.singlemission"
	id := mission.MissionsId

	stmt, err := s.db.Prepare("SELECT * FROM missions WHERE id = $1")
	if err != nil {
		return app.Missions{}, fmt.Errorf("%s: %w", op, err)
	}
	row := stmt.QueryRow(id)

	Mission := app.Missions{}
	err = row.Scan(&Mission.MissionsId, &Mission.Cat_id, &Mission.Target_id, &Mission.CompleteState)

	return Mission, nil

}

func AssignCat(s Storage, mission app.Missions) error {

	const op = "storage.postgres.assigncat"
	var count int

	err := s.db.QueryRow(`SELECT COUNT(*) FROM missions WHERE cat_id = $1 AND complete_state = false`, mission.Cat_id).Scan(&count)
	if err != nil {
		return fmt.Errorf("failed to check cat missions: %w", err)
	}

	if count > 0 {
		return fmt.Errorf("%s: cat already has an active mission", op)
	}

	_, err = s.db.Exec("UPDATE missions SET cat_id = $1 WHERE id = $2  AND complete_state = false  ", mission.Cat_id, mission.MissionsId)
	if err != nil {
		return fmt.Errorf("%s: %w: ", op, err)
	}

	return nil
}
