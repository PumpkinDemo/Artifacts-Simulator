package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// 掷骰子
func dice(probability float32) bool {
	randInt := rand.Intn(100)
	return randInt < int(100*probability)
}

// 生成指定圣遗物
func newArtifact(stars int32, set SetType) Artifact {
	// 随机生成部位
	var slot SlotType = SlotType(rand.Intn(5))
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
		weight = rand.Intn(sandsMainStatWeightTotal)
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
		weight = rand.Intn(gobletMainStatWeightTotal)
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
		weight = rand.Intn(circletMainStatWeightTotal)
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
		var weight = rand.Intn(totalWeight)
		var subStatRank = rand.Intn(4)
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
	}
}

// 从某个秘境获取圣遗物
func gainArtifacts(domain string, resin int32) []Artifact {
	// artifacts := list.New()

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
	}

	garbages := make([]Artifact, 0, count)

	for i := 0; i < count; i++ {
		var setIndex = rand.Intn(2)
		var setType = artifactInDomains[domain][setIndex]
		// artifacts.PushBack(newArtifact(5, setType))
		garbages = append(garbages, newArtifact(5, setType))
	}

	// return artifacts
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

func (a Artifact) levelUp() {

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
