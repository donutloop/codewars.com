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


        if (sum == oldNum) {
            return 1;
        } else if (sum < oldNum) {
            return -1;
        }

        var digitCountSum = Math.log10(sum);
        var digitCountDiff = (int) digitCountSum - digitCount;
        var upper = Math.pow(10, digitCountDiff);
        var upperNum = oldNum * upper;

        var upK = (double) ((sum / upperNum) * (double) 100);
        if (upK * oldNum == sum) {
            return (long) upK;
        }

        var downK = upK;
        do {
            upK = upK * 10;
            downK = downK / 10;

            if (Math.ceil(upK) * oldNum == sum) {
                return (long) Math.ceil(upK);
            }
            if (Math.ceil(downK) * oldNum == sum) {
                return (long) Math.ceil(downK);
            }

        } while(upK * oldNum <= sum || downK * oldNum > sum);


        return -1;
    }

}