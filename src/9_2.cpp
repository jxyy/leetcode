#include<iostream>
#include<vector>
using namespace std;

class Solution {
public:
    bool isPalindrome(int x) {
        if(x < 0){
            return false;
        }
        int digits[20];
        int digit_num = 0;
        int n = x;
        while(n != 0){
            digits[digit_num] = n % 10;
            digit_num ++;
            n /= 10;
        }
        for(int i = 0; i < digit_num/2; i++){
            int j = digit_num - i - 1;
            if(digits[i] != digits[j]){
                return false;
            }
        }
        return true;
    }
};

int main(){
    Solution s;
    cout << s.isPalindrome(1234321) << endl;
    return 0;
}

