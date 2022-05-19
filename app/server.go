package main
import (
    "fmt"
    "net/http"

    "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
    "os"

	"github.com/joho/godotenv"
)

// Item is struct for DynamoDB
type Item struct {
	MyHashKey  string
	MyRangeKey int
	MyText     string
}

// 本来はenvから取得した方が良い
// const AWS_REGION = "ap-northeast-1"
// const DYNAMO_ENDPOINT = "http://192.168.64.5/"

func main() {
    http.HandleFunc("/", handler)
    http.HandleFunc("/dynamo_input", handlerDynamoInput)
    http.HandleFunc("/dynamo_feed_all", handlerDynamoFeedAll)
    http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Your url path is %s", r.URL.Path[1:])
}

func handlerDynamoFeedAll(w http.ResponseWriter, r *http.Request) {
}


func handlerDynamoInput(w http.ResponseWriter, r *http.Request) {


    disableSsl := false

	// DynamoDB Localを利用する場合はEndpointのURLを設定する
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	// デフォルトでは東京リージョンを指定
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("MyFirstTable")

    //////////////////////
	// 単純なCRUD - Create
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		MyText:     "My First Text",
	}

	err := table.Put(item).Run()
	if err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
        return
	}

    fmt.Fprintf("done input")
}    