export function LevDistance(string1, string2) {
  let wordOneLength = string1.length + 1;
  let wordTwoLength = string2.length + 1;
  let matrix = Array.from(Array(wordOneLength), () =>
    new Array(wordTwoLength).fill(0),
  );

  for (let i = 0; i < wordOneLength; i++) {
    matrix[i] = [i];
  }
  for (let j = 0; j < wordTwoLength; j++) {
    matrix[0][j] = j;
  }
  for (let column = 1; column < wordOneLength; column++) {
    for (let row = 1; row < wordTwoLength; row++) {
      if (string1[column - 1] === string2[row - 1]) {
        matrix[column][row] = matrix[column - 1][row - 1];
      } else {
        matrix[column][row] =
          1 +
          Math.min(
            matrix[column - 1][row - 1],
            matrix[column][row - 1],
            matrix[column - 1][row],
          );
      }
    }
  }
  return matrix[wordOneLength - 1][wordTwoLength - 1];
}
