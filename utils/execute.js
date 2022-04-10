import { spawnSync } from "child_process";

export const execute = (cmd, argv, cwd = process.cwd()) => {
  try {
    spawnSync(cmd, argv, { stdio: "inherit", cwd: cwd });
  } catch (e) {
    console.log(e);
    process.exit(1);
  }
};