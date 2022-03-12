package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// 获取随机数
func getRandomInteger(i int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(i)
}

// 掷骰子
func dice(probability float32) bool {
	randInt := getRandomInteger(100)
	return randInt < int(100*probability)
}

// 生成指定圣遗物
func newArtifact(stars int32, set SetType) Artifact {
	// 随机生成部位
	var slot SlotType = SlotType(getRandomInteger(5))
	var name string = setSlotName[set][slot]

	// 随机生成主词条
	var mainStat Stat
	var weight int
	switch slot {
	case Flower_of_Life:
		mainStat = Stat{
			Type:  HP,
			Value: mainStatMap[HP][0],
		}
	case Plume_of_Death:
		mainStat = Stat{
			Type:  ATK,
			Value: mainStatMap[ATK][0],
		}
	case Sands_of_Eon:
		weight = getRandomInteger(sandsMainStatWeightTotal)
		for k, v := range sandsMainStatWeightMap {
			if v <= weight {
				weight -= v
			} else {
				mainStat = Stat{
					Type:  k,
					Value: mainStatMap[k][0],
				}
				break
			}
		}
	case Goblet_of_Eonothem:
		weight = getRandomInteger(gobletMainStatWeightTotal)
		for k, v := range gobletMainStatWeightMap {
			if v <= weight {
				weight -= v
			} else {
				mainStat = Stat{
					Type:  k,
					Value: mainStatMap[k][0],
				}
				break
			}
		}
	case Circlet_of_Logos:
		weight = getRandomInteger(circletMainStatWeightTotal)
		for k, v := range circletMainStatWeightMap {
			if v <= weight {
				weight -= v
			} else {
				mainStat = Stat{
					Type:  k,
					Value: mainStatMap[k][0],
				}
				break
			}
		}
	}

	// 决定副词条个数
	var subStatCount int
	if dice(initialThreeSubStats) {
		subStatCount = 3
	} else {
		subStatCount = 4
	}

	// 随机生成副词条，遵循不放回抽样
	var subStats [4]Stat
	subStats[3] = Stat{
		Type:  NIL,
		Value: 0,
	}
	weightMap := make(map[StatType]int)
	var totalWeight = subStatWeightTotal

	for k, v := range subStatWeightMap {
		if k == mainStat.Type {
			totalWeight -= v
			continue
		}
		weightMap[k] = v
	}

	for i := 0; i < subStatCount; i++ {
		var weight = getRandomInteger(totalWeight)
		var subStatRank = getRandomInteger(4)
		var chosenStat StatType
		for k, v := range weightMap {
			if v <= weight {
				weight -= v
			} else {
				totalWeight -= v
				subStats[i] = Stat{
					Type:  k,
					Value: subStatMap[k][subStatRank],
				}
				chosenStat = k
				break
			}
		}
		delete(weightMap, chosenStat)
	}

	return Artifact{
		Stars:    stars,
		Lv:       0,
		Set:      set,
		Slot:     slot,
		Name:     name,
		MainStat: mainStat,
		SubStat:  subStats,
		Exp:      0,
	}
}

// 从某个秘境获取圣遗物
func gainArtifacts(domain string, resin int32) []Artifact {
	var count = 0

	if resin == 20 {
		count = 1
		if dice(doubleFiveStarArtifact) {
			count += 1
		}
	} else if resin == 40 {
		count = 2
		if dice(doubleFiveStarArtifact) {
			count += 1
		}
		if dice(doubleFiveStarArtifact) {
			count += 1
		}
	} else {
		count = 1
	}

	garbages := make([]Artifact, 0, count)

	for i := 0; i < count; i++ {
		var setIndex = getRandomInteger(2)
		var setType = artifactInDomains[domain][setIndex]
		garbages = append(garbages, newArtifact(5, setType))
	}

	return garbages
}

func (a Artifact) Stringify() string {
	var res = fmt.Sprintf("%d星圣遗物 %s , lv%d, ", a.Stars, a.Name, a.Lv)
	res += fmt.Sprintf("%s套%s\n", setOralHanzi[a.Set], slotHanzi[a.Slot])
	res += fmt.Sprintf(" 主词条: %s %.2f\n 副词条: \n", statHanzi[a.MainStat.Type], a.MainStat.Value)
	for i := 0; i <= 3; i++ {
		if a.SubStat[i].Type == NIL {
			break
		}
		res += fmt.Sprintf("  %d. %s %.2f\n", i+1, statHanzi[a.SubStat[i].Type], a.SubStat[i].Value)
	}
	return res
}

// 增加副词条
func (a *Artifact) addOneSubStat() {
	weightMap := make(map[StatType]int)
	var totalWeight = subStatWeightTotal
	var indexUninitializedSubStat = -1

	// 剔除主词条、副词条中的权重
	for k, v := range subStatWeightMap {
		if a.MainStat.Type == k {
			totalWeight -= v
			continue
		}
		weightMap[k] = v
	}

	for k, v := range a.SubStat {
		if v.Type != NIL {
			var weight = weightMap[v.Type]
			delete(weightMap, v.Type)
			totalWeight -= weight
		} else {
			indexUninitializedSubStat = k
		}
	}

	// 不满四词条，按权重增加一个词条
	var subStatRank = getRandomInteger(4)
	if indexUninitializedSubStat >= 0 {
		var weight = getRandomInteger(totalWeight)
		for k, v := range weightMap {
			if v <= weight {
				weight -= v
			} else {
				a.SubStat[indexUninitializedSubStat] = Stat{
					Type:  k,
					Value: subStatMap[k][subStatRank],
				}
				break
			}
		}
	} else {
		// 满四词条等概率增加某词条
		var chosenIndex = getRandomInteger(4)
		var stat = a.SubStat[chosenIndex]
		a.SubStat[chosenIndex].Value += subStatMap[stat.Type][subStatRank]
	}
}

// 圣遗物升级
func (a *Artifact) levelUp(dogFood []DogFood) {
	// 决定强化倍数
	var multiplier = 1
	var randInt = getRandomInteger(100)
	if randInt >= int(100*experienceOnce) {
		if randInt < int(100*(1-experienceFifth)) {
			multiplier = 2
		} else {
			multiplier = 5
		}
	}

	// 获得总经验数
	var exp = int(a.Exp)
	for _, v := range dogFood {
		if v.Stars == 5 {
			exp += experience5StarAsGarbage[v.Lv]
		} else {
			exp += experienceOfDifferentStarsLv0[v.Stars]
		}
	}
	exp *= multiplier

	// 圣遗物逐级升级
	for exp > 0 && a.Lv < 20 {
		var expToLevelUp = experiencesToLevelUp[a.Lv]
		if exp >= expToLevelUp {
			a.Lv += 1
			if a.Lv%4 == 0 {
				a.addOneSubStat()
			}
		} else {
			a.Exp = int32(exp)
		}
		exp -= expToLevelUp
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func enhance_handler() {
}

func main1() {
	for j := 1; j <= 4; j++ {
		fmt.Printf("\n第%v次刷圣遗物, 使用20个树脂\n", j)
		garbages := gainArtifacts("Peak of Vindagnyr", 20)
		for _, value := range garbages {
			fmt.Println(value.Stringify())
		}
	}
}
