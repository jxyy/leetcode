#include<iostream>
using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        if(x < 0){
            return false;
        }
        int index1, index2;
        int digit_num = 0;
        int n = x;
        while(n != 0){
            digit_num ++;
            n /= 10;
        }
        int i = 0;
        int sum = 0, unit = 1;
        n = x;
        for(; i < digit_num/2; i++){
            sum += (n%10) * unit;
            n /= 10;
            unit *= 10;
        }
        cout << sum << endl;
        if(digit_num%2 == 1){
            i++;
            n /= 10;
        }
        for(; i < digit_num; i++){
            unit /= 10;
            sum -= (n%10) * unit;
            n /= 10;
        }
        cout << sum << endl;
        return sum==0;
    }
};

int main(){
    Solution s;
    cout << s.isPalindrome(1234321) << endl;
    return 0;
}
