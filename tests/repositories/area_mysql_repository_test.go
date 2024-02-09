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

func TestAreaMySQLRepositoryPolimorfism(t *testing.T) {
	var _ repoInterfaces.IAreaRepository = &repositories.AreaMySQLRepository{}
}

func TestGetAreasIdsByOwnerId(t *testing.T) {
	var dbConn *sql.DB
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetSelectArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	area := models.Area{OwnerId: AreasObjs[0].OwnerId}
	ids, err := repoArea.GetAreasIdsByOwnerId(&area)
	if err != nil {
		t.Errorf("[GetAreasIdsByOwnerId] %v", err)
	}
	if len(ids) == 0 {
		t.Errorf("[GetAreasIdsByOwnerId] %v", err)
	}
}

func TestGetAreaById(t *testing.T) {
	var dbConn *sql.DB
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetSelectArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range AreasObjs {
		area := models.Area{Id: tc.Id}
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
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `user_has_area_access` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range UsersObjs {
		if err = repoArea.InsertUser(&AreasObjs[0], &tc); err != nil {
			t.Errorf("[InsertUser]%v", err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	var ids []int
	if ids, err = repoArea.GetUserIdsByAreaId(&AreasObjs[0]); err != nil {
		t.Errorf("[GetUsersIdsByArea] %v", err)
	}
	if len(ids) != len(UsersObjs) {
		t.Errorf("[GetUsersIdsByArea] ids: %v, users: %v", len(ids), len(UsersObjs))
	}
	for _, tc := range UsersObjs {
		if err = repoArea.GetPermission(&AreasObjs[0], &tc); err != nil {
			t.Errorf("[GetPermission]%v", err)
		}
	}
	for _, tc := range UsersObjs {
		if err = repoArea.RemoveUser(&AreasObjs[0], &tc); err != nil {
			t.Errorf("[RemoveUser] %v", err)
		}
	}
}

func TestInsertArea(t *testing.T) {
	var dbConn *sql.DB
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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
	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}

	for _, tc := range AreasObjs {
		if err = repoArea.Insert(&tc); err != nil {
			t.Errorf("[InsertArea] %v", err)
		}
	}

}

func TestUpdateArea(t *testing.T) {
	var dbConn *sql.DB
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetSelectArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range AreasObjs {
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
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetSelectArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	for _, tc := range AreasObjs {
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
	dC := environment.Environment().GetDatabaseConnections()

	err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
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

	if err = database.InitDBConnection(&dbConn, dC.GetSelectArea(), "mysql"); err != nil {
		t.Fatal(err)
	}

	var repoArea *repositories.AreaMySQLRepository
	if repoArea, err = repositories.NewAreaMySQLRepository(dbConn); err != nil {
		t.Error(err)
	}
	var areas []models.Area
	for i := 0; i < len(AreasObjs); i++ {
		areas = append(areas, models.Area{Id: AreasObjs[i].Id})
	}
	if err = repoArea.GetAreasByIds(areas); err != nil {
		t.Errorf("[GetAreasByIds], %v", err)
	}
	for _, tc := range AreasObjs {
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
