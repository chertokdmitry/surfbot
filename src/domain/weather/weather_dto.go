package weather

type Weather struct {
	SpotId int64
	Title  string
	Hourly string
}

type Hour struct {
	Clouds     int
	Wind_deg   int
	Wind_speed float32
	Temp       float32
}
