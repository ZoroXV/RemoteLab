package monitor

import (
	"fmt"
	"log"
	"regexp"
	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
	"github.com/citilinkru/libudev"
	"github.com/citilinkru/libudev/matcher"
	"golang.org/x/exp/slices"
)

const (
	DEVNAME = "DEVNAME"
	RULE_VENDOR_ID = "ID_VENDOR_ID"
	RULE_MODEL_ID = "ID_MODEL_ID"
	ARDUINO_VENDOR_ID gousb.ID = 0x2341
	STM32_VENDOR_ID gousb.ID = 0x0483
	WCH_VENDOR_ID gousb.ID = 0x1a86 //Correpond to CH340 Serial Converter
)

var (
	vendorIds = []gousb.ID{ARDUINO_VENDOR_ID, STM32_VENDOR_ID, WCH_VENDOR_ID}
)

func ListMicrocontrollers() []microControllerInfos {
	var microContInfos []microControllerInfos

	ctx := gousb.NewContext()
	defer ctx.Close()

	// List global informations
	_, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if slices.Contains(vendorIds, desc.Vendor) {
			vendor, _ := usbid.Vendors[desc.Vendor]
			product, _ := vendor.Product[desc.Product]

			microContInfos = append(microContInfos, microControllerInfos{
				VendorID:		uint16(desc.Vendor),
				ProductID:		uint16(desc.Product),
				VendorName:		vendor.String(),
				ProductName:	product.String(),
				Port:			"",
				SerialNumber:	"",
			})
		}
		
		return false
	})

	if err != nil {
		log.Fatalf("[MONITOR][LIST_CONTROLLERS][ERR] An error occured while listing usb devices.\n\t%v\n", err)
	}

	// Get port
	sc := libudev.NewScanner()
	err, udevDevices := sc.ScanDevices()

	if err != nil {
		log.Fatalf("[MONITOR][LIST_CONTROLLERS][ERR] Could not list used ports.\n\t%v\n", err)
	}

	for i, dev := range microContInfos {
		m := matcher.NewMatcher()
		m.SetStrategy(matcher.StrategyAnd)
		m.AddRule(matcher.NewRuleEnv(RULE_VENDOR_ID, fmt.Sprintf("%04x", uint16(dev.VendorID))))
		m.AddRule(matcher.NewRuleEnv(RULE_MODEL_ID, fmt.Sprintf("%04x", uint16(dev.ProductID))))
		filteredUdevDevices := m.Match(udevDevices)

		for _, fDev := range filteredUdevDevices {
			matched, err := regexp.MatchString(`tty.*`, fDev.Env[DEVNAME])
			if err == nil && matched {
				microContInfos[i].Port = fmt.Sprintf("/dev/%s", fDev.Env[DEVNAME])
				break
			}
		}
	}

	return microContInfos
}
