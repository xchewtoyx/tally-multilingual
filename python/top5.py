#!/usr/bin/env python

from collections import Counter
import sys

def readnumbers(filename):
    with open(filename) as input:
        for line in input:
            yield line.strip()

def counts(filename):
    counter = Counter()
    for number in readnumbers(filename):
        counter[number] += 1
    return counter

def top5(filename):
    return counts(filename).most_common(5)

def main():
    for pair in top5(sys.argv[1]):
        print "%4s %s" % pair

if __name__ == '__main__':
    main()
