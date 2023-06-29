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
	14: {
		33: []float64{20, 15, 15}, // Macetas de cerámica
		34: []float64{25, 20, 20}, // Macetas de madera
		35: []float64{15, 12, 12}, // Macetas de arcilla
		36: []float64{18, 16, 14}, // Macetas de plástico
		37: []float64{22, 18, 18}, // Macetas de vidrio
		38: []float64{20, 20, 20}, // Macetas de metal
	},
}
