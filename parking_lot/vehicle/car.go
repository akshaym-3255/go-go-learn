package main


type car struct {
	vehicleType string
	number string
	noofWheels int
}

func (c *car) GetType()(string) {
	return c.vehicleType
}

func (c *car) GetNumber()(string) {
	return c.number
}