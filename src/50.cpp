#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    double myPow(double x, int n) {
        if(n == INT_MIN){
            return 1 / myPow(x, INT_MAX) / x;
        }
        if(n < 0){
            return 1 / myPow(x, -n);
        }
        if(n == 0){
            return 1;
        }
        if(n == 1) {
            return x;
        }
        int half = n / 2;
        double a = myPow(x, half);
        return n % 2 == 0 ? a*a:a*a*x;
    }
};

int main(){
    Solution s;
    cout << s.myPow(2, 10) << endl;
    cout << s.myPow(2.1, 3) << endl;
    cout << s.myPow(2, -2) << endl;
    cout << s.myPow(1, -2147483648) << endl;
    return 0;
}
