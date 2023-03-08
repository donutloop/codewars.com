package kata

import "strings"
import "math"


func Derive(coefficient, exponent int) string {
   var sb strings.Builder
   sum := coefficient * exponent
   digitCount := int(math.Log10(float64(sum)))

   for i := float64(digitCount); i >= 0; i-- {
      digit := sum / int(math.Pow(10, i)) % 10;
      sb.WriteByte(byte(digit+48))
   }

   sb.WriteString("x^")
   exponent = exponent-1
   digitCount = int(math.Log10(float64(exponent)))
    for i := float64(digitCount); i >= 0; i-- {
      digit := exponent / int(math.Pow(10, i)) % 10;
      sb.WriteByte(byte(digit+48))
   }


   return sb.String()
}