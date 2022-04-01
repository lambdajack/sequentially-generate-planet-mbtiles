import fs from "fs";

const pwd = process.cwd();
export const binaryExists = fs.existsSync(
  `${pwd}/sequentially-generate-mbtiles.json`
);
