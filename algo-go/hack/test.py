import time
import sys
import re

FMT = '%Y/%m/%d %H:%M:%S'
STR = re.compile(
    r'^\d{4}(\-|\/|.)\d{1,2}(\-|\/|.)\d{1,2} \d{1,2}:\d{1,2}:\d{1,2}$')
NUM = re.compile(r'^\d+$')


def output(data):
    sys.stdout.write(str(data))


def main():
    if len(sys.argv) > 1:
        data = sys.argv[1].strip()
        if re.match(STR, data):
            t = time.strptime(data, FMT)
            output(int(time.mktime(t)))
            return 0
        if re.match(NUM, data):
            t = time.localtime(int(data))
            output(time.strftime(FMT, t))
            return 0

    output(int(time.time()))


if __name__ == '__main__':
    main()
