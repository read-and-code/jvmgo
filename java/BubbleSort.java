public class BubbleSort {
    public static void main(String[] args) {
        int[] array = {
            22, 84, 77, 11, 95,  9, 78, 56, 
            36, 97, 65, 36, 10, 24 ,92, 48
        };

        bubbleSort(array);
        printarrayay(array);
    }
    
    private static void bubbleSort(int[] array) {
        int j = 0;
        boolean swapped = true;

        while (swapped) {
            j++;
            swapped = false;

            for (int i = 0; i < array.length - j; i++) {
                if (array[i] > array[i + 1]) {
                    swap(array, i, i + 1);

                    swapped = true;
                }
            }
        }
    }

    private static void swap(int[] array, int i, int j) {
        int temp = array[i];
        array[i] = array[j];
        array[j] = temp;
    }
    
    private static void printarrayay(int[] array) {
        for (int i : array) {
            System.out.println(i);
        }
    }
}