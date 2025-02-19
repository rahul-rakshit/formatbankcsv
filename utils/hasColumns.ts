export function hasColumns(headerColumn: string[], requiredColumns: string[]): boolean {
  return requiredColumns.every((required) => headerColumn.includes(required));
}
