package repositories

import (
	"database/sql"
	"net"
	"strings"
	"testing"
	"time"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/environment"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repoInterfaces "github.com/milnner/b_modules/repositories/interfaces"
)

func init() {
	port := "3306"

	dC := environment.NewDatabaseConnections()

	dC.SetRootConnString(RootConnString)

	environment.Environment().InitDatabaseConnections(dC)
	target := "127.0.0.1:" + port
	conn, err := net.DialTimeout("tcp", target, 10*time.Second)
	if err != nil {
		panic(err)
	}
	conn.Close()
}

func TestPolimorfismContentMySQLRepository(t *testing.T) {
	var _ repoInterfaces.IContentRepository = &repositories.ContentMySQLRepository{}
}

func TestInsertContentMySQLRepository(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := ContentObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := ContentObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := ContentObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := ContentObjs
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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := ContentObjs
	contents := make([]models.Content, len(Contents))
	for i := 0; i < len(Contents); i++ {
		contents[i].Id = testcases[i].Id
	}

	if err = repo.GetContentsByIds(contents); err != nil {
		t.Errorf("[TestGetContentsByIds] %v", err)
	}

	for i := 0; i < len(Contents); i++ {
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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	testcases := ContentObjs

	area := AreasObjs[0]

	var contents []models.Content
	if contents, err = repo.GetContentsByAreaId(&area); err != nil {
		t.Errorf("[TestGetContentsByAreaId] %v", err)
	}

	for _, tc := range testcases {
		for _, content := range contents {

			if content.Id == tc.Id && (content.Activated != tc.Activated ||
				content.AreaId != tc.AreaId ||
				strings.Compare(content.CreationDate.String()[:19], tc.CreationDate.String()[:19]) != 0 ||
				strings.Compare(content.Title, tc.Title) != 0 ||
				strings.Compare(content.Description, tc.Description) != 0 ||
				strings.Compare(content.LastUpdate.String()[:19], tc.LastUpdate.String()[:19]) != 0) {
				t.Errorf("[TestGetContentsByAreaId] %v|=\n%v", tc, content)
			}
		}
	}
}

func TestContentAddActivity_TestContentGetActivityIds_TestContentUpdateActivity_TestContentRemoveActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
		repo   *repositories.ContentMySQLRepository
	)
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `content_see_activity` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteTxtAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `text_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteOneQuestionNAnswerActivity(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `one_question_n_answer_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err := dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Fatal(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Fatal(err)
		}
	}()

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInserImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(ImageActivity); i++ {
		_, err = dbConn.Exec(ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertTxtAct(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(TextActivity); i++ {
		_, err = dbConn.Exec(TextActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertOneQuestionNAnswerActivity(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(OneQuestionNAnswerActivity); i++ {
		_, err = dbConn.Exec(OneQuestionNAnswerActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	for _, act := range TextActivityObjs {
		if err = repo.AddActivity(&ContentObjs[0], &act); err != nil {
			t.Errorf("[TestContentAddActivity] %v", err)
		}
	}
	for _, act := range ImageActivityObjs {
		if err = repo.AddActivity(&ContentObjs[0], &act); err != nil {
			t.Errorf("[TestContentAddActivity] %v", err)
		}
	}
	for _, act := range OneQuestionNAnswerActivityObjs {
		if err = repo.AddActivity(&ContentObjs[0], &act); err != nil {
			t.Errorf("[TestContentAddActivity] %v", err)
		}
	}

	// GET
	if err = database.InitDBConnection(&dbConn, dC.GetSelectContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if actIds, err := repo.GetActivityIdsByContentId(&ContentObjs[0], &models.TextActivity{}); err != nil || len(actIds) == 0 {
		t.Errorf("[TestContentGetActivity] %v\n ids: %v", err, actIds)
	}

	if actIds, err := repo.GetActivityIdsByContentId(&ContentObjs[0], &models.ImageActivity{}); err != nil || len(actIds) == 0 {
		t.Errorf("[TestContentGetActivity] %v\n ids: %v", err, actIds)
	}

	if actIds, err := repo.GetActivityIdsByContentId(&ContentObjs[0], &models.OneQuestionNAnswerActivity{}); err != nil || len(actIds) == 0 {
		t.Errorf("[TestContentGetActivity] %v\n ids: %v", err, actIds)
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repo, err = repositories.NewContentMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}
	// update
	if err = database.InitDBConnection(&dbConn, dC.GetSelectContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	textUpdated := TextActivityObjs[0]
	textUpdated.Position = 1
	if err := repo.UpdateActivityPosition(&ContentObjs[0], &textUpdated); err != nil {
		t.Errorf("[TestContentUpdateActivity] %v\n", err)
	}
	imgUpdated := ImageActivityObjs[0]
	imgUpdated.Position = 1
	if err := repo.UpdateActivityPosition(&ContentObjs[0], &imgUpdated); err != nil {
		t.Errorf("[TestContentUpdateActivity] %v\n", err)
	}
	oneQNAswUpdate := OneQuestionNAnswerActivityObjs[0]
	oneQNAswUpdate.Position = 1
	if err := repo.UpdateActivityPosition(&ContentObjs[0], &oneQNAswUpdate); err != nil {
		t.Errorf("[TestContentUpdateActivity] %v\n", err)
	}
	// remove
	if err = database.InitDBConnection(&dbConn, dC.GetSelectContent(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if err := repo.RemoveActivity(&ContentObjs[0], &models.TextActivity{}); err != nil {
		t.Errorf("[TestContentRemoveActivity] %v\n", err)
	}

	if err := repo.RemoveActivity(&ContentObjs[0], &models.ImageActivity{}); err != nil {
		t.Errorf("[TestContentRemoveActivity] %v\n", err)
	}

	if err := repo.RemoveActivity(&ContentObjs[0], &models.OneQuestionNAnswerActivity{}); err != nil {
		t.Errorf("[TestContentRemoveActivity] %v\n", err)
	}
}
