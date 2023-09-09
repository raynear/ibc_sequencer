const crypto = require("crypto");

// This is the seed which we will give to the sequencer
const seed = "Hello World";
const timeInSeconds = 20;

// Create time-lock puzzle with the seed and desired time. This function will hash the seed until the time is up. The final hash will be our encryption-decryption key.
function createPuzzle(seed, timeInSeconds) {
  let key = seed;
  const startTime = Date.now();
  const endTime = startTime + timeInSeconds * 1000; // Convert seconds to milliseconds
  let iters = 0;
  while (Date.now() < endTime) {
    // Repetitively hash and count iterations
    key = crypto.createHash("sha256").update(key).digest("hex");
    iters++;
  }
  return { key, iters };
}

function solvePuzzle(seed, iters) {
  let key = seed;
  for (let i = 0; i < iters; i++) {
    key = crypto.createHash("sha256").update(key).digest("hex");
  }
  return key;
}

const { key, iters } = createPuzzle(seed, timeInSeconds);
console.log(key, iters);

// Test

// console.log(key === solvePuzzle(seed, iters));
