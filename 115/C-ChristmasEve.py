import numpy as np
N, K = map(int,input().split())
H = np.array([int(input()) for _ in range(0, N)])
H = np.sort(H)
H_diff = np.diff(H)
best = int(np.min(np.convolve((H_diff), np.ones(K-1), mode='valid')))
print(best)
