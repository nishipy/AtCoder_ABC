import numpy as np
N=int(input())
info = np.array([input().split() for _ in range(0, N)], dtype='int')
info = info[np.argsort(info[:,2])][::-1]
start_x = info[0,:][0]
start_y = info[0,:][1]
start_h = info[0,:][2]

for Cx in range(101):
    for Cy in range(101):
        H = start_h+abs(start_x-Cx)+abs(start_y-Cy)
        if all(max(0,H-abs(x-Cx)-abs(y-Cy))==h for x,y,h in info):
            print(Cx,Cy,H)
            break
        else:
            continue
        break