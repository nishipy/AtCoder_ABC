A,B,K=map(int,input().split())
l = []
for i in reversed(range(0, min(A, B)+1)):
    if (A%i==0)&(B%i==0):
        l.append(i)
    if len(l)==K:
        break        
print(l[-1])