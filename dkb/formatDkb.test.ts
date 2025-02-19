import { describe, it } from "@std/testing/bdd";
import { expect } from "@std/expect";
import { readCsv } from "../csv/csv.ts";
import { formatDkb } from "./formatDkb.ts";

describe("formatDkb", () => {
  it("formats dkb csv data correctly", async () => {
    const sampleDkbData = await readCsv("./dkb/fixtures/sampleDkb.csv", ";", 4);
    const formatted = formatDkb(sampleDkbData);
    const expectedResultCsv = await readCsv("./dkb/fixtures/expectedResult.csv", ",", 0);
    expect(formatted).toEqual(expectedResultCsv);
  });

  it("throws error for bad header", async () => {
    const badHeaderData = await readCsv("./dkb/fixtures/badHeaderDkb.csv", ";", 4);
    expect(() => formatDkb(badHeaderData)).toThrow("Unexpected header for dkb format");
  });

  it("throws error for old format", async () => {
    const oldFormatData = await readCsv("./dkb/fixtures/oldDkbFormat.csv", ";", 4);
    expect(() => formatDkb(oldFormatData)).toThrow("Unexpected header for dkb format");
  });
});
