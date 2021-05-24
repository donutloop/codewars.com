def circle_of_numbers(n, fst)
  sum = (n / 2) + fst
  if sum < n
    return sum
  end
  sum = (n / 2) - fst
  return sum.abs
end