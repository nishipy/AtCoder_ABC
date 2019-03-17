N_M = [int(x.strip()) for x in input().split()]
price_l = [[int(i) for i in input().split()] for i in range(0, N_M[0])]
price_l.sort(key=lambda x: x[0])
total = 0
remain = N_M[1]
while remain > 0:
    if len(price_l)<1:
        break
    price_n = price_l.pop(0)
    total += price_n[0] * min(price_n[1], remain)
    remain -= min(price_n[1], remain)
print(total)
