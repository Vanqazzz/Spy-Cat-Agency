package app

type Cats struct {
	Id                int    `json:"Id"`
	Name              string `json:"name"`
	YearsOfExperience int    `json:"YearsOfExperience"`
	Breed             string `json:"breed"`
	Salary            int    `json:"salary"`
}

type Missions struct {
	MissionsId    int  `json:"Id"`
	Cat_id        int  `json:"Cat_id"`
	Target_id     int  `json:"Target_id"`
	CompleteState bool `json:"complete_state"`
}

type Target struct {
	TargetId             int    `json:"Id"`
	Target_name          string `json:"Target_name"`
	Country              string `json:"country"`
	Notes                string `json:"notes"`
	CompleteState_target bool   `json:"complete_state_target"`
}
