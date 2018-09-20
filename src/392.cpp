#include<iostream>
#include<vector>
#include<string>
using namespace std;

class Solution {
public:
    bool isSubsequence(string s, string t) {
        int is = 0, it = 0;
        int sl = s.length(), tl = t.length();
        for(; is < sl && it < tl; it ++) {
            if(s[is] == t[it]) {
                is ++;
            }
        }
        return is == sl;
    }
};

int main(){
    Solution s;
    cout << s.isSubsequence("ace", "abcdefg") << endl;
    cout << s.isSubsequence("abc", "ahbgdc") << endl;
    cout << s.isSubsequence("axc", "ahbgdc") << endl;
    return 0;
}
