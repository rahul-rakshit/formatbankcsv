import { parseArgs } from "@std/cli";
import { handleOptions } from "./handleOptions.ts";
import { helpMessage } from "./help.ts";

if (import.meta.main) {
  const args = parseArgs(Deno.args, {
    string: ["input", "output", "format", "from", "to"],
    boolean: ["help"],
    alias: {
      i: "input",
      o: "output",
      F: "format",
      f: "from",
      t: "to",
      h: "help",
    },
  });

  if (args.help) {
    console.log(helpMessage);
    Deno.exit(0);
  }

  if (!args.input || !args.output || !args.format) {
    console.error(helpMessage);
    Deno.exit(1);
  }

  const format = args.format.toLowerCase();
  if (format !== "dkb" && format !== "n26") {
    console.error('Format must be either "dkb" or "n26"');
    Deno.exit(1);
  }

  try {
    await handleOptions(
      args.input,
      args.output,
      args.from || "",
      args.to || "",
      format
    );
  } catch (error: unknown) {
    if (error instanceof Error) {
      console.error(error.message);
    } else {
      console.error("An unknown error occurred");
    }
    Deno.exit(2);
  }
}
