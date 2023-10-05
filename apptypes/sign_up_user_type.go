package apptypes

type SignUpUserType struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	EntryDate string `json:"entry-date"`
	BournDate string `json:"bourn-date"`
	Sex       Sex    `json:"sex"`
	Password  string `json:"password"`
	PasswordR string `json:"password-r"`
}
