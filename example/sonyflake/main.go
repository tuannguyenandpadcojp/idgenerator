package main

import (
	"context"
	"fmt"
	"time"

	"github.com/tuannguyenandpadcojp/utils/base62"

	"github.com/tuannguyenandpadcojp/idgenerator/sonyflake"
)

func main() {
	idGen := sonyflake.NewGenerator(time.Now())

	newID, _ := idGen.NewID(context.Background())
	fmt.Printf("ID: %d\n", newID)
	fmt.Printf("ID string: %s\n", base62.Encode(newID))
}
