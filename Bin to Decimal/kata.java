public class Converter{
    public static int binToDecimal(String inp){
        var j = 0;
        var sum = 0;
        for (var i = inp.length()-1; i >= 0; i--) {
            var digit = (int) inp.codePointAt(i)-48;
            sum += digit * Math.pow(2, j);
            j++;
        }
        return sum;
    }
}