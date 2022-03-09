package machineId

import (
	"fmt"
	"log"

	"github.com/denisbrodbeck/machineid"
)

func test() {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)
}
