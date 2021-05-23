func findSum(_ n: Int) -> Int {
   var sum = 0
   for k in 1...n {
       if k % 3 == 0 || k % 5 == 0 {
           sum += k
       }
   }
   return sum
}