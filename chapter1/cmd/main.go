package main

import (
	"encoding/json"
	"fmt"
)

const response = `
{
  "SPSoftwareDataType" : [
    {
      "_name" : "os_overview",
      "boot_mode" : "normal_boot",
      "boot_volume" : "Macintosh HD",
      "kernel_version" : "Darwin 21.2.0",
      "local_host_name" : "Svetlin's MacBook Pro",
      "os_version" : "macOS 12.1 (21C52)",
      "secure_vm" : "secure_vm_enabled",
      "system_integrity" : "integrity_enabled",
      "uptime" : "up 20:22:28:31",
      "user_name" : "Svetlin Mladenov (svetlinmladenov)"
    }
  ]
}
`

type SystemProfiler[T any] struct{}

type Data interface {
	Deserialize()
}

type Software struct {
	KernelVersion string `json:"kernel_version,omitempty"`
	OSVersion     string `json:"os_version,omitempty"`
}

func (s Software) Deserialize() {

}

type IBridge struct {
	BootUUID string `json:"ibridge_boot_uuid"`
}

type Hardware struct {
	PlatformUUID string `json:"platform_UUID"`
}

func main() {
	// PrintAnything("da")
	sp := SystemProfiler[Software]{}
	s, _ := sp.GetFirst("SPSoftwareDataType")
	fmt.Println(s.KernelVersion)
}

func PrintAnything[T any](thing T) {
	fmt.Println(thing)
}

func (sp SystemProfiler[T]) Get(cmd string) ([]T, error) {
	var data = make(map[string]json.RawMessage)

	json.Unmarshal([]byte(response), &data)

	var t []T

	for _, msg := range data {
		if err := json.Unmarshal(msg, &t); err != nil {
			return t, err
		}
		return t, nil
	}

	return t, nil
}

func (sp SystemProfiler[T]) GetFirst(cmd string) (T, error) {
	all, err := sp.Get(cmd)

	if err != nil {
		return all[0], err
	}

	return all[0], err
}
