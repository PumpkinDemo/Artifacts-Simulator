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
