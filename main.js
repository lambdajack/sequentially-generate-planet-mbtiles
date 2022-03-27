#!/usr/bin/env node

import fs from "fs";
import { execSync } from "child_process";
import { generateMbtiles } from "./utils/generateMbtiles.js";
import { mergeMbtiles } from "./utils/mergeMbtiles.js";
import { execute } from "./utils/execute.js";

const pwd = process.cwd();
let config = {};

// Load config.json
if (process.argv.includes("-c")) {
  try {
    const c = process.argv.indexOf("-c");
    const configFile = process.argv[c + 1];
    config = JSON.parse(fs.readFileSync(configFile, "utf-8"));
  } catch (e) {
    console.log(e);
    console.log(
      "No config.json file found, or the supplied config has an unexpected format. Using default config."
    );
    try {
      config = JSON.parse(fs.readFileSync(`${pwd}/config.json`, "utf-8"));
    } catch (e) {
      console.log(e);
      console.log("Cannot find default config.json file... this is a problem.");
    }
  }
} else {
  try {
    config = JSON.parse(fs.readFileSync(`${pwd}/config.json`, "utf-8"));
  } catch (e) {
    console.log(e);
    console.log("Cannot find default config.json file... this is a problem.");
  }
}

if (!config.subRegions) {
  console.log("No sub regions found in config.json");
  process.exit(1);
}

const requiredRepos = [
  {
    dir: `${pwd}/openmaptiles`,
    github: "https://github.com/openmaptiles/openmaptiles",
  },
  {
    dir: `${pwd}/tippecanoe`,
    github: "https://github.com/mapbox/tippecanoe",
  },
];

// STEP 1 - Download openmaptiles/openmaptiles and mapbox/tippecanoe and install them
console.log("Downloading required packages...");
for (const repo of requiredRepos) {
  if (!fs.existsSync(repo.dir)) {
    fs.mkdirSync(repo.dir);
    console.log(`Cloning ${repo.github}...`);
    execute("git", ["clone", repo.github, repo.dir]);
  } else {
    console.log(`${repo.dir} already exists. Skipping download.`);
  }
}

// Install tippecanoe
console.log(
  "Globally installing tippecanoe which is required for combining .mbtiles files..."
);
execute("make", ["-j"], `${pwd}/tippecanoe`);
execute("sudo", ["make", "install"], `${pwd}/tippecanoe`);

// STEP 2 - Define how to divide up the work (default by sub region). Must follow the same naming convention as the geofabirk .osm.pbf files. e.g. "australia-oceania-latest.osm.pbf" should be "australia-oceania"; "chad-latest.osm.pbf" should be "africa/chad".
const subRegions = config.subRegions;
const keepDownloadedFiles = config.keepDownloadedFiles || false;
const keepSubRegionMbtiles = config.keepSubRegionMbtiles || false;
const tileZoomLevel = config.tileZoomLevel || 14;

console.log("Downloading the following sub regions:", subRegions);

// STEP 3 - Generate mbtiles files for each sub region
await generateMbtiles(subRegions, keepDownloadedFiles, tileZoomLevel);

// STEP 4 - Merge mbtiles into a single mbtiles file
await mergeMbtiles(subRegions, keepSubRegionMbtiles);

console.log(
  "--------------------------------------------------------------------------------"
);
console.log("PHEW......... You made it! That probably took a long time!");
console.log(
  `
  Thank you for using lambdajack/sequentially-generate-planet-mbtiles!
  If you have any questions, please contact me on github.com/lambdajack
  All contributions are welcome! Please share what you made with our help.
  Our thanks to the wonderful openmaptiles and tippecanoe for making this possible!
  `
);
console.log(
  `

  --------------------------------------------------------------------------------
  Your planet.mbtiles files is located at ${pwd}/mbtiles/merged/planet.mbtiles. Check the github repo for more information on serving them if you wish. 
  --------------------------------------------------------------------------------`
);
