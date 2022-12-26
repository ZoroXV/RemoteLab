package monitor

import (
	"log"
	"github.com/google/gousb"
	"github.com/google/gousb/usbid"
)

var (
	ARDUINO_VENDOR_ID gousb.ID = 0x2341
	STM32_VENDOR_ID gousb.ID = 0x0483

	//ARDUINO_PRODUCTS_ID gousb.ID[] = [0x01, 0x10, 0x3F, 0x42, 0x43, 0x44]
	//STM32_PRODUCTS_ID gousb.ID[] = []
)

func ListMicrocontrollers() []microControllerInfos {
	var microContInfos []microControllerInfos

	ctx := gousb.NewContext()
	defer ctx.Close()

	_, err := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		if desc.Vendor == ARDUINO_VENDOR_ID || desc.Vendor == STM32_VENDOR_ID {
			microContInfos = append(microContInfos, microControllerInfos{
				VendorID:	uint16(desc.Vendor),
				ProductID:	uint16(desc.Product),
				Description:	usbid.Describe(desc),
			})
			log.Println(desc)
		}

		return false
	})


	if err != nil {
		log.Fatalf("[MONITOR][LIST_CONTROLLERS][ERR] An error occured while listing usb devices.\n\t%v\n", err)
	}

	return microContInfos
}
