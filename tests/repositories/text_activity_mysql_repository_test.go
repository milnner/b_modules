package repositories

import (
	"bytes"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repositoryInterface "github.com/milnner/b_modules/repositories/interfaces"
	"github.com/milnner/b_modules/tests/config"
)

func TestTextActivityMySQLRepositoryPolimorfism(t *testing.T) {
	var _ repositoryInterface.ITextActivityRepository = &repositories.TextActivityMySQLRepository{}
}

func TestTextActivityMySQLRepositoryInsert(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql")
	var repo *repositories.TextActivityMySQLRepository
	if repo, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	testcase := config.TextActivityObjs
	for _, tc := range testcase {
		if err = repo.Insert(&tc); err != nil {
			t.Errorf("[TextActivity][insert] %v\n", err)
		}
	}
}

func TestTextActivityMySQLRepositoryDelete(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql")
	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.TextActivity); i++ {
		_, err = dbConn.Exec(config.TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	testcase := config.TextActivityObjs
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.TextActivityMySQLRepository
	if repo, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	for _, tc := range testcase {
		if err = repo.Delete(&tc); err != nil {
			t.Errorf("[TextActivity][delete] %v\n", err)
		}
	}
}
func TestTextActivityMySQLRepositoryUpdate(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql")
	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.TextActivity); i++ {
		_, err = dbConn.Exec(config.TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	testcase := config.TextActivityObjs
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.TextActivityMySQLRepository
	if repo, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	nowTime := time.Now()
	for _, tc := range testcase {
		tc.LastUpdate = nowTime
		if err = repo.Update(&tc); err != nil {
			t.Errorf("[TextActivity][update] %v\n", err)
		}
	}
}

func TestTextActivityMySQLRepositoryGetTextActivityById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql")
	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.TextActivity); i++ {
		_, err = dbConn.Exec(config.TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	testcase := config.TextActivityObjs
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.TextActivityMySQLRepository
	if repo, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	textActivity := models.TextActivity{}
	for _, tc := range testcase {
		textActivity.Id = tc.Id
		if err = repo.GetTextActivityById(&textActivity); err != nil {
			t.Errorf("[TextActivity][GetTextActivityById] %v\n", err)
		}
		if textActivity.Activated != tc.Activated ||
			textActivity.AreaId != tc.AreaId ||
			!bytes.Equal(textActivity.Blob, tc.Blob) ||
			textActivity.Title != tc.Title {
			t.Errorf("[TextActivity][GetTextActivityById] \n%v !=\n %v\n", tc, textActivity)
		}
	}
}

func TestTextActivityMySQLRepositoryGetTextActivitiesByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql")
	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.TextActivity); i++ {
		_, err = dbConn.Exec(config.TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	testcase := config.TextActivityObjs
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.TextActivityMySQLRepository
	if repo, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	textActivities := make([]models.TextActivity, len(testcase))
	for i, t := range testcase {
		textActivities[i].Id = t.Id
	}
	if err = repo.GetTextActivitiesByIds(textActivities); err != nil {
		t.Errorf("[TextActivity][GetTextActivityById] %v\n", err)
	}
	fmt.Println(textActivities)

	for i, j, end := 0, 0, len(testcase); i < end; {

		if textActivities[i].Activated != testcase[j].Activated ||
			textActivities[i].AreaId != testcase[j].AreaId ||
			!bytes.Equal(textActivities[i].Blob, testcase[j].Blob) ||
			textActivities[i].Title != testcase[j].Title {
			t.Errorf("[TextActivity][GetTextActivitiesById] \n%v !=\n %v\n", testcase[j], textActivities[i])
		}
		i++
		j++
	}
}

func TestGetTextActivitiesByAreaId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql")
	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.TextActivity); i++ {
		_, err = dbConn.Exec(config.TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	testcase := config.TextActivityObjs
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.TextActivityMySQLRepository
	if repo, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	var textActivities []models.TextActivity

	areaTest := config.AreasObjs[0]
	if textActivities, err = repo.GetTextActivitiesByAreaId(&areaTest); err != nil {
		t.Errorf("[TextActivity][GetTextActivityById] %v\n", err)
	}
	for i, j, end := 0, 0, len(testcase); i < end; {
		if textActivities[i].Activated != testcase[j].Activated ||
			textActivities[i].AreaId != testcase[j].AreaId ||
			!bytes.Equal(textActivities[i].Blob, testcase[j].Blob) ||
			textActivities[i].Title != testcase[j].Title {
			t.Errorf("[TextActivity][GetTextActivitiesById] \n%v !=\n %v\n", testcase[j], textActivities[i])
		}
		i++
		j++
	}
}

func TestGetTextActivityIdsByAreaId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

	}()

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Users); i++ {
		_, err = dbConn.Exec(config.Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.Area); i++ {
		_, err = dbConn.Exec(config.Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.TextActivity); i++ {
		_, err = dbConn.Exec(config.TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	var textRepository *repositories.TextActivityMySQLRepository
	if textRepository, err = repositories.NewTextActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	var txtActsIds []int

	areaTests := config.AreasObjs[0]
	if txtActsIds, err = textRepository.GetTextActivityIdsByAreaId(&areaTests); err != nil {
		t.Errorf("[GetTextActivityIdsByAreaId][%v]\n", err.Error())
	}

	if len(testCases) != len(txtActsIds) {
		t.Errorf("[GetTextActivityIdsByAreaId][len][%v]!=[%v]\n", len(testCases), len(txtActsIds))
	}
}
