import { formatBytes } from "./formatBytes.js";
import fs from "fs";
import https from "https";

const pwd = process.cwd();

export const downloadFile = async (url, saveFileName) => {
  const out = fs.createWriteStream(`${pwd}/pbf/${saveFileName}.osm.pbf`);
  let chunks = 0;
  let total = 0;
  let chunkCount = 0;

  return new Promise((resolve) => {
    https.get(url, (res) => {
      total = res.headers["content-length"];

      res.on("data", (chunk) => {
        chunks += chunk.length;
        chunkCount++;
        if (chunkCount === 200) {
          process.stdout.clearLine(0);
          process.stdout.cursorTo(0);
          process.stdout.write(
            "\r" + `${formatBytes(chunks)}/${formatBytes(total)}`
          );
          chunkCount = 0;
        }
      });
      res.pipe(out);
      res.on("end", () => {
        console.log(` ${url} Download Complete`);
        resolve();
      });
    });
  });
};
