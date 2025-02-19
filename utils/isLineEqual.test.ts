import { expect } from "@std/expect";
import { describe, it } from "@std/testing/bdd";
import { isLineEqual } from "./isLineEqual.ts";

describe("isLineEqual", () => {
  it("Equal case", () => {
    const line = ["1", "2", "3"];
    expect(isLineEqual(line, line)).toBe(true);
  });

  it("Length mismatch", () => {
    const line1 = ["1", "2", "3"];
    const line2 = ["1", "2", "3", "4"];
    expect(isLineEqual(line1, line2)).toBe(false);
  });

  it("Unequal case", () => {
    const line1 = ["1", "3", "2"];
    const line2 = ["1", "2", "3"];
    expect(isLineEqual(line1, line2)).toBe(false);
  });
});
