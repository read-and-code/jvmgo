public class MethodInvocation implements Runnable {
    public static void main(String[] args) {
        new MethodInvocation().test();
    }

    public void test() {
        // invokestatic
        MethodInvocation.staticMethod();

        // invokespecial
        MethodInvocation methodInvocation = new MethodInvocation();

        // invokespecial
        methodInvocation.instanceMethod();
        
        // invokespecial
        super.equals(null);

        // invokevirtual
        this.run();

        // invokeinterface
        ((Runnable)methodInvocation).run();
    }

    public static void staticMethod() {
    }

    private static void instanceMethod() {
    }

    @Override
    public void run() {
    }
}