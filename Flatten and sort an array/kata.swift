func flattenAndSort<T: Comparable>(_ arr: [[T]]) -> [T] {

  var flatten: Array<T> = Array()
  for nestedArr in arr {
      for element in nestedArr {
          flatten.append(element)
      }
  }

  flatten.sort()

  return flatten
}