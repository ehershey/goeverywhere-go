package proto

func (og *OldGeometry) GetLat() float64 {
	return float64(og.Coordinates[1])
}

func (og *OldGeometry) GetLon() float64 {
	return float64(og.Coordinates[0])
}

func (ob *OldBookmark) GetLon() float64 {
	return ob.GetLoc().GetLon()
}

func (ob *OldBookmark) GetLat() float64 {
	return ob.GetLoc().GetLat()
}

func (g *Geometry) GetLat() float64 {
	return float64(g.Coordinates.GetLatitude())
}

func (g *Geometry) GetLon() float64 {
	return float64(g.Coordinates.GetLongitude())
}

func (ob *Bookmark) GetLon() float64 {
	return ob.GetLoc().GetLon()
}

func (ob *Bookmark) GetLat() float64 {
	return ob.GetLoc().GetLat()
}
