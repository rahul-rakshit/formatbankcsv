import { parse, stringify } from "@std/csv";

/**
 * Reads a CSV file with the specified delimiter and number of lines to skip
 * @param inputPath Path to the input CSV file
 * @param delimiter The delimiter character used in the CSV
 * @param skipLines Number of lines to skip before parsing CSV content
 */
export async function readCsv(
  inputPath: string,
  delimiter: string,
  skipLines: number,
): Promise<string[][]> {
  const fileContent = await Deno.readTextFile(inputPath);
  const lines = fileContent.split("\n");
  const nonEmptyContent = lines.slice(skipLines).join("\n");

  const rows = parse(nonEmptyContent, {
    separator: delimiter,
  }).filter((row) => row.some((cell) => cell !== ""));

  if (rows.length === 0) return rows;
  const headerLength = rows[0].length;

  // Pad rows with empty strings if they have fewer columns than the header
  return rows.map(row => {
    if (row.length < headerLength) {
      return [...row, ...Array(headerLength - row.length).fill("")];
    }
    return row;
  });
}

/**
 * Writes CSV data to a file with the specified delimiter
 * @param data Array of string arrays representing CSV data
 * @param outputPath Path where the CSV file should be written
 * @param delimiter The delimiter character to use in the CSV
 */
export async function writeCsv(
  data: string[][],
  outputPath: string,
  delimiter: string,
): Promise<void> {
  const csvContent = stringify(data, {
    separator: delimiter,
  });

  await Deno.writeTextFile(outputPath, csvContent);
}
