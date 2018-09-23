#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int smallestRangeI(vector<int>& A, int K) {
        if(A.size() == 1) {
            return 0;
        }
        int max_value = 0, min_value = A[0];
        for(auto x : A) {

            max_value = x > max_value ? x : max_value;
            min_value = x < min_value ? x : min_value;
        }
        int dv = (max_value - K ) - (min_value + K);
        return dv > 0 ? dv : 0;
    }
};

int main(){
    Solution s;
    vector<int> A;
    A.push_back(1);
    cout << s.smallestRangeI(A, 0) << endl;

    A.clear();
    A.push_back(0);
    A.push_back(10);
    cout << s.smallestRangeI(A, 2) << endl;

    A.clear();
    A.push_back(1);
    A.push_back(3);
    A.push_back(6);
    cout << s.smallestRangeI(A, 3) << endl;

    A.clear();
    A.push_back(2);
    A.push_back(7);
    A.push_back(2);
    cout << s.smallestRangeI(A, 1) << endl;
    return 0;
}
