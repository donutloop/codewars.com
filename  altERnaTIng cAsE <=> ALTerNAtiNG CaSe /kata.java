public class StringUtils {

    public static String toAlternativeString(String string) {
        var sb = new StringBuilder(string);
        for (var i = 0; i < sb.length(); i++) {
            if (sb.charAt(i) == ' ') {
                continue;
            }

            if (Character.isLowerCase(sb.charAt(i))) {
                sb.setCharAt(i, Character.toUpperCase(sb.charAt(i)));
            } else {
                sb.setCharAt(i, Character.toLowerCase(sb.charAt(i)));
            }

        }
        return sb.toString();
    }
}