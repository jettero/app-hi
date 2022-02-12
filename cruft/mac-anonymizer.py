#!/usr/bin/env python

import re
import sys
import random

MDB = dict()
MRE = re.compile(r'\b([a-fA-F0-9:]{17})\b')

def anonymize_mac(x):
    if r := MDB.get(x):
        return r
    MDB[x] = r = random_mac()
    return r

def random_mac():
    return ':'.join([ f'{random.randint(0, 255):02x}' for _ in range(6) ])

def process_stream(lines):
    for line in lines:
        for mac in MRE.findall(line):
            pretend_mac = anonymize_mac(mac)
            line = line.replace(mac, pretend_mac)
        yield line

if __name__ == '__main__':
    for line in process_stream(sys.stdin):
        sys.stdout.write(line)
