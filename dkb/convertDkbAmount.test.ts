import { describe, it } from "@std/testing/bdd";
import { expect } from "@std/expect";
import { convertDkbAmount } from "./convertDkbAmount.ts";

describe("convertDkbAmount", () => {
  it("converts amount with currency symbol", () => {
    const input = "-55,94 €";
    const expectedAmount = "-55.94";
    expect(convertDkbAmount(input)).toBe(expectedAmount);
  });

  it("converts amount without currency symbol", () => {
    const input = "-55,94";
    const expectedAmount = "-55.94";
    expect(convertDkbAmount(input)).toBe(expectedAmount);
  });

  it("throws error for invalid input", () => {
    const input = "asdf";
    expect(() => convertDkbAmount(input)).toThrow("Failed to parse amount asdf");
  });
});
