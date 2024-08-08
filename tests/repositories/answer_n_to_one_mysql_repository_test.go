package repositories

import (
	"bytes"
	"database/sql"
	"testing"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repoInterfaces "github.com/milnner/b_modules/repositories/interfaces"
)

func TestPolimorfismAnswerNToOneMySQLRepository(t *testing.T) {
	var _ repoInterfaces.IAnswerNToOneRepository = &repositories.AnswerNToOneMySQLRepository{}
}
func TestInsertAnswerNToOneMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	testcase := AnswerNToOneObjs
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	for _, tc := range testcase {
		if err = repo.Insert(&tc); err != nil {
			t.Errorf("[AnswerNToOneMySQLRepository][Insert] %v", err)
		}
	}
}

func TestUpdateAnswerNToOneMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(AnswerNToOne); i++ {
		_, err = dbConn.Exec(AnswerNToOne[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	testcase := AnswerNToOneObjs
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	for _, tc := range testcase {
		tc.Answer = []byte("asdasdasd")
		tc.Correctness = 9
		if err = repo.Update(&tc); err != nil {
			t.Errorf("[TestUpdateAnswerNToOneMySQLRepository] %v", err)
		}
	}
}

func TestDeleteAnswerNToOneMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(AnswerNToOne); i++ {
		_, err = dbConn.Exec(AnswerNToOne[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	testcase := AnswerNToOneObjs
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	for _, tc := range testcase {
		if err = repo.Delete(&tc); err != nil {
			t.Errorf("[TestDeleteAnswerNToOneMySQLRepository] %v", err)
		}
	}
}

func TestGetAnswerNToOneById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(AnswerNToOne); i++ {
		_, err = dbConn.Exec(AnswerNToOne[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	testcase := AnswerNToOneObjs
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	var answer models.AnswerNToOne
	for _, tc := range testcase {
		answer.Id = tc.Id

		if err = repo.GetAnswerNToOneById(&answer); err != nil {
			t.Errorf("[TestGetAnswerNToOneById] %v", err)
		}
		if answer.Id != tc.Id ||
			answer.AreaId != tc.AreaId ||
			!bytes.Equal(answer.Answer, tc.Answer) ||
			answer.OneQuestionNAnswerActivityId != tc.OneQuestionNAnswerActivityId ||
			answer.Correctness != tc.Correctness {
			t.Errorf("[TestGetAnswerNToOneById]\n%v !=\n%v", answer, tc)
		}
	}
}

func TestGetAnswersNToOneByOneQuestionNAnswerActivityId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(AnswerNToOne); i++ {
		_, err = dbConn.Exec(AnswerNToOne[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	testcase := AnswerNToOneObjs
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	var answers []models.AnswerNToOne

	if answers, err = repo.GetAnswersNToOneByOneQuestionNAnswerActivityId(&OneQuestionNAnswerActivityObjs[0]); err != nil {
		t.Errorf("[TestGetAnswerNToOneById] %v", err)
	}

	for _, tc := range testcase {
		for _, answer := range answers {
			if tc.Id == answer.Id &&
				(answer.Id != tc.Id ||
					answer.AreaId != tc.AreaId ||
					!bytes.Equal(answer.Answer, tc.Answer) ||
					answer.OneQuestionNAnswerActivityId != tc.OneQuestionNAnswerActivityId ||
					answer.Correctness != tc.Correctness) {
				t.Errorf("[TestGetAnswerNToOneById]\n%v !=\n%v", answer, tc)
			}
		}
	}
}

func TestGetAnswersNToOneByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(AnswerNToOne); i++ {
		_, err = dbConn.Exec(AnswerNToOne[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	testcase := AnswerNToOneObjs
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	answers := make([]models.AnswerNToOne, len(testcase))
	for i := 0; i < len(answers); i++ {
		answers[i].Id = testcase[i].Id
	}
	if err = repo.GetAnswersNToOneByIds(answers); err != nil {
		t.Errorf("[GetAnswersNToOneIdsByOneQuestionNAnswerActivityId] %v", err)
	}

	for _, tc := range testcase {
		for _, answer := range answers {
			if tc.Id == answer.Id &&
				(answer.Id != tc.Id ||
					answer.AreaId != tc.AreaId ||
					!bytes.Equal(answer.Answer, tc.Answer) ||
					answer.OneQuestionNAnswerActivityId != tc.OneQuestionNAnswerActivityId ||
					answer.Correctness != tc.Correctness) {
				t.Errorf("[GetAnswersNToOneIdsByOneQuestionNAnswerActivityId]\n%v !=\n%v", answer, tc)
			}
		}

	}
}

func TestGetAnswersNToOneIdsByOneQuestionNAnswerActivityId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `answer_n_to_one` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.User.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	for i := 0; i < len(AnswerNToOne); i++ {
		_, err = dbConn.Exec(AnswerNToOne[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, DatabaseConn.AnswerNToOneActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var repo *repositories.AnswerNToOneMySQLRepository
	if repo, err = repositories.NewAnswerNToOneMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	var answerIds []int

	if answerIds, err = repo.GetAnswersNToOneIdsByOneQuestionNAnswerActivityId(&OneQuestionNAnswerActivityObjs[0]); err != nil {
		t.Errorf("[GetAnswersNToOneIdsByOneQuestionNAnswerActivityId] %v", err)
	}
	if len(answerIds) != len(AnswerNToOneObjs) {
		t.Errorf("[GetAnswersNToOneIdsByOneQuestionNAnswerActivityId] %v", err)
	}
}
