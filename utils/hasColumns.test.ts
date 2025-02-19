import { expect } from "@std/expect";
import { describe, it } from "@std/testing/bdd";
import { hasColumns } from "./hasColumns.ts";

describe("hasColumns", () => {
  it("positive case", () => {
    const headerColumn = ["Item", "Description", "Amount"];
    const requiredColumns = ["Amount", "Item"];
    expect(hasColumns(headerColumn, requiredColumns)).toBe(true);
  });
});

Deno.test("negative case", () => {
  const headerColumn = ["Item", "Description", "Amount"];
  const requiredColumns = ["Supermarket", "Item"];
  expect(hasColumns(headerColumn, requiredColumns)).toBe(false);
});
