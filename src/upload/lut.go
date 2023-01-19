package upload

import (
	"errors"
	"fmt"
	"github.com/google/gousb"
)

type VendorProduct struct {
	Vendor	uint16
	Product	uint16
}

const (
	STM32_VENDOR_ID gousb.ID = 0x0483
)

var (
	LutVendorProductFqbn = map[VendorProduct][]string{
		VendorProduct{0x2341, 0x0042}: []string{"arduino:avr:mega"},
		VendorProduct{0x2341, 0x0043}: []string{"arduino:avr:uno"},
		VendorProduct{0x1a86, 0x7523}: []string{"arduino:avr:mega", "arduino:avr:nano", "arduino:avr:uno"},
		VendorProduct{0x0403, 0x6001}: []string{"arduino:avr:mega", "arduino:avr:nano", "arduino:avr:uno"},

	}
)

func GetFqbn(VendorID uint16, ProductID uint16) ([]string, error) {
	if VendorID == uint16(STM32_VENDOR_ID) {
		return []string{}, nil
	}

	val, exist := LutVendorProductFqbn[VendorProduct{VendorID, ProductID}]

	if exist {
		return val, nil
	} else {
		return []string{}, errors.New(fmt.Sprintf("{0x%04x, 0x%04x} is an unknown {VendorID, ProductID} pair.", VendorID, ProductID))
	}
}
