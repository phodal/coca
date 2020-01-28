package tbs;

import java.util.regex.Pattern;

/**
 * Replaces critical characters in xml files
 */
public final class XmlSanitizer {

    private static final Pattern PURE_ASCII_STRING = Pattern.compile("^\\p{ASCII}*$"); // "[^\\p{ASCII}]+"

    /**
     * Checks if string contains &, <, >, ", ', non-ascii characters or anything other than A-Z, 0-9
     *
     * @param test String to test
     * @return true, if string only contains valid chars
     */
    public static boolean isValid(final String test) {
        // check we don't have xml chars in it
        boolean result = !test.contains("&") && !test.contains("<") && !test.contains(">") && !test.contains("\"") && !test.contains("'");
        // assure we only have ASCII chars
        result = result && PURE_ASCII_STRING.matcher(test).matches();
        // assure we really only A-Z and numbers in it
        result = result && (test.replaceAll("[^a-zA-Z0-9-+.!_\\s]", "").length() == test.length());
        return result;
    }

    private XmlSanitizer() {

    }
}
