func descendingOrder(of number: Int) -> Int {
    var number = number
    var digits: Array<Int> = Array()
    while number > 0 {
        let k = number % 10
        number = number / 10
        digits.append(k)
    }
   digits.sort(by: >)

    var sum = 0
    var inc = 1
    var count = digits.count-1
    while count >= 0 {
        sum += digits[count] * inc
        inc = inc * 10
        count -= 1
    }

    return sum
}