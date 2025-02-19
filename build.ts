import { ensureDir, existsSync } from "jsr:@std/fs@1";
import { join } from "jsr:@std/path@1";

const versionFile = "VERSION.txt";
const distDir = "./dist";
const entryPoint = "./main.ts";
const binaryName = "formatbankcsv";
let version: string;

try {
  version = (await Deno.readTextFile(versionFile)).trim();
} catch (error) {
  // @ts-ignore deno-ts(18046)
  console.error(`Error reading ${versionFile}:`, error.message);
  Deno.exit(1);
}

if (!/^v\d+\.\d+\.\d+$/.test(version)) {
  console.error(
    `Invalid version format in ${versionFile}: "${version}". Expected format: vX.Y.Z`
  );
  Deno.exit(1);
}

const targets = [
  { denoTarget: "x86_64-apple-darwin", arch: "darwin-amd64" },
  { denoTarget: "aarch64-apple-darwin", arch: "darwin-arm64" },
  { denoTarget: "x86_64-unknown-linux-gnu", arch: "linux-amd64" },
  { denoTarget: "aarch64-unknown-linux-gnu", arch: "linux-arm64" },
];

async function compile(denoTarget: string, arch: string) {
  const tempDir = join(distDir, denoTarget);
  const binaryPath = join(tempDir, binaryName);
  const zipOutputPath = join(distDir, `${arch}-${version}.zip`);

  await ensureDir(distDir);
  await ensureDir(tempDir);

  console.log(`Compiling for ${denoTarget}...`);

  const compileProcess = new Deno.Command("deno", {
    args: [
      "compile",
      "--target",
      denoTarget,
      "--output",
      binaryPath,
      entryPoint,
    ],
    stdout: "inherit",
    stderr: "inherit",
  });

  const compileStatus = await compileProcess.output();
  if (!compileStatus.success) {
    console.error(`Failed to compile for target ${denoTarget}`);
    await Deno.remove(tempDir, { recursive: true });
    return;
  }

  console.log(`Zipping ${zipOutputPath}...`);

  const zipProcess = new Deno.Command("zip", {
    args: ["-j", zipOutputPath, binaryPath],
    stdout: "inherit",
    stderr: "inherit",
  });

  const zipStatus = await zipProcess.output();
  if (!zipStatus.success) {
    console.error(`Failed to create zip for ${zipOutputPath}`);
    await Deno.remove(tempDir, { recursive: true });
    return;
  }

  console.log(`Successfully created ${zipOutputPath}`, { recursive: true });
  await Deno.remove(tempDir, { recursive: true });
}

existsSync(distDir) && Deno.removeSync(distDir, { recursive: true });
for (const { denoTarget, arch } of targets) {
  await compile(denoTarget, arch);
}
