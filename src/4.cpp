#include<iostream>
#include<vector>
using namespace std;

class Solution{
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2){
        if(nums1.size() == 0){
            return nums2.size()%2==1? nums2[nums2.size()/2] : (nums2[nums2.size()/2]+nums2[nums2.size()/2-1])/2.0f;
        }
        if(nums2.size() == 0){
            return nums1.size()%2==1? nums1[nums1.size()/2] : (nums1[nums1.size()/2]+nums1[nums1.size()/2-1])/2.0f;
        }
        int left1 = 0, left2 = 0;
        int right1 = nums1.size(), right2 = nums2.size();
        while(right1-left1>2 && right2-left2>2){
            int median_index1 = (right1-left1)%2==1? left1+(right1-left1)/2 : left1+(right1-left1)/2-1;
            int median_index2 = (right2-left2)%2==1? left2+(right2-left2)/2 : left2+(right2-left2)/2-1;
            if(nums1[median_index1] >= nums2[median_index2]){
                int cut_length1 = right1 - median_index1;
                int cut_length2 = median_index2 - left2;
                int cut_length = (cut_length1<cut_length2? cut_length1 : cut_length2);
                right1 -= cut_length;
                left2 += cut_length;
            }else{
                int cut_length1 = median_index1 - left1;
                int cut_length2 = right2 - median_index2;
                int cut_length = (cut_length1<cut_length2? cut_length1 : cut_length2);
                left1 += cut_length;
                right2 -= cut_length;
            }
        }
        if(right1-left1==1 && right2-left2==1){
            return (nums1[left1]+nums2[left2])/2.0f;
        }
        if(right1-left1==2 && right2-left2==2){
            int num1 = nums1[left1], num2 = nums1[left1+1], num3 = nums2[left2], num4 = nums2[left2+1];
            int max1 = num1>num2? num1 : num2, max2 = num3>num4? num3 : num4;
            int min1 = num1<num2? num1 : num2, min2 = num3<num4? num3 : num4;
            int max_min = min1>min2? min1 : min2;
            int min_max = max1<max2? max1 : max2;
            return (max_min+min_max)/2.0f;
        }

        int num = (right1-left1 == 1? nums1[left1] : nums2[left2]);
        vector<int> nums = (right1-left1 == 1? nums2 : nums1);
        int right = (right1-left1 == 1? right2 : right1);
        int left = (right1-left1 == 1? left2 : left1);
        if(right-left%2 == 1){
            if(num > nums[left+(right-left)/2+1]){
                return (nums[left+(right-left)/2]+nums[left+(right-left)/2+1])/2.0f;
            }else if(num < nums[left+(right-left)/2-1]){
                return (nums[left+(right-left)/2]+nums[left+(right-left)/2-1])/2.0f;
            }else{
                return (nums[left+(right-left)/2]+num)/2.0f;
            }
        }else{
            if(num > nums[left+(right-left)/2]){
                return nums[left+(right-left)/2];
            }else if(num > nums[left+(right-left)/2-1]){
                return num;
            }else{
                return nums[left+(right-left)/2-1];
            }
        }
    }
};

int main(){
    vector<int> v1, v2;
    v1.push_back(1);
    //v1.push_back(1);
    v2.push_back(1);
    v2.push_back(2);
    Solution s;
    cout << s.findMedianSortedArrays(v1, v2) << endl;
    return 0;
}
