#include<iostream>
#include<vector>
#include<math.h>
#include<bitset>
using namespace std;

class Solution {
    
public:
    vector<int> grayCode(int n) {
        int count = 1;
        int total = pow(2, n);
        int prev = 0;
        vector<bool> flag(total, false);
        flag[0] = true;
        vector<int> ret(1, 0);
        while(count < total) {
            for(int i = 0; i < n; i++) {
                int cur = prev ^ (1 << i);
                if (!flag[cur]) {
                    ret.push_back(cur);
                    flag[cur] = true;
                    prev = cur;
                    count ++;
                    break;
                }
            }
        }
        return ret;
    }
};

void print(const vector<int>& v) {
    for(int i = 0; i < v.size(); i++) {
        bitset<5> bs(v[i]);
        cout << bs << endl;
    }
    cout << endl;
}

int main(){
    Solution s;
    print(s.grayCode(2));
    print(s.grayCode(3));
    print(s.grayCode(4));
    print(s.grayCode(5));
    return 0;
}
