#include<iostream>
#include<vector>
#include<hash_map>
#include<map>

using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        vector<int> rtn;
        map<int, int> another;
        for(int i = 0; i < nums.size(); i++){
            if(another.find(nums[i]) != another.end()){
                rtn.push_back(another[nums[i]]+1);
                rtn.push_back(i+1);
                break;
            }else{
                another[target-nums[i]] = i;
            }
        }
        return rtn;
    }
};

int main(){
    return 0;
}
