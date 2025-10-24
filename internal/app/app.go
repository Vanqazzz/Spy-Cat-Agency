package app

type Cats struct {
	Name              string `json:"name"`
	YearsOfExperience int    `json:"YearsOfExperience"`
	Breed             string `json:"breed"`
	Salary            int    `json:"salary"`
}

type Missons struct {
	Cats
	Targets       string `json:"targets"`
	CompleteState string `json:"complete_state"`
}

type Targets struct {
	Name          string `json:"target_name"`
	Country       string `json:"country"`
	Notes         string `json:"notes"`
	CompleteState string `json:"complete_state"`
}
