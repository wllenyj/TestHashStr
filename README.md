# TestHashStr
Test your id field use which HashStr function is most appropriate

## Useage
Edit your own keys.txt file, the content for you to test the id field. Such as:
9787101003048
3529385239842
2398523985293
3469431928234

``` shell
$ sh run.sh 
=== RUN   TestBKDRHash1_16
--- PASS: TestBKDRHash1_16 (0.00s)
        hash_test.go:57: [1678 1659 1627 1711 1686 1690 1658 1930 1690 1712 1907 2024 1697 1982 1785 1739]
=== RUN   TestBKDRHash1_32
--- PASS: TestBKDRHash1_32 (0.00s)
        hash_test.go:57: [863 849 830 837 865 854 842 859 904 792 882 917 841 869 958 864 815 810 797 874 821 836 816 1071 786 920 1025 1107 856 1113 827 875]
=== RUN   TestBKDRHash2_16
--- PASS: TestBKDRHash2_16 (0.00s)
        hash_test.go:57: [1638 1674 1657 1987 1722 1815 1692 1752 1710 1710 1722 1970 1681 1844 1906 1695]
=== RUN   TestBKDRHash2_32
--- PASS: TestBKDRHash2_32 (0.00s)
        hash_test.go:57: [809 857 857 852 854 824 856 882 795 857 813 891 812 814 1065 827 829 817 800 1135 868 991 836 870 915 853 909 1079 869 1030 841 868]
=== RUN   TestBKDRHash3_16
--- PASS: TestBKDRHash3_16 (0.00s)
        hash_test.go:57: [1638 1674 1657 1987 1722 1815 1692 1752 1710 1710 1722 1970 1681 1844 1906 1695]
=== RUN   TestBKDRHash3_32
--- PASS: TestBKDRHash3_32 (0.00s)
        hash_test.go:57: [807 850 823 1072 846 924 854 889 814 895 909 915 899 818 850 852 831 824 834 915 876 891 838 863 896 815 813 1055 782 1026 1056 843]
=== RUN   TestSDBMHash_16
--- PASS: TestSDBMHash_16 (0.00s)
        hash_test.go:57: [1621 1715 1848 1673 1701 1689 1766 1764 1747 1656 1686 2062 1682 1983 1677 1905]
=== RUN   TestSDBMHash_32
--- PASS: TestSDBMHash_32 (0.00s)
        hash_test.go:57: [545 528 665 400 360 365 462 401 482 467 573 671 701 851 916 1186 1076 1187 1183 1273 1341 1324 1304 1363 1265 1189 1113 1391 981 1132 761 719]
=== RUN   TestAPHash_16
--- PASS: TestAPHash_16 (0.01s)
        hash_test.go:57: [1854 2079 1740 1682 1837 1649 1698 1661 1686 1708 1639 1639 1773 1780 1801 1949]
=== RUN   TestAPHash_32
--- PASS: TestAPHash_32 (0.01s)
        hash_test.go:57: [834 897 928 845 1049 843 823 823 845 852 848 845 844 909 960 817 1020 1182 812 837 788 806 875 838 841 856 791 794 929 871 841 1132]
=== RUN   TestMetroHash_16
--- PASS: TestMetroHash_16 (0.00s)
        hash_test.go:57: [1901 1714 1670 1664 1974 1733 1716 1970 1685 1709 1706 1693 1829 1716 1762 1733]
=== RUN   TestMetroHash_32
--- PASS: TestMetroHash_32 (0.00s)
        hash_test.go:57: [1066 885 829 849 1124 841 862 1125 879 821 863 841 792 863 855 837 835 829 841 815 850 892 854 845 806 888 843 852 1037 853 907 896]
=== RUN   TestJavaHash_16
--- PASS: TestJavaHash_16 (0.01s)
        hash_test.go:57: [1621 1715 1848 1673 1701 1689 1766 1764 1747 1656 1686 2062 1682 1983 1677 1905]
=== RUN   TestJavaHash_32
--- PASS: TestJavaHash_32 (0.01s)
        hash_test.go:57: [545 528 665 400 360 365 462 401 482 467 573 671 701 851 916 1186 1076 1187 1183 1273 1341 1324 1304 1363 1265 1189 1113 1391 981 1132 761 719]
=== RUN   TestDJBHash_16
--- PASS: TestDJBHash_16 (0.01s)
        hash_test.go:57: [2033 1670 1897 1718 1850 1841 1712 1683 1710 1637 1688 1666 1686 1872 1871 1641]
=== RUN   TestDJBHash_32
--- PASS: TestDJBHash_32 (0.01s)
        hash_test.go:57: [1639 1326 1599 1395 1421 1377 1297 1207 1165 1020 977 865 760 610 737 478 394 344 298 323 429 464 415 476 545 617 711 801 926 1262 1134 1163]
=== RUN   TestSuperFastHash_16
--- PASS: TestSuperFastHash_16 (0.01s)
        hash_test.go:57: [1637 1692 1664 1664 1911 1652 1697 1980 1734 1972 1776 1692 1753 1678 1906 1767]
=== RUN   TestSuperFastHash_32
--- PASS: TestSuperFastHash_32 (0.01s)
        hash_test.go:57: [794 800 844 871 1063 796 860 1155 891 844 919 848 806 835 1086 899 843 892 820 793 848 856 837 825 843 1128 857 844 947 843 820 868]
PASS
ok      _/.../hash/hashstr    0.121s
goos: linux
goarch: amd64
BenchmarkBKDRHash1-4                 500           3207708 ns/op
BenchmarkBKDRHash2-4                 500           3239999 ns/op
BenchmarkBKDRHash3-4                 500           3283682 ns/op
BenchmarkSDBMHash-4                  300           4476090 ns/op
BenchmarkAPHash-4                    200           6059254 ns/op
BenchmarkMetroHash-4                2000           1145394 ns/op
BenchmarkJavaHash-4                  200           8656975 ns/op
BenchmarkDJBHash-4                   200           8649133 ns/op
BenchmarkSuperFastHash-4             200           9277220 ns/op
PASS
ok      _/.../hash/hashstr    20.046s
```
### Dep
github.com/dgryski/dgohash
github.com/dgryski/go-metro
