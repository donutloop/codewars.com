function solve(a){

  let eventCount = 0
  let oddCount = 0
  for (x of a) {
      if (isNaN(x) == false) {
          if (x % 2 == 0) {
            eventCount++
          } else {
            oddCount++
          }
      }
  }

  return eventCount - oddCount;
};