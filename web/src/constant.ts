/* 常用数据/常用函数文件 */

import { artifactsData } from "./assets/artifacts";
import { ArtifactDomain, ArtifactStatType, ArtifactSet, ArtifactSlot, Artifact, ArtifactStat, ArtifactStatNonnullable } from "./type";

function isNil(value: any){
  return value === undefined || value === null;
}

export const allDomains: ArtifactDomain[] = [
  "Domain of Guyun",
  "Midsummer Courtyard",
  "Valley of Remembrance",
  "Hidden Palace of Zhou Formula",
  "Peak of Vindagnyr",
  "Ridge Watch",
  "Momiji-Dyed Court",
  "Slumbering Court",
  "Clear Pool and Mountain Cavern",
];

const allStar3Sets: ArtifactSet[] = [
  "Adventurer", 
  "Lucky Dog", 
  "Traveling Doctor"
]

const allStar4Sets: ArtifactSet[] = [
  "Resolution of Sojourner", 
  "Tiny Miracle", 
  "Berserker", 
  "Instructor", 
  "The Exile", 
  "Defender's Will", 
  "Brave Heart", 
  "Martial Artist", 
  "Gambler", 
  "Scholar", 
  "Prayers for Illumination", 
  "Prayers for Destiny", 
  "Prayers for Wisdom", 
  "Prayers to Springtime"
]

const allStar5Sets: ArtifactSet[] = [
  "Ocean Hued Clam",
  "Husk of Opulent Dreams",
  "Emblem of Severed Fate",
  "Shimenawas Reminiscence",
  "Pale Flame",
  "Tenacity of the Millelith",
  "Heart of Depth",
  "Blizzard Strayer",
  "Crimson Witch of Flames",
  "Lavawalker",
  "Thundering Fury",
  "Thundersmoother",
  "Retracing Bolide",
  "Archaic Petra",
  "Viridescent Venerer",
  "Maiden Beloved",
  "Bloodstained Chivalry",
  "Noblesse Oblige",
  "Wanderer's Troupe",
  "Gladiator's Finale",
];

export const allSets = [
  [],
  allStar3Sets,
  allStar3Sets,
  allStar3Sets.concat(allStar4Sets),
  allStar4Sets.concat(allStar5Sets),
  allStar5Sets
]

export const allSlots: ArtifactSlot[] = [
  "CIRCLET",
  "GOBLET",
  "PLUME",
  "FLOWER",
  "SANDS",
];

export const allStats: ArtifactStatType[] = [
  "ATK%",
  "ATK",
  "HP%",
  "HP",
  "DEF%",
  "DEF",
  "CRIT RATE",
  "CRIT DMG",
  "ER",
  "EM",
  "HYDRO BONUS",
  "PYRO BONUS",
  "GEO BONUS",
  "ANEMO BONUS",
  "ELECTRO BONUS",
  "CRYO BONUS",
  "PHYSICAL BONUS",
  "HEALING BONUS",
];

function statFrom(obj: any): ArtifactStat{
  const { stat = null, value = 239 } = obj;
  return stat? {
    statType: stat as ArtifactStatType,
    value: value as number
  }: null;
}

function statJSON(stat: ArtifactStat) {
  const { statType = null, value = 239 } = stat as ArtifactStatNonnullable;
  return {
    stat: statType || '',
    value
  };
}

export function isPercentageStat(stat: ArtifactStatType){
  const integralStats: ArtifactStatType[] = ["ATK", "HP", "DEF", "EM"]
  return !integralStats.includes(stat);
}

const slotMapper = {
  abbreviate: {
    "Circlet of Logos": "CIRCLET",
    "Goblet of Eonothem": "GOBLET",
    "Sands of Eon": "SANDS",
    "Plume of Death": "PLUME",
    "Flower of Life": "FLOWER"
  },
  deabbreviate: {
    "CIRCLET": "Circlet of Logos",
    "GOBLET": "Goblet of Eonothem",
    "SANDS": "Sands of Eon",
    "PLUME": "Plume of Death",
    "FLOWER": "Flower of Life"
  }
}

export function artifactFrom(obj: any): Artifact{
  return {
    stars: obj.Stars as number || 5,
    level: isNil(obj.Lv)? obj.Lv as number: 0,
    set: obj.Set as ArtifactSet || "Blizzard Strayer",
    slot: slotMapper.abbreviate[obj.Slot] as ArtifactSlot || "FLOWER",
    name: obj.Name as string || "Frost-Weaved Dignity",
    exp: isNil(obj.Exp)? obj.Exp as number: 0,
    mainStat: statFrom(obj.MainStat),
    subStats: (obj.SubStat || []).map(statFrom).filter(Boolean)
  }
}

export function artifactJSON(a: Artifact): any{
  return {
    Stars: a.stars || 5,
    Lv: a.level || 0,
    Set: a.set || "Blizzard Strayer",
    Slot: slotMapper.deabbreviate[a.slot] || "FLOWER",
    Name: a.name || "Frost-Weaved Dignity",
    Exp: a.exp || 0,
    MainStat: statJSON(a.mainStat),
    SubStat: (a.subStats || []).map(statJSON)
  }
}

export function getArtifactImageUrl(artifact: ArtifactSet, slot: ArtifactSlot){
  const camelize = (str: ArtifactSet) => {
   const bigCamel = str.replace('\'', '').replace(/\s(\w)/g, (space, letter) => letter.toUpperCase());
   return bigCamel.substring(0, 1).toLowerCase() + bigCamel.substring(1);
  }
  const setType = camelize(artifact);
  const slotType = slot.toLowerCase();
  return artifactsData[setType][slotType]['url'];
}