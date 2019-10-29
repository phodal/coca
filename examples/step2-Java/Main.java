import gateways.FakeRouter;
import domain.AggregateRootARepo;
import repositories.AggregateRootA;

public class Main {

    public static void main(String[] args) {
        System.out.println("main");
        Router router = new FakeRouter();
        AggregateRootARepo repo = new AggregateRootARepo();
        repo.save(new AggregateRootA(router));

    }
}
