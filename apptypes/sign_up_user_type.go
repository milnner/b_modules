package apptypes

type SignUpUserType struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Email     string `json:"email"`
	Professor uint8  `json:"professor"`
	EntryDate string `json:"entry-date"`
	BournDate string `json:"bourn-date"`
	Sex       Sex    `json:"sex"`
	Password  string `json:"password"`
}
