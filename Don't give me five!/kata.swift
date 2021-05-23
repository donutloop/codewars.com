func dontGiveMeFive(_ start: Int, _ end: Int) -> Int {
  print(start, end)

  var count = 0
  for num in start...end {

      if hasDigit(num) {
          continue
      }
      count += 1
  }
  return count
}

func hasDigit(_ num : Int) -> Bool {
   var num = num

   if num < 0 {
      num = num * -1
   }

   while num != 0 {
       let k = num % 10
       num = num / 10
       if k == 5 {
           return true
       }
   }
   return false
}