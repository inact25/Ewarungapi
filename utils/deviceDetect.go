package utils

import (
	"fmt"
	ua "github.com/mileusna/useragent"
)

func DeviceDetect(detect ua.UserAgent) string {
	device := GetDeviceType(detect)
	result := fmt.Sprintf("%v %v, OS : %v %v, Device : %v", detect.Name, detect.Version, detect.OS, detect.OSVersion, device)
	return result
}
