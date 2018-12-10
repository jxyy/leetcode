#include<iostream>
#include<vector>
#include<string>
#include<unordered_map>
using namespace std;

class Solution {
    bool isSameCharset(string s1, string s2) {
        int cs[256] = {0};
        for(char c : s1) {
            cs[c] ++;
        }
        for(char c : s2) {
            if(--cs[c] < 0){
                return false;
            }
        }
        return true;
    }

public:
    bool isScramble(string s1, string s2) {
        int len = s1.size();
        if(len <= 1) {
            return s1 == s2;
        }
        if(len == 2) {
            return s1 == s2 || s1[0] == s2[1] && s1[1] == s2[0];
        }
        if(!isSameCharset(s1, s2)){
            return false;
        }
        for(int i = 1; i < len; i++) {
            if(isScramble(s1.substr(0, i), s2.substr(0, i)) && isScramble(s1.substr(i, len-i), s2.substr(i, len-i)) || 
            isScramble(s1.substr(0, i), s2.substr(len-i, i)) && isScramble(s1.substr(i, len-i), s2.substr(0, len-i))){
                return true;
            }
        }
        return false;
    }
};

int main(){
    Solution s;
    cout << s.isScramble("great", "rgeat") << endl;
    cout << s.isScramble("abcde", "caebd") << endl;
    cout << s.isScramble("abc", "cab") << endl;

    cout << s.isScramble("abcdefghijklmnopq", "efghijklmnopqcadb") << endl;
    cout << s.isScramble("ccabcbabcbabbbbcbb", "bbbbabccccbbbabcba") << endl;
    return 0;
}
