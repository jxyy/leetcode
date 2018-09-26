#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    int smallestRangeII(vector<int>& A, int K) {
        int sum = 0;
        int max = 0, min = 10000;
        for(auto i : A) {
            sum += i;
            max = i > max ? i : max;
            min = i < min ? i : min;
        }
        int origin_range = max - min;
        if(origin_range <= K) {
            return origin_range;
        }

        float mean = float(sum) / A.size();
        max = 0;
        min = 10000;
        for(auto i : A) {
            if(i > mean) {
                i -= K;
                max = i > max ? i : max;
                min = i < min ? i : min;
            }else{
                i += K;
                max = i > max ? i : max;
                min = i < min ? i : min;
            }
        }
        return max - min;
    }
};

int main(){
    Solution s;
    vector<int> A;
    A.push_back(1);
    cout << s.smallestRangeII(A, 0) << endl;

    A.clear();
    A.push_back(0);
    A.push_back(10);
    cout << s.smallestRangeII(A, 2) << endl;

    A.clear();
    A.push_back(1);
    A.push_back(3);
    A.push_back(6);
    cout << s.smallestRangeII(A, 3) << endl;

    A.clear();
    A.push_back(7);
    A.push_back(8);
    A.push_back(8);
    cout << s.smallestRangeII(A, 5) << endl;

    A.clear();
    A.push_back(7);
    A.push_back(8);
    A.push_back(8);
    A.push_back(5);
    A.push_back(2);
    cout << s.smallestRangeII(A, 4) << endl;
    return 0;
}
