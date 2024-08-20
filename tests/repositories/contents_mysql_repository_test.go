package repositories

import (
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

func TestPolimorfismContentMySQLRepository(t *testing.T) {
	var _ repoInterfaces.IContentRepository = &repositories.ContentMySQLRepository{}
}

func TestInsertContentMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := config.ContentObjs

	for _, tc := range testcases {
		if err = repo.Insert(&tc); err != nil {
			t.Errorf("[TestInsertContentMySQLRepository] %v", err)
		}
	}
}

func TestUpdateContentMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Contents); i++ {
		_, err = dbConn.Exec(config.Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := config.ContentObjs

	for _, tc := range testcases {
		tc.Description = "ola"
		tc.LastUpdate = time.Now()
		if err = repo.Update(&tc); err != nil {
			t.Errorf("[TestUpdateContentMySQLRepository] %v", err)
		}
	}
}

func TestDeleteContentMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Contents); i++ {
		_, err = dbConn.Exec(config.Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := config.ContentObjs

	for _, tc := range testcases {
		if err = repo.Delete(&tc); err != nil {
			t.Errorf("[TestUpdateContentMySQLRepository] %v", err)
		}
	}
}

func TestGetContentById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Contents); i++ {
		_, err = dbConn.Exec(config.Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := config.ContentObjs
	var content models.Content
	for _, tc := range testcases {
		content.Id = tc.Id
		if err = repo.GetContentById(&content); err != nil {
			t.Errorf("[TestGetContentById] %v", err)
		}
		if content.Id != tc.Id ||
			content.Activated != tc.Activated ||
			content.AreaId != tc.AreaId ||
			strings.Compare(content.CreationDate.String()[:19], tc.CreationDate.String()[:19]) != 0 ||
			strings.Compare(content.Title, tc.Title) != 0 ||
			strings.Compare(content.Description, tc.Description) != 0 ||
			strings.Compare(content.LastUpdate.String()[:19], tc.LastUpdate.String()[:19]) != 0 {
			t.Errorf("[TestGetContentById] %v|=\n%v", tc, content)
		}
	}
}

func TestGetContentsByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Contents); i++ {
		_, err = dbConn.Exec(config.Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := config.ContentObjs
	contents := make([]models.Content, len(config.Contents))
	for i := 0; i < len(config.Contents); i++ {
		contents[i].Id = testcases[i].Id
	}

	if err = repo.GetContentsByIds(contents); err != nil {
		t.Errorf("[TestGetContentsByIds] %v", err)
	}

	for i := 0; i < len(config.Contents); i++ {
		contents[i].Id = testcases[i].Id

		if contents[i].Id != testcases[i].Id ||
			contents[i].Activated != testcases[i].Activated ||
			contents[i].AreaId != testcases[i].AreaId ||
			strings.Compare(contents[i].CreationDate.String()[:19], testcases[i].CreationDate.String()[:19]) != 0 ||
			strings.Compare(contents[i].Title, testcases[i].Title) != 0 ||
			strings.Compare(contents[i].Description, testcases[i].Description) != 0 ||
			strings.Compare(contents[i].LastUpdate.String()[:19], testcases[i].LastUpdate.String()[:19]) != 0 {
			t.Errorf("[TestGetContentById] %v|=\n%v", testcases[i], contents[i])
		}
	}
}

func TestGetContentsByAreaId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Contents); i++ {
		_, err = dbConn.Exec(config.Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	area := config.AreasObjs[0]

	var contents []int
	if contents, err = repo.GetContentIdsByAreaId(&area); err != nil && len(config.ContentObjs) != len(contents) {
		t.Errorf("[GetContentIdsByAreaId] %v", err)
	}

}

func TestContentAddActivity_TestContentGetActivityIds_TestContentUpdateActivity_TestContentRemoveActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `content_see_activity` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.TextActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.Contents); i++ {
		_, err = dbConn.Exec(config.Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
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
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.OneQuestionNAnswerActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(config.OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	for _, act := range config.TextActivityObjs {
		if err = repo.AddActivity(&config.ContentObjs[0], &act); err != nil {
			t.Errorf("[TestContentAddActivity] %v", err)
		}
	}
	for _, act := range config.ImageActivityObjs {
		if err = repo.AddActivity(&config.ContentObjs[0], &act); err != nil {
			t.Errorf("[TestContentAddActivity] %v", err)
		}
	}
	for _, act := range config.OneQuestionNAnswerActivityObjs {
		if err = repo.AddActivity(&config.ContentObjs[0], &act); err != nil {
			t.Errorf("[TestContentAddActivity] %v", err)
		}
	}

	// GET
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if actIds, err := repo.GetActivityIdsByContentId(&config.ContentObjs[0], &models.TextActivity{}); err != nil || len(actIds) == 0 {
		t.Errorf("[TestContentGetActivity] %v\n ids: %v", err, actIds)
	}

	if actIds, err := repo.GetActivityIdsByContentId(&config.ContentObjs[0], &models.ImageActivity{}); err != nil || len(actIds) == 0 {
		t.Errorf("[TestContentGetActivity] %v\n ids: %v", err, actIds)
	}

	if actIds, err := repo.GetActivityIdsByContentId(&config.ContentObjs[0], &models.OneQuestionNAnswerActivity{}); err != nil || len(actIds) == 0 {
		t.Errorf("[TestContentGetActivity] %v\n ids: %v", err, actIds)
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	// update
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	textUpdated := config.TextActivityObjs[0]
	textUpdated.Position = 1
	if err := repo.UpdateActivityPosition(&config.ContentObjs[0], &textUpdated); err != nil {
		t.Errorf("[TestContentUpdateActivity] %v\n", err)
	}
	imgUpdated := config.ImageActivityObjs[0]
	imgUpdated.Position = 1
	if err := repo.UpdateActivityPosition(&config.ContentObjs[0], &imgUpdated); err != nil {
		t.Errorf("[TestContentUpdateActivity] %v\n", err)
	}
	oneQNAswUpdate := config.OneQuestionNAnswerActivityObjs[0]
	oneQNAswUpdate.Position = 1
	if err := repo.UpdateActivityPosition(&config.ContentObjs[0], &oneQNAswUpdate); err != nil {
		t.Errorf("[TestContentUpdateActivity] %v\n", err)
	}
	// remove
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Content.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if err := repo.RemoveActivity(&config.ContentObjs[0], &models.TextActivity{}); err != nil {
		t.Errorf("[TestContentRemoveActivity] %v\n", err)
	}

	if err := repo.RemoveActivity(&config.ContentObjs[0], &models.ImageActivity{}); err != nil {
		t.Errorf("[TestContentRemoveActivity] %v\n", err)
	}

	if err := repo.RemoveActivity(&config.ContentObjs[0], &models.OneQuestionNAnswerActivity{}); err != nil {
		t.Errorf("[TestContentRemoveActivity] %v\n", err)
	}
}
