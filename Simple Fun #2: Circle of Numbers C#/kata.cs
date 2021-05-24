namespace myjinxin
{
    using System;

    public class Kata
    {
        public int CircleOfNumbers(int n, int FirstNumber){
          int sum = 0;
          sum = (n / 2) + FirstNumber;
          if (sum < n) {
             return sum;
          }

          sum = (n / 2) - FirstNumber;
          return Math.Abs(sum);
        }
    }
}