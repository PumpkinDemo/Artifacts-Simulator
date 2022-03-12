import chs from "./chs.json";
import { StringMapper, Language, AppOptions } from "../type";

const mapper = {
  CHS: chs as StringMapper,
};

export const languageOptions = [
  {
    value: "CHS",
    label: "简体中文",
  },
  {
    value: "EN",
    label: "English",
  },
];

export function makeOptions(texts: string[], language: Language): AppOptions {
  return texts.map((text) => {
    return {
      label: language === "EN" ? text : mapper[language][text] || text,
      value: text,
    };
  });
}

export function locale(text: string, language: Language): string {
  if (language === "EN") {
    return text;
  } else {
    return mapper[language][text] || text;
  }
}
