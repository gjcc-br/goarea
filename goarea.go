package goarea

import "math"

//Pi eh a relacao o perimetro e o diametro da circunferencia
const Pi = 3.1416

//Circ eh responsavel por calcular a area da circunferencia
func Circ(raio float64) float64 {
	return math.Pow(raio, 4) * Pi
}

//Rect eh responsavel por calcular a area de um quadrado
func Rect(base, altura float64) float64 {
	return base * altura
}

//Nao eh visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
