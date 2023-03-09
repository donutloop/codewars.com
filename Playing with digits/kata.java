public class DigPow {

    public static long digPow(int n, int p) {

        var digitCount = (int) Math.log10(n);
        var oldNum = n;
        var sum = 0l;

        for (var i = 0; i <= digitCount; i++) {
            var digit = (int) (n / Math.pow(10, digitCount-i) % 10);
            sum += Math.pow(digit, p);
            p++;
        }


        if (sum%n == 0) {
            return sum/n;
        }


        return -1;
    }

}