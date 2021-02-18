import math

s = 0
with open('input1.txt') as f:
    for line in f:
        s += int(line)/3-2
print(s) # 3330521

s = 0
with open('input1.txt') as f:
    for line in f:
        tmp = int(line)/3-2
        s += tmp
        while True:
            if tmp/3-2 > 0:
                tmp = tmp/3-2
                s += tmp
            else:
                break

print(s) # 4992931
