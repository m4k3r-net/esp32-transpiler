package worker

import (
	"github.com/andygeiss/esp32-transpiler/api/worker"
)

var (
	rules = map[string]string{
		"digital.Low":             "LOW",
		"digital.High":            "HIGH",
		"digital.ModeInput":       "INPUT",
		"digital.ModeOutput":      "OUTPUT",
		"digital.PinMode":         "pinMode",
		"digital.Write":           "digitalWrite",
		"random.Num":              "random",
		"random.NumBetween":       "random",
		"random.Seed":             "randomSeed",
		"serial.Available":        "Serial.available",
		"serial.BaudRate300":      "300",
		"serial.BaudRate600":      "600",
		"serial.BaudRate1200":     "1200",
		"serial.BaudRate2400":     "2400",
		"serial.BaudRate4800":     "4800",
		"serial.BaudRate9600":     "9600",
		"serial.BaudRate14400":    "14400",
		"serial.BaudRate28800":    "28800",
		"serial.BaudRate38400":    "38400",
		"serial.BaudRate57600":    "57600",
		"serial.BaudRate115200":   "115200",
		"serial.Begin":            "Serial.begin",
		"serial.Print":            "Serial.print",
		"serial.Println":          "Serial.println",
		"timer.Delay":             "delay",
		"wifi":                    "WiFi",
		"wifi.Begin":              "WiFi.begin",
		"wifi.BeginEncrypted":     "WiFi.begin",
		"wifi.BSSID":              "WiFi.BSSID",
		"wifi.Disconnect":         "WiFi.disconnect",
		"wifi.EncryptionType":     "WiFi.encryptionType",
		"wifi.EncryptionTypeAuto": "8",
		"wifi.EncryptionTypeCCMP": "4",
		"wifi.EncryptionTypeNone": "7",
		"wifi.EncryptionTypeTKIP": "2",
		"wifi.EncryptionTypeWEP":  "5",
		"wifi.LocalIP":            "WiFi.localIP",
		"wifi.RSSI":               "WiFi.RSSI",
		"wifi.ScanNetworks":       "WiFi.scanNetworks",
		"wifi.SetDNS":             "WiFi.setDNS",
		"wifi.SSID":               "WiFi.SSID",
		"wifi.StatusConnected":    "WL_CONNECTED",
		"wifi.StatusIdle":         "WL_IDLE",
		"Loop":                    "void loop",
		"Setup":                   "void setup",
	}
)

// Mapping specifies the api logic to apply transformation to a specific Golang identifier by reading simple JSON map.
type Mapping struct {}

// NewMapping creates a new mapping and returns its address.
func NewMapping() worker.Mapping {
	return &Mapping{}
}

// Apply checks the Golang identifier and transforms it to a specific representation.
func (*Mapping) Apply(ident string) string {
	for wanted := range rules {
		if ident == wanted {
			ident = rules[ident]
		}
	}
	return ident
}
