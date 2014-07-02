import array
import math

from collections import defaultdict


class Graph(object):
    def __init__(self, size):
        self.inEdges = [[] for _ in range(size)]
        self.outDegree = [0] * size
        self.emptyNodes = []
        self.nNode = size


def pagerank(graph, d, eps):
    glen = graph.nNode
    pg1, pg2 = [0.0]*glen, [0.0]*glen
    weight1, weight2 = [0.0]*glen, [0.0]*glen

    # initialize pagerank
    for i in range(glen):
        pg1[i] = 1.0
        if graph.outDegree[i] != 0:
            weight1[i] = 1.0 / graph.outDegree[i]

    # power iteration
    while True:
        totalE = 0.0
        for idx in graph.emptyNodes:
            totalE += pg1[idx]
        totalE /= glen

        diff = 0.0
        for i in range(glen):
            w = totalE
            for idx in graph.inEdges[i]:
                w += weight1[idx]
            pg2[i] = (1.0 - d) + d*w
            if graph.outDegree[i] != 0:
                weight2[i] = pg2[i] / graph.outDegree[i]
            diff += pow(pg1[i]-pg2[i], 2)

        if math.sqrt(diff) < eps:
            break

        pg1, pg2 = pg2, pg1
        weight1, weight2 = weight2, weight1

    return pg2
