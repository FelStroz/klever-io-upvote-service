package configDB

import (
	"fmt"
)

type DBSettings struct {
	MongoURI string
	MongoDB  string
}

var Db DBSettings

func init() {
	//mongoURI := fmt.Sprintf("mongodb://%v:%v", "localhost", "27017")
	mongoURI := "mongodb+srv://eu:1234qwer@cluster0.ogm0x.mongodb.net/cryptos?retryWrites=true&w=majority"
	fmt.Sprintln("")
	Db = DBSettings{
		MongoURI: mongoURI,
		MongoDB:  "cryptos",
	}
}
