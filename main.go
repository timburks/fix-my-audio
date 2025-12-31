package main

import (
	"encoding/json"
	"errors"
	"log"
	"os/exec"
	"strconv"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%s", err)
	}
}

func run() error {
	devices, err := getOutputDevices()
	if err != nil {
		return err
	}
	output := findOutputDevice(devices, "Digital Stereo (HDMI 2)")
	if output == -1 {
		return errors.New("device not found")
	}
	if err = setOutputDevice(output); err != nil {
		return err
	}
	return nil
}

type Device struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Info struct {
		Props map[string]interface{} `json:"props"`
	} `json:"info"`
}

func getOutputDevices() ([]Device, error) {
	cmd := exec.Command("pw-dump")
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var devices []Device
	err = json.Unmarshal(out, &devices)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func findOutputDevice(devices []Device, description string) int {
	for _, device := range devices {
		d, ok := device.Info.Props["device.profile.description"].(string)
		if ok && d == description {
			return device.Id
		}
	}
	return -1
}

func setOutputDevice(id int) error {
	cmd := exec.Command("wpctl", "set-default", strconv.Itoa(id))
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}
