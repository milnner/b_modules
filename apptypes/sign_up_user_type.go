package apptypes

type SignUpUser struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	Professor  uint8  `json:"professor"`
	EntryDate  string `json:"entryDate"`
	BournDate  string `json:"bournDate"`
	Permission UserClassPermission
	Sex        Sex    `json:"sex"`
	Password   string `json:"password"`
}
