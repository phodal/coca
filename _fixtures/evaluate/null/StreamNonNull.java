package nonnull;

import java.util.List;
import java.util.Objects;
import java.util.Optional;
import java.util.stream.Collectors;
import java.util.stream.Stream;

import org.apache.commons.lang3.StringUtils;
import javax.annotation.Nullable;

public class StreamNonNull {
    public static void main(String[] args) {
        Stream<String> language = Stream.of("java", "python", "node", null, "ruby", null, "php");
        List<String> result = language.filter(Objects::nonNull).collect(Collectors.toList());
        result.forEach(System.out::println);

        String a = testNull("3");
        a.charAt(0);
        String output = orElseNull();
        output.charAt(0);

        StringUtils.isNotEmpty(null);
    }

    public static String testNull(String input) {
        if (input.length() > 2) {
            return input;
        }

        return null;
    }

    @Nullable
    public static String orElseNull() {
        Name userName = new Name();
        userName.setName(null);

        return Optional.ofNullable(userName) // will be an empty optional if userName is null
                .map(Name::getName)   // will turn to empty optional if getName returns null
                .map("Name is: "::concat) // prepend "Name is: " (only when we have a name)
                .orElse(null); // get the result string or the alternative
    }

    private static class Name {
        private String name;

        public String getName() {
            return name;
        }

        public void setName(String name) {
            this.name = name;
        }
    }
}
