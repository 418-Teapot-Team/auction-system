package main

import "fmt"

func main() {

	//logx, err := logrusx.New("auction-system")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//DbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
	//	os.Getenv("PG_USERNAME"),
	//	os.Getenv("PG_PASSWORD"),
	//	os.Getenv("PG_HOST"),
	//	os.Getenv("PG_PORT"),
	//	os.Getenv("DB_NAME"),
	//)
	//db, err := repository.NewDB(DbUrl)
	//if err != nil {
	//	logx.Fatal("error while trying to create NewDB",
	//		logrusx.LogField{Key: "context", Value: err},
	//	)
	//}
	//
	//repo := repository.NewAuthRepository(db)
	//
	//fmt.Println(repo)

	fmt.Println("hello")
}
