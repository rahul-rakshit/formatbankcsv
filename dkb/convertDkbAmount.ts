export function convertDkbAmount(input: string): string {
  const cleanedInput = input.replace(/[^0-9,-]/g, "").replace(",", ".");
  const value = parseFloat(cleanedInput);

  if (isNaN(value)) {
    throw new Error(`Failed to parse amount ${input}`);
  }

  return value.toFixed(2);
}
