export function isLineEqual(line1: string[], line2: string[]): boolean {
  if (line1.length !== line2.length) {
    return false;
  }

  for (let i = 0; i < line1.length; i++) {
    if (line1[i] !== line2[i]) {
      return false;
    }
  }

  return true;
}
