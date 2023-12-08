//

package main

import (
	"fmt"

	"go.temporal.io/sdk/contrib/opentelemetry"
	tempTally "go.temporal.io/sdk/contrib/tally"
	"go.temporal.io/server/common/codec"
)

func _use_temporal() {
	prop := opentelemetry.TracerOptions{}
	fmt.Printf("prop: %v\n", prop)

	cs := tempTally.PrometheusSanitizeOptions.NameCharacters
	fmt.Printf("cs: %v\n", cs)

	cc := codec.NewJSONPBEncoder()
	fmt.Printf("cc: %v\n", cc)
}

func main() {
	fmt.Println("ohai")
}
