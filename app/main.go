package main
import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/aws/aws-lambda-go/lambda"
)
func Handler(ctx context.Context) {
	DBMS := "mysql"
	USER := "docker"
	PASS := "docker"
	PROTOCOL := "tcp(db:3306)"
	DBNAME := "test_database"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	conn, err := sql.Open(DBMS,  CONNECT)
	defer conn.Close()
	if err != nil {
		fmt.Println("Fail to connect db" + err.Error())
	}
	// 接続確認
	err = conn.Ping()
	if err != nil {
		fmt.Println("Failed to connect rds : %s", err.Error())
	} else {
		fmt.Println("Success to connect rds")
	}
	// 取得するレコード一行のデータ形式を構造体で定義する
	type UserData struct {
		UserID int
		FirstName string
		LastName string
		Email string
	}
	// DBからレコードを抽出
	rows, err := conn.Query("select user_id, first_name, last_name, email from user;")
	if err != nil {
		fmt.Println("Fail to query from db " + err.Error())
	}
	// データを構造体へ変換
	var UserDatas []UserData
	for rows.Next() {
		var tmpUserData UserData
		err := rows.Scan(&tmpUserData.UserID, &tmpUserData.FirstName, &tmpUserData.LastName, &tmpUserData.Email)
		if err != nil {
			fmt.Println("Fail to scan records " + err.Error())
		}
		UserDatas = append(UserDatas, UserData{
			UserID:    tmpUserData.UserID,
			FirstName: tmpUserData.FirstName,
			LastName:  tmpUserData.LastName,
			Email:     tmpUserData.Email,
		})
	}
	// 確認のための出力
	for _, userData := range UserDatas {
		fmt.Printf("%#v\n", userData)
	}
}
func main() {
	lambda.Start(Handler)
}