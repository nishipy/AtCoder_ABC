import numpy as np
N = int(input())
HPs = np.array([int(i) for i in input().split()])
np.sort(HPs)
while len(HPs) > 1:
    HPs[1:] %= HPs[0]
    HPs[HPs==0] = HPs[0]
    HPs=np.unique(HPs)
print(HPs[0])