import { expect } from "@std/expect";
import { describe, it, afterEach } from "@std/testing/bdd";
import { readCsv, writeCsv } from "./csv.ts";

async function createTestDir(dirPath: string) {
  try {
    await Deno.stat(dirPath);
  } catch (error) {
    if (error instanceof Deno.errors.NotFound) {
      await Deno.mkdir(dirPath);
    }
  }
}

describe("csv", () => {
  afterEach(async () => {
    try {
      await Deno.remove(".test_temp", { recursive: true });
    } catch (error) {
      if (!(error instanceof Deno.errors.NotFound)) {
        throw error;
      }
    }
  });

  it("can read and write CSV with custom delimiter", async () => {
    const data = [
      ["1", "2", "3"],
      ["a", "b", "c"],
    ];
    const filePath = ".test_temp/test.csv";
    await createTestDir(".test_temp");

    await writeCsv(data, filePath, "^");
    const fileContents = await readCsv(filePath, "^", 0);
    
    expect(fileContents).toEqual(data);
  });

  it("can skip lines when reading CSV", async () => {
    const fileLines = [
      "Line 1",
      "",
      "Line 3",
      "Some header",
      "1;2;3",
      "a;b;c",
    ];
    const filePath = ".test_temp/complex.csv";
    await createTestDir(".test_temp");

    await Deno.writeTextFile(filePath, fileLines.join("\n") + "\n");

    const linesToSkip = 4;
    const parsedData = await readCsv(filePath, ";", linesToSkip);

    const expectedData = [
      ["1", "2", "3"],
      ["a", "b", "c"],
    ];
    expect(parsedData).toEqual(expectedData);
  });

  it("adds empty strings for missing values", async () => {
    const fileLines = [
      "first*second*third*fourth",
      "1*2*3",
      "a*b*c*d"
    ];
    const filePath = ".test_temp/missingValues.csv";
    await createTestDir(".test_temp");
    await Deno.writeTextFile(filePath, fileLines.join("\n") + "\n");

    const parsedData = await readCsv(filePath, "*", 0);

    const expectedData = [
      ["first", "second", "third", "fourth"],
      ["1", "2", "3", ""], // adds an empty string here
      ["a", "b", "c", "d"]
    ];
    expect(parsedData).toEqual(expectedData);
  });
});
