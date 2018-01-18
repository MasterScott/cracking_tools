#!/usr/bin/env python

from operator import itemgetter
import string
import sys



def getMask(text):
    mask = ""
    for c in text:
        if c in string.digits:
            mask += "d"
        elif c in string.ascii_uppercase:
            mask += "u"
        elif c in string.ascii_lowercase:
            mask += "l"
        else:
            mask += "s"

    return mask


def printTopMasks(masks, top=10):
    i = 0
    for key, value in sorted(masks.iteritems(), key=lambda (k,v): (v,k), reverse=True):
        if i < top:
            print "%s: %s" % (key, value)
        else:
            break

        i += 1

if __name__ == '__main__':
    masks = dict()
    with open(sys.argv[1]) as fin:
        for line in fin:
            line = line.strip()

            mask = getMask(line)
            m = masks.get(mask, 0)
            m += 1
            masks[mask] = m

        printTopMasks(masks, 10)
