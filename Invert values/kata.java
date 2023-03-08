public class Kata {
    public static int[] invert(int[] array) {
        for (var i = 0; i < array.length; i++) array[i] *= -1;
        return array;
    }
}