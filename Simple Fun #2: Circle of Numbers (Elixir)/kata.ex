defmodule CircleNumbers do
  def circle_of_numbers(n, first_number) do
     half = (n / 2)
     sumPlus = half + first_number
     sumMinus = half - first_number
     cond do
      sumPlus < n -> sumPlus
      abs(sumMinus) < n -> abs(sumMinus)
     end
  end
end