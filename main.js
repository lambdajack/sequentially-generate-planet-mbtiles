#!/usr/bin/env node
import { execute } from "./utils/execute.js";
import { binaryExists } from "./utils/checkBinaryExists.js";

const cmdLineArguments = process.argv.slice(2);

const binaryName = "v2.0.0-sequentially-generate-planet-mbtiles.exe"

console.log(binaryExists(binaryName))

if (binaryExists(binaryName)) {
  execute("sudo", [
    `../release/${binaryName}`,
    ...cmdLineArguments,
  ]);
} else {
  console.log("sequentially-generate-planet-mbtiles binary not found. Exiting...");
  process.exit(1);
}
