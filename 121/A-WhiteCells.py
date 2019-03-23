H_W = [int(x.strip()) for x in input().split()]
h_w = [int(x.strip()) for x in input().split()]
H = H_W[0]
W = H_W[1]
h = h_w[0]
w = h_w[1]
 
print(H*W - h*W - (H - h)*w)
