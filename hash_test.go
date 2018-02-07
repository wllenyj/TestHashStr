package hashstr

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"testing"
)

var keysFile = flag.String("keys", "", "load keys datafile")

var gkeys []string

func loadKeys() []string {

	if *keysFile != "" {
		return loadBigKeys(*keysFile)
	}

	return []string{
		"foo",
		"bar",
		"baz",
		"qux",
		"zot",
		"frob",
		"zork",
		"zeek",
	}
}

func loadBigKeys(filename string) []string {
	f, err := os.Open(filename)
	if err != nil {
		//tb.Fatalf("unable to keys file: %v", err)
		panic(fmt.Sprintf("unable to keys file: %v", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var k []string
	for scanner.Scan() {
		k = append(k, scanner.Text())
	}

	return k
}

type HashFunc func(string) uint

func testHash(t *testing.T, keys []string, bucket uint, hf HashFunc) {
	count := make([]int, bucket)
	for _, k := range keys {
		count[hf(k)%bucket]++
	}
	t.Logf("%+v\n", count)
}

func TestBKDRHash1_16(t *testing.T) {
	testHash(t, gkeys, 16, Hash1)
}
func TestBKDRHash1_32(t *testing.T) {
	testHash(t, gkeys, 32, Hash1)
}
func TestBKDRHash2_16(t *testing.T) {
	testHash(t, gkeys, 16, Hash2)
}
func TestBKDRHash2_32(t *testing.T) {
	testHash(t, gkeys, 32, Hash2)
}
func TestBKDRHash3_16(t *testing.T) {
	testHash(t, gkeys, 16, Hash3)
}
func TestBKDRHash3_32(t *testing.T) {
	testHash(t, gkeys, 32, Hash3)
}
func TestSDBMHash_16(t *testing.T) {
	testHash(t, gkeys, 16, SDBMHash)
}
func TestSDBMHash_32(t *testing.T) {
	testHash(t, gkeys, 32, SDBMHash)
}
func TestAPHash_16(t *testing.T) {
	testHash(t, gkeys, 16, APHash)
}
func TestAPHash_32(t *testing.T) {
	testHash(t, gkeys, 32, APHash)
}
func TestMetroHash_16(t *testing.T) {
	testHash(t, gkeys, 16, MetroHash)
}
func TestMetroHash_32(t *testing.T) {
	testHash(t, gkeys, 32, MetroHash)
}
func TestJavaHash_16(t *testing.T) {
	testHash(t, gkeys, 16, JavaHash)
}
func TestJavaHash_32(t *testing.T) {
	testHash(t, gkeys, 32, JavaHash)
}
func TestDJBHash_16(t *testing.T) {
	testHash(t, gkeys, 16, DJBHash)
}
func TestDJBHash_32(t *testing.T) {
	testHash(t, gkeys, 32, DJBHash)
}
func TestSuperFastHash_16(t *testing.T) {
	testHash(t, gkeys, 16, SuperFastHash)
}
func TestSuperFastHash_32(t *testing.T) {
	testHash(t, gkeys, 32, SuperFastHash)
}

func benchmarkHash(b *testing.B, keys []string, hf HashFunc) {
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			hf(k)
		}
	}
}

func BenchmarkBKDRHash1(b *testing.B) {
	benchmarkHash(b, gkeys, Hash1)
}
func BenchmarkBKDRHash2(b *testing.B) {
	benchmarkHash(b, gkeys, Hash2)
}
func BenchmarkBKDRHash3(b *testing.B) {
	benchmarkHash(b, gkeys, Hash3)
}
func BenchmarkSDBMHash(b *testing.B) {
	benchmarkHash(b, gkeys, SDBMHash)
}
func BenchmarkAPHash(b *testing.B) {
	benchmarkHash(b, gkeys, APHash)
}
func BenchmarkMetroHash(b *testing.B) {
	benchmarkHash(b, gkeys, MetroHash)
}
func BenchmarkJavaHash(b *testing.B) {
	benchmarkHash(b, gkeys, JavaHash)
}
func BenchmarkDJBHash(b *testing.B) {
	benchmarkHash(b, gkeys, DJBHash)
}
func BenchmarkSuperFastHash(b *testing.B) {
	benchmarkHash(b, gkeys, SuperFastHash)
}

func TestMain(m *testing.M) {
	flag.Parse()
	gkeys = loadKeys()
	os.Exit(m.Run())
}
