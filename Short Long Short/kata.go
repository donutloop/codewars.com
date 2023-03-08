package kata

func Solution(a, b string) string {
  if len(a) > len(b) {
     return b + a + b
  }
  return a + b + a
}