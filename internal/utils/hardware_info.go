package utils

import (
	"fmt"
	"strings"

	"github.com/StackExchange/wmi"
)

type Win32_BaseBoard struct {
	Manufacturer string
	SerialNumber string
}

type Win32_BIOS struct {
	Manufacturer string
	SerialNumber string
}

type Win32_VideoController struct {
	PNPDeviceID string
}

type Win32_Processor struct {
	Manufacturer string
	ProcessorId  string
	Name         string
}

type Win32_LogicalDisk struct {
	DeviceID           string
	VolumeSerialNumber string
}

func queryWMI(query interface{}, class string) error {
	return wmi.Query(fmt.Sprintf("SELECT * FROM %s", class), query)
}

func GetHardwareInfo() (string, error) {
	var sb strings.Builder

	var boards []Win32_BaseBoard
	if err := queryWMI(&boards, "Win32_BaseBoard"); err != nil {
		return "", nil
	}
	if len(boards) > 0 {
		sb.WriteString(boards[0].Manufacturer + ";")
		sb.WriteString(boards[0].SerialNumber + ";")
	}

	var bios []Win32_BIOS
	if err := queryWMI(&bios, "Win32_BIOS"); err != nil {
		return "", nil
	}
	if len(bios) > 0 {
		sb.WriteString(bios[0].Manufacturer + ";")
		sb.WriteString(bios[0].SerialNumber + ";")
	}

	var disks []Win32_LogicalDisk
	if err := queryWMI(&disks, "Win32_LogicalDisk"); err != nil {
		return "", nil
	}
	for _, d := range disks {
		if d.DeviceID == "C:" {
			sb.WriteString(d.VolumeSerialNumber + ";")
		}
	}

	var video []Win32_VideoController
	if err := queryWMI(&video, "Win32_VideoController"); err != nil {
		return "", nil
	}
	if len(video) > 0 {
		sb.WriteString(video[0].PNPDeviceID + ";")
	}

	var cpu []Win32_Processor
	if err := queryWMI(&cpu, "Win32_Processor"); err != nil {
		return "", nil
	}
	if len(cpu) > 0 {
		sb.WriteString(cpu[0].Manufacturer + ";")
		sb.WriteString(cpu[0].ProcessorId + ";")
		sb.WriteString(cpu[0].Name + ";")
	}

	return sb.String(), nil
}
