package autoload

import "os"

func GetEnv() {

	os.Setenv("MYSQL_ROOT_USER", "root")
	os.Setenv("MYSQL_PORT", "3306")
	os.Setenv("MYSQL_ROOT_PASSWORD", "root")
	os.Setenv("MYSQL_DATABASE", "sky_assess")
	os.Setenv("MYSQL_HOST", "localhost")
}
