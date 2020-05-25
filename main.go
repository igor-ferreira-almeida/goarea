package goarea

import "math"

// PI é uma proporção numérica
const PI = 3.1415

// Circle calcula a área que é obtida pela fórmula PI r*r: raio ao quadrado vezes PI
func Circle(radius float64) float64 {
	return math.Pow(radius, 2) * PI
}

// Rectangle calcula a área que é obtida pela multiplicação entre a largura e o comprimento
func Rectangle(length, width float64) float64 {
	return length * width
}

func _Triangle(base, height float64) float64 {
	return base * height / 2
}
