package repositories

import (
	"bytes"
	"database/sql"
	"testing"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repositoryInterface "github.com/milnner/b_modules/repositories/interfaces"
	"github.com/milnner/b_modules/tests/config"
)

func TestImageActivityMySQLRepositoryPolimorfism(t *testing.T) {
	var _ repositoryInterface.IImageActivityRepository = &repositories.ImageActivityMySQLRepository{}
}

func TestInsertImageActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	for _, tc := range testCases {
		if err = imageRepository.Insert(&tc); err != nil {
			t.Errorf("[insert img] %v", err)
		}
	}
}

func TestUpdateImageActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	for _, tc := range testCases {
		tc.Title = "New"
		if err = imageRepository.Update(&tc); err != nil {
			t.Errorf("[ImageActivity][update][%v]\n", err.Error())
		}
	}
}

func TestGetImageActivityById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	for _, tc := range testCases {
		u := &models.ImageActivity{Id: tc.Id}
		if err = imageRepository.GetImageActivityById(u); err != nil {
			t.Errorf("[ImageActivity][update][%v]\n", err)

		}
		if u.AreaId != tc.AreaId ||
			!bytes.Equal(u.Blob, tc.Blob) ||
			u.Title != tc.Title ||
			u.Activated != tc.Activated {
			t.Errorf("[ImageActivity][update][%v]!=[%v]\n", u, tc)

		}
	}
}

func TestGetImageActivitiesByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	imgActs := []models.ImageActivity{{Id: 1}, {Id: 2}}

	if err = imageRepository.GetImageActivitiesByIds(imgActs); err != nil {
		t.Errorf("[ImageActivity][update][%v]\n", err.Error())
	}
	if len(testCases) != len(imgActs) {
		t.Errorf("[ImageActivity][update][len][%v]!=[%v]\n", len(testCases), len(imgActs))

	} else {
		for i, tc := range testCases {
			if tc.Id != imgActs[i].Id ||
				tc.AreaId != imgActs[i].AreaId ||
				tc.Title != imgActs[i].Title ||
				!bytes.Equal(tc.Blob, imgActs[i].Blob) ||
				tc.Activated != imgActs[i].Activated {
				t.Errorf("[ImageActivity][update][%v]!=[%v]\n", imgActs[i], tc)

			}
		}
	}
}
func TestDeleteImageActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	for _, tc := range testCases {
		if err := imageRepository.Delete(&tc); err != nil {
			t.Errorf("[ImageActivity][Delete][%v]\n", err.Error())
		}
	}
}

func TestGetImageActivitiesByAreaId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	var imgActs []models.ImageActivity

	areaTests := config.AreasObjs[0]
	if imgActs, err = imageRepository.GetImageActivitiesByAreaId(&areaTests); err != nil {
		t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][%v]\n", err.Error())
	}

	if len(testCases) != len(imgActs) {
		t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][len][%v]!=[%v]\n", len(testCases), len(imgActs))

	} else {
		((models.ImageActivities)(imgActs)).Sort("Id")
		((models.ImageActivities)(testCases)).Sort("Id")
		for i, tc := range testCases {
			if tc.Id != imgActs[i].Id ||
				tc.AreaId != imgActs[i].AreaId ||
				tc.Title != imgActs[i].Title ||
				!bytes.Equal(tc.Blob, imgActs[i].Blob) ||
				tc.Activated != imgActs[i].Activated {
				t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][%v]!=[%v]\n", imgActs[i], tc)

			}
		}
	}
}

func TestGetImageActivityIdsByAreaId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(config.ImageActivity); i++ {
		_, err = dbConn.Exec(config.ImageActivity[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.ImageActivity.GetUpdate(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := config.ImageActivityObjs

	var imgActsIds []int

	areaTests := config.AreasObjs[0]
	if imgActsIds, err = imageRepository.GetImageActivityIdsByAreaId(&areaTests); err != nil {
		t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][%v]\n", err.Error())
	}

	if len(testCases) != len(imgActsIds) {
		t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][len][%v]!=[%v]\n", len(testCases), len(imgActsIds))
	}
}
