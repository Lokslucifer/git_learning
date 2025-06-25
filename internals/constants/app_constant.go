package constants

import "os"

const (
	AppName     = "MyApp"
	Version     = "v1.0.0"
	DefaultPort = "0.0.0.0:8081"

	//Key for JWT Claim
	ClaimPrimaryKey = "id"

	SuccessMessage = "success"

	//Dburl
	// DBURL = "user=postgres password=yourpassword dbname=wetransfer host=db port=5432 sslmode=disable"
)

func DBURL() string {
	return "user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"
}
