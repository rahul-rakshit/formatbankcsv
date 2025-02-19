import { readCsv, writeCsv } from "../csv/csv.ts";
import { filterByDate } from "../utils/filterByDate.ts";
import { formatDkb } from "./formatDkb.ts";

export class DkbFormatter {
  async readCsv(inputPath: string): Promise<string[][]> {
    const delimiter = ";";
    const skipLines = 4;
    return await readCsv(inputPath, delimiter, skipLines);
  }

  format(csvData: string[][], fromDate: string, toDate: string): string[][] {
    const formatted = formatDkb(csvData);
    if (fromDate || toDate) {
      return filterByDate(formatted, fromDate, toDate);
    }
    return formatted;
  }

  async writeCsv(data: string[][], outputPath: string): Promise<void> {
    const delimiter = ",";
    return await writeCsv(data, outputPath, delimiter);
  }
}
