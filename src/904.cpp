#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int totalFruit(vector<int>& tree) {
        int tp1 = -1, tp2 = -1;
        int prev_tp = -1, prev_tp_count = 0;
        int max_amount = 0;
        int max_amount_yet = 0;
        for(int i = 0; i < tree.size(); i++) {
            int tp = tree[i];
            if (tp1 == -1 || tp1 == tp) {
                tp1 = tp;
                max_amount_yet ++;
            } else if(tp2 == -1 || tp2 == tp) {
                tp2 = tp;
                max_amount_yet ++;
            } else {
                if(max_amount_yet > max_amount) {
                    max_amount = max_amount_yet;
                }
                max_amount_yet = prev_tp_count + 1;
                tp1 = prev_tp;
                tp2 = tp;
            }
            if (tp == prev_tp) {
                prev_tp_count ++;
            } else {
                prev_tp_count = 1;
            }
            prev_tp = tp;
        }
        return max_amount > max_amount_yet ? max_amount : max_amount_yet;
    }
};

int main(){
    Solution s;
    int nums[] = {1,0,1,4,1,4,1,2,3};
    vector<int> input(nums, nums + 9);
    cout << s.totalFruit(input) << endl;
    return 0;
}
