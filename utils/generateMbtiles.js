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

  // Set up REPORT.txt and mbtiles folder
  if (!fs.existsSync(`${pwd}/mbtiles`)) {
    fs.mkdirSync(`${pwd}/mbtiles`);
  }
  if (!fs.existsSync(`${pwd}/mbtiles/REPORT.txt`)) {
    fs.writeFileSync(`${pwd}/mbtiles/REPORT.txt`, "");
  }

  for (const region of subRegions) {
    // Check to see if mbtiles already exist for the region. If so, skip their generation.
    const saveFileName = region.split("/").pop();
    if (
      fs.existsSync(`${pwd}/mbtiles/${saveFileName}.mbtiles`) &&
      fs.statSync(`${pwd}/mbtiles/${saveFileName}.mbtiles`).size > 0
    ) {
      console.log(`${region}.mbtiles already exists. Skipping generation.`);
      continue;
    }

    // Clean up openmaptiles for new sub region
    execute("sudo", ["make", "clean"], `${pwd}/openmaptiles`);
    execute("rm", ["-rf", "data"], `${pwd}/openmaptiles`);

    // Download sub region .osm.pbf file
    if (!fs.existsSync(`${pwd}/pbf/${saveFileName}.osm.pbf`)) {
      const downloadUrl = `https://download.geofabrik.de/${region}-latest.osm.pbf`;
      await downloadFile(downloadUrl, saveFileName);
    } else {
      console.log(`${region}.osm.pbf already exists. Skipping download.`);
    }

    // Set up openmaptiles folder for new sub region
    if (!fs.existsSync(`${pwd}/openmaptiles/data`)) {
      fs.mkdirSync(`${pwd}/openmaptiles/data`);
    }
    const mvOrCp = keepDownloadedFiles ? "cp" : "mv";
    execute(mvOrCp, [
      `${pwd}/pbf/${saveFileName}.osm.pbf`,
      `${pwd}/openmaptiles/data/${saveFileName}.osm.pbf`,
    ]);

    // Generate mbtiles file for sub region
    execute("sudo", ["./quickstart.sh", saveFileName], `${pwd}/openmaptiles`);

    // Move the generated .mbtiles file to safety
    if (!fs.existsSync(`${pwd}/mbtiles`)) {
      fs.mkdirSync(`${pwd}/mbtiles`);
    }
    execute("mv", [
      `${pwd}/openmaptiles/data/tiles.mbtiles`,
      `${pwd}/mbtiles/${saveFileName}.mbtiles`,
    ]);

    // Veryfy that the mbtiles file was generated successfully.
    const mbtilesFile = `${pwd}/mbtiles/${saveFileName}.mbtiles`;
    if (!fs.existsSync(mbtilesFile)) {
      console.log(
        `${mbtilesFile} does not exist. The process failed to complete for ${region}. Moving on to next process regardless.`
      );
      fs.appendFileSync(
        `${pwd}/mbtiles/REPORT.txt`,
        `GENERATE FAILED: ${region}: Failed to generate mbtiles.\n`
      );
    } else {
      fs.appendFileSync(
        `${pwd}/mbtiles/REPORT.txt`,
        `GENERATE SUCCESS: ${region}: Successfully generated mbtiles.\n`
      );
    }
  }
};
