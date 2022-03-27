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
  if (!fs.existsSync(`${pwd}/mbtiles/merged/planet.mbtiles`)) {
    fs.writeFileSync(`${pwd}/mbtiles/merged/planet.mbtiles`, "");
  }

  const mbtilesPreFileSize = fs.statSync(
    `${pwd}/mbtiles/merged/planet.mbtiles`
  ).size;

  execute(
    "tile-join",
    ["-o", "merged/planet.mbtiles", ...mbtilesToMerge],
    `${pwd}/mbtiles`
  );

  const mbtilesPostFileSize = fs.statSync(
    `${pwd}/mbtiles/merged/planet.mbtiles`
  ).size;

  console.log(`${mbtilesPreFileSize} => ${mbtilesPostFileSize}`);
  if (mbtilesPreFileSize < mbtilesPostFileSize) {
    console.log(
      "mbtiles file size increased. The process appears to have completed successfully."
    );
    fs.appendFileSync(
      `${pwd}/mbtiles/REPORT.txt`,
      `
    --------------------------------------------------
    SUCCESSFULLY MERGED MBTILES FILES - see above for any tilesets which failed to be created, and therefore failed to be merged.

    Remember to check the files before production.
    --------------------------------------------------
    `
    );
  } else {
    console.log(
      "mbtiles file size did not increase. The process appears to have failed."
    );
    fs.appendFileSync(
      `${pwd}/mbtiles/REPORT.txt`,
      `
      --------------------------------------------------
      FAILED TO MERGE MBTILES FILES - see above for any tilesets which failed to be created. 
      
      If you are seeing this message it is becuase either no (new) mbtiles were provided to be merged (i.e. there were no sub regions), or the joining process failed for some reason.
      --------------------------------------------------
      `
    );
  }

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
