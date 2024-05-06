package global

import (
	"math"
	"math/rand"
)

const (
	DailyLoginReward = 20 // 每日登陆奖励
	BasicOrderReward = 10 // 基础订单奖励
	StdDev           = 40 // 计算浮动金币的正态分布函数标准差
)

// 构造正态分布函数
func normalPDF(x, mean, stdDev float64) float64 {
	exponent := math.Exp(-(math.Pow(x-mean, 2) / (2 * math.Pow(stdDev, 2))))
	return (1 / (math.Sqrt(2*math.Pi) * stdDev)) * exponent
}

// 计算获取各金币值的概率（权重）
func calculateWeightsFromNormalDist(mean, stdDev float64) []float64 {
	points := []float64{0, 20, 40, 60, 80, 100} // 更改分数点
	weights := make([]float64, len(points))

	for i, point := range points {
		weights[i] = normalPDF(point, mean, stdDev)
	}

	// 权重归一化
	totalWeight := 0.0
	for _, w := range weights {
		totalWeight += w
	}
	for i := range weights {
		weights[i] /= totalWeight
	}

	return weights
}

func CalculateMoney(score int) int {
	baseMoney := BasicOrderReward
	stdDevScore := float64(StdDev)
	// 使用玩家分数的均值
	meanScore := float64(score)

	// 计算权重
	weights := calculateWeightsFromNormalDist(meanScore, stdDevScore)

	// 生成随机数并根据权重选择浮动金币值
	randomIndex := rand.Float64()
	accumulatedWeight := 0.0
	randomMoney := 0
	for i, weight := range weights {
		accumulatedWeight += weight
		if randomIndex <= accumulatedWeight {
			randomMoney = i
			break
		}
	}

	// 确保随机索引在有效范围内
	if randomIndex > accumulatedWeight {
		randomMoney = len(weights) - 1
	}

	totalMoney := baseMoney + randomMoney

	return totalMoney
}
