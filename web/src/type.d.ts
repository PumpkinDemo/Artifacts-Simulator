export type Language = "CHS" | "EN";

export interface StringMapper {
  [prop: string]: string;
}

interface AppOption {
  label: string;
  value: string;
  [prop: string]: string;
}

export type AppOptions = AppOption[];

export type ArtifactSlot = "CIRCLET" | "GOBLET" | "SANDS" | "PLUME" | "FLOWER";

export type ArtifactSet = "Ocean Hued Clam" | "Husk of Opulent Dreams" | "Emblem of Severed Fate" | "Shimenawas Reminiscence" | "Pale Flame" | "Tenacity of the Millelith" | "Heart of Depth" | "Blizzard Strayer" | "Crimson Witch of Flames" | "Lavawalker" | "Thundering Fury" | "Thundersoother" | "Retracing Bolide" | "Archaic Petra" | "Viridescent Venerer" | "Maiden Beloved" | "Bloodstained Chivalry" | "Noblesse Oblige" | "Wanderers Troupe" | "Gladiators Finale";

export type ArtifactDomain = "Domain of Guyun" | "Midsummer Courtyard" | "Valley of Remembrance" | "Hidden Palace of Zhou Formula" | "Peak of Vindagnyr" | "Ridge Watch" | "Momiji-Dyed Court" | "Slumbering Court" | "Clear Pool and Mountain Cavern";

export type ArtifactStat = "ATK%" | "ATK" | "HP%" | "HP" | "DEF%" | "DEF" | "CRIT RATE" | "CRIT DMG" | "ER" | "EM" | "HYDRO BONUS" | "PYRO BONUS" | "GEO BONUS" | "ANEMO BONUS" | "ELECTRO BONUS" | "CRYO BONUS" | "PHYSICAL BONUS" | "HEALING BONUS";
