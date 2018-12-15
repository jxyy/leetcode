#include<iostream>
#include<vector>
#include<algorithm>
#include<math.h>
using namespace std;

class Solution {
    int sumN(int n){
        return n*(n+1)/2;
    }
public:
    int reachNumber(int target) {
        target = abs(target);
        int estimate = int(sqrt(0.25 + 2.0f*target) - 0.5);
        if(sumN(estimate) == target) {
            return estimate;
        }
        int gap = sumN(estimate+1) - target;
        if(gap % 2 == 0){
            return estimate + 1;
        } else {
            return estimate + 3 - (estimate % 2);
        }
    }
};

int main(){
    Solution s;
    cout << s.reachNumber(3) << endl;
    cout << s.reachNumber(2) << endl;
    cout << s.reachNumber(-1) << endl;
    cout << s.reachNumber(1) << endl;
    cout << s.reachNumber(5) << endl;
    cout << s.reachNumber(4) << endl;
    cout << s.reachNumber(9) << endl;
    return 0;
}
