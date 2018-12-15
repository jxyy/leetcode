#include<iostream>
#include<vector>
using namespace std;

class Solution {
    bool isValidChar(char c){
        return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9';
    }
    bool isEqual(char c1, char c2){
        return c1 == c2 || c1-32 == c2 && c2 >= 'A' || c1+32 == c2 && c1 >= 'A';
    }
public:
    bool isPalindrome(string s) {
        int lo = 0, hi = s.size()-1;
        while(lo < hi) {
            while(!isValidChar(s[lo]) && lo < hi){
                lo++;
            }
            while(!isValidChar(s[hi]) && lo < hi){
                hi--;
            }
            if(!isEqual(s[lo], s[hi])){
                return false;
            }
            lo++;
            hi--;
        }
        return true;
    }
};

int main(){
    Solution s;

    cout << s.isPalindrome("") << endl;
    cout << s.isPalindrome("A man, a plan, a canal: Panama") << endl;
    cout << s.isPalindrome("race a car") << endl;
    cout << s.isPalindrome("0P") << endl;
    return 0;
}
