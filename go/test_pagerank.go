package main

import (
	"./pagerank"
	"bufio"
	"flag"
	"fmt"
	"os"
	"time"
)

const (
	DP    = 0.85
	EPS   = 1e-6
	OFILE = "output.pagerank"
)

func constructGraph(fnGraph string) *pagerank.Graph {
	fGraph, err := os.Open(fnGraph)
	if err != nil {
		fmt.Println("cannot open file to read! filename:", fnGraph)
		fmt.Println(err)
		os.Exit(1)
	}

	bGraph := bufio.NewReader(fGraph)
	size := 0
	for {
		var prefix string
		_, err := fmt.Fscan(bGraph, &prefix)
		if err != nil {
			panic(err)
		}
		if prefix[0] == '#' {
			_, err := fmt.Fscan(bGraph, &size)
			if err != nil {
				panic(err)
			}
			break
		}
	}

	g := pagerank.NewGraph(size)

	// read all edges
	for {
		var idx, outDeg int
		_, err := fmt.Fscan(bGraph, &idx)
		if err != nil {
			break
		} else {
			_, err := fmt.Fscanf(bGraph, ":%d", &outDeg)
			if err != nil {
				fmt.Println(err)
				break
			}
		}
		idx--
		g.OutDegree[idx] = outDeg
		for i := 0; i < outDeg; i++ {
			var oIdx int
			fmt.Fscan(bGraph, &oIdx)
			g.InEdges[oIdx-1] = append(g.InEdges[oIdx-1], idx)
		}
	}

	if err := fGraph.Close(); err != nil {
		fmt.Println("cannot close file!")
		fmt.Println(err)
		os.Exit(1)
	}

	// record all nodes with 0 out degree
	for idx, dg := range g.OutDegree {
		if dg == 0 {
			g.EmptyNodes = append(g.EmptyNodes, idx)
		}
	}
	return g
}

func outputPagerank(pg []float64, fnRank string) {
	fRank, err := os.Create(fnRank)
	if err != nil {
		fmt.Println("cannot create file to write! filename:", fnRank)
		fmt.Println(err)
		os.Exit(1)
	}
	for i, v := range pg {
		fmt.Fprintf(fRank, "%d:%.16f\n", i+1, v)
	}
	if err := fRank.Close(); err != nil {
		fmt.Println("cannot close file!")
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// process command line arguments

	dPtr := flag.Float64("d", DP, "damping factor")
	epsPtr := flag.Float64("e", EPS, "epsilon")
	outputPtr := flag.String("o", OFILE, "output file")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Println("please specify input file")
		os.Exit(1)
	}
	fnGraph := flag.Arg(0)
	fnRank := *outputPtr
	dNum := *dPtr
	eps := *epsPtr

	fmt.Printf("executing with damping factor %f, epsilon %f\n",
		dNum, eps)

	// read the file

	fmt.Println("constructing graph...")

	g := constructGraph(fnGraph)

	fmt.Printf("%d nodes loaded.\n", g.NNode)

	//

	fmt.Println("start pageranking...")

	startTime := time.Now()
	pg := pagerank.Pagerank(g, dNum, eps)
	endTime := time.Now()

	fmt.Println("Time spent:", endTime.Sub(startTime))

	//

	fmt.Println("store result...")

	outputPagerank(pg, fnRank)
}
