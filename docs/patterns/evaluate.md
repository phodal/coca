url: https://github.com/macrozheng/mall

```
+--------------------------------+-------+-----------------------+-------+-----------+
|              TYPE              | COUNT |         LEVEL         | TOTAL |   RATE    |
+--------------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null         |    22 | Method                | 12682 | 0.17%     |
| Utils                          |     2 | Class                 |   458 | 0.44%     |
| Static Method                  |     7 | Method                | 12682 | 0.02%     |
| Average Method Num.            | 12682 | Method/Class          |   458 | 27.689956 |
| Method Num. Std Dev / 标准差   | 12682 | Class                 | -     | 58.993564 |
| Average Method Length          | 45091 | Without Getter/Setter | 10205 |  4.418520 |
| Method Length Std Dev / 标准差 | 12682 | Method                | -     |  2.904016 |
+--------------------------------+-------+-----------------------+-------+-----------+
```

// issues: big data class / god service 

url: https://github.com/shuzheng/zheng

```
+--------------------------------+-------+-----------------------+-------+-----------+
|              TYPE              | COUNT |         LEVEL         | TOTAL |   RATE    |
+--------------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null         |     0 | Method                |  5256 | 0.00%     |
| Utils                          |    18 | Class                 |   366 | 4.92%     |
| Static Method                  |     0 | Method                |  5256 | 0.34%     |
| Average Method Num.            |  5256 | Method/Class          |   366 | 14.360656 |
| Method Num. Std Devv / 标准差   |  5256 | Class                 | -     | 32.965336 |
| Average Method Length          | 19644 | Without Getter/Setter |  4328 |  4.538817 |
| Method Length Std Devv / 标准差 |  5256 | Method                | -     |  4.467935 |
+--------------------------------+-------+-----------------------+-------+-----------+
```

// data class

big data

```
+------------------------+-------+-----------------------+-------+-----------+
|          TYPE          | COUNT |         LEVEL         | TOTAL |   RATE    |
+------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null |   128 | Method                |  3041 | 4.21%     |
| Utils                  |    18 | Class                 |   496 | 7.06%     |
| Static Method          |   400 | Method                |  3041 | 1.15%     |
| Average Method Num     |  3041 | Method/Class          |   496 | 6.13      |
| Method Num. Std Dev     |  3041 | Class                 | -     | 11.03 |
| Average Method Length  | 17730 | Without Getter/Setter |  1551 | 11.43     |
| Method Length Std Dev   |  3041 | Method                | -     |  14.22 |
+------------------------+-------+-----------------------+-------+-----------+
```

// null issues

Common Project

```
+------------------------+-------+-----------------------+-------+-----------+
|          TYPE          | COUNT |         LEVEL         | TOTAL |   RATE    |
+------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null |   234 | Method                | 16642 | 1.41%     |
| Utils                  |    26 | Class                 |  1007 | 7.06%     |
| Static Method          |  2062 | Method                | 16642 | 0.16%     |
| Average Method Num     | 16642 | Method/Class          |  1007 | 16.52     |
| Method Num. Std Dev     |  16642 | Class                 | -     | 29.86 |
| Average Method Length  | 69012 | Without Getter/Setter |  6020 | 11.46     |
| Method Length Std Dev   |  16642 | Method                | -     |  19.30 |
+------------------------+-------+-----------------------+-------+-----------+
```

// static projects

Algo Project

```
+------------------------+-------+-----------------------+-------+-----------+
|          TYPE          | COUNT |         LEVEL         | TOTAL |   RATE    |
+------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null |   101 | Method                |  4914 | 2.06%     |
| Utils                  |    43 | Class                 |   926 | 4.64%     |
| Static Method          |  542  | Method                |  4914 | 0.88%     |
| Average Method Num     | 4914  | Method/Class          |   927 | 5.30      |
| Method Num. Std Dev     |  4914 | Class                 | -     | 9.24 |
| Average Method Length  | 51056 | Without Getter/Setter |  2603 | 19.61     |
| Method Length Std Dev   |  4914 | Method                | -     |  24.39 |
+------------------------+-------+-----------------------+-------+-----------+
```

// longest method

CMS

url: https://github.com/sanluan/PublicCMS

```
+------------------------+-------+-----------------------+-------+-----------+
|          TYPE          | COUNT |         LEVEL         | TOTAL |   RATE    |
+------------------------+-------+-----------------------+-------+-----------+
| Nullable / Return Null |   131 | Method                |  2707 | 4.84%     |
| Utils                  |    21 | Class                 |   483 | 4.35%     |
| Static Method          |   246 | Method                |  2707 | 0.78%     |
| Average Method Num     |  2707 | Method/Class          |   483 |  5.604555 |
| Average Method Length  | 13622 | Without Getter/Setter |  1350 | 10.090370 |
+------------------------+-------+-----------------------+-------+-----------+
```

// to many null

ERP

url: https://github.com/ilscipio/scipio-erp

```
+------------------------+--------+-----------------------+-------+-----------+
|          TYPE          | COUNT  |         LEVEL         | TOTAL |   RATE    |
+------------------------+--------+-----------------------+-------+-----------+
| Nullable / Return Null |   1660 | Method                | 23461 | 7.08%     |
| Utils                  |     86 | Class                 |  2288 | 3.76%     |
| Static Method          |   6469 | Method                | 23461 | 0.37%     |
| Average Method Num     |  23461 | Method/Class          |  2288 | 10.253934 |
| Average Method Length  | 228550 | Without Getter/Setter | 14080 | 16.232244 |
+------------------------+--------+-----------------------+-------+-----------+
```

// longest method
// to many null (bugs)