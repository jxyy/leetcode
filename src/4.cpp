#include<iostream>
#include<vector>
using namespace std;

class Solution{
public:
    int start1, start2;

    double findKth(vector<int>& nums1, vector<int>& nums2, int k){
        if(nums1.size()-start1 == 0){
            start2 = start2+k-1;
            return nums2[start2];
        }else if(nums2.size()-start2 == 0){
            start1 = start1+k-1;
            return nums1[start1];
        }else if(k == 1){
            return nums1[start1]<nums2[start2]? nums1[start1] : nums2[start2];
        }
        int offset1 = k/2, offset2 = k - k/2;
        if(offset1 > nums1.size()-start1){
            offset1 = nums1.size()-start1;
            offset2 = k - offset1;
        }else if(offset2 > nums2.size()-start2){
            offset2 = nums2.size()-start2;
            offset1 = k - offset2;
        }
        int index1 = start1 + offset1 - 1, index2 = start2 + offset2 - 1;
        if(nums1[index1] > nums2[index2]){
            start2 += offset2;
            return findKth(nums1, nums2, k-offset2);
        }else{
            start1 += offset1;
            return findKth(nums1, nums2, k-offset1);
        }
    }

    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2){
        start1 = start2 = 0;
        if((nums1.size()+nums2.size())%2 == 1){
            return findKth(nums1, nums2, (nums1.size()+nums2.size())/2+1);
        }else{
            findKth(nums1, nums2, (nums1.size()+nums2.size())/2);
            if(nums1.size()-start1 == 0){
                return (nums2[start2] + nums2[start2+1])/2.0;
            }else if(nums2.size()-start2 == 0){
                return (nums1[start1] + nums1[start1+1])/2.0;
            }else{
                if(nums1[start1]<nums2[start2]){
                    if(start1+1 < nums1.size() && nums1[start1+1] < nums2[start2]){
                        return (nums1[start1] + nums1[start1+1])/2.0;
                    }else{
                        return (nums1[start1] + nums2[start2])/2.0;
                    }
                }else{
                    if(start2+1 < nums2.size() && nums1[start1] > nums2[start2+1]){
                        return (nums2[start2] + nums2[start2+1])/2.0;
                    }else{
                        return (nums2[start2] + nums1[start1])/2.0;
                    }
                }
            }
        }
    }
};

int main(){
    vector<int> v1, v2;
    v1.push_back(1);
    v2.push_back(2);
    v2.push_back(3);
    v2.push_back(4);
    v2.push_back(5);
    v2.push_back(6);
    Solution s;
    cout << s.findMedianSortedArrays(v1, v2) << endl;
    return 0;
}
