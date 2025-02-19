import { Header } from "./constants.ts";

interface DateValuePair {
  date: string;
  value: string[];
}

export function sortByDate(input: string[][]): string[][] {
  const data: DateValuePair[] = input.map((item) => ({
    date: item[0],
    value: item.slice(1),
  }));

  data.sort((a, b) => {
    if (a.date === Header[0]) return -1;
    if (b.date === Header[0]) return 1;

    try {
      const date1 = new Date(a.date);
      const date2 = new Date(b.date);

      if (isNaN(date1.getTime()) || isNaN(date2.getTime())) {
        return 0;
      }

      return date1.getTime() - date2.getTime();
    } catch {
      return 0;
    }
  });

  return data.map((item) => [item.date, ...item.value]);
}
