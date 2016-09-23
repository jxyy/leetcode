#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int index(int start, int offset, int length){
        return (start + offset) % length;
    }

    int search(vector<int>& nums, int target) {
        int start = 0;
        int size = nums.size();
        int v0 = nums[0];
        int l = 1;
        int r = size;
        while(l != r){
            int m = l + (r - l) / 2;
            int vm = nums[m];
            if(vm < v0){
                r = m;
            }else if(vm > v0){
                l = m + 1;
            }        
        }
        start = l;

        l = 0;
        r = size;
        while(l != r){
            int m = l + (r - l) / 2;
            int im = index(start, m, size);
            int vm = nums[im];
            if(vm < target){
                l = m + 1;
            }else if(vm > target){
                r = m;
            }else{
                return im;
            }
        }
        return -1;
    }
};

int main(){
    vector<int> v;
    v.push_back(4);
    v.push_back(5);
    v.push_back(6);
    v.push_back(7);
    v.push_back(0);
    v.push_back(1);
    v.push_back(2);

    Solution s;
    for(int i = 0; i < 7; i++)
        cout << s.search(v, i) << endl;
    return 0;
}
