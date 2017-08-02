package conf

import "os"

var (
	MgoAddr = os.Getenv("MONGO_PORT_27017_TCP_ADDR")
	MgoPort = os.Getenv("MONGO_PORT_27017_TCP_PORT")
)
