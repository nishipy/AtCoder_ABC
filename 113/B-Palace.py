import numpy as np
N = int(input())
T, A = map(int,input().split())
H = np.array([int(i) for i in input().split()])
H = T - H * 0.006
H = np.abs(H - A)
print(H.argmin()+1)
