public class MyClass {
    public static int staticVariable;

    public int instanceVariable;

    public static void main(String[] args) {
        // ldc
        int x = 32768;

        // new
        MyClass myObject = new MyClass();

        // putstatic
        myObject.staticVariable = x;

        // getstatic
        x = myObject.staticVariable;

        // putfield
        myObject.instanceVariable = x;

        // getfield
        x = myObject.instanceVariable;

        Object yourObject = myObject;

        // instanceof
        if (yourObject instanceof MyClass) {
            // checkcast
            myObject = (MyClass)yourObject;

            System.out.println(myObject.instanceVariable);
        }
    }
}