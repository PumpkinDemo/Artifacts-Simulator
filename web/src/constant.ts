import { artifactsData, artifactsIcon } from "./assets/artifacts";
import { ArtifactDomain, ArtifactStat, ArtifactSet, ArtifactSlot } from "./type";

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
  "Thundersoother",
  "Retracing Bolide",
  "Archaic Petra",
  "Viridescent Venerer",
  "Maiden Beloved",
  "Bloodstained Chivalry",
  "Noblesse Oblige",
  "Wanderers Troupe",
  "Gladiators Finale",
];

export const allStats: ArtifactStat[] = [
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

export function getArtifactImage(artifact: ArtifactSet | null, name: ArtifactSlot){
  if(artifact){

  }else{
    return 
  }
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