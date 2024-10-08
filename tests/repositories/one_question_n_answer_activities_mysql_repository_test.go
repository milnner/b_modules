package repositories

import (
	"bytes"
	"database/sql"
	"strings"
	"testing"
	"time"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repoInterfaces "github.com/milnner/b_modules/repositories/interfaces"
	"github.com/milnner/b_modules/tests/config"
)

func TestOneQuestionNAnswerActivityPolimorfism(t *testing.T) {
	var _ repoInterfaces.IOneQuestionNAnswerActivityRepository = &repositories.OneQuestionNAnswerActivityMySQLRepository{}
}
func TestInsertOneQuestionNAnswerActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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
	testcases := config.OneQuestionNAnswerActivityObjs

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	for _, tc := range testcases {
		if err = repo.Insert(&tc); err != nil {
			t.Error(err)
		}
	}
}

func TestGetOneQuestionNAnswerActivityById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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

	testcases := config.OneQuestionNAnswerActivityObjs
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	var oneQuesNAnsw models.OneQuestionNAnswerActivity
	for _, tc := range testcases {
		oneQuesNAnsw.Id = tc.Id
		if repo.GetOneQuestionNAnswerActivityById(&oneQuesNAnsw); err != nil {
			t.Errorf("[OneQuestionNAnswerActivityselect] %v", err)
		}
		if oneQuesNAnsw.AreaId != tc.AreaId ||
			!bytes.Equal(oneQuesNAnsw.Question, tc.Question) ||
			strings.Compare(oneQuesNAnsw.LastUpdate.String(), tc.LastUpdate.String()) != 0 {
			t.Errorf("[TestGetOneQuestionNAnswerActivityById]\n%v !=\n%v\n", oneQuesNAnsw, tc)
		}
	}
}

func TestGetOneQuestionNAnswerActivitiesByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	testcases := config.OneQuestionNAnswerActivityObjs
	oneQuestionNAnswerActivities := make([]models.OneQuestionNAnswerActivity, len(testcases))
	for i, t := range testcases {
		oneQuestionNAnswerActivities[i] = models.OneQuestionNAnswerActivity{Id: t.Id}
	}

	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	if err = repo.GetOneQuestionNAnswerActivitiesByIds(oneQuestionNAnswerActivities); err != nil {
		t.Errorf("[OneQuestionNAnswerActivity][select] %v", err)
	}
	for i := 0; i < len(oneQuestionNAnswerActivities); i++ {

		if oneQuestionNAnswerActivities[i].AreaId != testcases[i].AreaId ||
			!bytes.Equal(oneQuestionNAnswerActivities[i].Question, testcases[i].Question) ||
			strings.Compare(oneQuestionNAnswerActivities[i].LastUpdate.String(), testcases[i].LastUpdate.String()) != 0 {
			t.Errorf("[TestUpdateOneQuestionNAnswerActivity]\n%v !=\n%v\n", oneQuestionNAnswerActivities, testcases[i])
		}
	}
}

func TestGetOneQuestionNAnswerActivitiesByAreaId(t *testing.T) {
	var (
		dbConn                       *sql.DB
		err                          error
		oneQuestionNAnswerActivities []models.OneQuestionNAnswerActivity
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	testcases := config.OneQuestionNAnswerActivityObjs
	area := config.AreasObjs[0]
	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	if oneQuestionNAnswerActivities, err = repo.GetOneQuestionNAnswerActivitiesByAreaId(&area); err != nil {
		t.Errorf("[OneQuestionNAnswerActivity][select] %v", err)
	}

	for i := 0; i < len(oneQuestionNAnswerActivities); i++ {
		if oneQuestionNAnswerActivities[i].AreaId != testcases[i].AreaId ||
			!bytes.Equal(oneQuestionNAnswerActivities[i].Question, testcases[i].Question) ||
			strings.Compare(oneQuestionNAnswerActivities[i].LastUpdate.String(), testcases[i].LastUpdate.String()) != 0 {
			t.Errorf("[TestGetOneQuestionNAnswerActivitiesByAreaId]\n%v !=\n%v\n", oneQuestionNAnswerActivities, testcases[i])
		}
	}
}

func TestUpdateOneQuestionNAnswerActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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

	testcases := config.OneQuestionNAnswerActivityObjs
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	var oneQuesNAnsw models.OneQuestionNAnswerActivity
	for _, tc := range testcases {
		oneQuesNAnsw = tc
		oneQuesNAnsw.Question = []byte("Outra coisa")
		oneQuesNAnsw.LastUpdate = time.Now()
		if err = repo.Update(&oneQuesNAnsw); err != nil {
			t.Errorf("[TestUpdateOneQuestionNAnswerActivity] %v", err)
		}
	}
}

func TestDeleteOneQuestionNAnswerActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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

	testcases := config.OneQuestionNAnswerActivityObjs
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	for _, tc := range testcases {

		if err = repo.Delete(&tc); err != nil {
			t.Errorf("[DeleteOneQuestionNAnswerActivity] %v", err)
		}
	}
}

func TestGetOneQuestionNAnswerActivityIdsByAreaId(t *testing.T) {
	var (
		dbConn                          *sql.DB
		err                             error
		oneQuestionNAnswerActivitiesIds []int
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	area := config.AreasObjs[0]
	var repo *repositories.OneQuestionNAnswerActivityMySQLRepository
	if repo, err = repositories.NewOneQuestionNAnswerActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	if oneQuestionNAnswerActivitiesIds, err = repo.GetOneQuestionNAnswerActivityIdsByAreaId(&area); err != nil {
		t.Errorf("[GetOneQuestionNAnswerActivityIdsByAreaId] %v", err)
	}
	if len(oneQuestionNAnswerActivitiesIds) != len(config.OneQuestionNAnswerActivityObjs) {
		t.Errorf("[GetOneQuestionNAnswerActivityIdsByAreaId][len] %v !=\n%v", len(oneQuestionNAnswerActivitiesIds), len(config.OneQuestionNAnswerActivityObjs))
	}
}
