package repositories;

import domain.*;


public class AggregateRootARepo extends Repository {
    private AggregateRootA[] arList;
    public void save(AggregateRootA a){
        System.out.println("saved\n");
	};
}