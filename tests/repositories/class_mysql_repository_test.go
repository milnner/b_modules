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

func TestClassPolimorfism(t *testing.T) {
	var _ repoInterfaces.IClassRepository = &repositories.ClassMySQLRepository{}
}
func TestInsertClass(t *testing.T) {
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

		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
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

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}

	classRepository, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
	}

	testCases := ClassesObjs
	for _, tC := range testCases {

		if err := classRepository.Insert(&tC); err != nil {
			t.Errorf("[TestInsert] %v\n", err)
		}
	}
}

func TestUpdateClass(t *testing.T) {
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

		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
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

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}

	classRepository, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
	}

	testCases := ClassesObjs
	for _, tC := range testCases {
		tC.Description = "Olss"
		if err := classRepository.Update(&tC); err != nil {
			t.Errorf("[TestUpdate] %v\n", err)
		}
	}
}

func TestDeleteClass(t *testing.T) {
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

		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
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

	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}

	classRepository, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Fatal(err)
	}

	testCases := ClassesObjs
	for _, tC := range testCases {
		tC.Description = "Olss"
		if err := classRepository.Delete(&tC); err != nil {
			t.Errorf("[TestDelete] %v\n", err)
		}
	}
}
func TestGetClassById(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Error(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Error(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Error(err)
		}
	}()

	if err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Error(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Error(err)
		}
	}

	if err := database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Classes); i++ {
		_, err = dbConn.Exec(Classes[i])
		if err != nil {
			t.Error(err)
		}
	}

	testCases := ClassesObjs

	repo, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Error(err)
	}

	var class models.Class
	for _, tC := range testCases {
		class = tC

		if err = repo.GetClassById(&class); err != nil ||
			tC.Id != class.Id ||
			tC.AreaId != class.AreaId ||
			strings.Compare(tC.CreationDate.String()[:19], class.CreationDate.String()[:19]) != 0 ||
			strings.Compare(tC.LastUpdate.String()[:19], class.LastUpdate.String()[:19]) != 0 ||
			strings.Compare(tC.Title, class.Title) != 0 ||
			strings.Compare(tC.Description, class.Description) != 0 ||
			tC.UserCreatorId != class.UserCreatorId ||
			tC.Activated != class.Activated {
			t.Errorf("[GetClassById]%v  %v\n", tC, class)
		}
	}
}

func TestClassAddContent_TestClassUpdateContentPosition_TestClassGetContent_TestClassRemoveContent(t *testing.T) {
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
		if _, err = dbConn.Exec("DELETE FROM `class_see_content` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
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
	if err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err := database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Classes); i++ {
		_, err = dbConn.Exec(Classes[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}
	classRepository, err := repositories.NewClassMySQLRepository(dbConn)

	if err != nil {
		t.Fatal(err)
	}
	class := ClassesObjs[0]
	for _, tc := range ContentObjs {
		if err = classRepository.AddContent(&class, &tc); err != nil {
			t.Errorf("[AddContent] %v", err)
		}
	}

	if err = database.InitDBConnection(&dbConn, dC.GetUpdateClass(), "mysql"); err != nil {
		t.Fatal(err)
	}

	for _, tc := range ContentObjs {
		tc.Position = 1
		if err = classRepository.UpdateContentPosition(&class, &tc); err != nil {
			t.Errorf("[UpdateContentPosition] %v", err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetSelectClass(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var contentIds []int
	if contentIds, err = classRepository.GetContentIdsById(&class); err != nil || len(contentIds) != len(ContentObjs) {
		t.Errorf("[GetContentIdsById] %v, len = %v", err, len(contentIds))
	}

	for _, tc := range ContentObjs {
		if err = classRepository.RemoveContent(&class, &tc); err != nil {
			t.Errorf("[RemoveContent] %v", err)
		}
	}

}

func TestClassAddStudent_RemoveStudent(t *testing.T) {
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
		if _, err = dbConn.Exec("DELETE FROM `class_see_content` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteContent(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `contents` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Fatal(err)
		}
		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Fatal(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
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
	if err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Fatal(err)
		}
	}

	if err := database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Classes); i++ {
		_, err = dbConn.Exec(Classes[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertContent(), "mysql"); err != nil {
		t.Fatal(err)
	}
	for i := 0; i < len(Contents); i++ {
		_, err = dbConn.Exec(Contents[i])
		if err != nil {
			t.Fatal(err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Fatal(err)
	}
	classRepository, err := repositories.NewClassMySQLRepository(dbConn)

	if err != nil {
		t.Fatal(err)
	}
	class := ClassesObjs[0]
	for _, tc := range UsersObjs {
		if err = classRepository.AddStudentUser(&class, &tc); err != nil {
			t.Errorf("[AddStudent] %v", err)
		}
	}
	if err = database.InitDBConnection(&dbConn, dC.GetSelectClass(), "mysql"); err != nil {
		t.Fatal(err)
	}
	var userIds []int
	if userIds, err = classRepository.GetStudentIdsById(&class); err != nil || len(userIds) != len(UsersObjs) {
		t.Errorf("[GetStudentIdsById] %v, len = %v", err, len(userIds))
	}

	for _, tc := range UsersObjs {
		if err = classRepository.RemoveStudentUser(&class, &tc); err != nil {
			t.Errorf("[RemoveStudent] %v", err)
		}
	}
}

func TestGetClassesByIds(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Error(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Error(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Error(err)
		}
	}()

	if err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Error(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Error(err)
		}
	}

	if err := database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Classes); i++ {
		_, err = dbConn.Exec(Classes[i])
		if err != nil {
			t.Error(err)
		}
	}

	testCases := ClassesObjs

	repo, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Error(err)
	}

	var classes []models.Class
	for _, v := range ClassesObjs {
		classes = append(classes, models.Class{Id: v.Id})
	}
	for _, tC := range testCases {
		for _, class := range classes {
			if err = repo.GetClassById(&class); tC.Id == class.Id && (err != nil ||
				tC.Id != class.Id ||
				tC.AreaId != class.AreaId ||
				strings.Compare(tC.CreationDate.String()[:19], class.CreationDate.String()[:19]) != 0 ||
				strings.Compare(tC.LastUpdate.String()[:19], class.LastUpdate.String()[:19]) != 0 ||
				strings.Compare(tC.Title, class.Title) != 0 ||
				strings.Compare(tC.Description, class.Description) != 0 ||
				tC.UserCreatorId != class.UserCreatorId ||
				tC.Activated != class.Activated) {
				t.Errorf("[GetClassById]%v  %v\n", tC, class)
			}
		}
	}
}

func TestGetClassIdsByAreaId(t *testing.T) {
	var (
		dbConn *sql.DB
		err    error
	)
	dC := environment.Environment().GetDatabaseConnections()

	defer func() {
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteClass(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `classes` WHERE 1"); err != nil {
			t.Error(err)
		}
		if err = database.InitDBConnection(&dbConn, dC.GetDeleteArea(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `area` WHERE 1"); err != nil {
			t.Error(err)
		}

		if err = database.InitDBConnection(&dbConn, dC.GetDeleteUser(), "mysql"); err != nil {
			t.Error(err)
		}

		if _, err = dbConn.Exec("DELETE FROM `users` WHERE 1"); err != nil {
			t.Error(err)
		}
	}()

	if err := database.InitDBConnection(&dbConn, dC.GetInsertUser(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Users); i++ {
		_, err = dbConn.Exec(Users[i])
		if err != nil {
			t.Error(err)
		}
	}
	if err := database.InitDBConnection(&dbConn, dC.GetInsertArea(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Area); i++ {
		_, err = dbConn.Exec(Area[i])
		if err != nil {
			t.Error(err)
		}
	}

	if err := database.InitDBConnection(&dbConn, dC.GetInsertClass(), "mysql"); err != nil {
		t.Error(err)
	}
	for i := 0; i < len(Classes); i++ {
		_, err = dbConn.Exec(Classes[i])
		if err != nil {
			t.Error(err)
		}
	}

	repo, err := repositories.NewClassMySQLRepository(dbConn)
	if err != nil {
		t.Error(err)
	}
	var classIds []int
	if classIds, err = repo.GetClassIdsByAreaId(&AreasObjs[0]); err != nil {
		t.Errorf("[GetClassIdsByAreaId] %v", err)
	}
	if len(classIds) != len(ClassesObjs) {
		t.Errorf("[GetClassIdsByAreaId][len] %v != %v", len(classIds), len(ClassesObjs))
	}
}
