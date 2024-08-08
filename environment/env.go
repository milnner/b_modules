package environment

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	errapp "github.com/milnner/b_modules/errors"
)

var env *environment

func CreateEnvironment() {
	env = &environment{}
}

type environment struct {
	jwtSecretKey *string
	logPath      *string
	certFile     string
	keyFile      string
	addr         string
	debug        bool
}

type environmentVariables struct {
	logPathEnvVariable      string
	jwtSecretKeyEnvVariable string
}
type httpFiles struct {
	certFile string
	keyFile  string
}

func NewEnvironmentVariables(logPath string, jwtSecretKeyEnvVariable string) *environmentVariables {
	return &environmentVariables{logPathEnvVariable: logPath, jwtSecretKeyEnvVariable: jwtSecretKeyEnvVariable}
}

func NewHttpFiles(certFile string, keyFile string) *httpFiles {
	return &httpFiles{certFile: certFile, keyFile: keyFile}
}

func InitEnvironment(eV environmentVariables, hF httpFiles, addr string, debug bool) error {

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
