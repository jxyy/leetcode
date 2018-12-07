#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int singleNonDuplicate(vector<int>& nums) {
        int low = 0, high = nums.size();
        while(high - low > 1) {
            int half = (high - low) / 2; 
            int mid = low + half;
            if(half % 2 == 1){
                if(nums[mid] == nums[mid-1])
                    low = mid + 1;
                else
                    high = mid;
            } else {
                if(nums[mid] == nums[mid+1])
                    low = mid + 2;
                else
                    high = mid - 1;
            }
        }
        return nums[low];
    }
};

int main(){
    Solution s;
    int nums[] = {1,1,2,3,3,4,4,8,8};
    vector<int> input(nums, nums + 9);
    cout << s.singleNonDuplicate(input) << endl;

    int nums2[] = {3,3,7,7,10,11,11};
    vector<int> input2(nums2, nums2 + 7);
    cout << s.singleNonDuplicate(input2) << endl;

    int nums3[] = {1, 1, 2, 2, 4, 4, 5, 5,9};
    vector<int> input3(nums3, nums3 + 9);
    cout << s.singleNonDuplicate(input3) << endl;
    return 0;
}
