package environment

import (
	"fmt"
)

var env *environment

type environment struct {
	jwtSecretKey string
	certFile     string
	keyFile      string
	addr         string
	debug        bool
}

func GetEnvironment() *environment {
	if env == nil {
		env = &environment{}
	}
	return env
}
func (u *environment) SetJWTSecretKey(jwtSecretKey string) error {
	u.jwtSecretKey = jwtSecretKey

	if u.jwtSecretKey == "" {
		return fmt.Errorf("no jwt secret key")
	}
	return nil
}

func (u *environment) SetHTTPS(certFile string, keyFile string) error {
	u.certFile = certFile
	if u.certFile == "" {
		return fmt.Errorf("certificate file is required f or HTTPS")
	}

	u.keyFile = keyFile
	if u.keyFile == "" {
		return fmt.Errorf("key file is required for HTTPS")
	}
	return nil
}

func (u *environment) SetAddr(addr string) {
	u.addr = addr
}

func (u *environment) SetDebug(debug bool) {
	u.debug = debug
}

func (u *environment) GetAddr() string {
	return u.addr
}

func (u *environment) GetJWTSecretKey() string {
	return u.jwtSecretKey
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
