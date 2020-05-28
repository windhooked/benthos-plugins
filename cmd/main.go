package main

import (
	"github.com/Jeffail/benthos/v3/lib/service"

	// Add your plugin packages here
//	_ "github.com/windhooked/benthos-plugins/input"
	_ "github.com/windhooked/benthos-plugins/output"
//	_ "github.com/windhooked/benthos-plugins/processor"
)

func main() {
	service.Run()
}
