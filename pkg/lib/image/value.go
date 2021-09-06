package image

import (
	"fmt"
	"strconv"
)

// for example: 11010001
type ColorValue struct {
	OriginDecimalValue  uint8  // 209
	OriginBinaryValue   string // 11010001
	CompressBinaryValue string // 11010000
	GroupDecimalValue   uint8  // 6
	GroupBinaryValue    string // 110
	RepairBinaryValue   string // 001
	NewValue            uint8
	isChange            bool
}

func NewColorValue(v uint8) *ColorValue {
	binary := fmt.Sprintf("%08s", parseBinary(v))
	compressBinary := fmt.Sprintf("%s000", binary[:5])
	group := parseUint(compressBinary) / interval
	groupBinary := fmt.Sprintf("%03s", parseBinary(group))
	return &ColorValue{
		OriginDecimalValue:  v,
		OriginBinaryValue:   binary,
		CompressBinaryValue: compressBinary,
		GroupDecimalValue:   group,
		GroupBinaryValue:    groupBinary,
		RepairBinaryValue:   binary[5:],
	}
}

func (v *ColorValue) SetHideColor(hide *ColorValue) {
	compressBinary := fmt.Sprintf("%s%s", v.OriginBinaryValue[:5], hide.GroupBinaryValue)
	v.NewValue = parseUint(compressBinary)
	v.isChange = true
}

func (v *ColorValue) Authenticate(repairColor *ColorValue) bool {
	return repairColor.RepairBinaryValue == v.GroupBinaryValue
}

func (v *ColorValue) SetRepairColor(repairColor *ColorValue) {
	r := parseUint(repairColor.RepairBinaryValue) * interval
	v.NewValue = r
	v.isChange = true
}

func (v *ColorValue) SetBlackColor() {
	v.NewValue = 0
	v.isChange = true
}

func parseBinary(v uint8) string {
	return strconv.FormatUint(uint64(v), 2)
}

func parseUint(v string) uint8 {
	d, err := strconv.ParseUint(v, 2, 64)
	if err != nil {
		panic(fmt.Errorf("failed to parse uint: %v", err))
	}
	return uint8(d)
}
