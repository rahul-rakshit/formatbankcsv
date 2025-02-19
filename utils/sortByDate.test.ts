import { expect } from "@std/expect";
import { describe, it } from "@std/testing/bdd";
import { sortByDate } from "./sortByDate.ts";

describe("sortByDate", () => {
  it("sorts dates in ascending order", () => {
    const input = [
      ["Date", "Letter", "Number"],
      ["1983-01-07", "a", "1"],
      ["2007-12-19", "b", "2"],
      ["1998-08-27", "c", "3"],
    ];

    const sorted = sortByDate(input);
    const expected = [
      ["Date", "Letter", "Number"],
      ["1983-01-07", "a", "1"],
      ["1998-08-27", "c", "3"],
      ["2007-12-19", "b", "2"],
    ];

    expect(sorted).toEqual(expected);
  });

  it("sorts header to top", () => {
    const input = [
      ["1983-01-07", "a", "1"],
      ["2007-12-19", "b", "2"],
      ["1998-08-27", "c", "3"],
      ["Date", "Letter", "Number"],
    ];

    const sorted = sortByDate(input);
    const expected = [
      ["Date", "Letter", "Number"],
      ["1983-01-07", "a", "1"],
      ["1998-08-27", "c", "3"],
      ["2007-12-19", "b", "2"],
    ];

    expect(sorted).toEqual(expected);
  });
});
