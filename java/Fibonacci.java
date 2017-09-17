public class Fibonacci {
    public static void main(String[] args) {
        long result = fibonacci(10);

        System.out.println(result);
    }

    private static long fibonacci(long n) {
        if (n <= 1) {
            return n;
        }

        return fibonacci(n - 1) + fibonacci(n - 2);
    }
}