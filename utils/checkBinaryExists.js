import fs from "fs";

export const binaryExists = (binaryName) => {
  return fs.existsSync(
    `../release/${binaryName}`
  );
}