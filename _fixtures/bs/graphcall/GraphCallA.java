
package graphcall;

public class GraphCallA {
    private GraphCallB graphCallB;
    private GraphCallC graphCallC;

    public void sayHi(){
        graphCallB.sayHi();
        graphCallC.sayHi();
    }
}
