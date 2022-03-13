const _artifactsIcon = require("./icons/index.js");

let cache: any = {};

function importAll(r: any) {
  for (let path of r.keys()) {
    let name = r(path).default.eng;
    cache[name] = r(path).default;
  }
}

importAll(require.context("./data", true, /index\.js$/));

export const artifactsData = cache;
export const artifactsIcon = _artifactsIcon;
