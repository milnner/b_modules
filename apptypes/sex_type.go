package apptypes

type Sex string

func (u Sex) Equals(o Sex) bool {
	return string(u) == string(o)
}

const (
	male   Sex = "male"
	famale Sex = "famale"
	other  Sex = "other"
)

type Sexs struct {
	male   Sex
	famale Sex
	other  Sex
}

var sexs = Sexs{
	male:   male,
	famale: famale,
	other:  other,
}

func (u Sexs) Male() string {
	return string(sexs.male)
}

func (u Sexs) Female() string {
	return string(sexs.famale)
}
func (u Sexs) Other() string {
	return string(sexs.other)
}
