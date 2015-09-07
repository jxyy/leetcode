#include<iostream>
using namespace std;

class Solution{
public:
    bool isMatch(string s, string p){
        int m = p.length(), n = s.length();
        bool matches[m+1][n+1];
        matches[0][0] = true;
        bool maybe_empty = true;
        for(int i = 1; i < m+1; i++){
            matches[i][0] = false || (p[i-1] == '*' && matches[i-2][0]);
        }
        for(int j = 1; j < n+1; j++){
            matches[0][j] = false;
        }
        for(int i = 1; i < m+1; i++){
            for(int j = 1; j < n+1; j++){
                if(p[i-1] == '*'){
                    int pc = p[i-2];
                    matches[i][j] = matches[i-1][j] || matches[i-2][j] || (matches[i][j-1] && (pc == '.' || pc == s[j-1]));
                }else{
                    int pc = p[i-1];
                    matches[i][j] = matches[i-1][j-1] && (pc == '.' || pc == s[j-1]);
                }
            }
        }

        return matches[m][n];
    }
};

int main(){
    Solution s;
    cout << s.isMatch("aaa", ".*") << endl;
    return 0;
}
