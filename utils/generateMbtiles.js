import fs from "fs";
import { downloadFile } from "./downloadFile.js";
import { execute } from "./execute.js";
import { execSync, spawnSync } from "child_process";

const pwd = process.cwd();

export const generateMbtiles = async (
  subRegions,
  keepDownloadedFiles,
  tileZoomLevel
) => {
  if (!fs.existsSync(`${pwd}/pbf`)) {
    fs.mkdirSync(`${pwd}/pbf`);
  }

  // Set the tile-zoom level to use for the mbtiles generation.
  try {
    const env = fs.readFileSync(`${pwd}/openmaptiles/.env`, "utf-8");
    console.log(env);
    const result = env.replace(/MAX_ZOOM=[0-9]+/g, `MAX_ZOOM=${tileZoomLevel}`);
    console.log(result);
    fs.writeFileSync(`${pwd}/openmaptiles/.env`, result);
  } catch (e) {
    console.log(e);
    console.log(
      "Unable to find or write to openmaptiles/.env - using default zoom value of 7"
    );
  }

  // for (const region of subRegions) {
  //   // Clean up openmaptiles for new sub region
  //   execute("sudo", ["make", "clean"], `${pwd}/openmaptiles`);
  //   execute("rm", ["-rf", "data"], `${pwd}/openmaptiles`);

  //   // Download sub region .osm.pbf file
  //   const saveFileName = region.split("/").pop();
  //   if (!fs.existsSync(`${pwd}/pbf/${saveFileName}.osm.pbf`)) {
  //     const downloadUrl = `https://download.geofabrik.de/${region}-latest.osm.pbf`;
  //     await downloadFile(downloadUrl, saveFileName);
  //   } else {
  //     console.log(`${region}.osm.pbf already exists. Skipping download.`);
  //   }

  //   // Set up openmaptiles folder for new sub region
  //   if (!fs.existsSync(`${pwd}/openmaptiles/data`)) {
  //     fs.mkdirSync(`${pwd}/openmaptiles/data`);
  //   }
  //   const mvOrCp = keepDownloadedFiles ? "cp" : "mv";
  //   execute(mvOrCp, [
  //     `${pwd}/pbf/${saveFileName}.osm.pbf`,
  //     `${pwd}/openmaptiles/data/${saveFileName}.osm.pbf`,
  //   ]);

  //   // Generate mbtiles file for sub region
  //   execute("sudo", ["./quickstart.sh", saveFileName], `${pwd}/openmaptiles`);

  //   // Move the generated .mbtiles file to safety
  //   if (!fs.existsSync(`${pwd}/mbtiles`)) {
  //     fs.mkdirSync(`${pwd}/mbtiles`);
  //   }
  //   execute("mv", [
  //     `${pwd}/openmaptiles/data/tiles.mbtiles`,
  //     `${pwd}/mbtiles/${saveFileName}.mbtiles`,
  //   ]);
  // }
};
