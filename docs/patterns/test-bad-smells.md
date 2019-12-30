# Test Data

Source: 

| System | Description | Classes | Methods | KLOCs | Test Methods |
|--------|-------------|---------|---------|-------|-----------------|
Apache Ant 1.8.3     | Java library and command-line tool to build systems | 813 | 8,540 | 204 | 3,097
Apache Cassandra 1.1 | Scalable DB Management System | 586 | 5,730 | 111 | 586
Apache Derby 10.9    | Relational DB Management System | 1,929 | 28,119 | 734 | 426
Apache Hive 0.9 Data | Warehouse Software Facilities Provider | 1,115 | 9,572 | 204 | 58
Apache Ivy 2.1.0x    | Flexible Dependency Manager | 349 | 3,775 | 58 | 793
Apache Hbase 0.94    | Distributed DB System | 699 | 8,148 | 271  | 604
Apache Karaf 2.3     | Standalone Software Container | 470 | 2,678 | 56 | 199
Apache Lucene 3.6    | Search Engine | 2,246 | 17,021 | 466 | 3,895
Apache Nutch 1.4     | Web-search Software built on Lucene | 259 | 1,937 | 51 | 389
Apache Pig 0.8       | Large Dataset Query Maker | 922 | 7,619 | 184 | 449
Apache Qpid 0.18     | AMQP-based Messaging Tool | 922 | 9,777 | 193 | 786
Apache Struts 3.0    | MVC Framework | 1,002 | 7,506 | 152 | 1,751
Apache Wicket 1.4.20 | Java Serverside Web Framework | 825 | 6,900 | 179 | 1,553
Elastic Search 0.19  | RESTful Search Engine | 2,265 | 17,095 | 316 | 397
Hibernate 4          | Java Persistence Manager | 154 | 2,387 | 47 | 132
JHotDraw 7.6         | Java GUI Framework for Technical Graphics | 679 | 6,687 | 135 | 516
JFreeChart 1.0.14    | Java Chart Library | 775 | 8,746 | 231 | 3,842
HSQLDB 2.2.8         |  HyperSQL Database Engine | 444 | 8,808 | 260 | 59
Overall              | - | 16,454 | 161,045 |3,852 | 19,532

Issues

| Test Smell       | Rule                |
|------------------|---------------------|
Resource Optimism  | JUnit methods using an external resource without checking its status.
Indirect Testing   | JUnit methods invoking, besides methods of the corresponding production class, methods of other classes in the production code.
Test Run War       | JUnit methods that allocate resources that are also used by other test methods (e.g., temporary files)

Flaky Tests

| Category | Description | 
|----------|------------|
Async Wait      | A test method making an asynchronous call and that does not wait for the result of the call.
Concurrency     | Different threads interact in a non-desirable manner.
Test Order Dependency | The test outcome depends on the order of execution of the tests.
Resource Leak   | The test method does not properly acquire or release one or more of its resources.
Network Test    | execution depends on the network performance.
Time            | The test method relies on the system time.
IO              | The test method does not properly manage external resources.
Randomness      | The test method uses random number.
Floating Point Operation | The test method performs floating-point operations.
Unordered Collections | Test outcome depends on the order of collections.

