#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:

    int searchInsert(vector<int>& nums, int target) {
        int l = 0, r = nums.size();
        while(l < r){
            int m = (l + r) / 2;
            int vm = nums[m];
            if(target > vm){
                l = m + 1;
            }else if(target < vm){
                r = m;
            }else{
                return m;
            }
        }
        return l;
    }
};

int main(){
    vector<int> v;
    v.push_back(1);
    v.push_back(3);
    v.push_back(5);
    v.push_back(6);

    Solution s;
    int i = s.searchInsert(v, 99);
    cout << i << endl;
    return 0;
}
