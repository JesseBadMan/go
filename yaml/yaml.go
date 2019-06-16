package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Args Args
}

type FormulaLevel struct {
	Alpha float64 `json:"alpha"` // 分段公式第一个系数
	Beta  float64 `json:"beta"`  // 分段公式第二个系数
}

// TagBoost 按照标签区分不同的最终权重
type TagBoost struct {
	UVBoost      float64 `json:"uvBoost"`      // UV权重
	IncomeBoost  float64 `json:"incomeBoost"`  // 收入权重
	BarrageBoost float64 `json:"barrageBoost"` // 弹幕权重
}

// Args 计算中用到的系数
type Args struct {
	/*
		1. 5min数据各维度做加权平均
		令五个时间片数据为 a1-a5，其中a5为最近的时间片
		(a1+2*a2+3*a3+4*a4+5*a5)/(1+2+3+4+5)
		如果不足5个时间片，算式相应简化, 如只有3个时间片数据 算式变为(a1+2*a2+3*a3)/(1+2+3)
		加权平均系数
	*/
	WeightedMeanCoefficient float64 `json:"weightedMeanCoefficient"`

	/*
		2. 所有数据的维度乘以系数
		uv 1
		送礼人数 150
		送礼金额(元) 10
		弹幕人数 6
		弹幕条数 2
	*/
	BoostUV                    float64 `json:"boostUV"`                    // UV权重
	BoostNumberOfConsumer      float64 `json:"boostNumberOfConsumer"`      // 消费人数权重
	BoostIncome                float64 `json:"boostIncome"`                // 收入权重
	BoostNumberOfBarrageSender float64 `json:"boostNumberOfBarrageSender"` // 弹幕人数权重
	BoostNumberOfBarrage       float64 `json:"boostNumberOfBarrage"`       // 弹幕数量权重

	// 分段函数的参数
	Level1  FormulaLevel `json:"level1"`  // 0 ~ 1
	Level2  FormulaLevel `json:"level2"`  // 2 ~ 10
	Level3  FormulaLevel `json:"level3"`  // 11 ~ 50
	Level4  FormulaLevel `json:"level4"`  // 51~100
	Level5  FormulaLevel `json:"level5"`  // 101~500
	Level6  FormulaLevel `json:"level6"`  // 501~1000
	Level7  FormulaLevel `json:"level7"`  // 1001~5000
	Level8  FormulaLevel `json:"level8"`  // 5001~20000
	Level9  FormulaLevel `json:"level9"`  // 2-4万
	Level10 FormulaLevel `json:"level10"` // 4-20万
	Level11 FormulaLevel `json:"level11"` // 20-50万
	Level12 FormulaLevel `json:"level12"` // 50万以上

	// 最终权重
	GameBoost          TagBoost `json:"tagGameBoost"`          // 游戏
	EntertainmentBoost TagBoost `json:"tagEntertainmentBoost"` // 泛娱乐
}

func main() {
	conf := &Args{}

	f, err := os.OpenFile("config.json", os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(f)

	err = encoder.Encode(conf)
	if err != nil {
		panic(err)
	}
}
