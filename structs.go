package main

import "encoding/json"

type StatType int32

const (
	NIL                StatType = -1
	ATK_big            StatType = 0
	ATK                StatType = 1
	HP_big             StatType = 2
	HP                 StatType = 3
	DEF_big            StatType = 4
	DEF                StatType = 5
	CRIT_RATE          StatType = 6
	CRIT_DMG           StatType = 7
	ENERGY_RECHARGE    StatType = 8
	ELEMENTAL_MASTARY  StatType = 9
	HYDRO_DMG_BONUS    StatType = 10
	PYRO_DMG_BONUS     StatType = 11
	GEO_DMG_BONUS      StatType = 12
	ANEMO_DMG_BONUS    StatType = 13
	ELECTRO_DMG_BONUS  StatType = 14
	CRYO_DMG_BONUS     StatType = 15
	PHYSICAL_DMG_BONUS StatType = 16
	HEALING_BONUS      StatType = 17
)

type SetType int32

const (
	Ocean_Hued_Clam           SetType = 0
	Husk_of_Opulent_Dreams    SetType = 1
	Emblem_of_Severed_Fate    SetType = 2
	Shimenawas_Reminiscence   SetType = 3
	Pale_Flame                SetType = 4
	Tenacity_of_the_Millelith SetType = 5
	Heart_of_Depth            SetType = 6
	Blizzard_Strayer          SetType = 7
	Crimson_Witch_of_Flames   SetType = 8
	Lavawalker                SetType = 9
	Thundering_Fury           SetType = 10
	Thundersoother            SetType = 11
	Retracing_Bolide          SetType = 12
	Archaic_Petra             SetType = 13
	Viridescent_Venerer       SetType = 14
	Maiden_Beloved            SetType = 15
	Bloodstained_Chivalry     SetType = 16
	Noblesse_Oblige           SetType = 17
	Wanderers_Troupe          SetType = 18
	Gladiators_Finale         SetType = 19
)

type SlotType int32

const (
	Flower_of_Life     SlotType = 0
	Plume_of_Death     SlotType = 1
	Sands_of_Eon       SlotType = 2
	Goblet_of_Eonothem SlotType = 3
	Circlet_of_Logos   SlotType = 4
)

type Stat struct {
	Type  StatType `json:"词条"`
	Value float32  `json:"数值"`
}

type Artifact struct {
	Stars    int32
	Lv       int32
	Name     string
	Set      SetType
	Slot     SlotType
	MainStat Stat
	SubStat  [4]Stat
	Exp      int32
}

func (s *Stat) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  string  `json:"词条"`
		Value float32 `json:"数值"`
	}{
		Type:  statTypeZN[s.Type],
		Value: s.Value,
	})
}

func (a *Artifact) MarshalJSON() ([]byte, error) {
	type t Artifact
	return json.Marshal(&struct {
		Set string
		*t
	}{
		Set: artifactSetZN[a.Set],
		t:   (*t)(a),
	})
}
