import numpy as np
N, M = map(int,input().split())
target_list = np.array([int(i) for i in input().split()])
if N > M:
    print(0)
else:
    target_list = np.sort(target_list)
    diff_list = target_list[1:] - target_list[:-1]
    diff_list = np.sort(diff_list)
    if N > 1:
        diff_list = diff_list[0:-(N-1)]
    print(np.sum(diff_list))