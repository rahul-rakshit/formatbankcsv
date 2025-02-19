import { describe, it } from "@std/testing/bdd";
import { expect } from "@std/expect";
import { convertDkbDate } from "./convertDkbDate.ts";

describe("convertDkbDate", () => {
  it("converts date from dd.mm.yy to yyyy-mm-dd", () => {
    const input = "02.01.23";
    const expectedDate = "2023-01-02";
    expect(convertDkbDate(input)).toEqual(expectedDate);
  });

  it("throws error for invalid date format", () => {
    const input = "some.invalid.date";
    expect(() => convertDkbDate(input)).toThrow();
  });
});
