#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int pivotIndex(vector<int>& nums) {
        vector<int> left_sums, right_sums;
        int left_sum = 0, right_sum = 0;
        int size = nums.size();
        for (int i = 0; i < size; i ++) {
            left_sums.push_back(left_sum);
            left_sum += nums[i];
            right_sums.push_back(right_sum);
            right_sum += nums[size-1-i];
        }
        
        for (int i = 0; i < size; i++) {
            if (left_sums[i] == right_sums[size-1-i]) {
                return i;
            }
        }
        return -1;
    }
};

int main(){
    Solution s;
    int input[] = {1, 7, 3, 6, 5, 6};
    vector<int> nums(input, input+6);
    cout << s.pivotIndex(nums) << endl;

    nums.clear();
    nums.push_back(1);
    nums.push_back(2);
    nums.push_back(3);
    cout << s.pivotIndex(nums) << endl;
    return 0;
}
