#!/usr/bin/env node
import { execute } from "./utils/execute.js";
import { binaryExists } from "./utils/checkBinaryExists.js";

const cmdLineArguments = process.argv.slice(2);

if (binaryExists) {
  execute("sudo", [
    "./releases/sequentially-generate-planet-mbtiles",
    ...cmdLineArguments,
  ]);
} else {
  console.log("sequentially-generate-planet-mbtiles not found. Exiting...");
  process.exit(1);
}
