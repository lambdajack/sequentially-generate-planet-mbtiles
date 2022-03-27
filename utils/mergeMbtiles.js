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

  const mbtilesInitialFileSize = fs.statSync(
    `${pwd}/mbtiles/merged/planet.mbtiles`
  ).size;

  for (const tileset of mbtilesToMerge) {
    if (
      fs
        .readFileSync(`${pwd}/mbtiles/REPORT.txt`)
        .includes(
          `MERGE SUCCESS: ${tileset} successfully merged in planet.mbtiles.\n`
        )
    ) {
      console.log(`${tileset} already merged in planet.mbtiles. Skipping.`);
      fs.appendFileSync(
        `${pwd}/mbtiles/REPORT.txt`,
        `MERGE SKIPPED: ${tileset} appears to be already merged, based on the contents of this report.\n`
      );
      continue;
    }

    // Rename existing planet.mbtiles
    if (fs.existsSync(`${pwd}/mbtiles/merged/planet.mbtiles`)) {
      fs.renameSync(
        `${pwd}/mbtiles/merged/planet.mbtiles`,
        `${pwd}/mbtiles/merged/temp.mbtiles`
      );
    }

    if (!fs.existsSync(`${pwd}/mbtiles/merged/planet.mbtiles`)) {
      fs.writeFileSync(`${pwd}/mbtiles/merged/planet.mbtiles`, "");
    }

    const mbtilesPreFileSize = fs.statSync(
      `${pwd}/mbtiles/merged/planet.mbtiles`
    ).size;

    execute(
      "tile-join",
      ["-o", "merged/planet.mbtiles", tileset, "merged/temp.mbtiles"],
      `${pwd}/mbtiles`
    );

    const mbtilesPostFileSize = fs.statSync(
      `${pwd}/mbtiles/merged/planet.mbtiles`
    ).size;

    if (mbtilesPreFileSize === mbtilesPostFileSize) {
      console.log(`mbtiles file size is the same; failed to merge ${tileset}`);
      fs.appendFileSync(
        `${pwd}/mbtiles/REPORT.txt`,
        `MERGE FAILED: ${tileset} failed to be merged in planet.mbtiles.\n`
      );
    } else if (mbtilesPreFileSize < mbtilesPostFileSize) {
      console.log(
        `mbtiles file size is larger; successfully merged ${tileset}`
      );
      fs.rmSync(`${pwd}/mbtiles/merged/temp.mbtiles`);
      fs.appendFileSync(
        `${pwd}/mbtiles/REPORT.txt`,
        `MERGE SUCCESS: ${tileset} successfully merged in planet.mbtiles.\n`
      );
    } else {
      console.log(
        "mbtiles is smaller after merge. Something unexpected happened."
      );
    }
  }

  const mbtilesFinalFileSize = fs.statSync(
    `${pwd}/mbtiles/merged/planet.mbtiles`
  ).size;

  console.log(`${mbtilesInitialFileSize} => ${mbtilesFinalFileSize}`);
  if (mbtilesInitialFileSize < mbtilesFinalFileSize) {
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
      "mbtiles file size did not increase. The process appears to have failed, or no new tilesets were supplied based on the contents of REPORT.txt."
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
