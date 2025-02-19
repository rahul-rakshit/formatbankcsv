export interface CsvFormatter {
  readCsv(inputPath: string): Promise<string[][]>;
  format(csvData: string[][], fromDate: string, toDate: string): string[][];
  writeCsv(data: string[][], outputPath: string): Promise<void>;
}
