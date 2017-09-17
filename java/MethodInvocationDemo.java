public class MethodInvocationDemo implements Runnable {
    public static void main(String[] args) {
        new MethodInvocationDemo().test();
    }

    public void test() {
        // invokestatic
        MethodInvocationDemo.staticMethod();

        // invokespecial
        MethodInvocationDemo methodInvocationDemo = new MethodInvocationDemo();

        // invokespecial
        methodInvocationDemo.instanceMethod();
        
        // invokespecial
        super.equals(null);

        // invokevirtual
        this.run();

        // invokeinterface
        ((Runnable)methodInvocationDemo).run();
    }

    public static void staticMethod() {
    }

    private static void instanceMethod() {
    }

    @Override
    public void run() {
    }
}