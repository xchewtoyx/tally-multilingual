#!/usr/bin/env python

from collections import defaultdict
import sys

def readnumbers(filename):
    with open(filename) as input:
        for line in input:
            yield line.strip()

def tally_counts(filename):
    counter = defaultdict(int)
    for number in readnumbers(filename):
        counter[number] += 1
    for key, count in counter.items():
        yield count, key

def top5(filename):
    counts = tally_counts(filename)
    most_common = sorted(tally_counts(filename), reverse=True)
    return most_common[:5]

def main():
    for count, key in top5(sys.argv[1]):
        print "%4s %s" % (key, count)

if __name__ == '__main__':
    main()
