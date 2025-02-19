import { expect } from "@std/expect";
import { describe, it } from "@std/testing/bdd";
import { filterByDate } from "./filterByDate.ts";

describe("filterByDate", () => {
  it("Bounds are inclusive", () => {
    const data = [
      ["Date", "Name", "Amount"],
      ["1999-03-03", "Frederick", "17"],
      ["2000-07-18", "Howard", "19"],
      ["2004-02-28", "Reginald", "91"],
      ["2004-04-13", "Eduard", "95"],
      ["2006-05-15", "Arthur", "98"],
    ];

    const filteredData = filterByDate(data, "2000-07-18", "2004-04-13");

    const expectedData = [
      ["Date", "Name", "Amount"],
      ["2000-07-18", "Howard", "19"],
      ["2004-02-28", "Reginald", "91"],
      ["2004-04-13", "Eduard", "95"],
    ];

    expect(filteredData).toEqual(expectedData);
  });

  it("Bounds are optional", () => {
    const originalData = [
      ["Date", "Name", "Amount"],
      ["2000-07-18", "Gertrud", "19"],
      ["2004-02-28", "Hedwig", "91"],
      ["2004-04-13", "Brunhilde", "95"],
    ];

    const filteredData = filterByDate(originalData, "", "");

    expect(filteredData).toEqual(originalData);
  });

  it("Date validation", () => {
    const originalData = [
      ["Date", "Name", "Amount"],
      ["2000-07-18", "Geneviève", "102"],
      ["123-9-mauvais", "Colette", "98"],
      ["2004-04-13", "Éléonore", "97"],
    ];

    expect(() => filterByDate(originalData, "", "")).toThrow(
      "Invalid date format: 123-9-mauvais"
    );
  });
});
