package com.phodal.coca.abug;

public class ClassOne extends HasStatic {
    static final int x = 3;
    static int y = 4;
    public static void pr() {

    }
}

public class ClassTwo {
  final int z=10;

  public static void main(String[] args) {
    Outer outer = new Outer();
    System.out.println(outer.new Inner().y);
  }
}