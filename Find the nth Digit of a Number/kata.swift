func findDigit(_ num:Int, _ nth: Int) -> Int {
  if nth  <= 0 {
      return -1
  }

  var num = num

  if num < 0 {
     num = num * -1
  }

  var digits: Array<Int> = Array()
  while num > 0 {
      let k = num % 10
      num = num / 10
      digits.append(k)
  }

  if digits.count < nth {
      return 0
  }

  return digits[nth-1]
}