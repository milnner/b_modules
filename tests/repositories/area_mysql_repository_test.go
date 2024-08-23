package repositories

import (
	"database/sql"
	"strings"
	"testing"

	"github.com/milnner/b_modules/database"
	"github.com/milnner/b_modules/models"
	"github.com/milnner/b_modules/repositories"
	repoInterfaces "github.com/milnner/b_modules/repositories/interfaces"
	"github.com/milnner/b_modules/tests/config"
)

func TestAreaMySQLRepositoryPolimorfism(t *testing.T) {
	var _ repoInterfaces.IAreaRepository = &repositories.AreaMySQLRepository{}
}

func TestGetAreasIdsByOwnerId(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Errorf(config.DatabaseConn.User.GetInsert())
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Class.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	area := models.Area{OwnerId: config.AreasObjs[0].OwnerId}
	ids, err := repoArea.GetAreaIdsByOwnerId(&area)
	if err != nil {
		t.Errorf("[GetAreasIdsByOwnerId] %v", err)
	}
	if len(ids) == 0 {
		t.Errorf("[GetAreasIdsByOwnerId] %v", err)
	}
}

func TestGetAreaById(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Class.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range config.AreasObjs {
		area := models.Area{Id: tc.Id, OwnerId: tc.OwnerId}
		if err = repoArea.GetAreaById(&area); err != nil {
			t.Errorf("[GetAreaById] %v", err)
		}
		if area.Id != tc.Id ||
			area.OwnerId != tc.OwnerId ||
			strings.Compare(area.CreationDatetime.String()[:19], tc.CreationDatetime.String()[:19]) != 0 ||
			strings.Compare(area.Title, tc.Title) != 0 ||
			strings.Compare(area.Description, tc.Description) != 0 ||
			area.Activated != tc.Activated {
			t.Errorf("[GetAreaById] %v !=\n %v", area, tc)

		}
	}
}

func TestInsertUser_GetUsersIdsByArea_GetPermission_RemoveUser(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `user_has_area_access` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range config.UsersObjs {
		if err = repoArea.InsertUser(&config.AreasObjs[0], &tc); err != nil {
			t.Errorf("[InsertUser]%v", err)
		}
	}
	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetInsert(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	var ids []int
	if ids, err = repoArea.GetUserIdsByAreaId(&config.AreasObjs[0]); err != nil {
		t.Errorf("[GetUsersIdsByArea] %v", err)
	}
	if len(ids) != len(config.UsersObjs) {
		t.Errorf("[GetUsersIdsByArea] ids: %v, users: %v", len(ids), len(config.UsersObjs))
	}
	for _, tc := range config.UsersObjs {
		if err = repoArea.GetPermission(&config.AreasObjs[0], &tc); err != nil {
			t.Errorf("[GetPermission]%v", err)
		}
	}
	for _, tc := range config.UsersObjs {
		if err = repoArea.RemoveUser(&config.AreasObjs[0], &tc); err != nil {
			t.Errorf("[RemoveUser] %v", err)
		}
	}
}

func TestInsertArea(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Class.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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
	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}

	for _, tc := range config.AreasObjs {
		if err = repoArea.Insert(&tc); err != nil {
			t.Errorf("[InsertArea] %v", err)
		}
	}

}

func TestUpdateArea(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Class.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range config.AreasObjs {
		area := tc
		area.Description = "asdasd"
		area.Title = "asdasd"

		if err = repoArea.Update(&area); err != nil {
			t.Errorf("[Update] %v", err)
		}

	}
}

func TestDeleteArea(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Class.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range config.AreasObjs {
		area := tc
		area.Description = "asdasd"
		area.Title = "asdasd"

		if err = repoArea.Delete(&area); err != nil {
			t.Errorf("[Delete] %v", err)
		}
	}
}

func TestGetAreasByIds(t *testing.T) {
	var dbConn *sql.DB

	err := database.InitDatabaseConn(&dbConn, config.DatabaseConn.User.GetInsert(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Class.GetDelete(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDatabaseConn(&dbConn, config.DatabaseConn.Area.GetSelect(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	var areas []models.Area
	for i := 0; i < len(config.AreasObjs); i++ {
		areas = append(areas, models.Area{Id: config.AreasObjs[i].Id, OwnerId: config.AreasObjs[i].OwnerId})
	}
	if err = repoArea.GetAreasByIds(areas); err != nil {
		t.Errorf("[GetAreasByIds], %v", err)
	}
	for _, tc := range config.AreasObjs {
		for _, area := range areas {
			if area.Id == tc.Id && (area.OwnerId != tc.OwnerId ||
				strings.Compare(area.CreationDatetime.String()[:19], tc.CreationDatetime.String()[:19]) != 0 ||
				strings.Compare(area.Title, tc.Title) != 0 ||
				strings.Compare(area.Description, tc.Description) != 0 ||
				area.Activated != tc.Activated) {
				t.Errorf("[GetAreaById] %v !=\n %v", area, tc)
			}
		}
	}
}
