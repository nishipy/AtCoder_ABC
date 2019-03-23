import numpy as np
N, T = map(int,input().split())
c_t = np.array([input().split() for _ in range(0, N)])
c_t = c_t.astype(int)

if np.min(c_t[:,1]) <= T:
    print(np.sort(c_t[np.where(c_t[:,1] <= T)][:,0])[0])
else:
    print('TLE')
