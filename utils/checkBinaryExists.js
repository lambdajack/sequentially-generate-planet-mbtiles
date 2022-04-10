import fs from "fs";

const pwd = process.cwd();

export const binaryExists = (binaryName) => {
  return fs.existsSync(
    `${pwd}/release/${binaryName}`
  );
}