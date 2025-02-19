import { N26Formatter } from "./n26/n26.ts";
import { DkbFormatter } from "./dkb/dkb.ts";
import { CsvFormatter } from "./csvFormatter.ts";

export async function handleOptions(
  inputPath: string,
  outputPath: string,
  fromDate: string,
  toDate: string,
  format: string,
): Promise<void> {
  const lowerCaseFormat = format.toLowerCase();
  let formatter: CsvFormatter;

  if (lowerCaseFormat === "n26") {
    formatter = new N26Formatter();
  } else if (lowerCaseFormat === "dkb") {
    formatter = new DkbFormatter();
  } else {
    throw new Error('Format must be either "dkb" or "n26"');
  }

  const csvData = await formatter.readCsv(inputPath);
  const formatted = formatter.format(csvData, fromDate, toDate);
  await formatter.writeCsv(formatted, outputPath);
}
