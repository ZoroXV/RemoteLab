package upload

type VendorProduct struct {
	Vendor	uint16
	Product	uint16
}

var (
	LutVendorProductFqbn = map[VendorProduct]string{
		VendorProduct{0x2341, 0x0042}: "arduino:avr:mega",
		VendorProduct{0x2341, 0x0043}: "arduino:avr:uno",
	}
)
