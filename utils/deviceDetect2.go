package utils

import ua "github.com/mileusna/useragent"

func GetDeviceType(detect ua.UserAgent) string {
	var device string
	if detect.Mobile {
		device = "Mobile"
	}
	if detect.Tablet {
		device = "Tablet"
	}
	if detect.Desktop {
		device = "Desktop"
	}
	if detect.Bot {
		device = "Bot"
	}

	if detect.OS == "" {
		detect.OS = "null"
	}
	if device == "" {
		device = "null"
	}
	return device
}
