package environment

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	errapp "github.com/milnner/b_modules/errors"
)

var Environment *environment

func CreateEnvironment() {
	Environment = &environment{}
}

type environment struct {
	jwtSecretKey *string
	logPath      *string
	certFile     string
	keyFile      string
	addr         string
	debug        bool
}

func (u *environment) InitEnvironment(logPathEnvVariable, jwtSecretKeyEnvVariable, certFile, keyFile, addr string, debug bool) {
	u.addr = addr
	if addr == "" {
		panic("undefined addr")
	}
	u.debug = debug

	u.logPath = new(string)
	*(u.logPath) = os.Getenv(logPathEnvVariable)

	if *(u.logPath) == "" {
		panic(errapp.NewNotExistEnvironmentVariableError(jwtSecretKeyEnvVariable))
	} else if _, err := os.Stat(*(u.logPath)); os.IsNotExist(err) {
		panic(err)
	}

	u.jwtSecretKey = new(string)
	*(u.jwtSecretKey) = os.Getenv(jwtSecretKeyEnvVariable)

	if *(u.jwtSecretKey) == "" {
		panic(errapp.NewNotExistEnvironmentVariableError(jwtSecretKeyEnvVariable))
	}
	if !u.debug {
		u.certFile = certFile
		if u.certFile == "" {
			panic("Certificate file is required for HTTPS")
		}

		u.keyFile = keyFile
		if u.keyFile == "" {
			panic("Key file is required for HTTPS")

		}
	}

}

func (u *environment) GetAddr() string {
	return u.addr
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
