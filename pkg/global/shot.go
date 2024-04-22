package global

const (
	SingleShotCost = 20
	TenShotCost    = 200

	InitialUnlockDishAmount = -1
)

var (
	// PieceAmountMap 表示每种品质物品的对应兑换碎片的所需的数量
	PieceAmountMap = map[string]int{
		"R":   20,
		"SR":  30,
		"SSR": 50,
	}

	// ProbabilityMap 表示每种品质物品的对应抽取概率
	ProbabilityMap = map[string]float64{
		"R":   0.75,
		"SR":  0.2,
		"SSR": 0.05,
	}
)
