package kata


func SequenceSum(start, end, step int) int {
  var sum int
  for num := start; num < end+1; num += step {
     sum += num;
  }
  return sum
}