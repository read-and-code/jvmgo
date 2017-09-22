public class ParseIntTest {
    public static void main(String[] args) {
        try {
            int x = Integer.parseInt("a");

            System.out.println(x);
        } catch (NumberFormatException e) {
            System.out.println(e.getMessage());
        }
    }
}