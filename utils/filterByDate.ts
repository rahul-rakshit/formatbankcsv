import { Header } from "./constants.ts";

function validateDate(dateStr: string): boolean {
  if (dateStr === "" || dateStr === Header[0]) return true;
  const date = new Date(dateStr);
  return date instanceof Date && !isNaN(date.getTime());
}

function validateDates(row: string[], fromDate: string, toDate: string): void {
  const rowDate = row[0];
  
  if (!validateDate(rowDate) && rowDate !== Header[0]) {
    throw new Error(`Invalid date format: ${rowDate}`);
  }
  if (!validateDate(fromDate)) {
    throw new Error(`Invalid fromDate format: ${fromDate}`);
  }
  if (!validateDate(toDate)) {
    throw new Error(`Invalid toDate format: ${toDate}`);
  }
}

function isHeaderRow(row: string[]): boolean {
  return row[0] === Header[0];
}

function meetsFromDateBound(row: string[], fromDateString: string): boolean {
  if (fromDateString === "") {
    return true;
  }

  const rowDate = new Date(row[0]);
  const fromDate = new Date(fromDateString);

  return rowDate >= fromDate;
}

function meetsToDateBound(row: string[], toDateString: string): boolean {
  if (toDateString === "") {
    return true;
  }

  const rowDate = new Date(row[0]);
  const toDate = new Date(toDateString);

  return rowDate <= toDate;
}

export function filterByDate(
  data: string[][],
  fromDate: string,
  toDate: string
): string[][] {
  const filteredData: string[][] = [];

  for (const row of data) {
    try {
      validateDates(row, fromDate, toDate);
    } catch (error) {
      throw error;
    }

    if (isHeaderRow(row) || (meetsFromDateBound(row, fromDate) && meetsToDateBound(row, toDate))) {
      filteredData.push(row);
    }
  }

  return filteredData;
}
