# Formatbankcsv

## About

This repo holds a cli that I wrote to unify the formats of different banks' csv downloads. It accepts the file path to a csv downloaded from either your DKB or n26 online banking and formats it to the following format:

| Date       | Vendor             | Reference          | Amount             |
| :--------- | :----------------- | :----------------- | :----------------- |
| yyyy-mm-dd | transaction vendor | "Verwendungszweck" | number, e.g. 10.00 |

The cli also accepts optional "from" and "to" date inputs to filter out transactions outside of this range.

## Installation

You may find pre-compiled binaries for your operating system and architecture under the github releases of this repo. Just unzip and add it to your PATH.

If you're on macOS, downloading it from your browser will give you the usual warning `"formatbankcsv" can't be opened because Apple cannot check it for malicious software`. To avoid this, you may just download it from the command-line, e.g. for arm64 on macOS: `curl -OL https://github.com/rahul-rakshit/formatbankcsv/releases/download/v1.0.0/darwin-arm64-v1.0.0.zip`.
