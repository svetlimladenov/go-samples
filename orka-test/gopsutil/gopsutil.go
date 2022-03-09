package gopsutil

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

func test() {

	fmt.Println(host.KernelVersion())
	fmt.Println(host.KernelArch())

	fmt.Println(host.Info())

	fmt.Println(host.PlatformInformation())
}
