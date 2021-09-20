package main

type luxeCarV struct {
	price int
}

func (lcv *luxeCarV) accept(v visitor) {
	v.visitForLuxeCarVPrice(lcv)
}

func (c *luxeCarV) getType() string {
	return "Luxe car Price"
}
