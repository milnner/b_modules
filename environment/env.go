package environment

import (
	"fmt"
	"os"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	errapp "github.com/milnner/b_modules/errors"
)

var env *environment

func CreateEnvironment() {
	env = &environment{}
}

type environment struct {
	jwtSecretKey  *string
	logPath       *string
	certFile      string
	keyFile       string
	addr          string
	debug         bool
	dbConnections databaseConnections
}

type environmentVariables struct {
	logPathEnvVariable      string
	jwtSecretKeyEnvVariable string
}
type httpFiles struct {
	certFile string
	keyFile  string
}

type databaseConnections struct {
	// Class database connection strings
	selectClass string
	insertClass string
	deleteClass string
	updateClass string

	// User database connection strings
	insertUser string
	selectUser string
	deleteUser string
	updateUser string

	// Content database connection strings
	insertContent string
	selectContent string
	deleteContent string
	updateContent string

	// Area database connection
	insertArea string
	selectArea string
	deleteArea string
	updateArea string

	// Image Activities
	insertImgAct string
	selectImgAct string
	deleteImgAct string
	updateImgAct string

	// Text Activities
	insertTxtAct string
	selectTxtAct string
	deleteTxtAct string
	updateTxtAct string

	// One Question N Answer Activity
	insertOneQuestionNAnswerActivity string
	selectOneQuestionNAnswerActivity string
	deleteOneQuestionNAnswerActivity string
	updateOneQuestionNAnswerActivity string

	// Answer N To One Activity
	insertAnswerNToOneActivity string
	selectAnswerNToOneActivity string
	deleteAnswerNToOneActivity string
	updateAnswerNToOneActivity string
}

func NewEnvironmentVariables(logPath string, jwtSecretKeyEnvVariable string) *environmentVariables {
	return &environmentVariables{logPathEnvVariable: logPath, jwtSecretKeyEnvVariable: jwtSecretKeyEnvVariable}
}

func NewHttpFiles(certFile string, keyFile string) *httpFiles {
	return &httpFiles{certFile: certFile, keyFile: keyFile}
}

func NewDatabaseConnections() *databaseConnections {
	return &databaseConnections{}
}

func InitEnvironment(eV environmentVariables, hF httpFiles, addr string, debug bool, dC databaseConnections) error {

	if env == nil {
		CreateEnvironment()
	}

	env.addr = addr
	if addr == "" {
		return fmt.Errorf("undefined address")
	}

	env.debug = debug
	env.InitLogPath(eV.logPathEnvVariable)
	env.InitJWTSecretKey(eV.jwtSecretKeyEnvVariable)

	if !env.debug {
		env.InitHTTPS(hF)
	}

	env.InitDatabaseConnections(&dC)
	return nil
}

func (u *environment) InitLogPath(logPathEnvVariable string) error {
	if env == nil {
		CreateEnvironment()
	}

	u.logPath = new(string)
	*(u.logPath) = os.Getenv(logPathEnvVariable)

	if *(u.logPath) == "" {
		return errapp.NewNotExistEnvironmentVariableError(logPathEnvVariable)
	}

	fi, err := os.Stat(*(u.logPath))
	if err != nil {
		return err
	}
	if !fi.IsDir() {
		return errapp.NewNotExistEnvironmentVariableError(logPathEnvVariable)
	}
	return nil
}

func (u *environment) InitJWTSecretKey(jwtSecretKeyEnvVariable string) error {
	if env == nil {
		CreateEnvironment()
	}
	u.jwtSecretKey = new(string)
	*(u.jwtSecretKey) = os.Getenv(jwtSecretKeyEnvVariable)

	if *(u.jwtSecretKey) == "" {
		return errapp.NewNotExistEnvironmentVariableError(jwtSecretKeyEnvVariable)
	}
	return nil
}

func (u *environment) InitHTTPS(hF httpFiles) error {
	if env == nil {
		CreateEnvironment()
	}
	u.certFile = hF.certFile
	if u.certFile == "" {
		return fmt.Errorf("certificate file is required for HTTPS")
	}

	u.keyFile = hF.keyFile
	if u.keyFile == "" {
		return fmt.Errorf("key file is required for HTTPS")
	}
	return nil
}
func (u *environment) InitDatabaseConnections(
	dC *databaseConnections,
) error {
	if env == nil {
		CreateEnvironment()
	}
	u.dbConnections = *dC
	var result string
	val := reflect.ValueOf(env.dbConnections)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Verifica se o campo é uma string e se está vazio
		if field.Kind() == reflect.String && field.Len() == 0 {
			result += fieldName + " "
		}
	}
	if len(result) > 0 {
		return errapp.NewUnreachableDatabaseStringsError(result)
	}
	return nil
}

func (u *databaseConnections) SetDBConnectionString() error {

	// a implementar ainda

	// check strings
	var result string
	val := reflect.ValueOf(env.dbConnections)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := val.Type().Field(i).Name

		// Verifica se o campo é uma string e se está vazio
		if field.Kind() == reflect.String && field.Len() == 0 {
			result += fieldName + " "
		}
	}
	if len(result) > 0 {
		return errapp.NewUnreachableDatabaseStringsError(result)
	}
	return nil
}

func Environment() *environment {
	if env == nil {
		CreateEnvironment()
	}
	return env
}

func (u *environment) GetAddr() string {
	return u.addr
}

func (u *environment) SetAddr(addr string) {
	u.addr = addr
}

func (u *environment) GetJWTSecretKey() []byte {
	return []byte(*(u.jwtSecretKey))
}

func (u *environment) GetLogPath() string {
	return *(u.logPath)
}

func (u *environment) GetCertFile() string {
	return u.certFile
}

func (u *environment) GetKeyFile() string {
	return u.keyFile
}

func (u *environment) IsDebug() bool {
	return u.debug
}

func (u *environment) SetDebug(debug bool) {
	u.debug = debug
}

func (u *environment) GetDatabaseConnections() *databaseConnections {
	return &u.dbConnections
}

func (u *databaseConnections) GetSelectClass() string {
	return u.selectClass
}

func (u *databaseConnections) GetInsertClass() string {
	return u.insertClass
}

func (u *databaseConnections) GetDeleteClass() string {
	return u.deleteClass
}

func (u *databaseConnections) GetUpdateClass() string {
	return u.updateClass
}

func (u *databaseConnections) SetSelectClass(connString string) {
	u.selectClass = connString
}

func (u *databaseConnections) SetInsertClass(connString string) {
	u.insertClass = connString
}

func (u *databaseConnections) SetDeleteClass(connString string) {
	u.deleteClass = connString
}

func (u *databaseConnections) SetUpdateClass(connString string) {
	u.updateClass = connString
}

func (u *databaseConnections) GetSelectUser() string {
	return u.selectUser
}

func (u *databaseConnections) GetInsertUser() string {
	return u.insertUser
}

func (u *databaseConnections) GetDeleteUser() string {
	return u.deleteUser
}

func (u *databaseConnections) GetUpdateUser() string {
	return u.updateUser
}

func (u *databaseConnections) SetSelectUser(connString string) {
	u.selectUser = connString
}

func (u *databaseConnections) SetInsertUser(connString string) {
	u.insertUser = connString
}

func (u *databaseConnections) SetDeleteUser(connString string) {
	u.deleteUser = connString
}

func (u *databaseConnections) SetUpdateUser(connString string) {
	u.updateUser = connString
}

func (u *databaseConnections) GetInsertContent() string {
	return u.insertContent
}

func (u *databaseConnections) GetSelectContent() string {
	return u.selectContent
}

func (u *databaseConnections) GetDeleteContent() string {
	return u.deleteContent
}

func (u *databaseConnections) GetUpdateContent() string {
	return u.updateContent
}

func (u *databaseConnections) SetSelectContent(connString string) {
	u.selectContent = connString
}

func (u *databaseConnections) SetInsertContent(connString string) {
	u.insertContent = connString
}

func (u *databaseConnections) SetDeleteContent(connString string) {
	u.deleteContent = connString
}

func (u *databaseConnections) SetUpdateContent(connString string) {
	u.updateContent = connString
}

func (u *databaseConnections) GetInsertArea() string {
	return u.insertArea
}

func (u *databaseConnections) GetSelectArea() string {
	return u.selectArea
}

func (u *databaseConnections) GetDeleteArea() string {
	return u.deleteArea
}

func (u *databaseConnections) GetUpdateArea() string {
	return u.updateArea
}

func (db *databaseConnections) SetInsertArea(area string) {
	db.insertArea = area
}

func (db *databaseConnections) SetSelectArea(area string) {
	db.selectArea = area
}

func (db *databaseConnections) SetDeleteArea(area string) {
	db.deleteArea = area
}

func (db *databaseConnections) SetUpdateArea(area string) {
	db.updateArea = area
}

func (db *databaseConnections) GetInserImgAct() string {
	return db.insertImgAct
}
func (db *databaseConnections) GetSelectImgAct() string {
	return db.selectImgAct
}
func (db *databaseConnections) GetDeleteImgAct() string {
	return db.deleteImgAct
}
func (db *databaseConnections) GetUpdateImgAct() string {
	return db.updateImgAct
}

func (db *databaseConnections) SetInsertImgAct(imgAct string) {
	db.insertImgAct = imgAct
}
func (db *databaseConnections) SetSelectImgAct(imgAct string) {
	db.selectImgAct = imgAct
}
func (db *databaseConnections) SetDeleteImgAct(imgAct string) {
	db.deleteImgAct = imgAct
}
func (db *databaseConnections) SetUpdateImgAct(imgAct string) {
	db.updateImgAct = imgAct
}

func (db *databaseConnections) GetInsertTxtAct() string {
	return db.insertTxtAct
}
func (db *databaseConnections) GetSelectTxtAct() string {
	return db.selectTxtAct
}
func (db *databaseConnections) GetDeleteTxtAct() string {
	return db.deleteTxtAct
}
func (db *databaseConnections) GetUpdateTxtAct() string {
	return db.updateTxtAct
}

func (db *databaseConnections) SetInsertTxtAct(txtAct string) {
	db.insertTxtAct = txtAct
}
func (db *databaseConnections) SetSelectTxtAct(txtAct string) {
	db.selectTxtAct = txtAct
}
func (db *databaseConnections) SetDeleteTxtAct(txtAct string) {
	db.deleteTxtAct = txtAct
}
func (db *databaseConnections) SetUpdateTxtAct(txtAct string) {
	db.updateTxtAct = txtAct
}

func (db *databaseConnections) GetInsertOneQuestionNAnswerActivity() string {
	return db.insertOneQuestionNAnswerActivity
}
func (db *databaseConnections) GetSelectOneQuestionNAnswerActivity() string {
	return db.selectOneQuestionNAnswerActivity
}
func (db *databaseConnections) GetDeleteOneQuestionNAnswerActivity() string {
	return db.deleteOneQuestionNAnswerActivity
}
func (db *databaseConnections) GetUpdateOneQuestionNAnswerActivity() string {
	return db.updateOneQuestionNAnswerActivity
}

func (db *databaseConnections) SetInsertOneQuestionNAnswerActivity(oneQuestionNAnswerActivity string) {
	db.insertOneQuestionNAnswerActivity = oneQuestionNAnswerActivity
}
func (db *databaseConnections) SetSelectOneQuestionNAnswerActivity(oneQuestionNAnswerActivity string) {
	db.selectOneQuestionNAnswerActivity = oneQuestionNAnswerActivity
}
func (db *databaseConnections) SetDeleteOneQuestionNAnswerActivity(oneQuestionNAnswerActivity string) {
	db.deleteOneQuestionNAnswerActivity = oneQuestionNAnswerActivity
}
func (db *databaseConnections) SetUpdateOneQuestionNAnswerActivity(oneQuestionNAnswerActivity string) {
	db.updateOneQuestionNAnswerActivity = oneQuestionNAnswerActivity
}

func (db *databaseConnections) GetInsertAnswerNToOneActivity() string {
	return db.insertAnswerNToOneActivity
}
func (db *databaseConnections) GetSelectAnswerNToOneActivity() string {
	return db.selectAnswerNToOneActivity
}
func (db *databaseConnections) GetDeleteAnswerNToOneActivity() string {
	return db.deleteAnswerNToOneActivity
}
func (db *databaseConnections) GetUpdateAnswerNToOneActivity() string {
	return db.updateAnswerNToOneActivity
}

func (db *databaseConnections) SetInsertAnswerNToOneActivity(insertAnswerNToOneActivity string) {
	db.insertAnswerNToOneActivity = insertAnswerNToOneActivity
}
func (db *databaseConnections) SetSelectAnswerNToOneActivity(selectAnswerNToOneActivity string) {
	db.selectAnswerNToOneActivity = selectAnswerNToOneActivity
}
func (db *databaseConnections) SetDeleteAnswerNToOneActivity(deleteAnswerNToOneActivity string) {
	db.deleteAnswerNToOneActivity = deleteAnswerNToOneActivity
}
func (db *databaseConnections) SetUpdateAnswerNToOneActivity(updateAnswerNToOneActivity string) {
	db.updateAnswerNToOneActivity = updateAnswerNToOneActivity
}
func (u *databaseConnections) SetRootConnString(connString string) {
	u.SetSelectClass(connString)
	u.SetInsertClass(connString)
	u.SetDeleteClass(connString)
	u.SetUpdateClass(connString)

	u.SetSelectUser(connString)
	u.SetInsertUser(connString)
	u.SetDeleteUser(connString)
	u.SetUpdateUser(connString)

	u.SetSelectContent(connString)
	u.SetInsertContent(connString)
	u.SetDeleteContent(connString)
	u.SetUpdateContent(connString)

	u.SetSelectArea(connString)
	u.SetInsertArea(connString)
	u.SetDeleteArea(connString)
	u.SetUpdateArea(connString)

	u.SetInsertImgAct(connString)
	u.SetSelectImgAct(connString)
	u.SetDeleteImgAct(connString)
	u.SetUpdateImgAct(connString)

	u.SetInsertTxtAct(connString)
	u.SetSelectTxtAct(connString)
	u.SetDeleteTxtAct(connString)
	u.SetUpdateTxtAct(connString)

	u.SetInsertOneQuestionNAnswerActivity(connString)
	u.SetSelectOneQuestionNAnswerActivity(connString)
	u.SetDeleteOneQuestionNAnswerActivity(connString)
	u.SetUpdateOneQuestionNAnswerActivity(connString)

	u.SetInsertAnswerNToOneActivity(connString)
	u.SetSelectAnswerNToOneActivity(connString)
	u.SetDeleteAnswerNToOneActivity(connString)
	u.SetUpdateAnswerNToOneActivity(connString)
}
