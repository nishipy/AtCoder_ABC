import numpy as np
N = int(input())
h = np.array([int(i) for i in input().split()])
count=0
 
while(np.max(h)>0):
    start=0
    i=0
    while(h[i]==0):
        i+=1
    start=i
    while(h[i]!=0 and i<N-1):
        i+=1
    if h[i]==0:
        h[start:i]=h[start:i]-1
    else:
        h[start:]=h[start:]-1
    count+=1
 
print(count)
