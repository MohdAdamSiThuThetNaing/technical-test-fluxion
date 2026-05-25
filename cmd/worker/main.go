package main

import (
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/db"
	"github.com/MohdAdamSiThuThetNaing/technical-test-fluxion/internal/queue"
)

func main() {
    db.ConnectMongo()
    queue.Consume()
}