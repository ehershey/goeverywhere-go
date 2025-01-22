package proto

func (og *OldGeometry) GetLat() float64 {
	return float64(og.Coordinates[0])
}

func (og *OldGeometry) GetLon() float64 {
	return float64(og.Coordinates[1])
}

func (ob *OldBookmark) GetLon() float64 {
	return ob.GetLoc().GetLon()
}

func (ob *OldBookmark) GetLat() float64 {
	return ob.GetLoc().GetLat()
}
