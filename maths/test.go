package main

import (
	"github.com/NubeIO/plugin-example/maths/add"
	"github.com/NubeIO/plugin-example/maths/subtract"
	"github.com/NubeIO/rxlib/libs/pluginlib"
)

var name = "math-plugin"

func main() {
	factory := pluginlib.New(name)
	factory.AddPallet("add", add.New)
	factory.AddPallet("subtract", subtract.New)
	factory.Register(name)
}
