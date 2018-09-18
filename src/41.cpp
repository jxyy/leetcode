#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        int n= nums.size();
        for(int i = 0; i < n; i++) {
            while(nums[i] > 0 && nums[i]-1 < n && nums[i] != i+1 && nums[nums[i]-1] != nums[i]) {
                int t = nums[nums[i]-1];
                nums[nums[i]-1] = nums[i];
                nums[i] = t;
            }
        }
        for(int i = 0; i < n; i++) {
            if(nums[i] != i+1) {
                return i+1;
            }
        }
        return n+1;
    }
};

int main(){
    Solution s;
    vector<int> nums;
    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(0);
    cout << s.firstMissingPositive(nums) << endl;

    nums.clear();
    nums.push_back(3);
    nums.push_back(4);
    nums.push_back(-1);
    nums.push_back(1);
    cout << s.firstMissingPositive(nums) << endl;

    nums.clear();
    nums.push_back(7);
    nums.push_back(8);
    nums.push_back(9);
    nums.push_back(10);
    nums.push_back(11);
    cout << s.firstMissingPositive(nums) << endl;

    nums.clear();
    nums.push_back(1);
    nums.push_back(1);
    cout << s.firstMissingPositive(nums) << endl;
    return 0;
}
