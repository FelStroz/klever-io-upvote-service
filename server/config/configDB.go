package configDB

import (
	"fmt"
)

type DBSettings struct {
	MongoURI string
	MongoDB  string
}

var db DBSettings

func init() {
	mongoURI := fmt.Sprintf("mongodb://%v:%v", "localhost", "27017")
	db = DBSettings{
		MongoURI: mongoURI,
		MongoDB:  "cryptos",
	}
}
