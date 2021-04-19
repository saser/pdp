package server

import (
	"github.com/Saser/pdp/aip/resource"
	"github.com/Saser/pdp/aip/resourcename"

	taskspb "github.com/Saser/pdp/tasks2/tasks_go_proto"
)

var (
	taskPattern  *resourcename.Pattern
	eventPattern *resourcename.Pattern
)

func init() {
	taskRD := resource.DescriptorOf(&taskspb.Task{})
	taskPattern = resourcename.MustCompile(taskRD.GetPattern()[0])

	eventRD := resource.DescriptorOf(&taskspb.Event{})
	eventPattern = resourcename.MustCompile(eventRD.GetPattern()[0])
}
