package repositories

import (
	"net"
	"reflect"
	"time"

	"github.com/milnner/b_modules/tests/config"
)

func AreSlicesEqual(slice1, slice2 []int) bool {
	// Converte os slices para mapas
	map1 := make(map[int]int)
	map2 := make(map[int]int)

	for _, v := range slice1 {
		map1[v]++
	}

	for _, v := range slice2 {
		map2[v]++
	}

	// Compara os mapas resultantes
	return reflect.DeepEqual(map1, map2)
}

func init() {
	config.SetDBData()
	config.SetRootDatabaseConn()
	port := "3306"

	target := "127.0.0.1:" + port
	conn, err := net.DialTimeout("tcp", target, 10*time.Second)
	if err != nil {
		panic(err)
	}
	conn.Close()
}
