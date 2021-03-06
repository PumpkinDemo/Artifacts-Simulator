package main

var statTypeZN = map[StatType]string{
	NIL:                "",
	ATK_big:            "攻击力百分比",
	ATK:                "攻击力",
	HP_big:             "生命值百分比",
	HP:                 "生命值",
	DEF_big:            "防御力百分比",
	DEF:                "防御力",
	CRIT_RATE:          "暴击率",
	CRIT_DMG:           "暴击伤害",
	ENERGY_RECHARGE:    "充能效率",
	ELEMENTAL_MASTARY:  "元素精通",
	HYDRO_DMG_BONUS:    "水元素伤害加成",
	PYRO_DMG_BONUS:     "火元素伤害加成",
	GEO_DMG_BONUS:      "岩元素伤害加成",
	ANEMO_DMG_BONUS:    "风元素伤害加成",
	ELECTRO_DMG_BONUS:  "雷元素伤害加成",
	CRYO_DMG_BONUS:     "冰元素伤害加成",
	PHYSICAL_DMG_BONUS: "物理伤害加成",
	HEALING_BONUS:      "治疗加成",
}

var statTypeEN = map[StatType]string{
	NIL:                "",
	ATK_big:            "ATK%",
	ATK:                "ATK",
	HP_big:             "HP%",
	HP:                 "HP",
	DEF_big:            "DEF%",
	DEF:                "DEF",
	CRIT_RATE:          "CRIT RATE",
	CRIT_DMG:           "CRIT DMG",
	ENERGY_RECHARGE:    "ER",
	ELEMENTAL_MASTARY:  "EM",
	HYDRO_DMG_BONUS:    "HYDRO BONUS",
	PYRO_DMG_BONUS:     "PYRO BONUS",
	GEO_DMG_BONUS:      "GEO BONUS",
	ANEMO_DMG_BONUS:    "ANEMO BONUS",
	ELECTRO_DMG_BONUS:  "ELECTRO BONUS",
	CRYO_DMG_BONUS:     "CRYO BONUS",
	PHYSICAL_DMG_BONUS: "PHYSICAL BONUS",
	HEALING_BONUS:      "HEALING BONUS",
}

var ZNstatType = map[string]StatType{
	"":        NIL,
	"攻击力百分比":  ATK_big,
	"攻击力":     ATK,
	"生命值百分比":  HP_big,
	"生命值":     HP,
	"防御力百分比":  DEF_big,
	"防御力":     DEF,
	"暴击率":     CRIT_RATE,
	"暴击伤害":    CRIT_DMG,
	"充能效率":    ENERGY_RECHARGE,
	"元素精通":    ELEMENTAL_MASTARY,
	"水元素伤害加成": HYDRO_DMG_BONUS,
	"火元素伤害加成": PYRO_DMG_BONUS,
	"岩元素伤害加成": GEO_DMG_BONUS,
	"风元素伤害加成": ANEMO_DMG_BONUS,
	"雷元素伤害加成": ELECTRO_DMG_BONUS,
	"冰元素伤害加成": CRYO_DMG_BONUS,
	"物理伤害加成":  PHYSICAL_DMG_BONUS,
	"治疗加成":    HEALING_BONUS,
}

var ENstatType = map[string]StatType{
	"":        NIL,
	"ATK%":  ATK_big,
	"ATK":     ATK,
	"HP%":  HP_big,
	"HP":     HP,
	"DEF%":  DEF_big,
	"DEF":     DEF,
	"CRIT RATE":     CRIT_RATE,
	"CRIT DMG":    CRIT_DMG,
	"ER":    ENERGY_RECHARGE,
	"EM":    ELEMENTAL_MASTARY,
	"HYDRO BONUS": HYDRO_DMG_BONUS,
	"PYRO BONUS": PYRO_DMG_BONUS,
	"GEO BONUS": GEO_DMG_BONUS,
	"ANEMO BONUS": ANEMO_DMG_BONUS,
	"ELECTRO BONUS": ELECTRO_DMG_BONUS,
	"CRYO BONUS": CRYO_DMG_BONUS,
	"PHYSICAL BONUS":  PHYSICAL_DMG_BONUS,
	"HEALING BONUS":    HEALING_BONUS,
}

var artifactSetZN = map[SetType]string{
	Ocean_Hued_Clam:           "海染砗磲",
	Husk_of_Opulent_Dreams:    "华馆梦醒形骸记",
	Emblem_of_Severed_Fate:    "绝缘之旗印",
	Shimenawas_Reminiscence:   "追忆之注连",
	Pale_Flame:                "苍白之火",
	Tenacity_of_the_Millelith: "千岩牢固",
	Heart_of_Depth:            "沉沦之心",
	Blizzard_Strayer:          "冰风迷途的勇士",
	Crimson_Witch_of_Flames:   "炽烈的炎之魔女",
	Lavawalker:                "渡过烈火的贤人",
	Thundering_Fury:           "如雷的盛怒",
	Thundersmoother:            "平息鸣雷的尊者",
	Retracing_Bolide:          "逆飞的流星",
	Archaic_Petra:             "悠古的磐岩",
	Viridescent_Venerer:       "翠绿之影",
	Maiden_Beloved:            "被怜爱的少女",
	Bloodstained_Chivalry:     "染血的骑士道",
	Noblesse_Oblige:           "昔日宗室之仪",
	Wanderers_Troupe:          "流浪大地的乐团",
	Gladiators_Finale:         "角斗士的终幕礼",
}

var artifactSetEN = map[SetType]string{
	Ocean_Hued_Clam: "Ocean Hued Clam",
	Husk_of_Opulent_Dreams: "Husk of Opulent Dreams",
	Emblem_of_Severed_Fate: "Emblem of Severed Fate",
	Shimenawas_Reminiscence: "Shimenawa's Reminiscence",
	Pale_Flame: "Pale Flame",
	Tenacity_of_the_Millelith: "Tenacity of the Millelith",
	Heart_of_Depth: "Heart of Depth",
	Blizzard_Strayer: "Blizzard Strayer",
	Crimson_Witch_of_Flames: "Crimson Witch of Flames",
	Lavawalker: "Lavawalker",
	Thundering_Fury: "Thundering Fury",
	Thundersmoother: "Thundersmoother",
	Retracing_Bolide: "Retracing Bolide",
	Archaic_Petra: "Archaic Petra",
	Viridescent_Venerer: "Viridescent Venerer",
	Maiden_Beloved: "Maiden Beloved",
	Bloodstained_Chivalry: "Bloodstained Chivalry",
	Noblesse_Oblige: "Noblesse Oblige",
	Wanderers_Troupe: "Wanderer's Troupe",
	Gladiators_Finale: "Gladiator's Finale",
  }

var ZNartifactSet = map[string]SetType{
	"海染砗磲":    Ocean_Hued_Clam,
	"华馆梦醒形骸记": Husk_of_Opulent_Dreams,
	"绝缘之旗印":   Emblem_of_Severed_Fate,
	"追忆之注连":   Shimenawas_Reminiscence,
	"苍白之火":    Pale_Flame,
	"千岩牢固":    Tenacity_of_the_Millelith,
	"沉沦之心":    Heart_of_Depth,
	"冰风迷途的勇士": Blizzard_Strayer,
	"炽烈的炎之魔女": Crimson_Witch_of_Flames,
	"渡过烈火的贤人": Lavawalker,
	"如雷的盛怒":   Thundering_Fury,
	"平息鸣雷的尊者": Thundersmoother,
	"逆飞的流星":   Retracing_Bolide,
	"悠古的磐岩":   Archaic_Petra,
	"翠绿之影":    Viridescent_Venerer,
	"被怜爱的少女":  Maiden_Beloved,
	"染血的骑士道":  Bloodstained_Chivalry,
	"昔日宗室之仪":  Noblesse_Oblige,
	"流浪大地的乐团": Wanderers_Troupe,
	"角斗士的终幕礼": Gladiators_Finale,
}

var ENartifactSet = map[string]SetType{
	"Ocean Hued Clam": Ocean_Hued_Clam,
	"Husk of Opulent Dreams": Husk_of_Opulent_Dreams,
	"Emblem of Severed Fate": Emblem_of_Severed_Fate,
	"Shimenawa's Reminiscence": Shimenawas_Reminiscence,
	"Pale Flame": Pale_Flame,
	"Tenacity of the Millelith": Tenacity_of_the_Millelith,
	"Heart of Depth": Heart_of_Depth,
	"Blizzard Strayer": Blizzard_Strayer,
	"Crimson Witch of Flames": Crimson_Witch_of_Flames,
	"Lavawalker": Lavawalker,
	"Thundering Fury": Thundering_Fury,
	"Thundersmoother": Thundersmoother,
	"Retracing Bolide": Retracing_Bolide,
	"Archaic Petra": Archaic_Petra,
	"Viridescent Venerer": Viridescent_Venerer,
	"Maiden Beloved": Maiden_Beloved,
	"Bloodstained Chivalry": Bloodstained_Chivalry,
	"Noblesse Oblige": Noblesse_Oblige,
	"Wanderer's Troupe": Wanderers_Troupe,
	"Gladiator's Finale": Gladiators_Finale,
}

var artifactInDomains = map[string][]SetType{
	"Domain of Guyun": {
		Archaic_Petra, Retracing_Bolide,
	},
	"Midsummer Courtyard": {
		Thundering_Fury, Thundersmoother,
	},
	"Valley of Remembrance": {
		Viridescent_Venerer, Maiden_Beloved,
	},
	"Hidden Palace of Zhou Formula": {
		Crimson_Witch_of_Flames, Lavawalker,
	},
	"Peak of Vindagnyr": {
		Heart_of_Depth, Blizzard_Strayer,
	},
	"Ridge Watch": {
		Tenacity_of_the_Millelith, Pale_Flame,
	},
	"Momiji-Dyed Court": {
		Emblem_of_Severed_Fate, Shimenawas_Reminiscence,
	},
	"Slumbering Court": {
		Husk_of_Opulent_Dreams, Ocean_Hued_Clam,
	},
	"Clear Pool and Mountain Cavern": {
		Bloodstained_Chivalry, Noblesse_Oblige,
	},
}

var setSlotName = map[SetType][]string{
	Ocean_Hued_Clam: {
		"Sea-Dyed Blossom",
		"Deep Palace's Plume",
		"Cowry of Parting",
		"Pearl Cage",
		"Crown of Watatsumi",
	},
	Husk_of_Opulent_Dreams: {
		"Bloom Times",
		"Plume of Luxury",
		"Song of Life",
		"Calabash of Awakening",
		"Skeletal Hat",
	},
	Emblem_of_Severed_Fate: {
		"Magnificent Tsuba",
		"Sundered Feather",
		"Storm Cage",
		"Scarlet Vessel",
		"Ornate Kabuto",
	},
	Shimenawas_Reminiscence: {
		"Entangling Bloom",
		"Shaft of Remembrance",
		"Morning Dew's Moment",
		"Hopeful Heart",
		"Capricious Visage",
	},
	Pale_Flame: {
		"Stainless Bloom",
		"Wise Doctor's Pinion",
		"Moment of Cessation",
		"Surpassing Cup",
		"Mocking Mask",
	},
	Tenacity_of_the_Millelith: {
		"Flower of Accolades",
		"Ceremonial War-Plume",
		"Orichalceous Time-Dial",
		"Noble's Pledging Vessel",
		"General's Ancient Helm",
	},
	Heart_of_Depth: {
		"Gilded Corsage",
		"Gust of Nostalgia",
		"Copper Compass",
		"Goblet of Thundering Deep",
		"Wine-Stained Tricorne",
	},
	Blizzard_Strayer: {
		"Snowswept Memory",
		"Icebreaker's Resolve",
		"Frozen Homeland's Demise",
		"Frost-Weaved Dignity",
		"Broken Rime's Echo",
	},
	Crimson_Witch_of_Flames: {
		"Witch's Flower of Blaze",
		"Witch's Ever-Burning Plume",
		"Witch's End Time",
		"Witch's Heart Flames",
		"Witch's Scorching Hat",
	},
	Lavawalker: {
		"Lavawalker's Resolution",
		"Lavawalker's Salvation",
		"Lavawalker's Torment",
		"Lavawalker's Epiphany",
		"Lavawalker's Wisdom",
	},
	Thundering_Fury: {
		"Thunderbird's Mercy",
		"Survivor of Catastrophe",
		"Hourglass of Thunder",
		"Omen of Thunderstorm",
		"Thunder Summoner's Crown",
	},
	Thundersmoother: {
		"Thundersmoother's Heart",
		"Thundersmoother's Plume",
		"Hour of Soothing Thunder",
		"Thundersmoother's Goblet",
		"Thundersmoother's Diadem",
	},
	Retracing_Bolide: {
		"Summer Night's Bloom",
		"Summer Night's Finale",
		"Summer Night's Moment",
		"Summer Night's Waterballoon",
		"Summer Night's Mask",
	},
	Archaic_Petra: {
		"Flower of Creviced Cliff",
		"Feather of Jagged Peaks",
		"Sundial of Enduring Jade",
		"Goblet of Chiseled Crag",
		"Mask of Solitude Basalt",
	},
	Viridescent_Venerer: {
		"In Remembrance of Viridescent Fields",
		"Viridescent Arrow Feather",
		"Viridescent Venerer's Determination",
		"Viridescent Venerer's Vessel",
		"Viridescent Venerer's Diadem",
	},
	Maiden_Beloved: {
		"Maiden's Distant Love",
		"Maiden's Heart-stricken Infatuation",
		"Maiden's Passing Youth",
		"Maiden's Fleeting Leisure",
		"Maiden's Fading Beauty",
	},
	Bloodstained_Chivalry: {
		"Bloodstained Flower of Iron",
		"Bloodstained Black Plume",
		"Bloodstained Final Hour",
		"Bloodstained Chevalier's Goblet",
		"Bloodstained Iron Mask",
	},
	Noblesse_Oblige: {
		"Royal Flora",
		"Royal Plume",
		"Royal Pocket Watch",
		"Royal Silver Urn",
		"Royal Masque",
	},
	Wanderers_Troupe: {
		"Troupe's Dawnlight",
		"Bard's Arrow Feather",
		"Concert's Final Hour",
		"Wanderer's String-Kettle",
		"Conductor's Top Hat",
	},
	Gladiators_Finale: {
		"Gladiator's Nostalgia",
		"Gladiator's Destiny",
		"Gladiator's Longing",
		"Gladiator's Intoxication",
		"Gladiator's Triumphus",
	},
}

// 测试用汉字
var slotHanzi = [5]string{
	"生之花",
	"死之羽",
	"时之沙",
	"空之杯",
	"理之冠",
}

var slotEN = [5]string{
	"Flower of Life",
	"Plume of Death",
	"Sands of Eon",
	"Goblet of Eonothem",
	"Circlet of Logos",
}

var ZNslotType = map[string]SlotType{
	"生之花": Flower_of_Life,
	"死之羽": Plume_of_Death,
	"时之沙": Sands_of_Eon,
	"空之杯": Goblet_of_Eonothem,
	"理之冠": Circlet_of_Logos,
}

var ENslotType = map[string]SlotType{
	"Flower of Life": Flower_of_Life,
	"Plume of Death": Plume_of_Death,
	"Sands of Eon": Sands_of_Eon,
	"Goblet of Eonothem": Goblet_of_Eonothem,
	"Circlet of Logos": Circlet_of_Logos,
}

var setOralHanzi = [20]string{
	"海染", "华馆", "绝缘", "追忆", "苍白",
	"千岩", "水", "冰", "魔女", "渡火", "如雷",
	"平雷", "逆飞", "岩",
	"风", "少女", "骑士", "宗室", "乐团", "角斗",
}

var statHanzi = [18]string{
	"攻击力百分比", "攻击力", "生命值百分比", "生命值", "防御力百分比", "防御力",
	"暴击率", "暴击伤害", "元素充能效率", "元素精通",
	"水元素伤害加成", "火元素伤害加成", "岩元素伤害加成", "风元素伤害加成",
	"雷元素伤害加成", "冰元素伤害加成", "物理伤害加成", "治疗加成",
}

var domainZhToEn = map[string]string{
	"岩本":  "Domain of Guyun",
	"逆飞本": "Domain of Guyun",
	"雷本":  "Midsummer Courtyard",
	"平雷本": "Midsummer Courtyard",
	"风本":  "Valley of Remembrance",
	"少女本": "Valley of Remembrance",
	"火本":  "Hidden Palace of Zhou Formula",
	"渡火本": "Hidden Palace of Zhou Formula",
	"水本":  "Peak of Vindagnyr",
	"冰本":  "Peak of Vindagnyr",
	"苍白本": "Ridge Watch",
	"千岩本": "Ridge Watch",
	"绝缘本": "Momiji-Dyed Court",
	"追忆本": "Momiji-Dyed Court",
	"蚌埠本": "Slumbering Court",
	"海染本": "Slumbering Court",
	"华馆本": "Slumbering Court",
	"骑士本": "Clear Pool and Mountain Cavern",
	"宗室本": "Clear Pool and Mountain Cavern",
}

// 20树脂圣遗物双五星概率
const doubleFiveStarArtifact = 0.07

// 圣遗物升级经验翻倍概率
const experienceOnce = 0.9
const experienceTwice = 0.09
const experienceFifth = 0.01

// 副本五星圣遗物 - 初始副词条数量概率
const initialThreeSubStats = 0.8
const initialFourSubStats = 0.2

// 时之沙主词条权重
const sandsMainStatWeightTotal = 30

var sandsMainStatWeightMap = map[StatType]int{
	HP_big:            8,
	DEF_big:           8,
	ATK_big:           8,
	ENERGY_RECHARGE:   3,
	ELEMENTAL_MASTARY: 3,
}

// 空之杯主词条权重
const gobletMainStatWeightTotal = 80

var gobletMainStatWeightMap = map[StatType]int{
	HP_big:             17,
	DEF_big:            16,
	ATK_big:            17,
	ELEMENTAL_MASTARY:  2,
	CRYO_DMG_BONUS:     4,
	GEO_DMG_BONUS:      4,
	ELECTRO_DMG_BONUS:  4,
	ANEMO_DMG_BONUS:    4,
	PHYSICAL_DMG_BONUS: 4,
	PYRO_DMG_BONUS:     4,
	HYDRO_DMG_BONUS:    4,
}

// 理之冠主词条权重
const circletMainStatWeightTotal = 100

var circletMainStatWeightMap = map[StatType]int{
	HP_big:            22,
	DEF_big:           22,
	ATK_big:           22,
	ELEMENTAL_MASTARY: 4,
	HEALING_BONUS:     10,
	CRIT_RATE:         10,
	CRIT_DMG:          10,
}

// 初始副词条权重
const subStatWeightTotal = 44

var subStatWeightMap = map[StatType]int{
	HP:                6,
	DEF:               6,
	ATK:               6,
	HP_big:            4,
	DEF_big:           4,
	ATK_big:           4,
	CRIT_RATE:         3,
	CRIT_DMG:          3,
	ELEMENTAL_MASTARY: 4,
	ENERGY_RECHARGE:   4,
}

// 五星圣遗物主词条档位
var mainStatMap = map[StatType][]float32{
	ATK_big: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	ATK: {
		47, 60, 73, 86, 100, 113,
		126, 139, 152, 166, 179, 192,
		205, 219, 232, 245, 258, 272,
		285, 298, 311,
	},
	HP_big: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	HP: {
		717, 920, 1123, 1326,
		1530, 1733, 1936, 2139,
		2342, 2545, 2749, 2952,
		3155, 3358, 3561, 3764,
		3967, 4171, 4374, 4577,
		4780,
	},
	DEF_big: {
		8.7, 11.2, 13.7, 16.2,
		18.6, 21.1, 23.6, 26.1,
		28.6, 31, 33.5, 36,
		38.5, 41, 43.4, 45.9,
		48.4, 50.8, 53.3, 55.8,
		58.3,
	},
	DEF: {
		111, 143, 174, 206, 237,
		269, 300, 331, 363, 394,
		426, 457, 489, 520, 552,
		583, 615, 646, 678, 709,
		741,
	},
	CRIT_RATE: {
		4.7, 6, 7.3, 8.6, 9.9,
		11.3, 12.6, 13.9, 15.2, 16.6,
		17.9, 19.2, 20.5, 21.8, 23.2,
		24.5, 25.8, 27.1, 28.4, 29.8,
		31.1,
	},
	CRIT_DMG: {
		9.3, 12, 14.6, 17.3, 19.9,
		22.5, 25.2, 27.8, 30.5, 33.1,
		35.7, 38.4, 41, 43.7, 46.3,
		49, 51.6, 54.2, 56.9, 59.5,
		62.2,
	},
	ENERGY_RECHARGE: {
		7.8, 10, 12.2, 14.4, 16.6,
		18.8, 21, 23.2, 25.4, 27.6,
		29.8, 32, 34.2, 36.4, 38.6,
		40.8, 43, 45.2, 47.4, 49.6,
		51.8,
	},
	ELEMENTAL_MASTARY: {
		28, 35.9, 43.8, 51.8,
		59.7, 67.6, 75.7, 83.5,
		91.4, 99.3, 107.2, 115.2,
		123.1, 131, 138.9, 146.9,
		154.8, 162.7, 170.6, 178.6,
		186.5,
	},
	HYDRO_DMG_BONUS: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	PYRO_DMG_BONUS: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	GEO_DMG_BONUS: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	ANEMO_DMG_BONUS: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	ELECTRO_DMG_BONUS: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	CRYO_DMG_BONUS: {
		7, 9, 11, 12.9, 14.9,
		16.9, 18.9, 20.9, 22.8, 24.8,
		26.8, 28.8, 30.8, 32.8, 34.7,
		36.7, 38.7, 40.7, 42.7, 44.6,
		46.6,
	},
	PHYSICAL_DMG_BONUS: {
		8.7, 11.2, 13.7, 16.2, 18.6,
		21.1, 23.6, 26.1, 28.6, 31,
		33.5, 36, 38.5, 41, 43.4,
		45.9, 48.4, 50.8, 53.3, 55.8,
		58.3,
	},
	HEALING_BONUS: {
		5.4, 6.9, 8.4, 10, 11.5,
		13, 14.6, 16.1, 17.6, 19.1,
		20.6, 22.1, 23.7, 25.2, 26.7,
		28.2, 29.8, 31.3, 32.8, 34.3,
		35.9,
	},
}

// 五星圣遗物副词条档位
var subStatMap = map[StatType][]float32{
	HP:                {209.13, 239.00, 268.88, 298.75},
	DEF:               {16.2, 18.52, 20.83, 23.15},
	ATK:               {13.62, 15.56, 17.51, 19.45},
	HP_big:            {4.08, 4.66, 5.25, 5.83},
	DEF_big:           {5.10, 5.83, 6.56, 7.29},
	ATK_big:           {4.08, 4.66, 5.25, 5.83},
	CRIT_RATE:         {2.72, 3.11, 3.50, 3.89},
	CRIT_DMG:          {5.44, 6.22, 6.99, 7.77},
	ELEMENTAL_MASTARY: {16.32, 18.65, 20.98, 23.31},
	ENERGY_RECHARGE:   {4.53, 5.18, 5.83, 6.48},
}

// 五星圣遗物升级所需经验
var experiencesToLevelUp = [...]int{
	3000, 3725, 4425, 5150, 5900,
	6675, 7500, 8350, 9225, 10125,
	11050, 12025, 13025, 15150, 17600,
	20375, 23500, 27050, 31050, 35575,
}

// 不同星级零级圣遗物狗粮对应经验值
var experienceOfDifferentStarsLv0 = [...]int{
	0, 420, 840, 1260, 1680, 2100,
}

// 五星不同级别圣遗物狗粮化经验
var experience5StarAsGarbage = [...]int{
	3780, 6180, 9160, 12700, 16820,
	21540, 26880, 32880, 39560, 46940,
	55040, 63880, 73500, 83920, 96040,
	110120, 126420, 145220, 166860, 191700,
	220160,
}
