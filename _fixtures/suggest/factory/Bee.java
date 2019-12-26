package factory;

class Insect {
    public Insect(int size) {
        this(size, 123);
        System.out.println("Constructor: Insect size");
    }

    public Insect(int size, int height) {
        System.out.println("Constructor: Insect size, height");
    }

    public Insect(int size, int height, String color) {
        System.out.println("Constructor: Insect size, height");
    }
}

public class Bee extends Insect {
    public Bee(int size) {
        super(size);
        System.out.println("Constructor: Bee size");
    }

    public Bee(int size, int height) {
        super(size, height);
        System.out.println("Constructor: Bee size, height");
    }

    public Bee(int size, int height, String color) {
        super(size, height);
        System.out.println("Constructor: Bee size, height, color");
    }

    public Bee(int size, int height, String color, int x, int y, int z) {
        super(size, height);
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
        System.out.println("Constructor: Bee size, height, color");
    }
}