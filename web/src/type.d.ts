/* 类型和接口文件 */

export type Language = "CHS" | "EN";

export interface StringMapper {
  [prop: string]: string;
}

interface AppOption {
  label: string;
  value: string;
  [prop: string]: string;
}

export interface DogFood {
  stars: number;
  level: number;
}

export interface ArtifactStatNonnullable {
  statType: ArtifactStatType;
  value: number;
}

export type ArtifactStat = ArtifactStatNonnullable | null;
export interface Artifact {
	stars: number;
	level: number;
	name: string;
	exp: number;
	set: ArtifactSet;
	slot: ArtifactSlot;
	mainStat: ArtifactStat;
	subStats: ArtifactStat[];
}

export type AppOptions = AppOption[];

export type ArtifactSlot = "CIRCLET" | "GOBLET" | "SANDS" | "PLUME" | "FLOWER";

export type ArtifactSet = "Ocean Hued Clam" | "Husk of Opulent Dreams" | "Emblem of Severed Fate" | "Shimenawas Reminiscence" | "Pale Flame" | "Tenacity of the Millelith" | "Heart of Depth" | "Blizzard Strayer" | "Crimson Witch of Flames" | "Lavawalker" | "Thundering Fury" | "Thundersmoother" | "Retracing Bolide" | "Archaic Petra" | "Viridescent Venerer" | "Maiden Beloved" | "Bloodstained Chivalry" | "Noblesse Oblige" | "Wanderer's Troupe" | "Gladiator's Finale"
| "Adventurer" | "Lucky Dog" | "Traveling Doctor"
| "Resolution of Sojourner" | "Tiny Miracle" | "Berserker" | "Instructor" | "The Exile" | "Defender's Will" | "Brave Heart" | "Martial Artist" | "Gambler" | "Scholar" | "Prayers for Illumination" | "Prayers for Destiny" | "Prayers for Wisdom" | "Prayers to Springtime";

export type ArtifactDomain = "Domain of Guyun" | "Midsummer Courtyard" | "Valley of Remembrance" | "Hidden Palace of Zhou Formula" | "Peak of Vindagnyr" | "Ridge Watch" | "Momiji-Dyed Court" | "Slumbering Court" | "Clear Pool and Mountain Cavern";

export type ArtifactStatType = "ATK%" | "ATK" | "HP%" | "HP" | "DEF%" | "DEF" | "CRIT RATE" | "CRIT DMG" | "ER" | "EM" | "HYDRO BONUS" | "PYRO BONUS" | "GEO BONUS" | "ANEMO BONUS" | "ELECTRO BONUS" | "CRYO BONUS" | "PHYSICAL BONUS" | "HEALING BONUS" | "NIL";
