export function convertDkbDate(input: string): string {
  const [day, month, year] = input.split(".");
  
  if (!day || !month || !year || 
      isNaN(parseInt(day)) || 
      isNaN(parseInt(month)) || 
      isNaN(parseInt(year))) {
    throw new Error(`Invalid date format: ${input}`);
  }

  const fullYear = parseInt(year) + 2000;
  
  const paddedDay = day.padStart(2, "0");
  const paddedMonth = month.padStart(2, "0");
  
  return `${fullYear}-${paddedMonth}-${paddedDay}`;
}
