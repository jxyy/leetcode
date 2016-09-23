#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:

    vector<int> searchRange(vector<int>& nums, int target) {
        vector<int> range;
        
        int l = 0, r = nums.size();
        while(l < r){
            int m = (l + r) / 2;
            int vm = nums[m];
            if(target <= vm){
                r = m;
            }else{
                l = m + 1;
            }
        }
        int start = l;

        l = 0;
        r = nums.size();
        while(l < r){
            int m = (l + r) / 2;
            int vm = nums[m];
            if(target < vm){
                r = m;
            }else{
                l = m + 1;
            }
        }
        int end = l;
        if(start == end){
            range.push_back(-1);
            range.push_back(-1);
        }else{
            range.push_back(start);
            range.push_back(end - 1);
        }
        return range;
    }
};

int main(){
    vector<int> v;
    v.push_back(5);
    v.push_back(7);
    v.push_back(7);
    v.push_back(8);
    v.push_back(8);
    v.push_back(10);

    Solution s;
    vector<int> r = s.searchRange(v, 5);
    cout << r[0] << "  " << r[1] << endl;
    return 0;
}
