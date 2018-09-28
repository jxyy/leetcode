#include<iostream>
#include<vector>
#include<math.h>
using namespace std;

void print(vector<int>& v) {
    for(auto i : v) {
        cout << i << " ";
    }
    cout << endl;
}

class Solution {
public:
    vector<int> partitionLabels(string S) {
        vector<int> ends(26, -1);
        for(int i = 0; i < S.length(); i ++) {
            int ci = S[i] - 97;
            ends[ci] = ends[ci] > i ? ends[ci] : i;
        }
        int prev_break = -1;
        int m = 0;
        vector<int> ret;
        for(int i = 0; i < S.length(); i++) {
            m = m > ends[S[i]-97] ? m : ends[S[i]-97];
            if (i == m) {
                ret.push_back(i - prev_break);
                prev_break = i;
                m = 0;
            }
        }
        return ret;
    }
};

int main(){
    Solution s;
    print(s.partitionLabels("ababcbacadefegdehijhklij"));
    return 0;
}
