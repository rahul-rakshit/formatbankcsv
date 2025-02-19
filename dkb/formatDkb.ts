import { Header } from "../utils/constants.ts";
import { isLineEqual } from "../utils/isLineEqual.ts";
import { sortByDate } from "../utils/sortByDate.ts";
import { convertDkbAmount } from "./convertDkbAmount.ts";
import { convertDkbDate } from "./convertDkbDate.ts";

const expectedInputHeader = [
  "Buchungsdatum",
  "Wertstellung",
  "Status",
  "Zahlungspflichtige*r",
  "Zahlungsempfänger*in",
  "Verwendungszweck",
  "Umsatztyp",
  "IBAN",
  "Betrag (€)",
  "Gläubiger-ID",
  "Mandatsreferenz",
  "Kundenreferenz",
];

export function formatDkb(inputLines: string[][]): string[][] {
  const outputLines: string[][] = [[...Header]];

  for (let i = 0; i < inputLines.length; i++) {
    const inputLine = inputLines[i];

    if (i === 0) {
      if (!isLineEqual(inputLine, expectedInputHeader)) {
        throw new Error("Unexpected header for dkb format");
      }
      continue;
    }

    const date = convertDkbDate(inputLine[0]);
    const reference = inputLine[5];
    const umsatztyp = inputLine[6];
    const amount = convertDkbAmount(inputLine[8]);
    const vendor = umsatztyp === "Ausgang" ? inputLine[4] : inputLine[3];

    outputLines.push([date, vendor, reference, amount]);
  }

  return sortByDate(outputLines);
}
