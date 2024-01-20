package repositories

import (
	"bytes"
	"database/sql"
	"net"
	"testing"
	"time"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/environment"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repositoryInterface "github.com/milnner/b_modules/repositories/interfaces"
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

func TestImageActivityMySQLRepositoryPolimorfism(t *testing.T) {
	var _ repositoryInterface.IImageActivityRepository = &repositories.ImageActivityMySQLRepository{}
}

func TestInsertImageActivity(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDBConnection(&dbConn, dC.GetInserImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDBConnection(&dbConn, dC.GetUpdateImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDBConnection(&dbConn, dC.GetUpdateImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
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

	if err = database.InitDBConnection(&dbConn, dC.GetUpdateImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var imageRepository repositoryInterface.IImageActivityRepository
	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDBConnection(&dbConn, dC.GetUpdateImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDBConnection(&dbConn, dC.GetUpdateImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

	var imgActs []models.ImageActivity

	areaTests := AreasObjs[0]
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
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteImgAct(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `image_activities` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err := dbConn.Exec("DELETE FROM `area` WHERE 1")
		if err != nil {
			t.Fatal(err)
		}

		err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = dbConn.Exec("DELETE FROM `users` WHERE 1")
		if err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
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

	var imageRepository repositoryInterface.IImageActivityRepository
	if err = database.InitDBConnection(&dbConn, dC.GetUpdateImgAct(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if imageRepository, err = repositories.NewImageActivityMySQLRepository(dbConn); err != nil {
		t.Fatal(err)
	}

	testCases := ImageActivityObjs

	var imgActsIds []int

	areaTests := AreasObjs[0]
	if imgActsIds, err = imageRepository.GetImageActivityIdsByAreaId(&areaTests); err != nil {
		t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][%v]\n", err.Error())
	}

	if len(testCases) != len(imgActsIds) {
		t.Errorf("[ImageActivity][GetImageActivitiesByAreaId][len][%v]!=[%v]\n", len(testCases), len(imgActsIds))
	}
}
