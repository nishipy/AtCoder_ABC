import itertools

N, A, B, C = map(int,input().split())
length_list = [int(input()) for i in range(0, N)]

use_flag = ['0', '1', '2', '-1']
use_list = list(itertools.product(use_flag, repeat=len(length_list)))

costs = []

for ls in use_list:
    if len(set(ls)&{'0','1', '2'}) == 3:
        #print(ls)
        A_diff = 0
        B_diff = 0
        C_diff = 0
        for length, use_for in zip(length_list, list(ls)):
            #print(length, use_for)
            if use_for == '0':
                A_diff += length
                #print(A_diff)
            elif use_for == '1':
                B_diff += length
                #print(B_diff)
            elif use_for == '2':
                C_diff += length
                #print(C_diff)
            #print('-----')
        costs.append(abs(A-A_diff) + abs(B-B_diff) + abs(C-C_diff) + 10*max(ls.count('0') - 1, 0) + 10*max(ls.count('1') - 1, 0) + 10*max(ls.count('2') - 1, 0))
print(min(costs))
#print(len(costs))
