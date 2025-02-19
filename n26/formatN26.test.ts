import { describe, it } from "@std/testing/bdd";
import { expect } from "@std/expect";
import { readCsv } from "../csv/csv.ts";
import { formatN26 } from "./formatN26.ts";

describe("formatN26", () => {
  it("should format n26 csv data correctly", async () => {
    const sampleN26Data = await readCsv("./n26/fixtures/sampleN26.csv", ",", 0);
    const formatted = formatN26(sampleN26Data);
    const expectedResultCsv = await readCsv("./n26/fixtures/expectedResult.csv", ",", 0);
    expect(formatted).toEqual(expectedResultCsv);
  });

  it("should throw error for bad header", async () => {
    const badHeaderData = await readCsv("./n26/fixtures/badHeaderN26.csv", ",", 0);
    expect(() => formatN26(badHeaderData)).toThrow("Unexpected header for n26 format");
  });
});
