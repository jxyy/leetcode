#include<iostream>
using namespace std;

class Solution{
public:
    int lengthOfLongestSubstring(string s){
        int max = 0, max_end_here = 0, max_end_start = 0;
        int flags[256];
        for(int i = 0; i < 256; i++){
            flags[i] = -1;
        }
        for(int i = 0; i < s.size(); i++){
            int index = int(s[i]);
            if(flags[index] != -1){
                max_end_here = i - flags[index];
                for(; max_end_start < flags[index]+1; max_end_start++){
                    flags[int(s[max_end_start])] = -1;
                }
                flags[index] = i;
            }else{
                flags[index] = i;
                max_end_here++;
                if(max_end_here > max){
                    max = max_end_here;
                }
            }
        }
        return max;
    }
};

int main(){
    Solution s = Solution();
    return 0;
}
