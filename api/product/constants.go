package product

var prices = map[int]float64{
	1:  10.99,
	12: 5.49,
	13: 8.75,
	14: 5.99,
}

var measures = map[int]map[int][]float64{
	1: { //h,  w,  d
		2:  []float64{30, 25, 20}, // Moyogi
		3:  []float64{35, 30, 25}, // Fukinagashi
		4:  []float64{40, 35, 30}, // Kabudachi
		5:  []float64{25, 20, 15}, // Literati
		6:  []float64{35, 25, 20}, // Shakan
		7:  []float64{30, 30, 30}, // Ikadabuki
		8:  []float64{45, 40, 35}, // Han-Kengai
		9:  []float64{40, 30, 25}, // Yose-ue
		10: []float64{35, 30, 25}, // Neagari
		11: []float64{30, 25, 20}, // Ishitsuki
	},
	12: {
		15: []float64{26.5, 26.5, 26.5}, // Akadama
		16: []float64{28, 35.5, 35.5},   // Kiryuzuna
		17: []float64{35.5, 47, 47},     // Kanuma
		18: []float64{25, 25, 25},       // Sphagnum moss
		19: []float64{26.5, 26.5, 26.5}, // Tierra para macetas
		20: []float64{28, 28, 28},       // Piedra pómez
		21: []float64{35.5, 35.5, 35.5}, // Perlita
		22: []float64{28, 28, 28},       // Vermiculita
	},
	13: {
		23: []float64{23, 13, 0}, // Tijeras de podar
		24: []float64{24, 13, 0}, // Cizallas de hoja larga
		25: []float64{25, 13, 0}, // Cizallas de yunque
		26: []float64{26, 13, 0}, // Tijeras de defoliación
		27: []float64{27, 13, 0}, // Pinzas de agarrar
		28: []float64{28, 13, 0}, // Gancho de raíz
		29: []float64{29, 13, 0}, // Cepillo de raíz
		30: []float64{30, 13, 0}, // Cincel de raíz
		31: []float64{31, 13, 0}, // Sierra
		32: []float64{32, 13, 0}, // Alicate
	},
	14: {
		33: []float64{20, 15, 15}, // Macetas de cerámica
		34: []float64{25, 20, 20}, // Macetas de madera
		35: []float64{15, 12, 12}, // Macetas de arcilla
		36: []float64{18, 16, 14}, // Macetas de plástico
		37: []float64{22, 18, 18}, // Macetas de vidrio
		38: []float64{20, 20, 20}, // Macetas de metal
	},
}
