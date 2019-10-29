package gateways;

import domain.Router;

public class FakeRouter extends Router {
    public int select(){
        System.out.println("routed \n");
        return 1;
	};
}