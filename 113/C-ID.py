import numpy as np
N, M = map(int,input().split())
P_Y_list = np.array([input().split() for _ in range(0, M)], dtype='int')
counts = (np.zeros(N+1)).astype(int)
id_list = []

sort_ind = np.argsort(P_Y_list[:,1])
sorted_P_Y = P_Y_list[sort_ind]

for p, y in sorted_P_Y:
    counts[p] += 1
    id_list.append(str(p).zfill(6) + str(counts[p]).zfill(6))
id_list = np.array(id_list)
id_list = id_list[np.argsort(sort_ind)]
for id in id_list:
    print(id)
