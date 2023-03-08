import java.util.List;


public class MixedSum {

    /*
     * Assume input will be only of Integer o String type
     */
    public int sum(List<?> mixed) {
        var sum = 0;

        for (var i = 0; i < mixed.size(); i++) {
            var num = mixed.get(i);
            if( num instanceof String numStr) {
                var base = 1;
                for (var j = numStr.length()-1; j >= 0 ; j--) {
                    var digit = (int) numStr.charAt(j)-48;
                    sum += digit * base;
                    base *= 10;
                }
            } else {
                sum += (int) num;
            }
        }

        return sum;
    }
}