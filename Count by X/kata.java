public class Kata{
    public static int[] countBy(int x, int n){
        var nums = new int[n];
        var index = 0;
        var limit = n*x+1;
        for (var num = x; num < limit; num=num+x) {
            nums[index] = num;
            index++;
        }
        return nums;
    }
}