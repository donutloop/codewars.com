func inviteMoreWomen(_ arr: [Int]) -> Bool {
    var countMen = 0
    var countWomen = 0
    for num in arr {
       if num == 1 {
           countMen += 1
       } else {
           countWomen += 1
       }
    }

    if countMen > countWomen {
        return true
    }

    return false

}