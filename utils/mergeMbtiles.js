import fs from "fs";
import { execute } from "./execute.js";

export const mergeMbtiles = async (subRegions, keepSubRegionMbtiles) => {
  const pwd = process.cwd();

  const mbtilesToMerge = subRegions.map(
    (region) => region.split("/").pop() + ".mbtiles"
  );

  // Set up the mbtiles folder for the merged mbtiles file.
  if (!fs.existsSync(`${pwd}/mbtiles/merged`)) {
    fs.mkdirSync(`${pwd}/mbtiles/merged`);
  }

  // Set up the base planet.mbtiles file to merge all other files into. Overwrite existing file.
  fs.writeFileSync(`${pwd}/mbtiles/merged/planet.mbtiles`, "");

  execute(
    "tile-join",
    ["-o", "merged/planet.mbtiles", ...mbtilesToMerge],
    `${pwd}/mbtiles`
  );

  if (!keepSubRegionMbtiles) {
    // Delete the sub region mbtiles files if required.
    for (const region of subRegions) {
      fs.unlink(`${pwd}/mbtiles/${region.split("/").pop()}.mbtiles`, (err) => {
        if (err) {
          console.log(err);
        }

        console.log(`${region}.mbtiles deleted.`);
      });
    }
  }
};
