package foundation

type Color struct {
	Red float64
	Green float64
	Blue float64
}

func NewColor(r, g, b float64) Color {
	return Color { r, g, b }
}

func (c1 Color) Add(c2 Color) Color {
	return Color {c1.Red + c2.Red, c1.Green + c2.Green, c1.Blue + c2.Blue }
}

func (c1 Color) Sub(c2 Color) Color {
	return Color {c1.Red - c2.Red, c1.Green - c2.Green, c1.Blue - c2.Blue }
}


func (c1 Color) HadamardProduct(c2 Color) Color {	
	return Color {c1.Red * c2.Red, c1.Green * c2.Green, c1.Blue * c2.Blue }
}


