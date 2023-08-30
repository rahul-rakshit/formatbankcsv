# Formatbankcsv

## About

This repo holds a cli that I wrote to unify the formats of different banks' csv downloads. It accepts the file path to a csv downloaded from either your DKB or n26 online banking and formats it to the following format:

| Date       | Vendor             | Reference          | Amount             |
| :--------- | :----------------- | :----------------- | :----------------- |
| yyyy-mm-dd | transaction vendor | "Verwendungszweck" | number, e.g. 10.00 |

The cli also accepts optional "from" and "to" date inputs to filter out transactions outside of this range.

## Installation

You may find pre-compiled binaries for your operating system and architecture under the github releases of this repo. Just unzip and add it to your PATH.
