import argparse
import time

import pagerank

DP = 0.85
EPS = 1e-6
OFILE = 'output.pagerank'

def load_graph(fname):
    with open(fname, 'r') as f:
        size = int(f.readline().split()[1])

        g = pagerank.Graph(size)

        # read all edges
        for l in f:
            tks = l.split()
            idx, outDeg = tks[0].split(':')
            idx = int(idx) - 1
            g.outDegree[idx] = int(outDeg)
            for oIdx in tks[1:]:
                oIdx = int(oIdx) - 1
                g.inEdges[oIdx].append(idx)

    # record all nodes with 0 out degree
    for i in range(g.nNode):
            if g.outDegree[i] == 0:
                    g.emptyNodes.append(i)

    return g


def save_pagerank(pg, fname):
    with open(fname, 'w') as f:
        for i in range(len(pg)):
            f.write('{}:{}\n'.format(i+1, pg[i]))


if __name__ == '__main__':
    # process command line arguments

    parser = argparse.ArgumentParser(description='PageRank.')
    parser.add_argument('-d', default=DP, type=float,
            help='damping factor')
    parser.add_argument('-e', default=EPS, type=float,
            help='epsilon')
    parser.add_argument('-o', default=OFILE,
            help='output file')
    parser.add_argument('ifile', metavar='input-file', nargs=1,
            help='input file')
    args = parser.parse_args()

    print('executing with damping factor {}, epsilon {}'.format(
        args.d, args.e))

    # read the file

    print('constructing graph...')

    g = load_graph(args.ifile[0])

    print('{} nodes loaded.'.format(g.nNode))

    #

    print('start pageranking...')

    startTime = time.perf_counter()
    pg = pagerank.pagerank(g, args.d, args.e)
    endTime = time.perf_counter()

    print('Time spent: {}s'.format(endTime-startTime))

    #

    print('store result...')

    save_pagerank(pg, args.o)
