package constant

type Nutrition struct {
	Energy                float64
	Sugar                 float64
	SaturatedFatty        float64
	Salt                  float64
	Protein               float64
	Fibres                float64
	FruitVegetableLegumes float64
}

type NutrientPoints struct {
	Energy                []float64
	Sugar                 []float64
	SaturatedFatty        []float64
	Salt                  []float64
	Protein               []float64
	Fibres                []float64
	FruitVegetableLegumes []float64
}

var FoodPoints = NutrientPoints{
	Energy: []float64{
		80,
		160,
		240,
		320,
		400,
		480,
		560,
		640,
		720,
		800,
		9999999,
	},
	Sugar: []float64{
		4.5,
		9.0,
		13.5,
		18.0,
		22.5,
		27.0,
		31.0,
		36.0,
		40.0,
		45.0,
		9999999,
	},
	SaturatedFatty: []float64{
		1,
		2,
		3,
		4,
		5,
		6,
		7,
		8,
		9,
		10,
		9999999,
	},
	Salt: []float64{
		90,
		180,
		270,
		360,
		450,
		540,
		630,
		720,
		810,
		900,
		9999999,
	},
	Protein: []float64{
		1.6,
		3.2,
		4.8,
		6.4,
		8.0,
		9999999,
	},
	Fibres: []float64{
		0.7,
		1.4,
		2.1,
		2.8,
		3.5,
		9999999,
	},
	FruitVegetableLegumes: []float64{
		40,
		60,
		80,
		-1,
		-1,
		9999999,
	},
}

var BeveragePoints = NutrientPoints{
	Energy: []float64{
		7.2,
		14.3,
		21.5,
		28.5,
		35.9,
		43.0,
		50.2,
		57.4,
		64.5,
		9999999,
	},
	Sugar: []float64{
		0,
		1.5,
		3.0,
		4.5,
		6.0,
		7.5,
		9.0,
		10.5,
		12.0,
		13.5,
		9999999,
	},
	SaturatedFatty: []float64{
		10,
		16,
		22,
		28,
		34,
		40,
		46,
		52,
		58,
		64,
		9999999,
	},
	Salt: []float64{
		90,
		180,
		270,
		360,
		450,
		540,
		630,
		720,
		810,
		900,
		9999999,
	},
	Protein: []float64{
		1.6,
		3.2,
		4.8,
		6.4,
		8.0,
		9999999,
	},
	Fibres: []float64{
		0.7,
		1.4,
		2.1,
		2.8,
		3.5,
		9999999,
	},
	FruitVegetableLegumes: []float64{
		40,
		-1, // skip
		60,
		-1, // skip
		80,
		-1, // skip
		-1, // skip
		-1, // skip
		-1, // skip
		-1, // skip
		9999999,
	},
}

func GetPoint(value float64, thresholds []float64) int {
	for i, threshold := range thresholds {
		if threshold == -1 {
			continue
		}
		if value <= threshold {
			return i
		}
	}
	return len(thresholds)
}
