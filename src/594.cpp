#include<iostream>
#include<vector>
#include<map>
#include<algorithm>
#include<unordered_map>
using namespace std;

class Solution {
public:
    int findLHS(vector<int>& nums) {
        map<int, int> counter;
        for(int n : nums) {
            counter[n] ++;
        }
        
        map<int, int>::iterator prev = counter.begin();
        int max_ever = 0;
        for(auto curr = counter.begin(); curr != counter.end(); curr ++) {
            if(curr != counter.begin() && curr->first == prev->first + 1) {
                max_ever = max(max_ever, prev->second + curr->second);
            }
            prev = curr;
        }
        return max_ever;
    }
};

class Solution1 {
public:
    int findLHS(vector<int>& nums) {
        unordered_map<int, int> counter;
        for(int n : nums) {
            counter[n] ++;
        }
        
        int max_ever = 0;
        for(auto curr : counter) {
            if(counter.find(curr.first+1) != counter.end()) {
                max_ever = max(max_ever, curr.second + counter[curr.first+1]);
            }
        }
        return max_ever;
    }
};

int main(){
    Solution1 s;

    int input[] = {1,3,2,2,5,2,3,7};
    vector<int> nums(input, input+8);
    cout << s.findLHS(nums) << endl;

    int input2[] = {1,1,1,1};
    vector<int> nums2(input2, input2+4);
    cout << s.findLHS(nums2) << endl;
    return 0;
}
