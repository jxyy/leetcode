#include<iostream>
#include<vector>
#include<algorithm>
using namespace std;

class Solution {
public:
    vector<pair<int, int> > twoSum(vector<int>& nums, int l, int r, int target) {
        //asume nums is sorted
        vector<pair<int, int> > ret;
        while(r > l){
            int sum = nums[l] + nums[r];
            if(sum == target){
                pair<int, int> p(nums[l], nums[r]);
                ret.push_back(p);
                do{
                    l ++;
                }while(r > l && nums[l] == nums[l-1]);
                do{
                    r --;
                }while(r > l && nums[r] == nums[r+1]);
            }else if(sum < target){
                do{
                    l ++;
                }while(r > l && nums[l] == nums[l-1]);
            }else{
                do{
                    r --;
                }while(r > l && nums[r] == nums[r+1]);
            }
        }
        return ret;
    }

    vector<vector<int> > threeSum(vector<int>& nums) {
        vector<vector<int> > ret;
        if(nums.size() < 3){
            return ret;
        }
        sort(nums.begin(), nums.end());
        for(int i = 0; i < nums.size() - 2; ){
            vector<pair<int, int> > pairs = twoSum(nums, i + 1, nums.size() - 1, -nums[i]);
            for(int j = 0; j < pairs.size(); j++){
                vector<int> triplet;
                triplet.push_back(nums[i]);
                triplet.push_back(pairs[j].first);
                triplet.push_back(pairs[j].second);
                ret.push_back(triplet);
            }

            do{
                i ++;
            }while(i < nums.size() - 2 && nums[i] == nums[i-1]);
        }

        return ret;
    }
};

int main(){
    vector<int> nums;

    Solution s;
    /*
    //test twosum
    sort(nums.begin(), nums.end());
    vector<pair<int, int> > r = s.twoSum(nums, 0, nums.size() - 1, 0);
    for(int i = 0; i < r.size(); i++){
        cout << r[i].first << ' ' << r[i].second << endl;
    }
    */


    vector<vector<int> > triplets = s.threeSum(nums);
    for(int i = 0; i < triplets.size(); i++){
        cout << triplets[i][0] << ' ' << triplets[i][1] << ' ' << triplets[i][2] << endl;
    }
    return 0;
}
