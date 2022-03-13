import { artifactsData } from "./assets/artifacts";
import { ArtifactDomain, ArtifactStatType, ArtifactSet, ArtifactSlot, Artifact, ArtifactStat } from "./type";

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

export const allSets: ArtifactSet[] = [
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

export function isPercentageStat(stat: ArtifactStatType){
  const integralStats: ArtifactStatType[] = ["ATK", "HP", "DEF", "EM"]
  return !integralStats.includes(stat);
}

export const slotMapper = {
  abbreviate: {
    "Circlet of Logos": "CIRCLET",
    "Goblet of Eonothem": "GOBLET",
    "Sands of Eon": "SANDS",
    "Plume of Death": "PLUME",
    "Flower of Life": "FLOWER"
  }
}

export function artifactFrom(obj: any): Artifact{
  return {
    stars: obj.Stars as number || 5,
    level: isNil(obj.Lv)? obj.Lv as number: 0,
    set: obj.Set as ArtifactSet || "Blizzard Strayer",
    slot: slotMapper.abbreviate[obj.Slot] as ArtifactSlot || "FLOWER",
    name: obj.Name as string || "Frost-Weaved Dignity",
    exp: isNil(obj.Lv)? obj.Lv as number: 0,
    mainStat: statFrom(obj.MainStat),
    subStats: (obj.SubStat || []).map(statFrom).filter(Boolean)
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