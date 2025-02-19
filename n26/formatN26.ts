import { hasColumns } from "../utils/hasColumns.ts";
import { Header } from "../utils/constants.ts";

export function formatN26(inputLines: string[][]): string[][] {
  const requiredColumns = ["Booking Date", "Partner Name", "Payment Reference", "Amount (EUR)"];

  const header = inputLines[0];
  if (!hasColumns(header, requiredColumns)) {
    throw new Error("Unexpected header for n26 format");
  }

  const columnIndices = new Map<string, number>();
  for (let i = 0; i < header.length; i++) {
    columnIndices.set(header[i], i);
  }

  const requiredIndices: number[] = [];
  for (const col of requiredColumns) {
    for (let i = 0; i < header.length; i++) {
      if (header[i] === col) {
        requiredIndices.push(i);
        break;
      }
    }
  }

  const outputLines: string[][] = [[...Header]];
  
  for (const row of inputLines.slice(1)) {
    const extractedRow = requiredIndices.map(index => row[index]);
    outputLines.push(extractedRow);
  }

  return outputLines;
}
