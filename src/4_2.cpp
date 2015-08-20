#include<iostream>
#include<vector>
using namespace std;

class Solution{
public:
    double findKth(vector<int>& nums1, vector<int>& nums2, int k, int start1, int start2){
        if(nums1.size()-start1 == 0){
            return nums2[start2+k-1];
        }else if(nums2.size()-start2 == 0){
            return nums1[start1+k-1];
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
            return findKth(nums1, nums2, k-offset2, start1, start2+offset2);
        }else{
            return findKth(nums1, nums2, k-offset1, start1+offset1, start2);
        }
    }

    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2){
        if((nums1.size()+nums2.size())%2 == 1){
            return findKth(nums1, nums2, (nums1.size()+nums2.size())/2+1, 0, 0);
        }else{
            double m1 = findKth(nums1, nums2, (nums1.size()+nums2.size())/2, 0, 0);
            cout << m1 << endl;
            double m2 = findKth(nums1, nums2, (nums1.size()+nums2.size())/2+1, 0, 0);
            cout << m2 << endl;
            return (m1 + m2)/2;
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
