package builder;

public class BeeBuilder extends Insect {
    public BeeBuilder(int size) {
        super(size);
        System.out.println("Constructor: BeeBuilder size");
    }

    public BeeBuilder(int size, int height) {
        super(size, height);
        System.out.println("Constructor: BeeBuilder size, height");
    }

    public BeeBuilder(int size, int height, String color) {
        super(size, height);
        System.out.println("Constructor: BeeBuilder size, height, color");
    }

    public BeeBuilder(int size, int height, String color, int x, int y, int z) {
        super(size, height);
        System.out.println("Constructor: BeeBuilder size, height, color");
    }
}