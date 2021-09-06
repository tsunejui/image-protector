package image

type Color struct {
	R *ColorValue
	G *ColorValue
	B *ColorValue
	A *ColorValue
}

func (c *Color) Hide(hideData *Color) {
	c.R.SetHideColor(hideData.R)
	c.G.SetHideColor(hideData.G)
	c.B.SetHideColor(hideData.B)
	c.A.SetHideColor(hideData.A)
}

func (c *Color) Repair(repairData *Color) {
	c.R.SetRepairColor(repairData.R)
	c.G.SetRepairColor(repairData.G)
	c.B.SetRepairColor(repairData.B)
	c.A.SetRepairColor(repairData.A)
}

func (c *Color) Black() {
	blackColor(c.R, c.G, c.B, c.A)
}

func (c *Color) Check(repairData *Color) bool {
	if c.R.Authenticate(repairData.R) &&
		c.G.Authenticate(repairData.G) &&
		c.B.Authenticate(repairData.B) &&
		c.A.Authenticate(repairData.A) {
		return true
	}
	return false
}

func (c *Color) RGBA() (r, g, b, a uint8) {
	return getColorValue(c.R), getColorValue(c.G), getColorValue(c.B), getColorValue(c.A)
}

func blackColor(colors ...*ColorValue) {
	for _, c := range colors {
		c.SetBlackColor()
	}
}

func getColorValue(v *ColorValue) uint8 {
	if v.isChange {
		return v.NewValue
	}
	return v.OriginDecimalValue
}
