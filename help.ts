export const helpMessage = `
Usage: formatbankcsv -F <format> -i <input> -o <output> [-f <from-date>] [-t <to-date>]

Required:
  -i, --input   Path to the input file
  -o, --output  Path to the output file
  -F, --format  Format ("dkb" or "n26")

Optional:
  -h, --help    Show this help message
  -f, --from    Start date in yyyy-mm-dd format, inclusive
  -t, --to      End date in yyyy-mm-dd format, inclusive

Example:
  formatbankcsv main.ts -F dkb -i ./dkb.csv -o output.csv -f 2023-04-01 -t 2023-06-30
`;
